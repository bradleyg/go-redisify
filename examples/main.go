package main

import (
	"log"

	"github.com/bradleyg/go-redisify"
)

func main() {
	// client, err := goredisify.Conn("redis://user:pass@host.com")
	client, err := goredisify.Conn(nil)
	if err != nil {
		log.Fatal(err.Error())
	}

	err = client.Set("goredisify-test", []byte("test"))
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Println("Set key 'goredisify-test'")

	_, err = client.Del("goredisify-test")
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Println("Del key 'goredisify-test'")
}
