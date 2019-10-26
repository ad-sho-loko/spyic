package spyic

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_BubbleSort(t *testing.T) {
	array := []int{2, 1, 4, 0, 5}
	buf := &bytes.Buffer{}

	ob := NewSlice(array)
	ob.SetWriter(buf)
	ob.SetDuration(time.Nanosecond * 1)

	ob.Start()
	for i := 0; i < len(array)-1; i++ {
		for j := len(array) - 1; j > i; j-- {
			ob.Step()
			if array[j-1] < array[j] {
				array[j], array[j-1] = array[j-1], array[j]
			}
		}
	}
	ob.Finish()
	assert.Equal(t, array, []int{5, 4, 2, 1, 0})
}
