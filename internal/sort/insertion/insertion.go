package insertion

func sort(arr []int) {
    for i := 0; i < len(arr); i++ {
        j := i
        for j > 0 {
            if arr[j-1] > arr[j] {
                arr[j-1], arr[j] = arr[j], arr[j-1]
            }
            j = j - 1
        }
    }
}
