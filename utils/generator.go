package utils

import (
	"fmt"
	"io/fs"
	"log"
	"os"
)
const templateDir = "templates/grpc"

func GetCwd() (string, error) {
	mydir, err := os.Getwd()
    if err != nil {
		 log.Fatal(err)
    }
    return mydir, err
}

func CheckDirectory(dirName, path string ) bool {
	entries, err := os.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	for _, e := range entries{
		if e.Name() == dirName{
			return true
		}
	}
	return false

}

func GetDirectoryList(path string) []fs.DirEntry {
	entries, err := os.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}
	return entries
}

func CheckIsDir(path string) bool {
	pwd, _ := GetCwd()
	route := pwd + "/" + templateDir + "/" + path
	fmt.Println(route)
	fileInfo, err := os.Stat(route)
	if err != nil {
		return false
	}

	return fileInfo.IsDir()
}