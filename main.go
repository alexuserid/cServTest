package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var (
	address = flag.String("l", "http://localhost:8080", "Server address for test")
)

func main() {
	flag.Parse()
	url := *address

	for i, v := range mass {
		resp, err := http.Get(url + "?s=" + v.s)
		if err != nil {
			log.Fatalf("http.Get: %v", err)
		}
		bytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalf("ioutil.ReadAll: %v", err)
		}
		resp.Body.Close()

		if string(bytes) != v.n+"\n" {
			fmt.Printf(`Wrong answer. Test %d: %s.
			Server answer is: %sRight answer is: %s
			`, i, v.s, string(bytes), v.n)
			return
		}
	}
	fmt.Println("Success!")
}
