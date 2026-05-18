// package main

// import (
// 	"fmt"
// 	"time"
// )

// func renderData(newChan chan int) {

// 	for i := range newChan {
// 		fmt.Println(i)
// 		time.Sleep(time.Millisecond * 500)
// 	}
// }

// func main() {

// 	var newChan = make(chan int)

// 	go renderData(newChan)

// 	go func() {
// 		for i := 0; i < 20; i++ {
// 			newChan <- i
// 		}

// 		close(newChan)
// 	}()

// 	fmt.Println("After Done")

// 	time.Sleep(time.Second * 5)
// }

package main

import (
	"fmt"
	"time"
)

func task(id int) {
	fmt.Println("Doing Task", id)
}

func main() {

	for i := range 20 {
		go task(i)
	}

	time.Sleep(time.Second)

}
