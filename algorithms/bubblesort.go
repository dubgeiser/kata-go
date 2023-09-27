package algorithms

func BubbleSort(arr []int) {
	l := len(arr)
	for l > 0 {
		l--
		for i := 0; i < l; i++ {
			if arr[i] > arr[i+1] {
				arr[i+1], arr[i] = arr[i], arr[i+1]
			}
		}
	}
}
