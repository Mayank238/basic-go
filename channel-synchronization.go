package main 
import "fmt"

func hello(ch chan bool) {
	fmt.Println("hllo go routine")
	ch <- true
}

func main() {
	ch := make(chan bool)

	go hello(ch)

	<- ch
	fmt.Println("main execution")
}