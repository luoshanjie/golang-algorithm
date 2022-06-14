package bubble

import (
	"testing"
)

func TestSort(t *testing.T) {
	arr := []int{1, 3, 2, 7, 5, 8, 10, 9, 6, 4}
	sort(arr)
	for i := 1; i <= 10; i++ {
		get := arr[i-1]
		want := i
		if get != want {
			t.Errorf("get %d but want %d\n", get, want)
		}
	}
}
