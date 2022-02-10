package path

import (
	"container/list"
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

const (
	FileNotFound = "File Not Found"
)

// Exists
func Exists(file string) (bool, error) {
	_, err := os.Stat(file)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		} else {
			return false, err
		}
	}

	return true, nil
}

func MoveFile(src, dst string) error {
	if exist, _ := Exists(src); !exist {
		return fmt.Errorf("source file(%s) not exist when move it", src)
	}

	dstDIR := filepath.Dir(dst)
	if exist, _ := Exists(dstDIR); !exist {
		if err := os.MkdirAll(dstDIR, os.ModeDir); err != nil {
			return err
		}
	}

	if fileInfo, err := os.Stat(dst); err == nil && fileInfo.IsDir() {
		dst = filepath.Join(dst, filepath.Base(src))
	}

	err := os.Rename(src, dst)
	return err
}

func MoveDir(srcDir, dstDir string) error {
	srcInfo, err := os.Stat(srcDir)
	if err != nil {
		return err
	}
	if !srcInfo.IsDir() {
		return errors.New("Not a directory")
	}

	if err := os.MkdirAll(dstDir, 0666); err != nil {
		return err
	}

	//filepath.Base("/home") => home
	//filepath.Base("/home/") => home
	dirName := filepath.Base(srcDir)
	for srcFile := range WalkDirDepthFirstIterator(srcDir, nil) {
		dstFile := filepath.Join(dstDir, dirName, strings.TrimPrefix(srcFile, srcDir))
		if err := MoveFile(srcFile, dstFile); err != nil {
			return err
			//there is question: WalkDirDepthFirstIterator goroutine not stop.
		}
	}

	if err := os.RemoveAll(srcDir); err != nil {
		return err
	}
	return nil
}

func WalkDir(dir string, filter func(string) bool) []string {
	var files []string

	for _, entry := range Dirents(dir) {
		subpath := filepath.Join(dir, entry.Name())
		if entry.IsDir() {
			files = append(files, WalkDir(subpath, filter)...)
		} else {
			if filter == nil || filter(subpath) {
				files = append(files, subpath)
			}
		}
	}

	return files
}

func Dirents(dir string) []os.FileInfo {
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		return []os.FileInfo{}
	}

	return entries
}

func WalkDirDepthFirstIterator(dir string, filter func(string) bool) <-chan string {
	// Depth-First Traversal
	type walkDirType func(string, func(string) bool)

	c := make(chan string)
	var walkDir walkDirType
	var deep uint = 0

	walkDir = func(dir string, filter func(string) bool) {
		deep++
		defer func() { deep-- }()

		for _, entry := range Dirents(dir) {
			subpath := filepath.Join(dir, entry.Name())
			if entry.IsDir() {
				walkDir(subpath, filter)
			} else {
				if filter == nil || filter(subpath) {
					c <- subpath
				}
			}
		}
		if deep == 1 {
			close(c)
		}

	}

	go walkDir(dir, filter)

	return c
}

func WalkDirBreadthFirstIterator(ctx context.Context, dir string) <-chan string {
	//Breadth-first Traversal

	dirList := list.New()
	dirList.PushBack(dir)
	c := make(chan string)

	go func() {
		defer close(c)

		for {
			if dirList.Len() == 0 {
				return
			}

			e := dirList.Front()
			dir := e.Value.(string)
			dirList.Remove(e)

			for _, entry := range Dirents(dir) {
				subpath := filepath.Join(dir, entry.Name())
				if entry.IsDir() {
					dirList.PushBack(subpath)
					continue
				}

				select {
				case <-ctx.Done():
					return
				case c <- subpath:
				}
			}

		}
	}()

	return c

}

func FindFile(dirPath, fileName string) string {
	ctx, cancel := context.WithCancel(context.Background())

	for fPath := range WalkDirBreadthFirstIterator(ctx, dirPath) {
		if strings.HasSuffix(fPath, filepath.Join("/"+fileName)) {
			cancel()
			return fPath
		}
	}

	return FileNotFound
}

func Basename(s string) string {
	slash := strings.LastIndex(s, "/")
	s = s[slash+1:]
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}

	return s
}

// GetCurDir Get current program directory
func GetCurDir() (string, error) {
	return filepath.Abs(filepath.Dir(os.Args[0]))
}
