package main

import (
	"fmt"
	"io"
	"net"
)

func handleFunc(c net.Conn){

	defer c.Close()

	fmt.Println("client connected",c.RemoteAddr())

	data := make([]byte,1024)

	for {
		n,err :=  c.Read(data)
		if err != nil {
			if err == io.EOF{
				fmt.Println("Client disconnected gracefully.")
			}else {
				fmt.Println("error:", err)
			}
			return
		}
		fmt.Printf("Received: %s \n", string(data[:n]))
		
		fmt.Fprintf(c,"res : %s \n",data[:n])
	}
	
	
}


func main(){
	l,err := net.Listen("tcp",":8080")
	if err != nil {
		fmt.Println("error creating server",err)
		return
	}
	fmt.Println("Server is listening on port 8080...")

	for {
		c,err := l.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go handleFunc(c)

	}
}