package db

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"log"
)

func conn() redis.Conn {
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println(err)
	}
	return conn
}

func SET(name string, key string) {
	defer conn().Close()
	_, err := conn().Do("SET", name, key)
	if err != nil {
		log.Fatalf("set is err%v", err)
	}
}

func LSET(name string, id int, value []byte) {
	defer conn().Close()
	_, err := conn().Do("LSET", name, id, value)
	if err != nil {
		log.Fatalf("LSET is err%v", err)
	}
}

func RPUSH(name string, key []byte) interface{} {
	defer conn().Close()
	id, err := conn().Do("RPUSH", name, key)
	if err != nil {
		log.Fatalf("set is err%v", err)
	}
	return id
}

func GET(name string) []byte {
	defer conn().Close()
	rep, err := redis.Bytes(conn().Do("GET", name))
	if err != nil {
		log.Fatalf("get err%v", err)
	}
	return rep
}

func LRANGE(name string, numSTART int, numSTOP int) []string {
	defer conn().Close()
	rep, err := redis.Strings(conn().Do("LRANGE", name, numSTART, numSTOP))
	if err != nil {
		log.Fatalf("get err%v", err)
	}
	return rep
}

func DEL(name string) {
	_, err := conn().Do("DEL", name)
	if err != nil {
		log.Fatalf("del is err%v", err)
	}
}
