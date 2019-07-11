package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
	// "encoding/json"
	// "os"
)

// type logger struct

func main() {

	ad := websocket.Dialer{
	}
	header := http.Header{
		"ledum":[]string{"Code","Money"},
	}
	ws,_,_ := ad.Dial(
		"wss://ws.binaryws.com/websockets/v3?app_id=1089",
		header,
	)

	for{
		err := ws.WriteMessage(1,[]byte(`{"ping":1}`))
		if err != nil{
			fmt.Println("The error is massive")
			break
		}
		
		_,msg,err := ws.ReadMessage()

		if err != nil{
			fmt.Println(err)
		}

		// msgo := make([]byte, 512000)

		fmt.Println("Received: ", string(msg))

		break
	}
	
}

// func (logger) Write(bs []byte) (int, error) {
// 	fmt.Println(string(bs))
// 	return len(bs), nil
// }



