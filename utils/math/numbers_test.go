package math

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestPercent(t *testing.T) {
	assert.Equal(t, float64(34), Percent(34, 100))
	assert.Equal(t, float64(0), Percent(34, 0))
	assert.Equal(t, float64(-100), Percent(34, -34))
}

func TestElapsedTime(t *testing.T) {
	nt := time.Now().Add(-time.Second * 3)
	num := ElapsedTime(nt)

	assert.Equal(t, 3000, int(MustFloat(num)))
}

func TestDataSize(t *testing.T) {
	assert.Equal(t, "3.38K", DataSize(3456))
}

func TestHowLongAgo(t *testing.T) {
	assert.Equal(t, "57 mins", HowLongAgo(3456))
}
