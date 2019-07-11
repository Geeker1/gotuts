package main

import(
	"encoding/json"
	"fmt"
	"io/ioutil"
	// "strings"
)

type kil struct{
	Firstname int `json:"firstname"`
	Lastname string `json:"lastname"`
}

type cod struct{
	Ping string `json:"ping"`
	Codex string `json:"codex"`
}


func main() {
	
	byt := []byte(
		`{"ping":"43","codex":"9"}`)

	ioutil.WriteFile("filename",byt, 0666)
	bs, _ := ioutil.ReadFile("filename")

	json.Unmarshal(bs, &gh)

	fmt.Printf("%+v",gh)
}











