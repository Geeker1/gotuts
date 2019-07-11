package main

import (
	zmq "github.com/pebbe/zmq4"
	"fmt"
	"os"
)

var flag zmq.Flag = 0

func main() {
	context, err := zmq.NewContext()
	if err != nil {
		fmt.Println("We've got a basic error here.....exiting")
		return
	}

	fmt.Println("Socket Flag",flag)

	REQ_socket,err := context.NewSocket(zmq.REQ)
	REP_socket,err := context.NewSocket(zmq.REP)
	REP_socket.Bind("tcp://*:50000")
	REQ_socket.Connect("tcp://localhost:50000")

	defer REQ_socket.Close()
	defer REP_socket.Unbind("tcp://localhost:50000")
	defer REP_socket.Close()
	defer context.Term()

	c := make(chan string)

	go StartSocket(REQ_socket,REP_socket, c)

	for {
		msg := <- c
		fmt.Println("Message from channel is:",msg)
		break
	}


	os.Exit(1)

	// fmt.Println(socket)
}

func StartSocket(REQ,REP *zmq.Socket, c chan string) {
	REQ.Send("hello",flag)

	msg, _ := REP.Recv(flag)
	fmt.Println(msg)
	c <- msg
}











