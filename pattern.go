package go_learn

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func mutexLock() {
	var count int = 0
	lock := new(sync.RWMutex)
	for i := 1; i <= 5; i++ {
		go func() {
			for j := 1; j <= 10000; j++ {
				lock.Lock()
				count += 1
				lock.Unlock()
			}
		}()
	}
	time.Sleep(time.Second * 5)
	fmt.Println("Count: ", count)
}

func fetchAPI(model string) string {
	time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	return model
}

func patternFetchAPI() {
	responseChan := make(chan string)
	var results []string
	go func() { responseChan <- fetchAPI("users") }()
	go func() { responseChan <- fetchAPI("categories") }()
	go func() { responseChan <- fetchAPI("products") }()

	for i := 1; i <= 3; i++ {
		results = append(results, <-responseChan)
	}
	fmt.Println(results)
}

func query(server string) string {
	time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	return server
}
func queryFirst(servers ...string) <-chan string {
	c := make(chan string)
	for _, serv := range servers {
		go func(s string) { c <- query(s) }(serv)
	}
	// return first, data after ignore
	return c
}
func patternQuery() {
	// leak memory - because get one but size is true
	// solution: run for loop get all data
	result := queryFirst("server 1", "server 2", "server 3")
	fmt.Println(<-result)
}

func main() {
	mutexLock()
	patternFetchAPI()
	patternQuery()
}
