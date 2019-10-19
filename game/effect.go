// Time : 2019/10/19 14:49
// Author : MashiroC

// dragongame
package game

// effect.go something

type Effect interface {
	Do(card Card, self, other Character)    // 落地的时候要做什么
	Begin(card Card, self, other Character) // 每回合开始的时候做的事情
	End(card Card, self, other Character)   // 每回合结束做的事情
}

type BaseEffect struct {
}

func (e BaseEffect) Do(card Card, self, other Character) {}

func (e BaseEffect) Begin(card Card, self, other Character) {}

func (e BaseEffect) End(card Card, self, other Character) {}
