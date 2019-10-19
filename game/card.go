// Time : 2019/10/19 14:38
// Author : MashiroC

// dragongame
package game

import (
	"fmt"
)

// Card.go something

// 回合开始 回合结束 攻击
type Card interface {
	Begin()
	Attend(target Card)
	End()
	SetAttendNum(int)
	Hurt() int
	Die() bool

	ReduceHp(num int)
	Name() string
	Cost() int
	BeforeDeploy()
	Hp() int
}

func NewCard(name string, hp, cost, hurt int, eff ...Effect) *BaseCard {
	if cost < 0 || hp <= 0 || hurt < 0 {
		panic("some failed game status")
	}
	return &BaseCard{
		cardName: name,
		hp:       hp,
		maxHp:    hp,
		cost:     cost,
		hurtNum:  hurt,
		effect:   eff,
	}
}

type BaseCard struct {
	cardName  string
	hp        int
	maxHp     int
	cost      int
	attendNum int
	hurtNum   int
	effect    []Effect
}

func (b *BaseCard) Die() bool {
	return b.hp <= 0
}

func (b *BaseCard) ReduceHp(num int) {
	b.hp -= num
}

// Begin 回合开始的效果
func (b *BaseCard) Begin() {
	// 普通的随从每回合开始攻击次数回到1
	b.attendNum = 1
	for _, v := range b.effect {
		v.Begin(b, g.Player, g.dragon)
	}
}

func (b *BaseCard) BeforeDeploy() {
	for _, v := range b.effect {
		v.Do(b, g.Player, g.dragon)
	}
}

// Attend 攻击
func (b *BaseCard) Attend(target Card) {
	if b.attendNum > 0 {
		target.ReduceHp(b.hurtNum)
		b.ReduceHp(target.Hurt())
		b.attendNum--
	} else {
		fmt.Println("这位随从这回合不能再攻击了！")
	}

	return
}

// End 回合结束的效果
func (b *BaseCard) End() {
	// 普通怪没有效果
}

func (b *BaseCard) Hurt() int {
	return b.hurtNum
}

func (b *BaseCard) Name() string {
	return b.cardName
}
func (b *BaseCard) SetAttendNum(n int) {
	b.attendNum = n
}

func (b *BaseCard) Cost() int {
	return b.cost
}

func (b *BaseCard) Hp() int {
	return b.hp
}
