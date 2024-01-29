package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	now := time.Now()
	resChannel := make(chan string, 128)
	userId := "1"

	wg := &sync.WaitGroup{}

	go fetchUserData(userId, resChannel, wg)
	wg.Add(1)
	go fetchUserRecommendations(userId, resChannel, wg)
	wg.Add(1)
	go fetchUserLikes(userId, resChannel, wg)
	wg.Add(1)

	wg.Wait()

	close(resChannel)

	for res := range resChannel {
		fmt.Println(res)
	}

	elapsed := time.Since(now)
	fmt.Printf("took %s\n", elapsed)
}

func fetchUserData(userId string, resChannel chan string, wg *sync.WaitGroup) {
	time.Sleep(80 * time.Millisecond)

	resChannel <- "user data"

	wg.Done()
}

func fetchUserRecommendations(userId string, resChannel chan string, wg *sync.WaitGroup) {
	time.Sleep(120 * time.Millisecond)

	resChannel <- "user recommendations"

	wg.Done()
}

func fetchUserLikes(userId string, resChannel chan string, wg *sync.WaitGroup) {
	time.Sleep(50 * time.Millisecond)

	resChannel <- "user likes"

	wg.Done()
}
