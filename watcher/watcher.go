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

const FILE_TYPES = ".*\\.(json|txt|csv|log)"

type (
	// File is represents a file in the dir watched by a Watcher.
	File struct {
		Name     string
		Checksum string // MD5 checksum of the file's content.
	}

	// Watcher keeps an eye on files within a specified dir.
	Watcher struct {
		Dir   string  // The directory to watch.
		Files []*File // A list of files within the directory.
	}
)

func init() {
	files, err := createFileList(*config.Watch)
	if err != nil {
		log.Fatalln("Error:", err)
	}

	for _, v := range files {
		sum, err := calcHashSum(v)
		if err != nil {
			log.Println("Error:", err)
		}
		log.Printf("%x", sum)
	}
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
				log.Println(dir.Name() + "/" + f.Name())
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
//func NewWatcher(dir string) *Watcher {
//}
