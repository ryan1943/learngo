package pipeline

import (
	"bufio"
	"net"
)

//排序好的结果，启动一个服务端，等待客户端发起连接，然后通过连接把数据传递过去
func NetworkSink(addr string, in <-chan int) {
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}
	//一次处理完成及关闭连接和监听
	go func() {
		defer listener.Close()
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		defer conn.Close()

		writer := bufio.NewWriter(conn)
		defer writer.Flush()
		WriterSink(writer, in)
	}()
}

//客户端
func NetworkSource(addr string) <-chan int {
	out := make(chan int)
	go func() {
		conn, err := net.Dial("tcp", addr)
		if err != nil {
			panic(err)
		}

		r := ReaderSource(bufio.NewReader(conn), -1)
		for v := range r {
			out <- v
		}
		close(out)
	}()

	return out
}
