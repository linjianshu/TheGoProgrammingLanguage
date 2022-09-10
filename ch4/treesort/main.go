package main

import "fmt"

func main() {
	//用树来排序  递归
	ints := []int{2, 5, 3, 8, 6}
	var root *treeNode
	for _, v := range ints {
		root = addNode(v, root)
	}
	tree := popTree(ints[:0], root)
	fmt.Println(tree)
}

type treeNode struct {
	val   int
	left  *treeNode
	right *treeNode
}

func addNode(val int, root *treeNode) *treeNode {
	if root == nil {
		root = new(treeNode)
		root.val = val
		return root
	}
	if root.val > val {
		root.left = addNode(val, root.left)
	} else if root.val < val {
		root.right = addNode(val, root.right)
	}
	return root
}

func popTree(res []int, root *treeNode) []int {
	if root == nil {
		return res
	}
	res = popTree(res, root.left)
	res = append(res, root.val)
	res = popTree(res, root.right)
	return res
}
