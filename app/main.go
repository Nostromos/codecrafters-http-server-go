package main

import (
	"bytes"
	"fmt"
	"net"
	"os"
	"strings"

	"github.com/codecrafters-io/http-server-starter-go/logger"
	"github.com/codecrafters-io/http-server-starter-go/parser"
	"github.com/codecrafters-io/http-server-starter-go/response"
	"github.com/codecrafters-io/http-server-starter-go/router"
)

// Ensures gofmt doesn't remove the "net" and "os" imports above (feel free to remove this!)
var _ = net.Listen
var _ = os.Exit

func main() {
	l, err := net.Listen("tcp", "0.0.0.0:4221")
	if err != nil {
		fmt.Println("Failed to bind to port 4221")
		os.Exit(1)
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("ERROR: Error accepting connection: ", err.Error())
			os.Exit(1)
		}
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close(); // close the connection after this function returns
	/* On connection, createa  buffer and read into it */
	buf := make([]byte, 4096)
	n, err := conn.Read(buf)
	// write loop that populates and ensures all data gets written into buffer
	if err != nil {
		fmt.Println("ERROR: Something bad happened while readin the connection")
		os.Exit(1)
	}
}

func readAllBytes(conn) ([]byte, err) {
	buf := make([]byte, 4096) // arbitrary sizing. not sure what is best here
	n, err := conn.Read(buf)

	for {

	}
}

	

	// conn.Write() returns n, err where n is the number of bytes from your byteslice that weresuccessfully written to the connection
	okStatus := "HTTP/1.1 200 OK\r\n\r\n"
	notFoundStatus := "HTTP/1.1 404 Not Found\r\n\r\n"
	total := 0

	for total < len(okStatus) {
		if (len(target) > 1) && (target != "/") {
			_, err := conn.Write([]byte(notFoundStatus))
			if err != nil {
				fmt.Println("ERROR: Cannot write to connection")
				os.Exit(1)
			}
		} else {
			_, err := conn.Write([]byte(okStatus))
			if err != nil {
				fmt.Println("ERROR: Cannot write to connection")
				os.Exit(1)
			}
		}
		// total += n
		os.Exit(0)
	}
// listening accepting writing in main

// short circuiting / request validation


type RawHTTPRequest struct {
	raw []byte
}

type ParsedHTTPRequest struct {
	method string
	target string
	version string
	host string
	headers string
	body string
}
// takes buffer and returns parsed request
func parseRequest(buf []byte, n int) *ParsedHTTPRequest {
	headerEnd := bytes.Index(buf[:n], []byte("\r\n\r\n"))
	headerBytes := buf[:headerEnd]
	// bodyBytes := buf[headerEnd + 4 : n] // if body exists

	headerText := string(headerBytes)
	lines := strings.Split(headerText, "\r\n")

	request := strings.Split(lines[0], " ");

	return &ParsedHTTPRequest{
		method: request[0],
		target: request[1],
		version: request[2],
		host: strings.Split(lines[1], " ")[1],
		headers: strings.Join(lines[2:], "\r\n"),
		body: "",
	}
}

type ResponseOptions struct {
	StatusCode int
	Headers map[string]string
	Body []byte
}

type HTTPResponse struct {
	ResponseOptions
}

type HTTPHeader struct {

}

// Takes an options struct
func buildResponse(opts ResponseOptions) *HTTPResponse {
	status := opts.StatusCode
	if status == 0 {
		status = 200
	}

	resp := &HTTPResponse{
		StadtusCode: status,
		Header: make(HTTPHeader),
		Body: opts.Body,
	}
}

// routing

// write test specs - table testing

