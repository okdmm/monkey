package parser

import (
	"fmt"
	"strings"
)

var taraceLevel int = 0

const tarceIdentPlaceholder string = "\t"

func identLevel() string {
	return strings.Repeat(tarceIdentPlaceholder, taraceLevel-1)
}

func tracePrint(fs string) {
	fmt.Printf("%s%s\n", identLevel(), fs)
}

func incIdent() { taraceLevel = taraceLevel + 1 }
func decIdent() { taraceLevel = taraceLevel - 1 }

func trace(msg string) string {
	incIdent()
	tracePrint("BEGIN " + msg)
	return msg
}

func untrace(msg string) {
	tracePrint("END" + msg)
	decIdent()
}
