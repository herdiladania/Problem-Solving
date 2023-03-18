package problemsolving

func pickingNumbers(a []int32) int32 {
	var result int32 = 0

	for i := 0; i < len(a); i++ {
		var subArray []int32 = []int32{}
		for j := 0; j < len(a); j++ {
			if a[i]-a[j] == 1 || a[i]-a[j] == 0 {
				subArray = append(subArray, a[j])
			}
		}
		if int32(len(subArray)) > result {
			result = int32(len(subArray))
		}
	}

	return result
}
