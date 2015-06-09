package watcher

import (
	"bufio"
	"crypto/md5"
	"io"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/smotti/tad/config"
)

type (
	// File is represents a file in the dir watched by a Watcher.
	File struct {
		Name     string
		Checksum []byte // MD5 checksum of the file's content.
	}

	// Watcher keeps an eye on files within a specified dir.
	Watcher struct {
		Dir   string  // The directory to watch.
		Files []*File // A list of files within the directory.
	}
)

const FILE_TYPES = ".*\\.(json|txt|csv|log)"

func init() {
	watcher := NewWatcher(*config.Watch)
}

// createFileList creates a list of files within the given directory and returns
// a string slice.
func createFileList(s string) ([]string, error) {
	// Open the dir so we actually can get any info about it and its childs.
	dir, err := os.Open(strings.Trim(s, "/"))
	if err != nil {
		return nil, err
	}
	defer dir.Close()

	// Get FileInfo about the dir.
	di, err := dir.Stat()
	if err != nil {
		return nil, err
	}

	// Read the dirs content.
	var fi []os.FileInfo
	if di.IsDir() {
		fi, err = dir.Readdir(-1)
		if err != nil {
			return nil, err
		}
	}

	// Create the list of filenames (full path to file).
	files := make([]string, 0)
	for _, f := range fi {
		if !f.IsDir() {
			matched, err := regexp.MatchString(FILE_TYPES, f.Name())
			if err != nil {
				log.Println("Error:", err)
				continue
			}
			if matched {
				files = append(files, dir.Name()+"/"+f.Name())
			}
		}
	}

	return files, nil
}

// calcHashSum uses the MD5 hash algorithm to calculate a hash sum of the
// content for the given file. It returns the sum as a byte slice.
func calcHashSum(s string) ([]byte, error) {
	// Open the file.
	file, err := os.Open(s)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Get a new hash value.
	h := md5.New()

	// Read line by line and add to the hash sum.
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			return nil, err
		}
		io.WriteString(h, scanner.Text())
	}

	return h.Sum(nil), nil
}

// NewWatcher create a new watcher for the given directory.
func NewWatcher(dir string) *Watcher {
	files, err := createFileList(dir)
	if err != nil {
		log.Fatalln("Error:", err)
	}

	watcher := &Watcher{Dir: dir}
	for _, f := range files {
		h, err := calcHashSum(f)
		if err != nil {
			log.Println("Error:", err)
		}
		watcher.Files = append(watcher.Files, &File{Name: f, Checksum: h})
	}

	return watcher
}
