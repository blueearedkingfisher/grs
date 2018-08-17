package main

import (
    "fmt"
    "net"
    "os"
    "bufio"
    "io"
)

type TcpServer struct {
    listener       *net.TCPListener
    hawkServer  *net.TCPAddr
}

func main() {

    hawkServer, err := net.ResolveTCPAddr("tcp", "127.0.0.1:6969")
    checkErr(err)

    listen, err := net.ListenTCP("tcp", hawkServer)
    checkErr(err)

    defer listen.Close()
    tcpServer := &TcpServer{
        listener:listen,
        hawkServer:hawkServer,
    }
    fmt.Println("start server successful...")

    for {
        conn, err := tcpServer.listener.Accept()
        fmt.Println("accept tcp client %s",conn.RemoteAddr().String())
        checkErr(err)

        go Handle(conn)
    }
}

func Handle(conn net.Conn) {

    defer conn.Close()
    buf := make([]byte, 40)
    bufferReader := bufio.NewReader(conn)

    for {
        recvByte,err := bufferReader.Read(buf)
        if err != nil {

            if err == io.EOF {
                fmt.Printf("client %s is close!\n",conn.RemoteAddr().String())
            }

            return
        }
        fmt.Printf("data: %s \n",string(buf[:recvByte]))
    }
}

func checkErr(err error) {
    if err != nil {
        fmt.Println(err)
        os.Exit(-1)
    }
}