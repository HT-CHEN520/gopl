package BinarySearch

// 基本二分查找
func BinarySearch(a []int, v int) int {
	n := len(a)
	if n == 0 {
		return -1
	}

	lo := 0
	hi := n - 1

	for lo <= hi {
		mid := (lo + hi) / 2
		if a[mid] == v {
			return mid
		} else if a[mid] < v {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}

	return -1
}
