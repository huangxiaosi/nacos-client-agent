package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func listFilesInDirectory(dirPath string) ([]string, error) {
	// 读取目录下的所有文件和子目录
	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return nil, fmt.Errorf("读取目录失败：%v", err)
	}

	// 存储文件名的切片
	fileNames := []string{}

	// 遍历文件列表并添加文件名到切片
	for _, file := range files {
		// 忽略子目录
		if !file.IsDir() {
			fileNames = append(fileNames, file.Name())
		}
	}

	return fileNames, nil
}

func getCurrentAbPathByCaller() string {
	//var abPath string
	//_, filename, _, ok := runtime.Caller(0)
	//if ok {
	//	abPath = path.Dir(filename)
	//}
	abPath, _ := os.Getwd()
	return abPath
}
