package main

import (
	"context"
	"fmt"
	pb "goProject/study/capMicro/gRPC/proto"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:9005", grpc.WithInsecure())
	if err != nil {
		fmt.Println("链接异常 err:", err)
		return
	}
	defer conn.Close()
	client := pb.NewCaptureLicServiceClient(conn)
	req := new(pb.CapRequest)

	resp, err := client.CapLicenseInfo(context.Background(), req)
	if err != nil {
		fmt.Println("响应异常 err:", err)
		return
	}
	fmt.Println("响应结果：", resp)
}
