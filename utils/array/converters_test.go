package array

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestToInt64s(t *testing.T) {
	is := assert.New(t)

	ints, err := ToInt64s([]string{"1", "2"})
	is.Nil(err)
	is.Equal("[]int64{1, 2}", fmt.Sprintf("%#v", ints))

	ints = MustToInt64s([]string{"1", "2"})
	is.Equal("[]int64{1, 2}", fmt.Sprintf("%#v", ints))

	ints = MustToInt64s([]interface{}{"1", "2"})
	is.Equal("[]int64{1, 2}", fmt.Sprintf("%#v", ints))

	ints = SliceToInt64s([]interface{}{"1", "2"})
	is.Equal("[]int64{1, 2}", fmt.Sprintf("%#v", ints))

	_, err = ToInt64s([]string{"a", "b"})
	is.Error(err)
}

func TestToStrings(t *testing.T) {
	is := assert.New(t)

	ss, err := ToStrings([]int{1, 2})
	is.Nil(err)
	is.Equal(`[]string{"1", "2"}`, fmt.Sprintf("%#v", ss))

	ss = MustToStrings([]int{1, 2})
	is.Equal(`[]string{"1", "2"}`, fmt.Sprintf("%#v", ss))

	ss = MustToStrings([]interface{}{1, 2})
	is.Equal(`[]string{"1", "2"}`, fmt.Sprintf("%#v", ss))

	ss = SliceToStrings([]interface{}{1, 2})
	is.Equal(`[]string{"1", "2"}`, fmt.Sprintf("%#v", ss))

	_, err = ToStrings("b")
	is.Error(err)
}

func TestStringsToInts(t *testing.T) {
	is := assert.New(t)

	ints, err := StringsToInts([]string{"1", "2"})
	is.Nil(err)
	is.Equal("[]int{1, 2}", fmt.Sprintf("%#v", ints))

	_, err = StringsToInts([]string{"a", "b"})
	is.Error(err)
}
