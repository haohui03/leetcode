package array

func intersect(nums1 []int, nums2 []int) []int {
	if len(nums1) > len(nums2) {
		return intersect(nums2, nums1)
	}
	nums2Map := make(map[int]int)
	outPutMap := make(map[int]int)
	for _, i := range nums2 {
		nums2Map[i] += 1
	}
	for _, i := range nums1 {
		if count, ok := nums2Map[i]; ok {
			if outPutMap[i]+1 < count {
				outPutMap[i] += 1
			} else {
				outPutMap[i] = count
			}
		}
	}
	index := 0
	for value, count := range outPutMap {
		for i := 0; i < count; i++ {
			nums1[index+i] = value
		}
		index += count
	}
	return nums1[:index]
}
