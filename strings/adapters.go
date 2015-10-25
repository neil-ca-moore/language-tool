package strings

import (
	"bytes"
	"unicode"
)

type Transformer func(string) string

func swapCase(str string) string {
	var buffer bytes.Buffer
	for c := range str {
		if unicode.IsUpper(rune(c)) {
			buffer.WriteRune(unicode.ToLower(rune(c)))
		} else if unicode.IsLower(rune(c)) {
			buffer.WriteRune(unicode.ToUpper(rune(c)))
		}
	}
	return buffer.String()
}

type PairMember struct {
	next        string
	other       *PairMember
	transformer Transformer
	maker       Class
}

func (p PairMember) make() string {
	retval := p.next
	p.other.next = p.transformer(p.maker.Make())
	return retval
}

func makePairs(maker Class, transformer Transformer) (PairMember, PairMember) {
	a := PairMember{
		transformer: transformer,
		maker:       maker,
	}
	b := a
	*a.other = b
	*b.other = a
	a.make()
	return a, b
}
