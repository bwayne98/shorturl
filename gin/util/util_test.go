package util

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRandString(t *testing.T) {

	length := 10

	res := RandString(length)
	require.Equal(t, length, len(res))

	pass, err := regexp.MatchString("^[\\s\\w]+$", res)

	require.NoError(t, err)
	require.True(t, pass)
}

func TestRandUrl(t *testing.T) {
	url := RandUrl()

	pass, err := regexp.MatchString("^https?://", url)

	require.NoError(t, err)
	require.True(t, pass)
}

func TestRandInt(t *testing.T) {
	min, max := int64(100), int64(1000)

	for i := 0; i < 10; i++ {
		res := RandInt(min, max)

		require.True(t, min <= res)
		require.True(t, max > res)
	}
}
