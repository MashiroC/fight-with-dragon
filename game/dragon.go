// Time : 2019/10/19 15:27
// Author : MashiroC

// character
package game

import (
	"fmt"
	"mashiroc.fun/dragongame/color"
)

// dragon.go something

type dragon struct {
	baseCharacter
}



func NewDragon() Character {
	return &dragon{
		baseCharacter: baseCharacter{
			name:  "凶猛大恶龙",
			maxHp: 99,
			hp:    99,
			cost:  0,
			hurtNum:99,
			servant: []Card{NewCard("巨肉恶龙", 49, 0, 0), NewCard("巨猛恶龙", 1, 0, 49), NewCard("巨强恶龙", 49, 0, 49)},
		},
	}
}

func (d *dragon) Do(target Character) {

}

func (d *dragon) Print() {
	str :=
		"---------------------------------------------------\n" +
			"     %10s hp:%s damage:%s          \n" +
			"---------------------------------------------------\n"
	fmt.Printf(str, d.name, color.Green(d.hp),color.Red(d.hurtNum))
	printCard(d.Servant(), d.name+"的随从")

}
