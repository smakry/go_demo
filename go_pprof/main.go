package main

import (
	"fmt"
	"go_demo/go_pprof/data"
	"go_demo/go_pprof/pprof"
	"net/http"
	_ "net/http/pprof"
	"time"
)

func main() {
	go func() {
		begin := time.Now().Unix()
		for {
			if time.Now().Unix()-begin > 500 {
				fmt.Println("quit")
				break
			}
			fmt.Println(data.Add("https://github.com/EDDYCJY"))
		}
	}()

	pprof.WaitCommand()
	http.ListenAndServe("0.0.0.0:6060", nil)
	time.Sleep(time.Second * 1000)
}
