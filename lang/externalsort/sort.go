package main

import (
	"bufio"
	"fmt"
	"learngo/lang/pipeline"
	"os"
	"strconv"
)

func main() {
	inFilename := "large.in"
	outFilename := "large.out"
	//p := createNetworkPipeline(inFilename, 80000000, 4) //网络版
	p := createPipeline(inFilename, 80000000, 4)
	writeToFile(p, outFilename)
	printFile(outFilename)

}

//读取数据源，排序和归并之后通过返回的<-chan int传递
//此处打开的多个*File没有关闭，返回可以增加一个[]*File在外部关闭
func createPipeline(filename string, fileSize, chunkCount int) <-chan int {
	chunkSize := fileSize / chunkCount
	pipeline.Init()
	var sortResults []<-chan int

	for i := 0; i < chunkCount; i++ {
		file, err := os.Open(filename)
		if err != nil {
			panic(err)
		}
		//Seek设置下一次读/写的位置。offset为相对偏移量
		// 而whence决定相对位置：
		// 0为相对文件开头，1为相对当前位置，2为相对文件结尾
		file.Seek(int64(i*chunkSize), 0)
		source := pipeline.ReaderSource(bufio.NewReader(file), chunkSize)

		sortResults = append(sortResults, pipeline.InMemSort(source))
	}
	return pipeline.MergeN(sortResults...)
}

//打印
func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	p := pipeline.ReaderSource(file, -1)
	count := 0
	for v := range p {
		fmt.Println(v)
		count++
		if count >= 10 {
			break
		}
	}
}

//把归并后的结果写入到文件中
func writeToFile(p <-chan int, filename string) {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	pipeline.WriterSink(writer, p)
}

//网络版
func createNetworkPipeline(filename string, fileSize, chunkCount int) <-chan int {
	chunkSize := fileSize / chunkCount
	pipeline.Init()

	var sortAddr []string
	for i := 0; i < chunkCount; i++ {
		file, err := os.Open(filename)
		if err != nil {
			panic(err)
		}
		file.Seek(int64(i*chunkSize), 0)
		source := pipeline.ReaderSource(bufio.NewReader(file), chunkSize)

		addr := ":" + strconv.Itoa(7000+i)
		pipeline.NetworkSink(addr, pipeline.InMemSort(source))
		sortAddr = append(sortAddr, addr)
	}
	//return nil
	var sortResults []<-chan int
	for _, addr := range sortAddr {
		sortResults = append(sortResults,
			pipeline.NetworkSource(addr))
	}
	return pipeline.MergeN(sortResults...)
}
