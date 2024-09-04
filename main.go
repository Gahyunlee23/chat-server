package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

// 클라이언트 구조체
type Client struct {
	conn net.Conn
	name string
}

var clients = make(map[net.Conn]string)

func main() {
	// TCP 서버 시작
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
	defer listener.Close()

	fmt.Println("Server started on port 8080")

	for {
		// 새로운 클라이언트 연결 대기
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		// 클라이언트를 처리하는 고루틴 실행
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	// 클라이언트 이름 설정
	conn.Write([]byte("Enter your name: "))
	name, _ := bufio.NewReader(conn).ReadString('\n')
	name = strings.TrimSpace(name)
	clients[conn] = name

	fmt.Printf("%s joined the chat\n", name)

	// 채팅 메시지 수신 및 방송
	for {
		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Printf("%s disconnected\n", name)
			delete(clients, conn)
			return
		}

		broadcastMessage(name, message)
	}
}

func broadcastMessage(senderName string, message string) {
	for conn, name := range clients {
		if name != senderName {
			conn.Write([]byte(fmt.Sprintf("%s: %s", senderName, message)))
		}
	}
}
