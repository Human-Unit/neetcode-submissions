func twoSum(nums []int, target int) []int {

	n:= make(map[int]int);

	for i, f := range nums {
		t := target - f;
		if ind, t := n[t]; t{
			return []int{ind, i}
		}
		n[f] = i;
	}

	return nil;
    
}

