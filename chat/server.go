package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

const (
	host = "127.0.0.1"
	port = "9220"
)

// 存放所有连接的客户端
var (
	clients []net.Conn
	data    = make([]byte, 1024)
)

func main() {
	// 创建一个 tcp 服务端连接
	listener, err := net.Listen("tcp", host+":"+port)
	if err != nil {
		log.Println(err)
		os.Exit(0)
	}
	defer listener.Close()
	log.Println("tcp server start", data)

	// 处理客户端的连接
	for {

		// 客户端初次连接
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			os.Exit(1)
		}
		// 保存客户端的连接
		clients = append(clients, conn)

		// todo for 嵌套 for 需要 协程？？？
		go func(conn net.Conn) {
			// 当前连接的客户端地址
			ClientRemoteAddr := conn.RemoteAddr().String()

			log.Println("客户端连接来自:", ClientRemoteAddr)

			// 读取客户端的数据包
			_, errRead := conn.Read(data)
			if errRead != nil {
				fmt.Printf("Client %v quit.\n", conn.RemoteAddr())
				conn.Close()
				disconnect(conn, conn.RemoteAddr().String())
				return
			}
			//name := string(data[:length])

			name := string(data) + "(" + ClientRemoteAddr + ")"

			name = strings.Replace(name, "\r\n", "", -1)

			conn.Write([]byte("欢迎你，" + name))

			notify(conn, name+" 上线了")

			log.Println("客户端：" + name + "上线了")

			// 处理消息交互
			for {
				_, err := conn.Read(data)
				if err != nil {
					fmt.Printf("Client %s quit.\n", name)
					conn.Close()
					disconnect(conn, name)
					return
				}
				res := string(data)

				sprdMsg := name + " said：" + res
				fmt.Println(sprdMsg)
				res = "我：" + res
				conn.Write([]byte(res))
				notify(conn, sprdMsg)
			}
		}(conn)

	}
}

func notify(conn net.Conn, msg string) {
	for _, con := range clients {
		if con.RemoteAddr() != conn.RemoteAddr() {
			con.Write([]byte(msg))
		}
	}
}

func disconnect(conn net.Conn, name string) {
	for index, con := range clients {
		if con.RemoteAddr() == conn.RemoteAddr() {
			disMsg := name + " 离开了."
			fmt.Println(disMsg)
			clients = append(clients[:index], clients[index+1:]...)
			notify(conn, disMsg)
		}
	}
}
