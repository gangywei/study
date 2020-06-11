package gee

import "strings"

type node struct {
	pattern string
	part string
	children []*node
	isWild bool	//是否精确匹配
}

//查询是否有匹配节点
func (n *node) matchChild(part string) *node {
	for _, child := range n.children{
		if child.part == part || child.isWild {
			return child
		}
	}
	
	return nil
}

//结点下所有匹配成功的节点，用于查询匹配成功的结点
func (n *node) matchChildren(part string) []*node {
	nodes := make([]*node, 0) //初始化数组切片
	for _, child := range n.children {
		if child.part == part || child.isWild {
			nodes = append(nodes, child)
		}
	}
	return nodes
}

//递归查找每一层的节点，如果没有匹配到当前part的节点
func (n *node) insert(pattern string, parts []string, height int) {
	if len(parts) == height {
		n.pattern = pattern
		return
	}

	part := parts[height]
	child := n.matchChild(part)
	if child == nil {
		child = &node{part: part, isWild: part[0] == ':' || part[0] == '*'}
		n.children = append(n.children, child)
	}
	child.insert(pattern, parts, height+1)
}

func (n *node) search(parts []string, height int) *node {
	if len(parts) == height || strings.HasPrefix(n.part, "*") {
		if n.pattern == "" {
			return nil
		}
		return n
	}

	part := parts[height]
	children := n.matchChildren(part)

	for _, child := range children {
		result := child.search(parts, height+1)
		if result != nil {
			return result
		}
	}

	return nil
}

