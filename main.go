package main

import (
	anstest "mysrc/cServTest/anstest"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var (
	link = flag.String("l", "http://localhost:8080", "Server link for test")
)

func main() {
	flag.Parse()
	url := *link

	for i, v := range anstest.Mass {
		resp, err := http.Get(url + "?s=" + v.S)
		if err != nil {
			log.Fatalf("http.Get: %v", err)
		}
		bytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalf("ioutil.ReadAll: %v", err)
		}
		resp.Body.Close()

		if string(bytes) != v.N+"\n" {
			fmt.Printf("Wrong answer. Test %d: %s.\nServer answer is: %sRight answer is: %s\n", i, v.S, string(bytes), v.N)
			return
		}
	}
	fmt.Println("Success!")
}
