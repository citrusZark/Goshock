// main
package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"sync"
	"time"
)

func main() {
	delay := flag.Duration("delay", 0, "Delay Request.")
	url := flag.String("api", "", "The addr of the application.")
	threadCount := flag.Int("thread", 1, "Number of goroutines.")
	filename := flag.String("datafile", "", "filename.")
	flag.Parse() // parse the flags

	if *url == "" {
		fmt.Println("url cannot be empty.")
		fmt.Println("usage: goshock -api http://localhost:9090/users/register -thread 4 -datafile data.json -delay 100ms")
		return
	}

	if *filename == "" {
		fmt.Println("datafile cannot be empty.")
		fmt.Println("usage: goshock -api http://localhost:9090/users/register -thread 4 -datafile data.json -delay 100ms")
		return
	}

	fmt.Println("Goshock\nDummy REST client")
	fmt.Println("usage: goshock -api http://localhost:9090/users/register -thread 4 -datafile data.json -delay 100ms")
	fmt.Println("p muh thoha")

	dataByte, err := ioutil.ReadFile(*filename)
	if err != nil {
		fmt.Println("Error: Couldn't read json.")
		return
	}

	var wg sync.WaitGroup
	wg.Add(*threadCount)
	for i := 0; i < *threadCount; i++ {
		go func() {
			for {
				time.Sleep(*delay)
				req, err := http.NewRequest("POST", *url, bytes.NewBuffer(dataByte))
				req.Header.Set("Content-Type", "application/json")

				client := &http.Client{}
				resp, err := client.Do(req)
				jsonData, _ := ioutil.ReadAll(resp.Body)
				fmt.Println(os.Stdout, string(jsonData))

				if err != nil {
					panic(err)
				}
				defer resp.Body.Close()
			}
		}()
	}
	wg.Wait()
}
