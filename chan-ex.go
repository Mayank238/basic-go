package main 

import "fmt"

func calcsqr(num int, chsqr chan int) {
	sum := 0
	for num!=0{
		digit := num %10
		sum += digit * digit
		num /= 10
	}
	chsqr <- sum
}

func calccube(num int,chcube chan int) {
	sum:= 0
	for num != 0 {
		digit := num %10
		sum += digit*digit*digit
		num /= 10
	}
	chcube <- sum
}

func main() {
	
	num := 2
	chsqr := make(chan int)
	chcube := make(chan int)

	go calcsqr(num, chsqr)
	go calccube(num, chcube)

	sqr , cube := <-chsqr , <-chcube

	fmt.Println("sum of digit sqr: ",sqr ,"sum of digit cube: ", cube)
}