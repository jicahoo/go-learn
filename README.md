# go-learn
Go programming language learn tips.

## Examples
* https://gobyexample.com

## Go echo server
* http://blog.csdn.net/d_guco/article/details/75150696
* http://blog.leanote.com/post/ljie-pi/goroutine-%E5%92%8C-go-channel-%E7%AE%80%E6%98%93%E5%AE%9E%E7%8E%B0
    * 这个材料使用C语言的ucontext机制实现了类似Go的routine. 
* https://groups.google.com/forum/#!topic/golang-china/q4pFH-AGnfs
```go
//

//A echo server with max-connections limit and interval connection show

//

package main


import (

    "fmt"

    "net"

    "os"

    "time"

)


const (

    MAX_CONN_NUM = 5

)


//echo server Goroutine

func EchoFunc(conn net.Conn, conn_close_flag chan int) {

    defer conn.Close()

    defer func() {

        conn_close_flag <- -1

    }()


    buf := make([]byte, 1024)

    for {

        _, err := conn.Read(buf)

        if err != nil {

            //println("Error reading:", err.Error())

            return

        }

        //send reply

        _, err = conn.Write(buf)

        if err != nil {

            //println("Error send reply:", err.Error())

            return

        }

    }

}


//initial listener and run

func main() {


    listener, err := net.Listen("tcp", "0.0.0.0:8088")


    if err != nil {

        println("error listening:", err.Error())

        os.Exit(1)

    }


    defer listener.Close()


    fmt.Printf("running ...\n")


    var cur_conn_num float64 = 0


    ch_conn_change := make(chan int, MAX_CONN_NUM)


    tick := time.Tick(1e8)


    for {

        //read all close flags berfor accept new connection

        //TODO: better code to handle batch close?

        readmore := 1

        for readmore > 0 {

            select {

            case conn_change := <-ch_conn_change:

                cur_conn_num = cur_conn_num + float64(conn_change)

            default:

                readmore = 0

            }

        }

        //FIXME: tick block by listener.Accept()

        select {

        case <-tick:

            fmt.Printf("cur conn num: %f\n", cur_conn_num)

        default:

        }

        if cur_conn_num >= MAX_CONN_NUM {

            //reach MAX_CONN_NUM, waiting for exist connection close

            time.Sleep(time.Second)

        } else {

            //accept new connetion

            conn, err := listener.Accept()

            if err != nil {

                println("Error accept:", err.Error())

                return

            }

            cur_conn_num++

            go EchoFunc(conn, ch_conn_change)

        }

    }

}
```
