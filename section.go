package section

import "math/rand"

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

// Each 遍历 slice，并对每个元素执行 fn
func Each[T any](slice []T, fn func(e T)) {
	for _, e := range slice {
		fn(e)
	}
}

// Filter 遍历 slice，并对每个元素执行 fn，返回符合 fn 的元素
func Filter[T any](slice []T, fn func(e T) bool) []T {
	var result []T
	for _, e := range slice {
		if fn(e) {
			result = append(result, e)
		}
	}
	return result
}

// First 返回 slice 中第一个元素
func First[T any](slice []T) T { return slice[0] }

// IsEmpty 判断 slice 是否为空
func IsEmpty[T any](slice []T) bool { return len(slice) == 0 }

// Pop 删除 slice 中的最后一个元素，并返回该元素
func Pop[T any](slice []T) (e T) {
	e = slice[len(slice)-1]
	slice = slice[:len(slice)-1]
	return
}

// Push 在 slice 的末尾添加元素
func Push[T any](slice []T, e T) []T {
	return append(slice, e)
}

//Shuffle 打乱 slice
func Shuffle[T any](slice []T) []T {
	for i := len(slice) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		slice[i], slice[j] = slice[j], slice[i]
	}
	return slice
}

// Sort 对 slice 排序
func Sort[T any](slice []T, less func(i, j T) bool) []T {
	for i := len(slice) - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if less(slice[j], slice[j+1]) {
				slice[j], slice[j+1] = slice[j+1], slice[j]
			}
		}
	}
	return slice
}

// MinOrMax 返回 slice 中最小或者最大的元素
func MinOrMax[T any](slice []T, less func(i, j T) bool) T {
	var min = slice[0]
	for _, e := range slice {
		if less(e, min) {
			min = e
		}
	}
	return min
}
