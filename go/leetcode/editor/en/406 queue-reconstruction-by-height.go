package main

import (
	"container/list"
	"sort"
)

//You are given an array of people, people, which are the attributes of some
//people in a queue (not necessarily in order). Each people[i] = [hi, ki] represents
//the iᵗʰ person of height hi with exactly ki other people in front who have a
//height greater than or equal to hi.
//
// Reconstruct and return the queue that is represented by the input array
//people. The returned queue should be formatted as an array queue, where queue[j] = [
//hj, kj] is the attributes of the jᵗʰ person in the queue (queue[0] is the person
//at the front of the queue).
//
//
// Example 1:
//
//
//Input: people = [[7,0],[4,4],[7,1],[5,0],[6,1],[5,2]]
//Output: [[5,0],[7,0],[5,2],[6,1],[4,4],[7,1]]
//Explanation:
//Person 0 has height 5 with no other people taller or the same height in front.
//
//Person 1 has height 7 with no other people taller or the same height in front.
//
//Person 2 has height 5 with two persons taller or the same height in front,
//which is person 0 and 1.
//Person 3 has height 6 with one person taller or the same height in front,
//which is person 1.
//Person 4 has height 4 with four people taller or the same height in front,
//which are people 0, 1, 2, and 3.
//Person 5 has height 7 with one person taller or the same height in front,
//which is person 1.
//Hence [[5,0],[7,0],[5,2],[6,1],[4,4],[7,1]] is the reconstructed queue.
//
//
// Example 2:
//
//
//Input: people = [[6,0],[5,0],[4,0],[3,2],[2,2],[1,4]]
//Output: [[4,0],[5,0],[2,2],[3,2],[1,4],[6,0]]
//
//
//
// Constraints:
//
//
// 1 <= people.length <= 2000
// 0 <= hi <= 10⁶
// 0 <= ki < people.length
// It is guaranteed that the queue can be reconstructed.
//
//
// Related Topics Array Greedy Binary Indexed Tree Segment Tree Sorting 👍 6546
//👎 656

//leetcode submit region begin(Prohibit modification and deletion)
// 貪心法：分成身高以及順序兩個維度分別判斷
// 排序完的people： [[7,0], [7,1], [6,1], [5,0], [5,2]，[4,4]]
// 插入的过程：
// 插入[7,0]：[[7,0]]
// 插入[7,1]：[[7,0],[7,1]]
// 插入[6,1]：[[7,0],[6,1],[7,1]]
// 插入[5,0]：[[5,0],[7,0],[6,1],[7,1]]
// 插入[5,2]：[[5,0],[7,0],[5,2],[6,1],[7,1]]
// 插入[4,4]：[[5,0],[7,0],[5,2],[6,1],[4,4],[7,1]]

// 解法二：鏈表實現
func reconstructQueue(people [][]int) [][]int {
	// 排序：高度 then 順序
	sort.Slice(people, func(i, j int) bool {
		return people[i][0] > people[j][0] || (people[i][0] == people[j][0] && people[i][1] < people[j][1])
	})
	res := make([][]int, 0)

	queue := list.New()
	for i := 0; i < len(people); i++ {
		position := people[i][1]
		mark := queue.PushBack(people[i]) //插入元素
		node := queue.Front()
		for position != 0 { //获取相对位置
			position--
			node = node.Next()
		}
		queue.MoveBefore(mark, node) //移动位置
	}

	for node := queue.Front(); node != nil; node = node.Next() {
		res = append(res, node.Value.([]int))
	}

	return res
}

// 解法一：數組實現（數組對於插入效能較差且麻煩）
func reconstructQueue1(people [][]int) [][]int {
	// 排序：高度 then 順序
	sort.Slice(people, func(i, j int) bool {
		return people[i][0] > people[j][0] || (people[i][0] == people[j][0] && people[i][1] < people[j][1])
	})

	// 再按照K进行插入排序，优先插入K小的
	for i, p := range people {
		copy(people[p[1]+1:i+1], people[p[1]:i+1]) // 空出一个位置
		people[p[1]] = p
	}
	return people
}

//leetcode submit region end(Prohibit modification and deletion)
