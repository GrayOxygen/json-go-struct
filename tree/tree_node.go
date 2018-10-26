package tree


import "github.com/GrayOxygen/json-go-struct/model"


type TreeNode struct {
	Value    *model.StructObj
	Children []*TreeNode
	Level    int
	MaxLevel int
}

//仅获取所有直接子节点
func GetSonObjs(root *TreeNode) []*model.StructObj {
	nodes := GetChildren(root)
	result := make([]*model.StructObj, 0)
	for _, item := range nodes {
		if item.Level == root.Level+1 {
			result = append(result, item.Value)
		}
	}
	return result
}

//获取所有直接子节点的孩子节点
func GetSonChildrenObjs(root *TreeNode) []*model.StructObj {
	nodes := GetChildren(root)
	result := make([]*model.StructObj, 0)
	for _, item := range nodes {
		if item.Level > root.Level+1 {
			result = append(result, item.Value)
		}
	}
	return result
}
func GetSonChildren(root *TreeNode) []*TreeNode {
	nodes := GetChildren(root)
	result := make([]*TreeNode, 0)
	for _, item := range nodes {
		if item.Level > root.Level+1 {
			result = append(result, item)
		}
	}
	return result
}
func GetSonNodes(root *TreeNode) []*TreeNode {
	nodes := GetChildren(root)
	result := make([]*TreeNode, 0)
	for _, item := range nodes {
		if item.Level == root.Level+1 {
			result = append(result, item)
		}
	}
	return result
}

//获取所有孩子节点
func GetChildren(root *TreeNode) []*TreeNode {
	if root == nil {
		return nil
	}
	result := make([]*TreeNode, 0)
	result = append(result, root)
	for _, item := range root.Children {
		temp := GetChildren(item)
		if temp != nil {
			result = append(result, temp...)
		}
	}
	return result
}
func GetFatherNode(root *TreeNode, curNode *TreeNode) *TreeNode {
	nodes := GetSonNodes(root)
	for _, item := range nodes {
		if item.Value.Id == curNode.Value.Id {
			return root
		}
	}
	res := new(TreeNode)
	for _, item := range nodes {
		res = GetFatherNode(item, curNode)
	}
	return res
}
