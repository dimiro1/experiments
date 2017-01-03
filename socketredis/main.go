package main

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"net"
	"strconv"
	"time"
)

var (
	InvalidConnectionErr = errors.New("Invalid connection")
	WriteErr             = errors.New("Could not write INFO command")
	ReadErr              = errors.New("Could not read Redis response")
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:6379")

	if err != nil {
		panic(err)
	}

	defer conn.Close()

	if output, err := info(conn); err == nil {
		fmt.Println(output)
	} else {
		fmt.Println(err)
	}
}

// This is only ready to exec the INFO command on Redis
func info(conn net.Conn) (string, error) {
	if conn == nil {
		return "", InvalidConnectionErr
	}

	conn.SetWriteDeadline(time.Now().Add(1 * time.Second))
	_, err := conn.Write([]byte("INFO\n"))

	if err != nil {
		return "", WriteErr
	}

	reader := bufio.NewReader(conn)

	line, err := reader.ReadString('\n')

	if err != nil {
		return "", ReadErr
	}

	// Empty line?
	if len(line) == 0 {
		return "", ReadErr
	}

	// Info command
	if line[0] != '$' {
		return "", ReadErr
	}

	count, err := strconv.ParseInt(line[1:len(line)-2], 10, 64)

	if err != nil {
		return "", ReadErr
	}

	var buf bytes.Buffer
	io.CopyN(&buf, reader, count)

	return buf.String(), nil
}
