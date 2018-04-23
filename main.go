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
	checkTime = flag.Duration("t", 3*time.Second, "Test duration")
)

func main() {
	flag.Parse()
	tstart := time.Now()
	var i int

	for ; time.Since(tstart) < *checkTime; i++ {
		for j, v := range mass {
			resp, err := http.Get(*address + "?s=" + v.s)
			if err != nil {
				log.Fatalf("http.Get: %v", err)
			}
			bytes, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Fatalf("ioutil.ReadAll: %v", err)
			}

			ans := strings.TrimSpace(string(bytes))
			if string(ans) != v.n {
				fmt.Printf("Wrong answer. Test %d: %q.\n Server answer is: %q.\n Right answer is: %s\n", j, v.s, string(bytes), v.n)
				return
			}
		}
	}
	fmt.Printf("Timeout. Test was repited %d times. Duration %v. %v rps.\n", i, checkTime, i/int(checkTime.Seconds()))
}
