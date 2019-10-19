// Time : 2019/10/19 14:51
// Author : MashiroC

// dragongame
package game

import (
	"fmt"
)

// g.go something



var g *Game

type Game struct {
	Player Character
	dragon Character
}

func Start(player Character) {
	g.Player=player

	for {
		// 回合开始进行准备
		g.Player.Begin()
		g.dragon.Begin()

		// 回合中 进行操作
		g.Player.Do(g.dragon)
		g.dragon.Do(g.dragon)

		// 回合后 判断游戏结束
		if g.Player.Die() {
			fmt.Println("你的挑战失败了，请重整实例再来挑战大恶龙")
			break
		}
		if g.dragon.Die() {
			fmt.Println("恭喜你打败了恶龙！")
			break
		}
	}

}

func init() {
	g = &Game{
		dragon: NewDragon(),
	}
}
