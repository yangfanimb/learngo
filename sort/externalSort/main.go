package main

import (
	"bufio"
	"fmt"
	"learngo/pipeline"
	"os"
	"strconv"
)

func main() {
	p := createNetworkPipeline("small.in", 512, 4)
	// p := createPipeline("large.in", 8000000, 2)
	writeToFile(p, "small.out")
	printFile("small.out")
}

func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	p := pipeline.ReaderSource(file, -1)

	count := 0
	for v := range p {
		fmt.Printf("%016x\n", v)
		count++
		if count > 10 {
			break
		}
	}

}

func writeToFile(p <-chan int, filename string) {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	pipeline.WriteSink(writer, p)
}

func createPipeline(filename string, filesize, chunkCount int) <-chan int {
	sortResult := []<-chan int{}
	chunkSize := filesize / chunkCount

	pipeline.Init()

	for i := 0; i < chunkCount; i++ {
		file, err := os.Open(filename)
		if err != nil {
			panic(err)
		}

		file.Seek(int64(i*chunkSize), 0)
		source := pipeline.ReaderSource(file, chunkSize)
		sortResult = append(sortResult, pipeline.InMemSort(source))
	}

	return pipeline.MergeN(sortResult...)
}

func createNetworkPipeline(filename string, filesize, chunkCount int) <-chan int {
	chunkSize := filesize / chunkCount
	sortAddr := []string{}

	pipeline.Init()
	for i := 0; i < chunkCount; i++ {
		file, err := os.Open(filename)
		if err != nil {
			panic(err)
		}

		file.Seek(int64(i*chunkSize), 0)
		source := pipeline.ReaderSource(file, chunkSize)

		addr := ":" + strconv.Itoa(7000+i)
		pipeline.NetworkShink(addr, pipeline.InMemSort(source))
		// sortResult = append(sortResult, pipeline.InMemSort(source))
		sortAddr = append(sortAddr, addr)
	}

	sortResult := []<-chan int{}
	for _, addr := range sortAddr {
		// fmt.Println("collect from ", addr)
		sortResult = append(sortResult, pipeline.NetworkSource(addr))
	}
	return pipeline.MergeN(sortResult...)
}
