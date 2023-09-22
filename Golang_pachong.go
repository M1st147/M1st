package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
)

func HttpGet(url string) (result string, err error) {
	resp, err1 := http.Get(url)
	if err1 != nil {
		err = err1 //将封装函数内部的错误传出给调用者，和err1一样
		return
	}
	defer resp.Body.Close()

	//循环读取网页数据，传出给调用者
	buf := make([]byte, 4096)
	for {
		n, err2 := resp.Body.Read(buf)
		if n == 0 {
			fmt.Println("网页读取完成")
			break
		}
		if err2 != nil && err2 != io.EOF {
			err = err2
			return
		}
		//累加每一次循环读到的buf数据，存入result一次性返回
		result += string(buf[:n])
	}
	return
}

// 爬取页面操作
func working(start, end int) {
	fmt.Printf("正在爬取第%d页到第%d页...\n", start, end)

	//循环爬取每一页数据，循环一次是一页数据
	for i := start; i <= end; i++ {
		url := "https://tieba.baidu.com/f?kw=%E7%BB%9D%E5%9C%B%E6%B1%82%E7%94%9F&ie=utf-8&pn=" + strconv.Itoa((i-1)*50)
		result, err := HttpGet(url)
		if err != nil {
			fmt.Println("HttpGet err:", err)
			continue
			//调用者处理情况，继续运行
		}
		//fmt.Println("result=", result)
		f, err := os.Create("第" + strconv.Itoa(i) + "页" + ".html")
		if err != nil {
			fmt.Println("Create err:", err)
			continue
		}
		f.WriteString(result)
		f.Close() //保存一个结果关闭一个结果
	}
}

func main() {
	//指定
	var start, end int
	fmt.Print("请输入爬取的起始页(>=1):")
	fmt.Scan(&start)
	fmt.Print("请输入爬取的结束页(>=start):")
	fmt.Scan(&end)

	working(start, end)
}
