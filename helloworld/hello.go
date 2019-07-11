package main
import "fmt"



// Declaring Functions in go

func printf(str string, args ...interface{}) (int,error){
	_, err := fmt.Printf(str,args...)
	
	return len(args), err
}

func main(){
	count := 1
	closure := func(msg string){
		printf("%d %s\n", count, msg)
		count++
	}
	fmt.Printf("Hello World!\n")
	closure("A Message")
	closure("Another message")
}