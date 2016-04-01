package main

import (
	"log"
	"strings"

	"github.com/garyburd/redigo/redis"
)

func parseInfo(in string) map[string]string {
	info := map[string]string{}
	lines := strings.Split(in, "\r\n")

	for _, line := range lines {
		values := strings.Split(line, ":")

		if len(values) > 1 {
			info[values[0]] = values[1]
		}
	}
	return info
}

func main() {
	conn, err := redis.Dial("tcp", ":6379")

	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	data, err := redis.String(conn.Do("INFO"))

	if err != nil {
		log.Fatal(err)
	}

	info := parseInfo(data)

	log.Printf("%+v", info["redis_version"])
}
