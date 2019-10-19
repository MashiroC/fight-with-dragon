// Time : 2019/10/19 15:46
// Author : MashiroC

// character
package game

import (
	"fmt"
	"mashiroc.fun/dragongame/color"
	"time"
)

var (
	pause = 1 * time.Second
)

// character.go something

func NewPlayer(name string, handCard []Card) Character {
	return &player{
		baseCharacter: baseCharacter{
			hurtNum: 2,
			name:    name,
			maxHp:   30,
			hp:      30,
			servant: []Card{},
		},
		handCard: handCard,
	}
}

type player struct {
	baseCharacter
	handCard []Card
}

func (p *player) Print() {
	printCard(p.Servant(), p.name+"的随从")
	str :=
		"---------------------------------------------------\n" +
			"     %10s hp:%s damage:%s          \n" +
			"---------------------------------------------------\n"
	fmt.Printf(str, p.name, color.Green(p.hp), color.Red(p.hurtNum))
	printCard(p.handCard, p.name+"的手牌 "+color.Yellow(p.cost))
}

func (p *player) Begin() {
	p.baseCharacter.Begin()
	for _, v := range p.servant {
		v.Begin()
	}
}

func (p *player) Do(target Character) {

	for {
		time.Sleep(pause)
		//  打印图案

		target.Print()
		fmt.Printf("---------------------------------------------------\n---------------------------------------------------\n")
		p.Print()

		// 判断是否存活
		if p.Die() {
			fmt.Println("you are die")
			return
		}

		// 输入命令

		fmt.Println("\n你的回合，请开始你的行动：\n输入'a x y'代表使用第x张随从攻击对方第y张随从，0代表恶龙。\n" +
			"输入's x y'代表将你手牌第x张放到第y个随从之后，当y大于在场随从时放到最后一个位置\n" +
			"输入'e 0 0'结束当前回合")

		var flag rune
		var x, y int
		n, err := fmt.Scanf("%c%d%d", &flag, &x, &y)
		if n != 3 || err != nil || !(flag == 'a' || flag == 's' || flag == 'e') {
			panic("输入错误 请重新开始游戏")
		}

		// 解析命令

		if flag == 'a' {
			// 攻击

			// 判断一些错误输入的情况
			if x > len(p.servant) {
				fmt.Println("你没那么多随从")
				continue
			}
			if y > len(target.Servant()) {
				fmt.Println("对方没那么多随从")
				continue
			}

			// 找攻击目标 开始攻击
			// self是你的随从或者英雄 other是对方的随从或者英雄

			var self, other Card
			if x == 0 {
				self = p
			} else {
				self = p.servant[x-1]
			}

			if y == 0 {
				other = target
			} else {
				other = target.Servant()[y-1]
			}

			// 攻击
			self.Attend(other)

			// 攻击完了检查一下随从状态 删掉死去的随从
			p.CheckServant()
			target.CheckServant()

		} else if flag == 's' {

			// 下随从

			// 判断一些错误情况
			if x >= len(p.handCard) {
				fmt.Println("输入错误")
				continue
			}
			if len(p.servant) == 3 {
				fmt.Println("场上位置满了")
				continue
			}

			servant := p.handCard[x]
			servant.BeforeDeploy()

			// 这里判断一下费用够不够
			if p.cost < servant.Cost() {
				fmt.Println("你的费不够")
				continue
			}
			p.cost -= servant.Cost()

			if y >= len(p.servant) {
				// y大于等于随从数 放到最后一个

				p.servant = append(p.servant,servant)

				p.handCard = append(p.handCard[:x], p.handCard[x+1:]...)

			} else {
				// 否则就放在中间

				tmp := append([]Card{}, p.servant[y:]...)

				p.servant = append(append(p.servant[:y], servant), tmp...)

				p.handCard = append(p.handCard[:x], p.handCard[x+1:]...)

			}


		} else {
			break
		}

	}

}
