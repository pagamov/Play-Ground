package main

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	dir1 := "K:\\Коммерческий отдел"
	dir2 := "D:\\Ком отдел"

	diffDirs(dir1, dir2)
}

func diffDirs(dir1, dir2 string) {
	files1, err := filepath.Glob(filepath.Join(dir1, "*"))
	if err != nil {
		fmt.Println(err)
		return
	}

	files2, err := filepath.Glob(filepath.Join(dir2, "*"))
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, file1 := range files1 {
		relPath, err := filepath.Rel(dir1, file1)
		if err != nil {
			fmt.Println(err)
			continue
		}

		file2 := filepath.Join(dir2, relPath)
		if _, err := os.Stat(file2); os.IsNotExist(err) {
			fmt.Printf("File %s exists in %s but not in %s\n", relPath, dir1, dir2)
			continue
		}

		data1, err := os.ReadFile(file1)
		if err != nil {
			fmt.Println(err)
			continue
		}

		data2, err := os.ReadFile(file2)
		if err != nil {
			fmt.Println(err)
			continue
		}

		if !bytes.Equal(data1, data2) {
			fmt.Printf("File %s differs between %s and %s\n", relPath, dir1, dir2)
		}
	}

	for _, file2 := range files2 {
		relPath, err := filepath.Rel(dir2, file2)
		if err != nil {
			fmt.Println(err)
			continue
		}

		file1 := filepath.Join(dir1, relPath)
		if _, err := os.Stat(file1); os.IsNotExist(err) {
			fmt.Printf("File %s exists in %s but not in %s\n", relPath, dir2, dir1)
		}
	}
}
