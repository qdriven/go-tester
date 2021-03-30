package str

import (
	"fmt"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestTrim(t *testing.T) {
	is := assert.New(t)

	// Trim
	tests := map[string]string{
		"abc ":  "",
		" abc":  "",
		" abc ": "",
		"abc,,": ",",
		"abc,.": ",.",
	}
	for sample, cutSet := range tests {
		is.Equal("abc", Trim(sample, cutSet))
	}

	is.Equal("abc", Trim("abc,.", ".,"))
	// is.Equal("", Trim(nil))

	// TrimLeft
	is.Equal("abc ", Ltrim(" abc "))
	is.Equal("abc ,", TrimLeft(", abc ,", " ,"))
	is.Equal("abc ,", TrimLeft(", abc ,", ", "))

	// TrimRight
	is.Equal(" abc", Rtrim(" abc "))
	is.Equal(", abc", TrimRight(", abc ,", ", "))
}

func TestURLEnDecode(t *testing.T) {
	is := assert.New(t)

	is.Equal("a.com/?name%3D%E4%BD%A0%E5%A5%BD", URLEncode("a.com/?name=你好"))
	is.Equal("a.com/?name=你好", URLDecode("a.com/?name%3D%E4%BD%A0%E5%A5%BD"))
	is.Equal("a.com", URLEncode("a.com"))
	is.Equal("a.com", URLDecode("a.com"))
}

func TestFilterEmail(t *testing.T) {
	is := assert.New(t)
	is.Equal("THE@inhere.com", FilterEmail("   THE@INHere.com  "))
	is.Equal("inhere.xyz", FilterEmail("   inhere.xyz  "))
}

func TestSimilarity(t *testing.T) {
	is := assert.New(t)
	_, ok := Similarity("hello", "he", 0.3)
	is.True(ok)
}

func TestSplit(t *testing.T) {
	ss := Split("a, , b,c", ",")
	assert.Equal(t, `[]string{"a", "b", "c"}`, fmt.Sprintf("%#v", ss))

	ss = Split(" ", ",")
	assert.Nil(t, ss)
}

func TestSubstr(t *testing.T) {
	assert.Equal(t, "abc", Substr("abcDef", 0, 3))
	assert.Equal(t, "cD", Substr("abcDef", 2, 2))
	assert.Equal(t, "", Substr("abcDEF", 23, 5))
}

func TestRepeat(t *testing.T) {
	assert.Equal(t, "aaa", Repeat("a", 3))
	assert.Equal(t, "D", Repeat("D", 1))
	assert.Equal(t, "D", Repeat("D", 0))
	assert.Equal(t, "D", Repeat("D", -3))
}

func TestRepeatRune(t *testing.T) {
	tests := []struct {
		want  []rune
		give  rune
		times int
	}{
		{[]rune("bbb"), 'b', 3},
		{[]rune("..."), '.', 3},
		{[]rune("  "), ' ', 2},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.want, RepeatRune(tt.give, tt.times))
	}
}

func TestReplaces(t *testing.T) {
	assert.Equal(t, "tom age is 20", Replaces(
		"{name} age is {age}",
		map[string]string{
			"{name}": "tom",
			"{age}":  "20",
		}))
}

func TestPadding(t *testing.T) {
	tests := []struct {
		want, give, pad string
		len             int
		pos             uint8
	}{
		{"ab000", "ab", "0", 5, PosRight},
		{"000ab", "ab", "0", 5, PosLeft},
		{"ab012", "ab012", "0", 4, PosLeft},
		{"ab   ", "ab", "", 5, PosRight},
		{"   ab", "ab", "", 5, PosLeft},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.want, Padding(tt.give, tt.pad, tt.len, tt.pos))

		if tt.pos == PosRight {
			assert.Equal(t, tt.want, PadRight(tt.give, tt.pad, tt.len))
		} else {
			assert.Equal(t, tt.want, PadLeft(tt.give, tt.pad, tt.len))
		}
	}
}

func TestPrettyJSON(t *testing.T) {
	tests := []interface{}{
		map[string]int{"a": 1},
		struct {
			A int `json:"a"`
		}{1},
	}
	want := `{
    "a": 1
}`
	for _, sample := range tests {
		got, err := PrettyJSON(sample)
		assert.NoError(t, err)
		assert.Equal(t, want, got)
	}
}
