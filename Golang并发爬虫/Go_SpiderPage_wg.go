package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"sync" // 添加 sync 包
)

func HttpGet(url string) (result string, err error) {
	resp, err1 := http.Get(url)
	if err1 != nil {
		err = err1
		return
	}
	defer resp.Body.Close()

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
		result += string(buf[:n])
	}
	return
}

func SpiderPage(i int, page chan int) {
	url := "https://tieba.baidu.com/f?kw=%E7%BB%9D%E5%9C%B%E6%B1%82%E7%94%9F&ie=utf-8&pn=" + strconv.Itoa((i-1)*50)
	result, err := HttpGet(url)
	if err != nil {
		fmt.Println("HttpGet err:", err)
		return
	}

	// 保存网页内容到文件
	f, err := os.Create("第" + strconv.Itoa(i) + "页" + ".html")
	if err != nil {
		fmt.Println("Create err:", err)
		return
	}
	f.WriteString(result)
	f.Close()

	page <- i
}

func working2(start, end int) {
	fmt.Printf("正在爬取第%d页到第%d页...\n", start, end)

	page := make(chan int)
	var wg sync.WaitGroup // 创建 WaitGroup

	for i := start; i <= end; i++ {
		wg.Add(1) // 每启动一个协程，增加 WaitGroup 的计数
		go func(i int) {
			defer wg.Done() // 协程完成时减少计数
			SpiderPage(i, page)
		}(i)
	}

	go func() {
		wg.Wait()   // 等待所有协程完成
		close(page) // 关闭通道
	}()

	for p := range page {
		fmt.Printf("第%d个页面爬取完成\n", p)
	}
}

func main() {
	var start, end int
	fmt.Print("请输入爬取的起始页(>=1):")
	fmt.Scan(&start)
	fmt.Print("请输入爬取的结束页(>=start):")
	fmt.Scan(&end)

	working2(start, end)
}
