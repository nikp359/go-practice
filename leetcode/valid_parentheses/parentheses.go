package parentheses

import (
	"errors"
	"unicode"
)

func isValid(s string) bool {
	return true
}

type stack []rune

func (s *stack) Push(r rune) {
	*s = append(*s, r)
}

func (s stack) Len() int {
	return len(s)
}

func (s *stack) Pop() rune {
	if s.Len() == 0 {
		return rune(-1)
	}

	n := s.Len() - 1
	sl := *s
	top := sl[n]
	*s = sl[:n]
	return top
}

func isParenthes(r rune) bool {
	return unicode.In(r, rangeParenthes())
}

func isOpenParenthes(r rune) bool {
	return unicode.In(r, rangeOpenParenthes())
}

func isCloseParenthes(r rune) bool {
	return unicode.In(r, rangeCloseParenthes())
}

//rune "(", ")", "{", "}", "[", "]"
func rangeParenthes() *unicode.RangeTable {
	return &unicode.RangeTable{
		R16: []unicode.Range16{
			unicode.Range16{Lo: 40, Hi: 41, Stride: 1},
			unicode.Range16{Lo: 91, Hi: 93, Stride: 2},
			unicode.Range16{Lo: 123, Hi: 125, Stride: 2}},
		R32: []unicode.Range32(nil), LatinOffset: 3}
}

//rune (, {, [
func rangeOpenParenthes() *unicode.RangeTable {
	return &unicode.RangeTable{
		R16: []unicode.Range16{
			unicode.Range16{Lo: 40, Hi: 40, Stride: 1},
			unicode.Range16{Lo: 91, Hi: 91, Stride: 1},
			unicode.Range16{Lo: 123, Hi: 123, Stride: 1}},
		R32: []unicode.Range32(nil), LatinOffset: 3}
}

//rune ), }, ]
func rangeCloseParenthes() *unicode.RangeTable {
	return &unicode.RangeTable{
		R16: []unicode.Range16{
			unicode.Range16{Lo: 41, Hi: 41, Stride: 1},
			unicode.Range16{Lo: 93, Hi: 93, Stride: 1},
			unicode.Range16{Lo: 125, Hi: 125, Stride: 1}},
		R32: []unicode.Range32(nil), LatinOffset: 3}
}

func runeStr(s string) (rune, error) {
	if len(s) != 1 {
		return rune(-1), errors.New("only one character is allowed")
	}
	r := []rune(s)
	return r[0], nil
}
