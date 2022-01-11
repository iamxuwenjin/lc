package code

func containsDuplicate(nums []int) bool {
	flagMap := map[int]struct{}{}
	for _, v := range nums {
		_, ok := flagMap[v]
		if ok {
			return true
		}
		flagMap[v] = struct{}{}
	}
	return false
}
