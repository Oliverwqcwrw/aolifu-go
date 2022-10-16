package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {

	file, err := os.Open("./file/a.txt")
	defer file.Close()
	if err != nil {
		fmt.Println("open file  error")
	}
	fmt.Println(file)

	//err2 := file.Close()
	//if err2 != nil {
	//	fmt.Println("close file  error")
	//}
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		fmt.Print(str)
	}
	fmt.Println("read file finished")

	readAllFile()

	writeToFile()

	overrideFile()

	appendFile()

	copyFileContent()

	pathIsExist()

	sourceFile := "./file/a.txt"
	destFile := "./file/d.txt"
	written, copyFileErr := copyFile(sourceFile, destFile)
	if copyFileErr != nil {
		fmt.Println(copyFileErr.Error())
	} else {
		fmt.Println(written)
	}

}

func readAllFile() {
	fmt.Println("read all file content")
	content, err := ioutil.ReadFile("./file/a.txt")
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(string(content))
}

func writeToFile() {
	filePath := "./file/b.txt"
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	str := "Hello World \n"
	for i := 0; i < 5; i++ {
		writer.WriteString(str)
	}
	flushErr := writer.Flush()
	if flushErr != nil {
		fmt.Println(err.Error())
	}
}

func overrideFile() {
	filePath := "./file/a.txt"
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	str := "override file \t\n"
	for i := 0; i < 5; i++ {
		writer.WriteString(str)
	}
	flushErr := writer.Flush()
	if flushErr != nil {
		fmt.Println(flushErr.Error())
	}
}

func appendFile() {
	filePath := "./file/a.txt"
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	str := "append file \t\n"
	for i := 0; i < 5; i++ {
		writer.WriteString(str)
	}
	flushErr := writer.Flush()
	if flushErr != nil {
		fmt.Println(flushErr.Error())
	}
}

// override content
func copyFileContent() {
	filePath1 := "./file/a.txt"
	filePath2 := "./file/b.txt"

	data, err := ioutil.ReadFile(filePath1)
	if err != nil {
		fmt.Println(err.Error())
	}
	err2 := ioutil.WriteFile(filePath2, data, 0)
	if err2 != nil {
		fmt.Println(err.Error())
	}
}

func pathIsExist() (bool, error) {
	filePath1 := "./file/a.txt"
	fileInfo, err := os.Stat(filePath1)
	if err == nil {
		fmt.Println(fileInfo)
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, err
	}
	return false, err

}

func copyFile(srcFilePath string, descFilePath string) (written int64, err error) {

	sourceFile, err1 := os.Open(srcFilePath)
	if err1 != nil {
		fmt.Println(err.Error())
		return
	}
	defer sourceFile.Close()

	sourceReader := bufio.NewReader(sourceFile)

	descFile, err2 := os.OpenFile(descFilePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err2 != nil {
		fmt.Println(err2.Error())
	}
	defer descFile.Close()
	destWriter := bufio.NewWriter(descFile)

	written, err3 := io.Copy(destWriter, sourceReader)
	if err3 != nil {
		fmt.Println(err3.Error())
		return
	}
	return written, nil
}
