package section

// Index 从 slice 查找符合 eq 的第一个元素并 返回在数组中的元素
func Index[T any](slice []T, eq func(e T) bool) (index int) {
	for i, e := range slice {
		if eq(e) {
			return i
		}
	}
	return -1
}

// Exists 判断 slice 中是否存在符合 eq 的元素
func Exists[T any](slice []T, eq func(e T) bool) bool {
	return Index(slice, eq) != -1
}

// Unique 去除 slice 中重复的元素
func Unique[T any](slice []T, eq func(i, j T) bool) []T {
	if len(slice) < 2 {
		return slice
	}
	var result []T
	for i := 0; i < len(slice); i++ {
		if !Exists(result, func(e T) bool {
			return eq(slice[i], e)
		}) {
			result = append(result, slice[i])
		}
	}
	return result
}

// Count 统计 slice 中符合 eq 的元素的个数
func Count[T any](slice []T, eq func(e T) bool) int {
	var count int
	for _, e := range slice {
		if eq(e) {
			count++
		}
	}
	return count
}

// Cut 去除 slice 中符合 eq 的元素
func Cut[T any](slice []T, eq func(e T) bool) []T {
	var result []T
	for _, e := range slice {
		if !eq(e) {
			result = append(result, e)
		}
	}
	return result
}

// Delete 删除 slice 中符合 eq 的元素
func Delete[T any](slice []T, eq func(e T) bool) []T {
	var result []T
	for _, e := range slice {
		if !eq(e) {
			result = append(result, e)
		}
	}
	return result
}
