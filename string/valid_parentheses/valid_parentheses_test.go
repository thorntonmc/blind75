package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsValid(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{"()", true},
		{"()[]{}", true},
		{"(]", false},
	}

	for _, test := range tests {
		got := isValid(test.input)
		assert.Equal(t, test.want, got)
	}
}
