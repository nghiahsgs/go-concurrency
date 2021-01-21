package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	mutext  sync.Mutex
	balance int
)

func deposit(value int, wg *sync.WaitGroup) {
	mutext.Lock()
	time.Sleep(5 * time.Second)
	fmt.Printf("Depositing %d to account with balance %d\n", value, balance)
	balance += value
	mutext.Unlock()
	wg.Done()

}
func withDraw(value int, wg *sync.WaitGroup) {
	mutext.Lock()
	time.Sleep(5 * time.Second)
	fmt.Printf("Withdraw %d to account with balance %d\n", value, balance)
	balance -= value
	mutext.Unlock()
	wg.Done()
}
func main() {
	balance = 1000
	fmt.Println("Start")

	var wg sync.WaitGroup
	wg.Add(2)
	go deposit(500, &wg)
	go withDraw(700, &wg)
	wg.Wait()

	fmt.Println("New Balance", balance)
}
