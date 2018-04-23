package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

var (
	address   = flag.String("l", "http://localhost:8080", "Server address for test")
	checkTime = flag.Duration("t", 200*time.Millisecond, "Test durationi in millisecond")
)

func main() {
	flag.Parse()
	url := *address

	timeout := time.After(*checkTime)
	for i := 0; ; i++ {

		select {
		case <-timeout:
			fmt.Printf("Timeout. Test was repited %d times. Duration: %v.", i, checkTime)
			return
		default:
			for j, v := range mass {
				resp, err := http.Get(url + "?s=" + v.s)
				if err != nil {
					log.Fatalf("http.Get: %v", err)
				}
				bytes, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					log.Fatalf("ioutil.ReadAll: %v", err)
				}
				resp.Body.Close()

				ans := strings.TrimSpace(string(bytes))
				if string(ans) != v.n {
					fmt.Printf(`Wrong answer. Test %d: %s.
Server answer is: %q
Right answer is: %s
`, j, v.s, string(bytes), v.n)
					return
				}
			}
		}
	}
}
