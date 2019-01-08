package main 
import (
       "fmt"
       "time"
)

func process(ch chan string) {
	time.Sleep(10*time.Millisecond)
	ch <- "process done"
}

func main() {
	ch := make(chan string)

	go process(ch)

	for{
      //time.Sleep(time.Second)
		select{
		case v := <- ch :
			fmt.Println("received value: ", v)
			return
		default :
		    fmt.Println("no value received ")	
		}
	} 
	
}