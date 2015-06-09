package watcher

import (
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/smotti/tad/config"
)

const FILE_TYPES = ".*\\.(json|txt|csv|log)"

// watcher keeps an eye on files within a specified dir.
type Watcher struct {
	Dir   string   // The directory to watch.
	Files []string // A list of files within the directory.
}

func init() {
	_, err := createFileList(*config.Watch)
	if err != nil {
		log.Fatalln("Error:", err)
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

// NewWatcher create a new watcher for the given directory.
//func NewWatcher(dir string) *Watcher {
//}
