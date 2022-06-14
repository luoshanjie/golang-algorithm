package bubble

func sort(arr []int) {
	for j := len(arr) - 1; j >= 1; j-- {
		for i := 1; i <= j; i++ {
			if arr[i-1] > arr[i] {
				arr[i-1], arr[i] = arr[i], arr[i-1]
			}
		}
	}
}
