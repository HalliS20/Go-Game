package main

import (
	"fmt"
	"go-game/player"
)

var playerObj *player.Player

func main() {
	// get input for the name from user
	name := getName()
	playerObj = player.NewPlayer(name)
	playerObj.RotateRight()
	playerObj.MoveForward()
}

func getName() string {
	var name string = ""
	for name == "" {
		fmt.Print("Enter players name: ")
		_, err := fmt.Scanln(&name)
		if err != nil {
			fmt.Println("Error scanning name: ", err)
		}
	}
	return name
}

// ┌─────┬─────┬─────┬─────┬─────┬─────┬─────┐
// │     │     │     │     │     │     │     │
// ├─────┼─────┼─────┼─────┼─────┼─────┼─────┤
// │     │     │     │     │     │     │     │
// ├─────┼─────┼─────┼─────┼─────┼─────┼─────┤
// │     │     │     │     │     │     │     │
// ├─────┼─────┼─────┼─────┼─────┼─────┼─────┤
// │     │     │     │  ■  │     │     │     │
// ├─────┼─────┼─────┼─────┼─────┼─────┼─────┤
// │     │     │     │     │     │     │     │
// ├─────┼─────┼─────┼─────┼─────┼─────┼─────┤
// │     │     │     │     │     │     │     │
// ├─────┼─────┼─────┼─────┼─────┼─────┼─────┤
// │     │     │     │     │     │     │     │
// └─────┴─────┴─────┴─────┴─────┴─────┴─────┘
