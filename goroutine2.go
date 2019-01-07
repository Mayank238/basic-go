package main 
import "fmt"
//import "time"

func f(from string) {
	for i := 0; i<3; i++{
		fmt.Println(from, ":", i)
		//time.Sleep(300*time.Millisecond)
	}
}

func main() {
	
	go f("direct")

	 f("goroutine")

	go func (msg string) {
		fmt.Println(msg)
	}("going")

	//fmt.Scanln()
	//fmt.Println("done")
}