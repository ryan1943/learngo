package pipeline

import (
	"encoding/binary"
	"fmt"
	"io"
	"math/rand"
	"sort"
	"time"
)

/*
小端口诀: 高位在高地址, 低位在低地址
大端口诀: 高位在低地址, 低位在高地址
long test = 0x313233334;
小端机器:
低地址 -->　高地址
00000010: 34 33 32 31
大端机器:
低地址 -->　高地址
00000010: 31 32 33 34
*/

var startTime time.Time

func Init() {
	startTime = time.Now()
}

//读取数据源
func ArraySource(a ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, v := range a {
			out <- v
		}
		close(out)
	}()
	return out
}

//内排序
func InMemSort(in <-chan int) <-chan int {
	out := make(chan int, 1024)
	go func() {
		//Read into memory
		var a []int
		for v := range in {
			a = append(a, v)
		}
		fmt.Println("Read done:", time.Now().Sub(startTime))
		sort.Ints(a)
		fmt.Println("InMemSort done:", time.Now().Sub(startTime))

		for _, v := range a {
			out <- v
		}
		close(out)
	}()
	return out
}

//归并
func Merge(in1, in2 <-chan int) <-chan int {
	out := make(chan int, 1024)
	go func() {
		v1, ok1 := <-in1 //检查channel是否被关闭
		v2, ok2 := <-in2
		for ok1 || ok2 {
			// !ok2为true时，ok1=true, ok2=false,不执行右边的判断
			// !ok2为false时, ok2 = true,然后判断(ok1 && v1 <= v2)
			if !ok2 || (ok1 && v1 <= v2) {
				out <- v1
				v1, ok1 = <-in1
			} else {
				out <- v2
				v2, ok2 = <-in2
			}
		}
		close(out)
		fmt.Println("Merge done:", time.Now().Sub(startTime))
	}()
	return out
}

//读取数据源
//chunkSize=-1表示可以读取全部，否则读取的内容大小超过chunkSize即停止
func ReaderSource(reader io.Reader, chunkSize int) <-chan int {
	out := make(chan int, 1024)
	go func() {
		buffer := make([]byte, 8)
		bytesRead := 0
		for {
			n, err := reader.Read(buffer)
			bytesRead += n
			if n > 0 {
				//按照大端序的格式将到buffer的数据反序列化为uint64类型
				v := int(binary.BigEndian.Uint64(buffer))
				out <- v
			}
			if err != nil || (chunkSize != -1 && bytesRead >= chunkSize) {
				break
			}
		}
		close(out)
	}()
	return out
}

//写入数据
func WriterSink(writer io.Writer, in <-chan int) {
	for v := range in {
		buffer := make([]byte, 8)
		//按照大端序的格式将uint64类型的数据序列化到buffer
		binary.BigEndian.PutUint64(buffer, uint64(v))
		writer.Write(buffer)
	}
}

//产生随机数
func RandomSource(count int) <-chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < count; i++ {
			out <- rand.Int()
		}
		close(out)
	}()
	return out
}

//多路的两两归并
func MergeN(inputs ...<-chan int) <-chan int {
	if len(inputs) == 1 {
		return inputs[0]
	}
	m := len(inputs) / 2
	return Merge(
		MergeN(inputs[:m]...),
		MergeN(inputs[m:]...))
}

//Merge->MergeN(inputs[0:2])->Merge(MergeN(inputs[0:1]), MergeN(inputs[1:2]))->Merge(inputs[0], inputs[1])
//	   ->MergeN(inputs[2:5])->Merge(MergeN(inputs[2:3]), MergeN(inputs[3:5]))->Merge(inputs[2], Merge(MergeN(inputs[3:4]), MergeN(inputs[4:5])))
