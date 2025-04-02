package swaggokratos

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMustGetPortNum(t *testing.T) {
	portNum := MustGetPortNum("0.0.0.0:8000")
	require.Equal(t, "8000", portNum)
}
