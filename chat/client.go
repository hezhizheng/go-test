package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

var (
	writeStr, readStr = make([]byte, 1024), make([]byte, 1024)
)

func main() {

	var (
		host = "127.0.0.1"
		port = "9220"
		reader = bufio.NewReader(os.Stdin)
	)


	conn, err := net.Dial("tcp", host+":"+port)

	if err != nil {
		fmt.Println("Error connecting:", err)
		os.Exit(1)
	}

	defer conn.Close()

	fmt.Println("连接成功 Connecting to " + "127.0.0.1:9220")

	fmt.Printf("输入用户名进行聊天: \n")

	fmt.Scanf("%s", &writeStr)

	in, errWrite := conn.Write(writeStr)
	if errWrite != nil {
		fmt.Printf("Error when send to server: %d\n", in)
		os.Exit(0)
	}

	// 一定得用协程！！！ 获取服务端的消息
	go read(conn)

	// todo main 下不能有两个或以上 for{} !!!
	// 将客户端的输入通知服务端
	for {
		//fmt.Printf("：")
		writeStr, _, _ = reader.ReadLine()
		if string(writeStr) == "quit" {
			fmt.Println("Communication terminated.")
			os.Exit(1)
		}

		in, err := conn.Write(writeStr)
		if err != nil {
			fmt.Printf("Error when send to server: %d\n", in)
			os.Exit(0)
		}
	}


}

func read(conn net.Conn)  {
	for  {
		_, err := conn.Read(readStr)
		if err != nil {
			fmt.Printf("Client quit.\n")
			conn.Close()
			return
		}

		msg := string(readStr)
		msg = strings.Replace(msg, "\r\n", "", -1)
		fmt.Println(msg)
	}
}
