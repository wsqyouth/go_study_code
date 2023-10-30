package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"strings"
	"sync"
)

const (
	maxClients = 1000
	maxNickLen = 32
)

var (
	serverPort = flag.Int("p", 8972, "server port")
)

type Client struct {
	conn net.Conn
	nick string
}

type ChatState struct {
	listener net.Listener

	clientsLock sync.RWMutex
	clients     map[net.Conn]*Client
	numClients  int
}

var chatState = &ChatState{
	clients: make(map[net.Conn]*Client),
}

func initChat() {
	var err error
	chatState.listener, err = net.Listen("tcp", fmt.Sprintf(":%d", *serverPort))
	if err != nil {
		fmt.Println("listen error:", err)
		os.Exit(1)
	}
}

func handleClient(client *Client) {
	// 发送欢迎信息
	welcomeMsg := "Welcome Simple Chat! Use /nick to change nick name.\n"
	client.conn.Write([]byte(welcomeMsg))

	buf := make([]byte, 256)
	for {
		n, err := client.conn.Read(buf)
		if err != nil {
			fmt.Printf("client left: %s\n", client.conn.RemoteAddr())
			chatState.clientsLock.Lock()
			delete(chatState.clients, client.conn)
			chatState.numClients--
			chatState.clientsLock.Unlock()
			return
		}

		msg := string(buf[:n])
		msg = strings.TrimSpace(msg)
		if len(msg) > 0 && msg[0] == '/' {
			// 处理命令
			parts := strings.SplitN(msg, " ", 2)
			cmd := parts[0]
			if cmd == "/nick" && len(parts) > 1 {
				client.nick = parts[1]
			}
			continue
		}

		fmt.Printf("\n%s: %s\n", client.nick, msg)

		// 将消息转发给其他客户端
		chatState.clientsLock.RLock()
		for conn, cl := range chatState.clients {
			if cl != client {
				conn.Write([]byte(client.nick + ": " + msg))
			}
		}
		chatState.clientsLock.RUnlock()
	}
}

func main() {
	flag.Parse()

	initChat()

	for {
		conn, err := chatState.listener.Accept()
		if err != nil {
			fmt.Println("accept error:", err)
			continue
		}

		client := &Client{conn: conn}
		client.nick = fmt.Sprintf("user%d", conn.RemoteAddr().(*net.TCPAddr).Port)

		chatState.clientsLock.Lock()
		chatState.clients[conn] = client
		chatState.numClients++
		chatState.clientsLock.Unlock()

		go handleClient(client)

		fmt.Printf("new client: %s\n", conn.RemoteAddr())
	}
}

/*
总结：
测试鸟叔代码: https://mp.weixin.qq.com/s/o8DM0HtQAuG_lsvL8wJFhA
好的代码读起来非常优雅. 获取唯一全局变量listener,accept阻塞监听获取conn,将每个conn放到map中维护,使用协程实现读写,通过唯一全局变量找到map集合遍历群发

相对原作者cpp代码，cpp代码有大量底层描述符、字符串操作、网络套接字设置，代码没有那么简洁
*/
