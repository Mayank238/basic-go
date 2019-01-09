package main 
import "fmt"

func send(chsend chan <- string) {
	chsend <-"message passed"
}

func receive(chreceive <- chan string,chsend chan <- string) {
	msg := <- chreceive
	chsend <- msg
}

func main() {
	chsend := make(chan string)
	chreceive := make(chan string)


	go send(chsend)
	go receive(chreceive,chsend)

	fmt.Println(<-chsend)

}