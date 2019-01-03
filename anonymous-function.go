package main 
import "fmt"

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
	
}

func main() {

	//this is another way to define anonymous function

	abc := func () int{
		
		return 3
	}
	
	nextInt := intSeq()

	//print the address

	fmt.Println(intSeq())

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
    fmt.Println(abc())

	newInt := intSeq()
	fmt.Println(newInt())
}