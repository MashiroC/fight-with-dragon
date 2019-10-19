// Time : 2019/10/19 21:13
// Author : MashiroC

// effect
package effect

import "mashiroc.fun/dragongame/game"

// charge.go something

type Charge struct {
	game.BaseEffect
}

func (c Charge) Do(card game.Card, self, other game.Character) {
	card.SetAttendNum(1)
}
