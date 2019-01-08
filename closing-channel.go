package main 

import "fmt"

func display(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
}

func main() {
	 ch := make(chan int)

	 go display(ch)

	 for {

	 	 v, ok := <- ch
	 	 if ok == false{
	 	 	break
	 	 } 
	 	 fmt.Println(v , ok )
	 }
}