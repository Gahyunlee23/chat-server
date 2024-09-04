package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	// 서버에 연결
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return
	}
	defer conn.Close()

	fmt.Println("Connected to server. Type a message and press enter to send.")

	// 서버에서 메시지 수신하는 고루틴
	go func() {
		for {
			message, _ := bufio.NewReader(conn).ReadString('\n')
			fmt.Print("Server: " + message)
		}
	}()

	// 사용자 입력을 서버로 전송
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("You: ")
		text, _ := reader.ReadString('\n')

		// 메시지 보내기
		fmt.Fprintf(conn, text+"\n")

		// "/quit" 입력 시 연결 종료
		if strings.TrimSpace(text) == "/quit" {
			fmt.Println("Exiting...")
			return
		}
	}
}
