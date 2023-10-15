package main

import (
	"fmt"

	"github.com/jlaffaye/ftp"
)

func main() {
	conn, err := ftp.Connect("22i554n646.imwork.net:21")
	if err != nil {
		fmt.Println("connect error")
	} else {
		fmt.Println("connect ok")
	}

	err = conn.Login("ftpuser_lfx", "fengxuan8743")
	if err != nil {
		fmt.Println("login ")
	}
	nameList, err := conn.List("/")
	if err != nil {
		fmt.Println("List error")
	}

	for _, name := range nameList {
		fmt.Println(name.Name, name.Type)
	}

}
