package main 
import "fmt"

func main() {
	
	ch := make(chan string, 3)

	ch <- "mayank"
	ch <- "aniket"
	ch <- "rishi"

	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}