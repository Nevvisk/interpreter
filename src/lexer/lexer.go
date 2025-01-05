package lexer

import "Interpreter/token"

type Lexer struct {
    input   string
    position int
    readPositon int
    ch      byte
}

func New(input string) *Lexer {
    l := &Lexer{input: input}
    l.readChar()
    return l
}

func (l *Lexer)readChar(*Lexer) {
    if l.readPosition >= len(l.input) {
        l.ch = 0
    } else {
        l.ch = l.input[l.readPosition]
    }
    l.position = l.readPosition
    l.readPositon += 1
}
