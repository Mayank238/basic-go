package main 
import "fmt"

type person struct {
	name string
	age int
}

func main() {

	

	fmt.Println(person{"mayank",20})

	fmt.Println(person{name: "aniket", age: 40})

	fmt.Println(person{name: "Fred"})

	fmt.Println(&person{name: "Ann", age: 50})
    
    p1:=person{"Bob",20}
	fmt.Println(p1)
	fmt.Println(p1.name)

	s := &p1
	fmt.Println(s.age)

	s.age=51
	fmt.Println(s.age)
}