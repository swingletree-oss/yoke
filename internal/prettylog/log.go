package prettylog

import "fmt"

const infoChar = "\u2139"
const checkChar = "\u2713"
const failChar = "\u2717"
const warnChar = "\u203c"

func Info(s string, data ...interface{}) {
	fmt.Printf(" %s %s\n", infoChar, fmt.Sprintf(s, data...))
}

func Check(s string, data ...interface{}) {
	fmt.Printf(" %s %s\n", checkChar, fmt.Sprintf(s, data...))
}

func Fail(s string, data ...interface{}) {
	fmt.Printf(" %s %s\n", failChar, fmt.Sprintf(s, data...))
}

func Warn(s string, data ...interface{}) {
	fmt.Printf(" %s %s\n", warnChar, fmt.Sprintf(s, data...))
}
