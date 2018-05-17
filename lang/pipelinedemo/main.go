package main

import (
	"bufio"
	"fmt"
	"learngo/lang/pipeline"
	"os"
)

func main() {
	const filename = "large.in"
	const n = 10000000 //生成的文件大小是1000000*8字节
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	p := pipeline.RandomSource(n)
	writer := bufio.NewWriter(file)
	pipeline.WriterSink(writer, p) //写入到文件的内容是二进制
	writer.Flush()
	file, err = os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	p = pipeline.ReaderSource(bufio.NewReader(file), -1)
	count := 0
	for v := range p {
		fmt.Println(v)
		count++
		if count >= 100 {
			break
		}
	}

}

func mergeDemo() {
	in1 := pipeline.InMemSort(pipeline.ArraySource(3, 2, 7, 8, 4))
	in2 := pipeline.InMemSort(pipeline.ArraySource(9, 0, 12, 6, 1))
	p := pipeline.Merge(in1, in2)
	for v := range p {
		fmt.Println(v)
	}
}
