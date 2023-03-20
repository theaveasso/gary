package main

import (
	"fmt"
)

type Player struct {
	health int
}

func main() {
	gary := &Player{health: 100}
	pete := &Player{health: 250}

	fmt.Printf("before explosion %v\n", gary)
	fmt.Printf("before explosion %v\n", pete)
	takeDamage(gary, 10)
	takeDamage(pete, 20)
	fmt.Printf("after explosion %v\n", gary)
	fmt.Printf("after explosion %v\n", pete)

	gary.takeDamageFromExplosion(20)
	fmt.Printf("after after explosion %v\n", gary)
}

func takeDamage(player *Player, dmg int) {
	player.health -= dmg

	fmt.Println("player is taking damage from explosion")
}

func (player *Player) takeDamageFromExplosion(dmg int) {
	player.health -= dmg

	fmt.Println("player is taking damage from explosion")
}
