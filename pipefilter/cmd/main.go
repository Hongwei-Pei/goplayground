package main

import (
	"fmt"
	"time"

	"github.com/pointernil/goplayground/pipefilter"
)

func main() {
	var cin chan interface{} = make(chan interface{})
	var cout chan interface{} = make(chan interface{})
	var quit chan struct{} = make(chan struct{})

	pl := pipefilter.NewPipeLine(&cin, &cout)
	u := pipefilter.UselessFilter{}
	pf := pipefilter.NewPipeFilter()
	pf.RegisterFilter(&u)
	pf.RegisterFilter(&u)
	pl.RegisterFilter(&u)
	pl.RegisterFilter(pf)
	go func() {
		i := 0
		for {
			data, ok := <-cout
			if ok {
				fmt.Println(i, data)
				i++
			}

		}
	}()
	go func() {
		for {
			select {
			case <-quit:
				return
			case <-time.After(1 * time.Second):
				cin <- "test"
				fmt.Println("send data")
			}
		}
	}()

	go pl.Run()
	go quitFunc(&quit)
	<-quit
}

func quitFunc(quit *chan struct{}) {
	time.Sleep(10 * time.Second)
	close(*quit)
}
