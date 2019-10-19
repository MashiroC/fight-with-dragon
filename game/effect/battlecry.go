// Time : 2019/10/19 20:26
// Author : MashiroC

// effect
package effect

import "mashiroc.fun/dragongame/game"

// Charge.go something

func NewBattleCry(fun func(game.Card, game.Character, game.Character)) BattleCry {
	return BattleCry{
		BaseEffect: game.BaseEffect{},
		isDo:       false,
		fun:        fun,
	}
}

type BattleCry struct {
	game.BaseEffect
	isDo bool
	fun  func(game.Card, game.Character, game.Character)
}

func (b BattleCry) Do(card game.Card, self, other game.Character) {
	if !b.isDo {
		b.fun(card, self, other)
	}
}
