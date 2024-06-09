package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/Nikhil-Giramkar/Pub-Sub-Go/contracts"
	"github.com/Nikhil-Giramkar/Pub-Sub-Go/implementations"
)

func main() {
	creator1 := implementations.NewCreator("Creator 1")

	user1 := implementations.NewUser("User 1", creator1.SubscribeToCreator())
	user2 := implementations.NewUser("User 2", creator1.SubscribeToCreator())
	actionsChannel := make(chan string)

	var wg sync.WaitGroup
	wg.Add(1)
	go actionsByCreator(creator1, actionsChannel, &wg)
	wg.Add(1)
	go actionsByUser(user1, user2, actionsChannel, &wg)
	wg.Wait()
	fmt.Println("------------------------")
}

func actionsByCreator(c contracts.Creator, actionsChannel chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	c.AddVideo("Title1", "Desc1")
	actionsChannel <- "New Video Added"

	time.Sleep(time.Second*2) 
	
	// Creator adds another video
	c.AddVideo("Title2", "Desc2")
	actionsChannel <- "New Video Added"
	
	close(actionsChannel) 
}

func actionsByUser(u1, u2 contracts.User, actionsChannel <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for action := range actionsChannel {
		fmt.Println(action)
		u1.ListenUpdates()
		u2.ListenUpdates()
	}
}