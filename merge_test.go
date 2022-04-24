package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHelloWorld(t *testing.T) {
	a := ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 6,
			},
		},
	}

	b := ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val: 4,
			},
		},
	}

	r := mergeTwoLists(&a, &b)

	rA := listToArray(r)
	expected := []int{1, 1, 3, 4, 4, 6}
	require.Equal(t, rA, expected)
}

func TestListToArray(t *testing.T) {
	a := ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 6,
			},
		},
	}

	aA := listToArray(&a)
	require.Equal(t, aA, []int{1, 4, 6})
}
