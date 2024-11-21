package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAdd(m *testing.T) {
	expected := 4
	result := Add(2, 2)
	require.Equal(m, expected, result)
}
