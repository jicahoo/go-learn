# go-learn
Go programming language learn tips.

## Examples
* https://gobyexample.com

## Go echo server
* Go语言中的异步式的echo server是非常简洁的。参见repo中的echo.go. 它利用Go语言的特性，用Channel和Go Routine做并发编程，在加上netpoller，负责监听网络上的IO事件。
* 在Liunx上, 执行命令 `go run echo.go`，用`telnet 127.0.0.1 8888`去连接server, 发一些消息。
* 如何看出在Linux上`go run echo.go`其实是使用了epoll的。 
    * `ps -L -p <EchoServerPid>` #找出所有LWP
    * `strace -p <LWP_PID>` 尝试监视这些LWP调用了哪些系统调用。你会发现epoll相关的系统调用：
    ```shell
stack@ubuntu16-dev:~/go$ sudo strace -p 30412
strace: Process 30412 attached
epoll_wait(4,


[{EPOLLIN, {u32=2678080184, u64=140216180425400}}], 128, -1) = 1
clock_gettime(CLOCK_MONOTONIC, {1331012, 46757911}) = 0
futex(0x65fe90, FUTEX_WAKE, 1)          = 1
futex(0x65fdd0, FUTEX_WAKE, 1)          = 1
accept4(3, {sa_family=AF_INET6, sin6_port=htons(53446), inet_pton(AF_INET6, "::ffff:127.0.0.1", &sin6_addr), sin6_flowinfo=0, sin6_scope_id=0}, [28], SOCK_CLOEXEC|SOCK_NONBLOCK) = 6
epoll_ctl(4, EPOLL_CTL_ADD, 6, {EPOLLIN|EPOLLOUT|EPOLLRDHUP|EPOLLET, {u32=2678079800, u64=140216180425016}}) = 0
getsockname(6, {sa_family=AF_INET6, sin6_port=htons(8888), inet_pton(AF_INET6, "::ffff:127.0.0.1", &sin6_addr), sin6_flowinfo=0, sin6_scope_id=0}, [28]) = 0
setsockopt(6, SOL_TCP, TCP_NODELAY, [1], 4) = 0
write(1, "[acceptConns.co] After Accept()\n", 32) = 32
write(1, "[acceptConns.co] Accept new conn"..., 71) = 71
futex(0xc82002a908, FUTEX_WAKE, 1)      = 1
write(1, "[acceptConns.co] Before Accept()"..., 33) = 33
accept4(3, 0xc820039b10, 0xc820039b0c, SOCK_CLOEXEC|SOCK_NONBLOCK) = -1 EAGAIN (Resource temporarily unavailable)
futex(0xc82002b108, FUTEX_WAIT, 0, NULL

    ```
* http://blog.csdn.net/d_guco/article/details/75150696
* http://blog.leanote.com/post/ljie-pi/goroutine-%E5%92%8C-go-channel-%E7%AE%80%E6%98%93%E5%AE%9E%E7%8E%B0
    * 这个材料使用C语言的ucontext机制实现了类似Go的routine. 
* https://groups.google.com/forum/#!topic/golang-china/q4pFH-AGnfs
