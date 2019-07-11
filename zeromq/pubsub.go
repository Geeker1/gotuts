package main

import (
	zmq "github.com/pebbe/zmq4"
	"fmt"
	"os"
	"time"
)



var flag zmq.Flag = 0

func main() {
	context, err := zmq.NewContext()
	if err != nil {
		fmt.Println("We've got a basic error here.....exiting")
		return
	}

	fmt.Println("Socket Flag",flag)

	PUB_socket,err := context.NewSocket(zmq.PUB)
	SUB_socket,err := context.NewSocket(zmq.SUB)
	PUB_socket.Bind("tcp://*:50000")
	SUB_socket.Connect("tcp://localhost:50000")
	SUB_socket.SetSubscribe("hg")

	defer SUB_socket.Close()
	defer PUB_socket.Unbind("tcp://localhost:50000")
	defer PUB_socket.Close()
	defer context.Term()

	c := make(chan string)

	go StartSocket(SUB_socket,c)
	go Print(c)

	fmt.Println("Created goroutine moving on to messages")

	for i := 0; i < 10; i++ {
		PUB_socket.Send("hg fnje",flag)
		time.Sleep(time.Second)
	}

	PUB_socket.Send("hg kill", flag)

	time.Sleep(5 * time.Second)

	os.Exit(1)

}

func Print(c chan string) {
	for i := 0; i < 10; i++ {
		fmt.Println("Monitoring values sent into channel >",<- c,"\n")
	}
}

func StartSocket(SUB *zmq.Socket, c chan string) {

	fmt.Println("Listening on SUB socket")
	for{
		msg, err := SUB.Recv(flag)
		if err != nil{
			fmt.Println("Error in code >> exiting now")
			return
		}
		fmt.Println("Received message sending into channel",msg)
		
		if msg == "kill"{
			fmt.Println("Kill message received, shutting down socket now..")
			return
		}
		c <- msg
	}
}











