package main

import (
	"fmt"
	"os"
	"path/filepath"
)

// ROBOCOPY "K:Коммерческий отдел" "D:Ком отдел" /e /l /ns /njs /njh /ndl /fp

func main() {
	// rootDir := "K:\\Коммерческий отдел"
	rootDir := "D:\\Ком отдел\\FleetOne"
	fileTypes := make(map[string]int)
	allTypes := make(map[string]int)

	to_arc := []string{".exe", ".dll", ".bat", ".js", ".msi", ".msp", ".msu"}

	to_delete := []string{".zip@SynoEAStream", ".TXT@SynoEAStream", ".XLS@SynoEAStream", ".gif@SynoEAStream", ".d91@SynoEAStream", ".mp4@SynoEAStream", ".xml@SynoEAStream",
		".vcf@SynoEAStream", ".inf@SynoEAStream", ".xlt@SynoEAStream", ".js@SynoEAStream", ".odt@SynoEAStream", ".xls@SynoEAStream", ".tmp@SynoEAStream", ".jpeg@SynoEAStream",
		".EXE@SynoEAStream", ".pptx@SynoEAStream", ".bat@SynoEAStream", ".pdf@SynoEAStream", ".docx@SynoEAStream", ".HTM@SynoEAStream", ".wmz@SynoEAStream", ".eml@SynoEAStream",
		".ppt@SynoEAStream", ".dll@SynoEAStream", ".tif@SynoEAStream", ".DOC@SynoEAStream", ".exe@SynoEAStream", ".chk@SynoEAStream", ".DLL@SynoEAStream", ".dbd@SynoEAStream",
		".db@SynoEAStream", ".Pdf@SynoEAStream", ".JPG@SynoEAStream", ".TIF@SynoEAStream", ".DOCX@SynoEAStream", ".7z@SynoEAStream", ".css@SynoEAStream", ".rtf@SynoEAStream",
		".txt@SynoEAStream", ".dat@SynoEAStream", ".htm@SynoEAStream", ".msu@SynoEAStream", ".jpg@SynoEAStream", ".png@SynoEAStream", ".bmp@SynoEAStream", ".html@SynoEAStream",
		".D91@SynoEAStream", ".doc@SynoEAStream", ".csv@SynoEAStream", ".Zip@SynoEAStream", ".060@SynoEAStream", ".RAR@SynoEAStream", ".PDF@SynoEAStream", ".msp@SynoEAStream",
		".tiff@SynoEAStream", ".p7s@SynoEAStream", ".jp2@SynoEAStream", ".wmf@SynoEAStream", ".PNG@SynoEAStream", ".msi@SynoEAStream", ".sig@SynoEAStream", ".rar@SynoEAStream",
		".BMP@SynoEAStream", ".xlsx@SynoEAStream"}

	err := filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			ext := filepath.Ext(path)
			allTypes[ext] += 1

			for to_del := 0; to_del < len(to_delete); to_del++ {
				if ext == to_delete[to_del] {
					err := os.Remove(path)
					fmt.Printf("deleted : %s\n", path)
					if err != nil {
						fmt.Println(err)
						return nil
					}
				}
			}

			if info.Name() == "Thumbs" {
				err := os.Remove(path)
				if err != nil {
					fmt.Println(err)
				}
			}

			// here add to archive

			// for to_ar := 0; to_ar < len(to_arc); to_ar++ {
			// 	if ext == to_arc[to_ar] {
			// 		err := os.Remove(path)
			// 		fmt.Printf("added to arc : %s\n", path)
			// 		if err != nil {
			// 			fmt.Println(err)
			// 			return nil
			// 		}
			// 	}
			// }

			switch ext {
			case ".txt", ".csv", ".log":
				fileTypes["text"]++
			case ".jpg", ".jpeg", ".png", ".gif", ".bmp":
				fileTypes["image"]++
			case ".doc", ".docx", ".pdf", ".xls", ".xlsx":
				fileTypes["document"]++
			case ".exe", ".bin", ".sh":
				fileTypes["executable"]++
			default:
				fileTypes["binary"]++
			}
		} else {
			if info.Name() == "@eaDir" {
				err := os.RemoveAll(path)
				fmt.Printf("deleted dir : %s\n", path)
				if err != nil {
					fmt.Println(err)
					return nil
				}
			}
			fmt.Println(path)
		}
		return nil
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("def types")

	for k, v := range fileTypes {
		fmt.Printf("%s: %d\n", k, v)
	}

	fmt.Println("all types")

	for k, v := range allTypes {
		fmt.Printf("%s: %d\n", k, v)
	}
}
