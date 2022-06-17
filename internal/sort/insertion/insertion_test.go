package insertion

import "testing"

func TestSort(t *testing.T) {
    items := []int{9, 7, 6, 4, 5, 8, 2, 1, 3, 10}
    sort(items)
    for i := 0; i < len(items); i++ {
        get := items[i]
        want := i + 1
        if get != want {
            t.Errorf("get %d but want %d\n", get, want)
        }
    }
}
