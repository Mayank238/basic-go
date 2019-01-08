package main 
import (
        "fmt"
        
)

func main() {


	 ch := make(chan string)
	go a(ch)
	go b()
	fmt.Println(<-ch)
    // time.Sleep(3*time.Second)

}

func a(d chan string) {
	
	for i := 0; i < 3; i++ {
		fmt.Println("hello")
		 // time.Sleep(1*time.Second)
		d <- "mayank"
	}
}

func b() {
	
	for i := 0; i < 3; i++ {
		 //time.Sleep(1*time.Second)
		fmt.Println("hii")
	}
}