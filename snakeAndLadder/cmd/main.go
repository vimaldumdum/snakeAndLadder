package main

import (
	"fmt"
	"snakeAndLadder/model"
	"snakeAndLadder/service"
)

var snake, ladder map[int]int

func main() {

	var snakes, ladders, players, start, end int
	fmt.Printf("Enter the number of snakes: ")
	fmt.Scanf("%d", &snakes)
	snake = make(map[int]int)
	for i := 0; i < snakes; i++ {
		fmt.Scanf("%d %d", &start, &end)
		snake[start] = end
	}

	fmt.Printf("Enter the number of ladders: ")
	fmt.Scanf("%d", &ladders)
	ladder = make(map[int]int)
	for i := 0; i < ladders; i++ {
		fmt.Scanf("%d %d", &start, &end)
		ladder[start] = end
	}
	var name string
	fmt.Printf("Enter the number of players: ")
	fmt.Scanf("%d", &players)
	names := make([]string, players)

	for i := 0; i < players; i++ {
		fmt.Printf("Enter player name: ")
		fmt.Scanf("%s", &name)
		names[i] = name
	}

	playerList := model.GenerateAllPlayers(names, players)
	fmt.Println(playerList)

	service.Play(playerList, snake, ladder)
}
