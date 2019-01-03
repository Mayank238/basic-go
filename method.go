package main 
import "fmt"


type rect struct{
	width,hight int
}

func (r *rect) area()int {
	return r.width*r.hight
}

func (r rect) perim() int {
	return 2*r.width + 2*r.hight
}

func main() {
	r:= rect{10,5}

	fmt.Println("area:", r.area())
	fmt.Println("perim:", r.perim())

	rp := &r
	fmt.Println("area:", rp.area())
	fmt.Println("perim:", rp.perim())
}