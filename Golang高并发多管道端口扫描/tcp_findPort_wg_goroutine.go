package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
)

func worker(ports, results chan int, target string) {
	for p := range ports {
		address := fmt.Sprintf("%s:%d", target, p)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			results <- 0
			continue
		}
		conn.Close()
		results <- p
	}
}

func scanPorts(target string, ports []int) []int {
	var wg sync.WaitGroup
	openports := make([]int, 0)
	portsChan := make(chan int, len(ports))
	resultsChan := make(chan int)

	for i := 0; i < cap(portsChan); i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			worker(portsChan, resultsChan, target)
		}()
	}

	go func() {
		for _, port := range ports {
			portsChan <- port
		}
		close(portsChan)
	}()

	go func() {
		wg.Wait()
		close(resultsChan)
	}()

	for port := range resultsChan {
		if port != 0 {
			openports = append(openports, port)
		}
	}

	sort.Ints(openports)
	return openports
}

func main() {
	var (
		target     string
		outputFile string
		portRange  string
		ipFile     string
		showHelp   bool
		openports  []int
	)

	flag.StringVar(&target, "u", "", "目标IP或URL")
	flag.StringVar(&outputFile, "o", "", "保存扫描结果的文件")
	flag.StringVar(&portRange, "p", "1-65535", "扫描的端口范围 (例如，'80,443' 或 '1-1024')")
	flag.StringVar(&ipFile, "f", "", "包含要扫描的IP地址的文件")
	flag.BoolVar(&showHelp, "h", false, "显示帮助信息")

	flag.Parse()

	if showHelp {
		fmt.Println("用法: portscanner -u <目标> -p <端口范围> -o <输出文件> -f <IP文件> -h")
		flag.PrintDefaults()
		return
	}

	if target == "" {
		fmt.Println("错误: 请使用 -u 标志指定目标.")
		return
	}

	if portRange != "" {
		ports, err := parsePortRange(portRange)
		if err != nil {
			fmt.Println("错误:", err)
			return
		}
		openports = scanPorts(target, ports)
	} else {
		fmt.Println("错误: 请使用 -p 标志指定端口范围.")
		return
	}

	if outputFile != "" {
		err := writeOpenPortsToFile(outputFile, openports)
		if err != nil {
			fmt.Println("错误:", err)
		} else {
			fmt.Printf("扫描结果已保存到文件 %s\n", outputFile)
		}
	} else {
		for _, port := range openports {
			fmt.Printf("%d 开放\n", port)
		}
	}
}

func parsePortRange(portRange string) ([]int, error) {
	// 解析端口范围字符串 (例如，'80,443' 或 '1-1024') 为整数切片.
	// 如果范围无效，则返回错误.
	var ports []int

	// 使用逗号分隔输入字符串以获取单个端口或范围条目.
	entries := strings.Split(portRange, ",")

	for _, entry := range entries {
		if strings.Contains(entry, "-") {
			// 处理范围 (例如，'1-1024').
			rangeParts := strings.Split(entry, "-")
			if len(rangeParts) != 2 {
				return nil, fmt.Errorf("无效的端口范围: %s", entry)
			}

			start, err := strconv.Atoi(rangeParts[0])
			if err != nil {
				return nil, fmt.Errorf("无效的端口范围: %s", entry)
			}

			end, err := strconv.Atoi(rangeParts[1])
			if err != nil {
				return nil, fmt.Errorf("无效的端口范围: %s", entry)
			}

			for port := start; port <= end; port++ {
				ports = append(ports, port)
			}
		} else {
			// 处理单个端口条目 (例如，'80').
			port, err := strconv.Atoi(entry)
			if err != nil {
				return nil, fmt.Errorf("无效的端口: %s", entry)
			}
			ports = append(ports, port)
		}
	}

	return ports, nil
}

func writeOpenPortsToFile(filename string, openports []int) error {
	// 将开放的端口写入文件.
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	for _, port := range openports {
		_, err := fmt.Fprintf(file, "%d 开放\n", port)
		if err != nil {
			return err
		}
	}

	return nil
}
