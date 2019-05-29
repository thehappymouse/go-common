package utils

import (
	"errors"
	"testing"
)

func TestCheckError(t *testing.T) {
	CheckError(errors.New("testError"))
}
