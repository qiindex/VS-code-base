package go_demo

// ListNode 定义链表节点
type ListNode struct {
	Val  int
	Next *ListNode
}

// HasCycle 使用快慢指针判断链表是否有环
// 时间复杂度：O(n)
// 空间复杂度：O(1)
func HasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	// 初始化快慢指针
	slow, fast := head, head.Next

	// 快指针每次移动两步，慢指针每次移动一步
	for fast != nil && fast.Next != nil {
		if slow == fast {
			return true // 快慢指针相遇，说明有环
		}
		slow = slow.Next
		fast = fast.Next.Next
	}

	return false // 快指针到达链表末尾，说明无环
}

// HasCycleHash 使用哈希表判断链表是否有环
// 时间复杂度：O(n)
// 空间复杂度：O(n)
func HasCycleHash(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	// 使用map记录已访问的节点
	visited := make(map[*ListNode]bool)
	current := head

	for current != nil {
		if visited[current] {
			return true // 节点已访问过，说明有环
		}
		visited[current] = true
		current = current.Next
	}

	return false // 遍历完整个链表，说明无环
}
