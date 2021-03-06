package main 
import (
	"fmt"
	//"runtime"
	"sync"
)

var wg=sync.WaitGroup{}
var counter= 0
var m = sync.RWMutex{}

func main() {
	for i := 0; i < 10; i++ {
		wg.Add(2)
		m.RLock()
		go sayHello()
		m.Lock()
		go increment()
	}
	wg.Wait()

}

func sayHello() {
   //	m.RLock()
	fmt.Printf("Heloo #%v\n",counter)
	m.RUnlock()
	wg.Done()
}

func increment() {
	//m.Lock()
	counter++
	m.Unlock()
	wg.Done()
}
