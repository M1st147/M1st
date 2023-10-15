package picture

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"gocv.io/x/gocv"
)

//SavePicture 保存图片
func SavePicture(deviceID int) string {
	webcam, err := gocv.OpenVideoCapture(deviceID)
	if err != nil {
		fmt.Printf("Error opening video capture device: %v\n", deviceID)
		return ""
	}
	defer webcam.Close()
	ch := make(chan gocv.Mat, 1)
	img := gocv.NewMat()
	defer img.Close()
	if ok := webcam.Read(&img); !ok {
		fmt.Printf("cannot read device %v\n", deviceID)
		return ""
	}
	if img.Empty() {
		fmt.Printf("no image on device %v\n", deviceID)
		return ""
	}
	ch <- img
	now := time.Now()
	appPath, err := getCurrentPath()
	if err != nil {
		fmt.Printf("get app path failed,err: %v\n", err)
		return ""
	}
	imgFile := fmt.Sprintf("%simages\\%s%s", appPath, now.Format("2006-01-02-15-04-05"), ".jpg")
	//gocv.IMWrite(imgFile, img)
	img2 := <-ch
	gocv.IMWrite(imgFile, img2)
	return imgFile
}

//getCurrentPath 获取exe根目录
func getCurrentPath() (string, error) {
	file, err := exec.LookPath(os.Args[0])
	if err != nil {
		return "", err
	}
	path, err := filepath.Abs(file)
	if err != nil {
		return "", err
	}
	i := strings.LastIndex(path, "/")
	if i < 0 {
		i = strings.LastIndex(path, "\\")
	}
	if i < 0 {
		return "", errors.New(`error: Can't find "/" or "\".`)
	}
	return string(path[0 : i+1]), nil
}
