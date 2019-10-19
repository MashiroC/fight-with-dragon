// Time : 2019/10/19 15:47
// Author : MashiroC

// character
package game

import (
	"fmt"
	"mashiroc.fun/dragongame/color"
)

// character.go something

const (
	maxServant = 3
)

type Character interface {
	Card
	Do(target Character)
	Servant() []Card
	CheckServant()
	Print()
	AttendNum() int
}

type baseCharacter struct {
	servant   []Card
	name      string
	maxHp     int
	hp        int
	cost      int
	maxCost   int
	hurtNum   int
	attendNum int
}

func (c *baseCharacter) Begin() {
	c.maxCost++
	c.cost = c.maxCost
	c.attendNum = 1
}

func (c *baseCharacter) end() {
}

func (c *baseCharacter) CheckServant() {
	for i, v := range c.servant {
		if v.Die() {
			c.servant = append(c.servant[:i], c.servant[i+1:]...)
			return
		}
	}
}

func (c *baseCharacter) Attend(target Card) {
	if c.attendNum == 0 {
		fmt.Println("你这回合不能进行攻击了。")
		return
	}
	c.attendNum--
	target.ReduceHp(c.hurtNum)
	c.ReduceHp(target.Hurt())
	return
}

func printCard(cards []Card, prefix string) {
	str := ""

	strName := fmt.Sprintf("%-8s\t", prefix)

	for i := 0; i < len(cards); i++ {
		strName += fmt.Sprintf("%-15s\t", cards[i].Name())
	}

	strName += "\n"
	str += strName

	strHp := "damage,cost,hp\t"
	for i := 0; i < len(cards); i++ {
		strHp += fmt.Sprintf("%s\t%s\t%s\t\t\t", color.Red(cards[i].Hurt()), color.Yellow(cards[i].Cost()), color.Green(cards[i].Hp()))
	}

	for i := len(cards); i < maxServant; i++ {
		strHp += "   \t\t\t"
	}
	strHp += "\n"
	str += strHp
	fmt.Print(str)
}

func (c *baseCharacter) Hp() int {
	return c.hp
}

func (c *baseCharacter) Name() string {
	return c.name
}

func (c *baseCharacter) Hurt() int {
	return c.hurtNum
}

func (c *baseCharacter) Cost() int {
	return c.cost
}

func (c *baseCharacter) ReduceHp(num int) {
	c.hp -= num
}

func (c *baseCharacter) Die() bool {
	return c.hp <= 0
}

func (c *baseCharacter) Servant() []Card {
	return c.servant
}

func (c *baseCharacter) AttendNum() int {
	return c.attendNum
}

func (c *baseCharacter) End() {
}

func (c *baseCharacter) SetAttendNum(n int) {
	c.attendNum = n
}

func (c *baseCharacter) BeforeDeploy() {
}
