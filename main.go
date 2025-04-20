package main

import (
	"fmt"
	"gochecklink/helper"
	"net/http"
	"os"
	"sync"
)

const MAX_URL uint8 = 5

func CheckEndpoint(endpoint string, smap *sync.Map, g *sync.WaitGroup) {
	defer g.Done()
	if _, ok := smap.Load(endpoint); ok {
		return
	}

	smap.Store(endpoint, true)
	resp, err := http.Get(endpoint)
	if err != nil {
		fmt.Printf("[ERR] %s (lookup: %s)\n", endpoint, err)
	} else {
		fmt.Printf("[OK] %s (%d %s)\n", endpoint, resp.StatusCode, resp.Status)
	}
}

func main() {
	args := os.Args[1:]

	if len(args) > int(MAX_URL) {
		panic(fmt.Sprintf("Max URL is %d", MAX_URL))
	}

	var wg sync.WaitGroup
	smap := sync.Map{}
	for _, arg := range args {
		if !helper.IsValidURL(arg) {
			panic(fmt.Sprintf("Not valid URL: %s", arg))
		}
		wg.Add(1)
		go CheckEndpoint(arg, &smap, &wg)
	}

	wg.Wait()
}
