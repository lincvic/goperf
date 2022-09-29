package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"github.com/spf13/cast"
	. "goperf/goperf_struct"
	"log"
	"os"
	"runtime"
	"time"
)

func configLoader() *GoperfConfig {
	config := new(GoperfConfig)
	flag.IntVar(&config.GoroutineWorkSize, "workers", 10, "The size (number) of workers for a single goroutine, default is 10")
	flag.IntVar(&config.ChannelBufferSize, "buffer", 1, "The size of Channel buffer, default is 1")
	flag.IntVar(&config.MessageSize, "msg", 100, "The size of message send to Channel, default is 100")
	flag.IntVar(&config.FactNumber, "fact", 1234, "The number to calculate fact in each worker, default is 1234")
	flag.IntVar(&config.RunningTime, "time", 1, "The number of how many times goperf will run the test, default is 1")
	flag.StringVar(&config.ResultPath, "o", "result.csv", "The path of result file, default is 'result.csv'")
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

func saveData(fileName *string, result *[][]string) {
	file, err := os.Create(*fileName)
	defer file.Close()
	if err != nil {
		fmt.Println("Creating File err !")
		log.Fatal(err)
	}

	csvWriter := csv.NewWriter(file)
	defer csvWriter.Flush()

	csvWriter.WriteAll(*result)
}

func runTest(configPtr *GoperfConfig, cpuCores *int, result *[][]string) {
	channel := make(chan string, configPtr.ChannelBufferSize)
	timeLocker := make(chan int, configPtr.MessageSize)
	startTime := time.Now()

	for i := 0; i < configPtr.MessageSize; i++ {
		go func(c chan string) {
			<-c
			for i := 0; i < configPtr.GoroutineWorkSize; i++ {
				fact(configPtr.FactNumber)
			}
			timeLocker <- 1
		}(channel)
	}

	sendMessage(channel, configPtr.MessageSize)

	for i := 0; i < configPtr.MessageSize; i++ {
		<-timeLocker
	}

	timeInSec := time.Since(startTime).Seconds()
	processSpeed := float64(configPtr.MessageSize) / timeInSec
	fmt.Printf("Total time is %f s\n", timeInSec)
	fmt.Printf("%d CPU cores process %f msg/s with %d goroutines \n", *cpuCores, processSpeed, configPtr.MessageSize)
	*result = append(*result, []string{cast.ToString(timeInSec), cast.ToString(processSpeed)})
}

func main() {
	configPtr := configLoader()
	fmt.Println("goperf MultiCore CPU Test starting ...")
	cpuCores := runtime.NumCPU()
	fmt.Printf("CPU Cores %d, Message Size %d, %d workers inside each goroutine, Channel Buffer Size %d, Test will run %d times\n",
		cpuCores, configPtr.MessageSize, configPtr.GoroutineWorkSize, configPtr.ChannelBufferSize, configPtr.RunningTime)
	runtime.GOMAXPROCS(cpuCores)

	result := [][]string{
		{"Total Time", "msg/s"},
	}
	for i := 0; i < configPtr.RunningTime; i++ {
		fmt.Printf("Testing Round %d/%d :\n", i+1, configPtr.RunningTime)
		runTest(configPtr, &cpuCores, &result)
		fmt.Println("======================================================")
	}

	fmt.Printf("Testing complete, Saving result into %s\n", configPtr.ResultPath)
	saveData(&configPtr.ResultPath, &result)

}
