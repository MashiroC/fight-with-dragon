// Time : 2019/10/19 17:27
// Author : MashiroC

// color
package color

import "fmt"

// color.go something

const (
	yellow = 33
	green  = 32
	red    = 31
)

func Red(s int) string {
	return fmt.Sprintf("\x1b[0;%dm%d\x1b[0m", red, s)
}

func Yellow(s int) string {
	return fmt.Sprintf("\x1b[0;%dm%d\x1b[0m", yellow, s)
}

func Green(s int) string {
	return fmt.Sprintf("\x1b[0;%dm%d\x1b[0m", green, s)
}
