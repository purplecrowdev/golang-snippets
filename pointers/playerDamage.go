package main

import "fmt"

type Player struct {
	health int
}

func takeDamage(player *Player) {
	player.health -= 10
}

func main() {

	player := Player{
		health: 100,
	}

	takeDamage(&player)

	fmt.Println("player's health", player.health)
}
