package main

import (
    "fmt"
    "os"
    "github.com/aliyun/aliyun-oss-go-sdk/oss"
)

func main() {
    // 创建OSSClient实例。
    client, err := oss.New("oss-cn-hangzhou.aliyuncs.com", "OSS_ACCESS_KEY_ID", "OSS_ACCESS_KEY_SECRET")
    if err != nil {
        fmt.Println("Error:", err)
        os.Exit(-1)
    }

    // 获取存储空间。
    bucket, err := client.Bucket("Bucket名称")
    if err != nil {
        fmt.Println("Error:", err)
        os.Exit(-1)
    }

    // 文件压缩下载。
    err = bucket.GetObjectToFile("文件名", "LocalFile.gzip", oss.AcceptEncoding("gzip"))
    if err != nil {
        fmt.Println("Error:", err)
        os.Exit(-1)
    }
}