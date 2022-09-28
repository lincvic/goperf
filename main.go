package main

import (
	"flag"
	"fmt"
	. "goperf/goperf_struct"
	"runtime"
	"time"
)

func configLoader() *GoperfConfig {
	config := new(GoperfConfig)
	flag.IntVar(&config.GoroutineSize, "goroutines", 10, "The size (number) of goroutines, default is 10")
	flag.IntVar(&config.GoroutineWorkSize, "workers", 10, "The size (number) of workers for a single goroutine, default is 10")
	flag.IntVar(&config.ChannelBufferSize, "buffer", 1, "The size of Channel buffer, default is 1")
	flag.IntVar(&config.MessageSize, "msg", 100, "The size of message send to Channel, default is 100")
	flag.IntVar(&config.FactNumber, "fact", 1234, "The number to calculate fact in each worker, default is 1234")
	flag.Parse()
	return config
}

func sendMessage(c chan string, messageSize int) {
	msg := "abc"
	for i := 0; i < messageSize; i++ {
		c <- msg
	}
}

func fact(n int) int {
	var result int
	if n == 0 || n == 1 {
		return 1
	}
	result = fact(n-1) * n
	return result
}

func main() {
	configPtr := configLoader()
	fmt.Println("goperf MultiCore CPU Test starting ...")
	cpuCores := runtime.NumCPU()
	fmt.Printf("CPU Cores %d, Goroutine Size %d, %d workers inside each goroutine, Channle Buffer Size %d\n",
		cpuCores, configPtr.GoroutineSize, configPtr.GoroutineWorkSize, configPtr.ChannelBufferSize)
	runtime.GOMAXPROCS(cpuCores)

	channel := make(chan string, configPtr.ChannelBufferSize)
	startTime := time.Now()
	for i := 0; i < configPtr.GoroutineSize; i++ {
		go func(c chan string) {
			<-c
			for {
				for i := 0; i < configPtr.GoroutineWorkSize; i++ {
					fact(configPtr.FactNumber)
				}
			}
		}(channel)
	}

	sendMessage(channel, configPtr.MessageSize)

	timeInMS := time.Since(startTime).Milliseconds()
	fmt.Printf("Total time is %d ms\n", timeInMS)
	fmt.Printf("%d CPU cores process %f message/ms with %d goroutines", cpuCores, float64(configPtr.MessageSize)/float64(timeInMS), configPtr.GoroutineSize)
}
