// Time : 2019/10/19 20:17
// Author : MashiroC

// dragongame
package main

import (
	"mashiroc.fun/dragongame/game"
	"mashiroc.fun/dragongame/game/card"
)

// main.go something

func main() {

	//player := game.NewPlayer("yrx", []game.Card{
	//	game.NewCard("工具人2", 5, 3, 3),
	//	game.NewCard("工具人3", 5, 3, 3),
	//	game.NewCard("工具人1", 5, 3, 3)},
	//)
	player := game.NewPlayer("your name", []game.Card{card.NewWildBoar(),card.NewElvenArcher()})

	game.Start(player)

}
