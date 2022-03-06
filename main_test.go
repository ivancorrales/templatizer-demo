package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_greet(t *testing.T) {
	got := greet()

	assert.Equal(t, "{{.DefaultMessage}}", got, "should properly greet")
}
