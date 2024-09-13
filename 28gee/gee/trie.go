package gee

import "strings"

type node struct {
	pattern  string
	part     string
	children []*node
	isWild   bool
}

// 向着前缀树中添加数据
func (n *node) insert(pattern string, parts []string, height int) {

	// 结束添加，已经到叶子结点了
	if len(parts) == height {
		n.pattern = pattern
		return
	}

	// 从当前结点的所有子结点中寻找该结点，如果存在，就直接使用
	// 不存在则需要创建
	part := parts[height]
	item := n.matchChild(part)

	if item == nil {
		// 新建结点，判断是否模糊匹配
		item = &node{
			part:     part,
			isWild:   part[0] == ':' || part[0] == '*',
			children: make([]*node, 0),
		}
		// 当前结点加入父结点的children中
		n.children = append(n.children, item)
	}

	// 向子结点内添加
	item.insert(pattern, parts, height+1)
}

// 匹配当前节点的所有字节点，查询是否有相同的 part.
func (n *node) matchChild(part string) *node {
	for _, child := range n.children {
		if child.part == part || child.isWild {
			return child
		}
	}
	return nil
}

// 递归搜索当前路径，没有匹配则返回nil. 匹配则返回对应匹配的叶子结点
func (n *node) search(parts []string, height int) *node {

	// 结束搜索
	if len(parts) == height || strings.HasPrefix(n.part, "*") {

		// 中间结点
		if n.pattern == "" {
			return nil
		}

		return n
	}

	part := parts[height]

	// 可能匹配到多个结点
	children := n.matchChildren(part)

	for _, item := range children {

		target := item.search(parts, height+1)

		if target == nil {
			continue
		}

		return target
	}

	return nil
}

func (n *node) matchChildren(part string) []*node {

	nodes := make([]*node, 0)

	for _, child := range n.children {
		if child.part == part || child.isWild {
			nodes = append(nodes, child)
		}
	}
	return nodes
}
