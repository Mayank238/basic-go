package main 
import "fmt"

func main() {
	
	ch := make(chan string)

	go func() {
		ch <- "Mayank"
	}()

	name := <- ch
	fmt.Println(name)
}


