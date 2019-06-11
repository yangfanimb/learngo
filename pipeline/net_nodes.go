package pipeline

import (
	"bufio"
	"net"
)

func NetworkShink(addr string, in <-chan int) {
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}

	// fmt.Println("listen on: ", addr)
	go func() {
		defer listener.Close()

		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		defer conn.Close()

		writer := bufio.NewWriter(conn)
		defer writer.Flush()
		WriteSink(writer, in)
	}()
}

func NetworkSource(addr string) <-chan int {
	out := make(chan int)

	go func() {
		conn, err := net.Dial("tcp", addr)
		if err != nil {
			panic(err)
		}
		// fmt.Println("connect to: ", addr)
		r := ReaderSource(bufio.NewReader(conn), -1)
		for v := range r {
			out <- v
		}
		// fmt.Println("collect ok")
		close(out)
	}()

	return out
}
