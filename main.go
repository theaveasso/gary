package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("Gary the GO guys")

	start := time.Now()
	username := fetchUser()
	// likes := fetchUserLikes()
	// comments := fetchUserComments()
	// shares := fetchUserShares()

	userChan := make(chan any, 3)
	wg := &sync.WaitGroup{}

	wg.Add(3)
	go fetchUserLikes(username, userChan, wg)
	go fetchUserComments(username, userChan, wg)
	go fetchUserShares(username, userChan, wg)

	wg.Wait()
	close(userChan)

	for user := range userChan {
		fmt.Println("user", user)
	}

	fmt.Printf("time took %v\n", time.Since(start))
}

func fetchUser() string {
	time.Sleep(time.Millisecond * 100)

	return "gary"
}

func fetchUserLikes(username string, userCh chan any, wg *sync.WaitGroup) {
	time.Sleep(time.Millisecond * 150)

	userCh <- 100
	wg.Done()
}

func fetchUserComments(username string, userCh chan any, wg *sync.WaitGroup) {
	time.Sleep(time.Millisecond * 200)

	userCh <- 15
	wg.Done()
}

func fetchUserShares(username string, userCh chan any, wg *sync.WaitGroup) {
	time.Sleep(time.Millisecond * 50)

	userCh <- 8
	wg.Done()
}
