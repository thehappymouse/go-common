package utils

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestUssNoGenerate(t *testing.T) {
	last := UssNoGenerate("")
	for i := 0; i < 2220; i++ {
		t.Log(last)
		last = UssNoGenerate(last)
	}

}

func TestBHex2Num(t *testing.T) {
	assert.Equal(t, BHex2Num("1E6K", 36), 65036)
	assert.Equal(t, BHex2Num("FE0C", 16), 65036)

	assert.Equal(t, NumToBHex(65036, 36), "1E6K")
	assert.Equal(t, NumToBHex(65036, 16), "FE0C")

	t.Log(BHex2Num("7H", 16))
}

