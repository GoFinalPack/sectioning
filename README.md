# sectioning
[![Build Status](https://github.com/issue9/sliceutil/workflows/Go/badge.svg)](https://github.com/GoFinalPack/sectioning/actions?query=workflow%3AGo)
[![license](https://img.shields.io/badge/license-MIT-brightgreen.svg?style=flat)](https://opensource.org/licenses/MIT)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/GoFinalPack/sectioning)](https://pkg.go.dev/github.com/GoFinalPack/sectioning)
[![Go version](https://img.shields.io/github/go-mod/go-version/GoFinalPack/sectioning)](https://golang.org)

----

sectioning 提供了针对数组和切片的功能

- `Index` 查找符合条件元素在数组中的位置
- `Delete` 删除符合条件的切片元素
- `Unique` 提取数组中的唯一元素
- `Count` 统计数组或切片中包含指定什的数量
- `Cut` Cut 去除 slice 中符合 eq 的元素
- `Each` 遍历 slice，并对每个元素执行 fn
- `Filter` 遍历 slice，并对每个元素执行 fn，返回符合 fn 的元素
- `First` 返回 slice 中第一个元素
- `IsEmpty` 判断 slice 是否为空
- `Pop` 删除 slice 中的最后一个元素，并返回该元素
- `Push` 在 slice 的末尾添加元素
- `Shuffle` 打乱 slice
- `Sort` 排序 slice
- `MinOrMax` 返回 slice 中最小或者最大的元素

GO 从 0.8.0 开始采用泛型技术，仅支持 go1.18 以上版本

```go

func TestShuffle(t *testing.T) {
	intSlice := []int{1, 2, 3, 7, 0, 4, 7}
	b := Shuffle(intSlice)
	t.Log(b)

	c := Shuffle(objSlice)
	t.Log(c)
}

func TestSort(t *testing.T) {
	intSlice := []int{1, 2, 3, 7, 0, 4, 7}
	b := Sort(intSlice, func(i, j int) bool {
		return i > j
	})
	t.Log(b)

	c := Sort(objSlice, func(i, j *obj) bool {
		return i.ID > j.ID
	})
	t.Log(c)
}

```
