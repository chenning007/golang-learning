package mygoroutine

import (
	"fmt"
	"runtime"
	"sync"
)

func init() {

}

func DoGoroutine() {

	runtime.GOMAXPROCS(1)
	var wg sync.WaitGroup
	wg.Add(2)

	fmt.Println("start Goroutines")

	go func() {
		defer wg.Done()

		for count := 0; count < 3; count++ {
			for char := 'a'; char < 'a'+26; char++ {
				fmt.Println("%c", char)
			}
		}
	}()

	go func() {
		defer wg.Done()

		for count := 0; count < 3; count++ {
			for char := 'A'; char < 'A'+26; char++ {
				fmt.Println("%c", char)
			}
		}
	}()

	fmt.Println("Waiting To Finish")
	wg.Wait()

	fmt.Println("\nTerminating Program")
}
