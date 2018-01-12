package main

import (
    "fmt"
    "bufio"
    "log"
    "net"
)

func main() {
    listener, error := net.Listen("tcp", ":8080")
    if error != nil {
        log.Fatalln(error)
    }

    defer listener.Close()

    fmt.Println("Awaiting request...")
    for {
        connection, error := listener.Accept()
        if error != nil {
            log.Println(error)
            continue
        }

        go handle(connection)
    }
} 

func handle(connection net.Conn) {
    scanner := bufio.NewScanner(connection)
    for scanner.Scan() {
        line := scanner.Text()
        fmt.Println(line)
    }

    defer connection.Close()

    fmt.Println("Code made it here...")
}