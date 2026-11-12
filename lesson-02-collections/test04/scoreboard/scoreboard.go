package scoreboard

import "sort"

// Totals 把每位学生的多次成绩累加到一张新 map。
// records 本身不能被修改；某个学生没有成绩时，总分为 0。
func Totals(records map[string][]int) map[string]int {
	// TODO: 创建结果 map，遍历 records，把每条成绩加起来。
	// 不要直接改 records 里的切片。
	totals := make(map[string]int, len(records))
	for i, grades := range records {
		totals[i] = 0
		for _, grade := range grades {
			totals[i] += grade
		}
	}
	return totals
}

// Entry 是排行榜上的一行。
type Entry struct {
	Name  string
	Score int
}

// Top 返回分数最高的前 n 名。
// 分数相同时按姓名升序；n <= 0 时返回空结果；n 大于人数时返回全部。
func Top(totals map[string]int, n int) []Entry {
	// TODO: 把 map 转成切片，按「分数降序、姓名升序」排序，再截取前 n 个。
	if n <= 0 {
		return nil
	}
	entry := make([]Entry, 0, len(totals))
	for name, score := range totals {
		entry = append(entry, Entry{Name: name, Score: score})
	}
	sort.Slice(entry, func(i, j int) bool {
		if entry[i].Score != entry[j].Score {
			return entry[i].Score > entry[j].Score
		}
		return entry[i].Name < entry[j].Name
	})
	/*slices.SortFunc(students, func(a, b Student) int {
	    // 返回负数表示 a 排前面，正数表示 b 排前面，0 表示相等
	    return cmp.Or(
	        cmp.Compare(b.Score, a.Score), // 分数降序
	        cmp.Compare(a.Name, b.Name),   // 姓名升序
	    )
	})*/
	if n > len(entry) {
		return entry
	}
	return entry[:n]
}

/*sort.Slice 是 Go 语言中最灵活、最常用的排序函数。它能对任意类型的切片（结构体切片、整数切片、字符串切片）进行排序，而且完全原地修改。
 sort.Slice 的定义极其简单，只有两个参数：
func Slice(x any, less func(i, j int) bool)
x any：你要排序的切片。any 代表它可以接收任何类型的切片。
less func(i, j int) bool：这是最核心的“裁判规则”。
它接收两个索引 i 和 j。
它必须返回一个布尔值 true 或 false。
核心口诀：如果索引为 i 的元素，应该排在索引为 j 的元素前面，你就返回 true；否则返回 false。*/
