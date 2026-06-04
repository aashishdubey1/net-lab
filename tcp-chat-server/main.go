package main

import (
	"fmt"
	"net"
)

var connList []net.Conn

func brodcast(msg string,conns []net.Conn){
	for _,c := range conns {
		fmt.Fprintf(c,"msg -> %s",msg)
	}
}

func handleConn(c net.Conn){

	defer c.Close()

	data := make([]byte,1024)

	for {
		n,err := c.Read(data)
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println(string(data[:n]))

		brodcast(string(data[:n]),connList)
	}

}





func main(){
	listener , err := net.Listen("tcp",":8080")
	if err != nil {
		fmt.Println("error creating server",err)
	}
	fmt.Println("server is running...")	

	for {
		conn , err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}

		connList = append(connList, conn)
		go handleConn(conn)
	}

}