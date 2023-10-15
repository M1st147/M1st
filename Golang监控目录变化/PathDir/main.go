package main

import (
	"container/list"
	"encoding/csv"
	_ "encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/fsnotify/fsnotify"
)

type Fileinfo struct {
	Filepath string `json:"Filepath"`
	FileName string `json:"FileName"`
	ModTime  string `json:"ModTime"`
	FileSize string `json:"FileSize"`
}

var outputFileName string = "filesName.txt"

func CheckErr(err error) {
	if nil != err {
		panic(err)
	}
}

func GetFullPath(path string) string {
	absolutePath, _ := filepath.Abs(path)
	return absolutePath
}

func ScanPath(path string) error {
	fullPath := GetFullPath(path)

	//listStr := list.New()
	f, err := os.Create(outputFileName)
	CheckErr(err)
	defer f.Close()
	f.WriteString("\xEF\xBB\xBF")

	filepath.Walk(fullPath, func(path string, fi os.FileInfo, err error) error {

		if nil == fi {
			return err
		}
		if fi.IsDir() { //忽略目录
			fmt.Println(fullPath + fi.Name())
			return nil
		}

		name := fi.Name()
		if outputFileName != name {

			go func() {
				f.WriteString(path + "|" + name + "|" + string(fi.Size()) + "|" + fi.ModTime().Format("2006-01-02 15:04:05") + "\n")
			}()

			//listStr.PushBack(path + "|" + name + "|" + string(fi.Size()) + "|" + fi.ModTime().Format("2006-01-02 15:04:05"))
		}

		return nil
	})
	return nil
	//OutputFilesName(listStr)
}

func ConvertToSlice(listStr *list.List) []string {
	sli := []string{}
	for el := listStr.Front(); nil != el; el = el.Next() {
		sli = append(sli, el.Value.(string))

	}
	return sli
}

func OutputFilesName(listStr *list.List) {
	files := ConvertToSlice(listStr)
	//sort.StringSlice(files).Sort()// sort

	f, err := os.Create(outputFileName)
	//f, err := os.OpenFile(outputFileName, os.O_APPEND | os.O_CREATE, os.ModeAppend)
	CheckErr(err)
	defer f.Close()

	f.WriteString("\xEF\xBB\xBF")
	writer := csv.NewWriter(f)
	length := len(files)
	for i := 0; i < length; i++ {
		writer.Write([]string{files[i]})
	}

	writer.Flush()
}

func main() {
	start := time.Now()
	fmt.Println(ScanPath("E:\\微"))
	fmt.Println("处理时间约 %f f% 秒", time.Now().Sub(start).Seconds())
	fmt.Println("done!")
}
