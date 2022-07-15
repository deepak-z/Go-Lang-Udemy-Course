package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)
			 
			 
func main() {

	fmt.Println("Number of CPU : ",runtime.NumCPU())
	fmt.Println("Number of Go Routines : ",runtime.NumGoroutine())
	counter := 0;
	const gs = 100;

	var wg sync.WaitGroup
	wg.Add(100)

	var mu sync.Mutex


	for i:=0;i < gs;i++{
		go func ()  {
			mu.Lock()
			v := counter
			time.Sleep(time.Second)
			runtime.Gosched()
			v++
			counter = v
			mu.Unlock()
			wg.Done()
		}()
		fmt.Println("Number of Go Routines : ",runtime.NumGoroutine())
	}
	wg.Wait()

	fmt.Println("Number of CPU : ",runtime.NumCPU())
	fmt.Println("Number of Go Routines : ",runtime.NumGoroutine())
	fmt.Println("Counter : ",counter)
			 
			 
}