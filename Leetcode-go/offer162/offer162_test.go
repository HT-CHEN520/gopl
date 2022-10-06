package offer162

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_constructArr(t *testing.T) {
	require.Equal(t, []int{120, 60, 40, 30, 24}, constructArr([]int{1, 2, 3, 4, 5}))
}
