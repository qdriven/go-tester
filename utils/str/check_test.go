package str

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsAlphabet(t *testing.T) {
	assert.False(t, IsAlphabet('9'))
	assert.False(t, IsAlphabet('+'))

	assert.True(t, IsAlphabet('A'))
	assert.True(t, IsAlphabet('a'))
	assert.True(t, IsAlphabet('Z'))
	assert.True(t, IsAlphabet('z'))
}

func TestIsAlphaNum(t *testing.T) {
	assert.False(t, IsAlphaNum('+'))

	assert.True(t, IsAlphaNum('9'))
	assert.True(t, IsAlphaNum('A'))
	assert.True(t, IsAlphaNum('a'))
	assert.True(t, IsAlphaNum('Z'))
	assert.True(t, IsAlphaNum('z'))
}

func TestEquals(t *testing.T) {
	assert.True(t, Equal("a", "a"))
	assert.False(t, Equal("a", "b"))
}

func TestLen(t *testing.T) {
	str := "Hello, 世界"

	assert.Equal(t, 7, Len("Hello, "))
	assert.Equal(t, 13, Len(str))
	assert.Equal(t, 9, Utf8len(str))
}

func TestStrPos(t *testing.T) {
	// StrPos
	assert.Equal(t, -1, StrPos("xyz", "a"))
	assert.Equal(t, 0, StrPos("xyz", "x"))
	assert.Equal(t, 2, StrPos("xyz", "z"))

	// RunePos
	assert.Equal(t, -1, RunePos("xyz", 'a'))
	assert.Equal(t, 0, RunePos("xyz", 'x'))
	assert.Equal(t, 2, RunePos("xyz", 'z'))
	assert.Equal(t, 5, RunePos("hi时间", '间'))

	// BytePos
	assert.Equal(t, -1, BytePos("xyz", 'a'))
	assert.Equal(t, 0, BytePos("xyz", 'x'))
	assert.Equal(t, 2, BytePos("xyz", 'z'))
	// assert.Equal(t, 2, BytePos("hi时间", '间')) // will build error
}

func TestIsStartOf(t *testing.T) {
	tests := []struct {
		give string
		sub  string
		want bool
	}{
		{"abc", "a", true},
		{"abc", "d", false},
	}

	for _, item := range tests {
		assert.Equal(t, item.want, HasPrefix(item.give, item.sub))
		assert.Equal(t, item.want, IsStartOf(item.give, item.sub))
	}
}

func TestIsEndOf(t *testing.T) {
	tests := []struct {
		give string
		sub  string
		want bool
	}{
		{"abc", "c", true},
		{"abc", "d", false},
		{"some.json", ".json", true},
	}

	for _, item := range tests {
		assert.Equal(t, item.want, HasSuffix(item.give, item.sub))
		assert.Equal(t, item.want, IsEndOf(item.give, item.sub))
	}
}

func TestIsSpace(t *testing.T) {
	assert.True(t, IsSpace(' '))
	assert.True(t, IsSpace('\n'))
	assert.True(t, IsSpaceRune('\n'))
	assert.True(t, IsSpaceRune('\t'))

	assert.False(t, IsBlank(" a "))
	assert.True(t, IsBlank(" "))
	assert.True(t, IsBlank("   "))

	assert.False(t, IsBlankBytes([]byte(" a ")))
	assert.True(t, IsBlankBytes([]byte(" ")))
	assert.True(t, IsBlankBytes([]byte("   ")))
}

func TestIsSymbol(t *testing.T) {
	assert.False(t, IsSymbol('a'))
	assert.True(t, IsSymbol('●'))
}
