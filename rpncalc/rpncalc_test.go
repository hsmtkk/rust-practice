package rpncalc_test

import (
	"testing"

	"github.com/hsmtkk/rust-practice/rpncalc"
	"github.com/stretchr/testify/assert"
)

func TestCalculate(t *testing.T) {
	expr := "1 1 1 + +"
	want := 3
	got, err := rpncalc.New().Calculate(expr)
	assert.Nil(t, err)
	assert.Equal(t, want, got)
}
