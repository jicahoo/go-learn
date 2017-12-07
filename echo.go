// $ go run echo.go
// $ ./echo
//
//  ~ in another terminal ~
//
// $ telnet 127.0.0.1 8888

package main

import (
    "net"
    "bufio"
    "strconv"
    "fmt"
)

const PORT = 8888

func main() {
    server, err := net.Listen("tcp", ":" + strconv.Itoa(PORT))
    if server == nil {
        panic("couldn't start listening: " + err.Error())
    } else {
	fmt.Printf("[main] Listen on port: %d \n", PORT);
    }

    conns := acceptConns(server)
    rtn_nr := 0
    for {
        rtn_nr++
        fmt.Printf("[main] Before go handleConn(): %d\n", rtn_nr);
        go handleConn(<-conns)
        fmt.Printf("[main] After go handleConn(): %d\n", rtn_nr);
    }
}

func acceptConns(listener net.Listener) chan net.Conn {
    ch := make(chan net.Conn)
    i := 0
    fmt.Printf("[acceptConns] Enter.\n")
    go func() {
        fmt.Printf("[acceptConns.co] Start to Accept new connections.\n")
	//It seems like, there may be several conns.
        for {
	    fmt.Printf("[acceptConns.co] Before Accept()\n");
            client, err := listener.Accept()
	    fmt.Printf("[acceptConns.co] After Accept()\n");
            if client == nil {
                fmt.Printf("couldn't accept: " + err.Error())
                continue
            }
            i++
            fmt.Printf("[acceptConns.co] Accept new conn %d: %v <-> %v\n", i, client.LocalAddr(), client.RemoteAddr())
            ch <- client
        }
    }()
    fmt.Printf("[acceptConns] Exit.\n")
    return ch
}

func handleConn(client net.Conn) {
    fmt.Printf("[handleConn] Enter.\n")
    b := bufio.NewReader(client)
    for {
        fmt.Printf("[handleConn] Before ReadBytes\n")
        line, err := b.ReadBytes('\n')
        fmt.Printf("[handleConn] After ReadBytes\n")
        if err != nil { // EOF, or worse
	    fmt.Printf("[handleConn] Read error: %s\n", err.Error())
            break
        }
        client.Write(line)
    }
    fmt.Printf("[handleConn] Exit.\n")
}
