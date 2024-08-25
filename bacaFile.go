package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func clearTerminal() {
	fmt.Print("\033[H\033[2J")
}

func main() {
	var namaFile, sesuai, input string
	var benar, benar2 bool

	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	for !benar {
		clearTerminal()

		namaFolder := filepath.Base(dir)
		fmt.Println("\nDirektori saat ini: ", namaFolder)
		fmt.Println("\nApakah folder sudah sesuai? (Ketik 'next' jika tidak sesuai atau 'null' jika sesuai)")
		fmt.Print("Masukkan input: ")
		fmt.Scan(&sesuai)

		if sesuai == "next" {
			benar = true
			for !benar2 {
				clearTerminal()

				namaFolder = filepath.Base(dir)
				fmt.Println("Direktori saat ini: ", namaFolder)
				fmt.Println("\nBerikut adalah kumpulan subfolder dalam direktori ini:")
				files, err := os.ReadDir(dir)
				if err != nil {
					fmt.Println("Error: ", err)
					return
				}

				for _, file := range files {
					if file.IsDir() {
						fmt.Println(" -", file.Name())
					}
				}

				fmt.Println("\nPerintah:")
				fmt.Println("1. Ketik 'up' untuk berpindah ke folder sebelumnya")
				fmt.Println("2. Ketik 'nama subFolder' untuk masuk ke sub folder tersebut")
				fmt.Println("3. Ketik 'ok' ketika direktori sudah sesuai")
				fmt.Print("\nMasukkan perintah atau nama subFolder: ")
				fmt.Scan(&input)
				input = strings.TrimSpace(input)

				if input == "up" {
					parentDir := filepath.Dir(dir)
					if parentDir == dir {
						fmt.Println("Anda sudah berada di direktori root")
					} else {
						dir = parentDir
					}
				} else if input == "ok" {
					break
				} else {
					subFolderPath := filepath.Join(dir, input)
					info, err := os.Stat(subFolderPath)
					if err != nil {
						fmt.Println("Error: ", err)
						continue
					}
					if !info.IsDir() {
						fmt.Println("Ini bukan folder")
						continue
					}
					dir = subFolderPath
					clearTerminal()
				}
			}
		} else if sesuai == "null" {
			benar = true
		} else {
			fmt.Println("\nInput tidak sesuai")
		}
	}

	fmt.Print("\nMasukkan nama file: ")
	fmt.Scan(&namaFile)

	file, err := os.Open(filepath.Join(dir, namaFile))
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	data := make([]byte, fileInfo.Size())

	_, err = file.Read(data)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	clearTerminal()
	fmt.Println(string(data))
}
