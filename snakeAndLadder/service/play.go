package service

import (
	"fmt"
	"math/rand"
	"snakeAndLadder/model"
)

func Play(players []*model.Player, snake, ladder map[int]int) {
	n := len(players)
	var count int = 0
	fmt.Printf("Inside play service\n")

	for true {
		if turn(players[count%n], snake, ladder) {
			return
		}
		count++
	}

}

func turn(player *model.Player, snake, ladder map[int]int) bool {
	roll := rand.Intn(100000) % 6
	roll++

	newPos := player.GetPos() + roll

	_, s := snake[newPos]
	_, l := ladder[newPos]

	if s {
		newPos = snake[newPos]
	}

	if l {
		newPos = ladder[newPos]
	}

	if newPos == 100 {
		fmt.Printf("%s wins the game\n", player.GetName())
		return true
	}

	if newPos > 100 {
		fmt.Printf("%s rolls over 100\n", player.GetName())
		return false
	}
	fmt.Printf("%s rolled a %d and moved from %d to %d\n", player.GetName(), roll, player.GetPos(), newPos)
	player.UpdatePos(newPos)
	return false
}
