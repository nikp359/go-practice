package parentheses

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRuneStr(t *testing.T) {
	got, err := runeStr("a")
	want := rune(97)

	assert.Nil(t, err)
	assert.Equal(t, want, got)

	got, err = runeStr("aaa")
	assert.Error(t, err)
}

func TestStack(t *testing.T) {
	st := stack{}
	st.Push(rune(97))
	st.Push(rune(97))
	st.Push(rune(98))

	assert.Equal(t, "aab", string(st))
	assert.Equal(t, 3, st.Len())

	r := st.Pop()

	assert.Equal(t, "aa", string(st))
	assert.Equal(t, 2, st.Len())
	assert.Equal(t, "b", string(r))
}

func TestIsParenthes(t *testing.T) {
	testCase := []struct {
		r        rune
		expected bool
	}{
		{rune(40), true},
		{rune(41), true},
		{rune(91), true},
		{rune(93), true},
		{rune(123), true},
		{rune(125), true},
		{rune(122), false},
		{rune(124), false},
		{rune(97), false},
	}

	for _, c := range testCase {
		isP := isParenthes(c.r)
		assert.Equal(t, c.expected, isP)
		t.Logf("Rune: %v isParenthes: %v \n", string(c.r), isP)
	}
}

func TestIsOpenParenthes(t *testing.T) {
	testCase := []struct {
		r        rune
		expected bool
	}{
		{rune(40), true},
		{rune(41), false},
		{rune(91), true},
		{rune(93), false},
		{rune(123), true},
		{rune(125), false},
		{rune(122), false},
		{rune(124), false},
		{rune(97), false},
	}

	for _, c := range testCase {
		isP := isOpenParenthes(c.r)
		assert.Equal(t, c.expected, isP)
		t.Logf("Rune: %v isParenthes: %v \n", string(c.r), isP)
	}
}

func TestIsCloseParenthes(t *testing.T) {
	testCase := []struct {
		r        rune
		expected bool
	}{
		{rune(40), false},
		{rune(41), true},
		{rune(91), false},
		{rune(93), true},
		{rune(123), false},
		{rune(125), true},
		{rune(122), false},
		{rune(124), false},
		{rune(97), false},
	}

	for _, c := range testCase {
		isP := isCloseParenthes(c.r)
		assert.Equal(t, c.expected, isP)
		t.Logf("Rune: %v isParenthes: %v \n", string(c.r), isP)
	}
}
