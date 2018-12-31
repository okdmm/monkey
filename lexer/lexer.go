package lexer

type Lexer struct {
	input        string
	position     int  // これから読み込む位置
	ReadPosition int  // これから読み込む位置
	ch           byte // 現在捜査中の文字
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

//func (l *Lexer) readChar() {
//	if l.ReadPosition >= len(l.input) {
//		l.ch = 0
//	} else {
//		l.ch = l.input[l.ReadPosition]
//	}
//	l.position = l.ReadPosition
//	l.ReadPosition += 1
//}
