package go_demo

import "testing"

func TestHasCycle(t *testing.T) {
	// 创建测试用例
	tests := []struct {
		name     string
		createFn func() *ListNode
		want     bool
	}{
		{
			name: "空链表",
			createFn: func() *ListNode {
				return nil
			},
			want: false,
		},
		{
			name: "单个节点无环",
			createFn: func() *ListNode {
				return &ListNode{Val: 1}
			},
			want: false,
		},
		{
			name: "两个节点无环",
			createFn: func() *ListNode {
				head := &ListNode{Val: 1}
				head.Next = &ListNode{Val: 2}
				return head
			},
			want: false,
		},
		{
			name: "有环链表",
			createFn: func() *ListNode {
				head := &ListNode{Val: 1}
				node2 := &ListNode{Val: 2}
				node3 := &ListNode{Val: 3}
				node4 := &ListNode{Val: 4}
				head.Next = node2
				node2.Next = node3
				node3.Next = node4
				node4.Next = node2 // 形成环
				return head
			},
			want: true,
		},
		{
			name: "自环",
			createFn: func() *ListNode {
				head := &ListNode{Val: 1}
				head.Next = head // 自环
				return head
			},
			want: true,
		},
	}

	// 测试 HasCycle 函数
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := tt.createFn()
			if got := HasCycle(head); got != tt.want {
				t.Errorf("HasCycle() = %v, want %v", got, tt.want)
			}
		})
	}

	// 测试 HasCycleHash 函数
	for _, tt := range tests {
		t.Run(tt.name+"_hash", func(t *testing.T) {
			head := tt.createFn()
			if got := HasCycleHash(head); got != tt.want {
				t.Errorf("HasCycleHash() = %v, want %v", got, tt.want)
			}
		})
	}
}

// 基准测试
func BenchmarkHasCycle(b *testing.B) {
	// 创建一个长链表
	head := &ListNode{Val: 1}
	current := head
	for i := 2; i <= 1000; i++ {
		current.Next = &ListNode{Val: i}
		current = current.Next
	}
	// 创建环
	current.Next = head.Next.Next

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		HasCycle(head)
	}
}

func BenchmarkHasCycleHash(b *testing.B) {
	// 创建一个长链表
	head := &ListNode{Val: 1}
	current := head
	for i := 2; i <= 1000; i++ {
		current.Next = &ListNode{Val: i}
		current = current.Next
	}
	// 创建环
	current.Next = head.Next.Next

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		HasCycleHash(head)
	}
}
