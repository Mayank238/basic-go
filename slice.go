package main 
import "fmt"

func main() {
	//create slice

	s:= make([]string, 3)
	fmt.Println("emp:",s)

	s[0]="a"
	s[1]="b"
	s[2]="c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	fmt.Println("len:", len(s))

	//append methode used in slice.this is not used in array.

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	c:=make([]string, len(s))
	copy(c,s)
	fmt.Println("cpy:", c)

	l:=s[2:5]
	fmt.Println("sl1:", l)

	l=s[:5]
	fmt.Println("sl2:", l)

	l=s[2:]
	fmt.Println("sl3:", l)


	array:=[]string{"mayank","Rishi","aniket"}
	fmt.Println(array)

	var d [3]string
	d[0]="mayank"
	d[1]="Rishi"
	d[2]="aniket"

    fmt.Println(d)

    twoD := make([][]int, 3)
    for i := 0; i < 3; i++ {
    	innerLen:=i+1
    	twoD[i] = make([]int, innerLen)
    	for j := 0; j < innerLen; j++ {
    		twoD[i][j] = i+j
    	}
    }

   fmt.Println("2d:", twoD)
}   
