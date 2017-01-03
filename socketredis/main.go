package main

import (
	"net"
	"bufio"
	"fmt"
	"io"
	"bytes"
	"strconv"
	"regexp"
	"errors"
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
		return "", errors.New("Invalid connection")
	}

	_, err := conn.Write([]byte("INFO\n"))

	if err != nil {
		return "", errors.New("Could not exec command")
	}

	reader := bufio.NewReader(conn)

	message, err := reader.ReadString('\n')

	if err != nil {
		return "", errors.New("Could not read command output")
	}

	re := regexp.MustCompile("\\d+")
	matches := re.FindAllString(message, -1)

	if len(matches) != 1 {
		return "", errors.New("Malformed output")
	}

	count, err := strconv.ParseInt(matches[0], 10, 64)

	if err != nil {
		return "", errors.New("Malformed output")
	}

	var buf bytes.Buffer
	io.CopyN(&buf, reader, count)

	return buf.String(), nil
}
