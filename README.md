# Go TCP Chat Server

This is a simple TCP chat server written in Go. It allows multiple clients to connect to the server, send messages, and receive responses from the server. The server runs on TCP port 8080, and clients can connect using TCP to send and receive messages.

## Features
- Concurrent client handling using Go's goroutines.
- TCP-based message broadcasting.
- A simple client that connects to the server and sends messages.

## Requirements
- Go 1.19 or later
- macOS, Linux, or Windows

## Project Structure
Server-side implementation 
chat-server/ 
├── main.go 
Simple client implementation
  client/
  └── client.go 

## How to Run

1. Clone the repository
git clone https://github.com/yourusername/chat-server.git
cd chat-server

2. Initialize the project
If go.mod is not present, initialize the Go module:
go mod init chat-server

3. Run the Server
To run the TCP chat server on port 8080:

go run main.go
The server will start and listen on port 8080 for incoming client connections.

4. Run the Client
Open a new terminal window and run the client to connect to the server:
go run client.go
You can open multiple terminal windows and run multiple clients to simulate a chat room.

5. Sending and Receiving Messages
Once the client is connected, you can type messages and send them to the server. The server will respond to each message with "Message received".

Example interaction:
Client sends: Hello server!
Server responds: Message received

6. Quit the Client
To exit the client, simply type /quit and press Enter.

Troubleshooting
Port 8080 Already in Use
If port 8080 is already in use, you can check which process is using the port and terminate it
