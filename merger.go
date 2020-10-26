package json2struct

import (
	"fmt"
	"os"
)

// n1, n2 node type must be SliceNode
func mergeSliceNode(n1, n2 *SliceNode) Node {
	var node Node
	var newVal []interface{}
	jsonName := n1.GetBaseNode().JsonName
	if n1.JsonName != n2.JsonName {
		fmt.Fprint(os.Stderr, "mergeSliceNode n1, n2 node jsonname is not same <" + n1.JsonName + "," + n2.JsonName + ">")
		jsonName = ""
	}
	// 1: merge two slice to one slice
	newVal = make([]interface{}, 0, len(n1.Val.([]interface{})) + len(n2.Val.([]interface{})))
	for _, v := range n1.Val.([]interface{}) {
		newVal = append(newVal, v)
	}
	for _, v := range n2.Val.([]interface{}) {
		newVal = append(newVal, v)
	}
	node = NewNode(jsonName, newVal, n1.GenType, SliceType, n1.print)
	
	// 2: merge all slice element to a full element
	node.fillSubNode()
	return fullNode(node)
}

func mergeMapNode(n1, n2 *MapNode) *MapNode {
	var node Node
	jsonName := n1.GetBaseNode().JsonName
	if n1.JsonName != n2.JsonName {
		fmt.Fprint(os.Stderr, "mergeMapNode n1, n2 node jsonname is not same <" + n1.JsonName + "," + n2.JsonName + ">")
		jsonName = ""
	}
	// it can't merge map/struct directly, in case of having the save map key but not save value type
	node = NewNode(jsonName, nil, n1.GetBaseNode().GenType, MapType, n1.GetBaseNode().print)
	v0 := n1.SubNode()[0]
	for _, v1 := range n1.SubNode()[1:] {
		v0 = mergeNode(v0, v1)
	}
	for _, v1 := range n2.SubNode() {
		v0 = mergeNode(v0, v1)
	}
	node.GetBaseNode().Val = map[string]interface{}{v0.GetBaseNode().JsonName: v0.GetBaseNode().Val}
	node.(*MapNode).KeyType = StringType
	node.(*MapNode).ValType = v0.GetBaseNode().Type
	node.(*MapNode).ValNode = []Node{v0}
	return node.(*MapNode)
}

func mergeMapObjectNode(n1, n2 Node) Node {
	var node Node
	jsonName := n1.GetBaseNode().JsonName
	if n1.GetBaseNode().JsonName != n2.GetBaseNode().JsonName {
		fmt.Fprint(os.Stderr, "mergeMapObjectNode n1, n2 node jsonname is not same <" + n1.GetBaseNode().JsonName + "," + n2.GetBaseNode().JsonName + ">")
		jsonName = ""
	}
	// it can't merge map/struct directly, in case of having the save map key but not save value type
	node = NewNode(jsonName, nil, n1.GetBaseNode().GenType, MapType, n1.GetBaseNode().print)
	v0 := n1.SubNode()[0]
	for _, v1 := range n1.SubNode()[1:] {
		v0 = mergeNode(v0, v1)
	}
	for _, v1 := range n2.SubNode() {
		v0 = mergeNode(v0, v1)
	}
	node.GetBaseNode().Val = map[string]interface{}{v0.GetBaseNode().JsonName: v0.GetBaseNode().Val}
	node.(*MapNode).KeyType = StringType
	node.(*MapNode).ValType = v0.GetBaseNode().Type
	node.(*MapNode).ValNode = []Node{v0}
	return node
}

func mergeObjectNode(n1, n2 *ObjectNode) *ObjectNode {
	jsonName := n1.GetBaseNode().JsonName
	if n1.JsonName != n2.JsonName {
		fmt.Fprint(os.Stderr, "mergeSliceNode n1, n2 node jsonname is not same <" + n1.JsonName + "," + n2.JsonName + ">")
		jsonName = ""
	}
	n1Map := make(map[string]Node, len(n1.SubNode()))
	nodeMap := make(map[string]Node, len(n2.SubNode()))
	node := NewNode(jsonName, nil, n1.GenType, StructType, n1.print)
	for _, v := range n1.SubNode() {
		n1Map[v.GetBaseNode().JsonName] = v
	}
	for _, v := range n2.SubNode() {
		tagName := v.GetBaseNode().JsonName
		nodeMap[tagName] = v
		if n1Map[tagName] != nil {
			// merge node, tagName in n1 and n2
			nodeMap[tagName] = mergeNode(n1Map[tagName], v)
			delete(n1Map, tagName)
			continue
		}
		// tagName in n2, but not in n1
		v.GetBaseNode().Optional = true
	}
	for k, v := range n1Map {
		// tagName in n1, but not in n2
		v.GetBaseNode().Optional = true
		nodeMap[k] = v
	}
	
	subNode := make([]Node, 0, len(nodeMap))
	newVal := make(map[string]interface{}, len(nodeMap))
	for k, v := range nodeMap {
		newVal[k] = v.GetBaseNode().Val
		subNode = append(subNode, v)
	}
	node.GetBaseNode().Val = newVal
	node.(*ObjectNode).KeyType = KeywordStringType
	node.(*ObjectNode).ValNode = subNode
	return node.(*ObjectNode)
}

func fullNode(node Node) Node {
	if node.isMerged() {
		return node
	}
	if node.IsSimple() {
		node.GetBaseNode().Merged = true
		return node
	}
	switch node.(type) {
	case *BaseNode:
		return node
	case *SliceNode:
		if isSimpleType(node.(*SliceNode).ValType) {
			return node
		}
		v0 := node.SubNode()[0]
		for _, v1 := range node.SubNode()[1:] {
			v0 = mergeNode(v0, v1)
		}
		v0.GetBaseNode().Merged = true
		node.GetBaseNode().Val = []interface{}{v0.GetBaseNode().Val}
		node.GetBaseNode().Merged = true
		node.(*SliceNode).ValType = v0.GetBaseNode().Type
		node.(*SliceNode).ValNode = []Node{v0}
		return node
	case *MapNode:
		if isSimpleType(node.(*MapNode).ValType) {
			return node
		}
		v0 := node.SubNode()[0]
		for _, v1 := range node.SubNode()[1:] {
			v0 = mergeNode(v0, v1)
		}
		v0.GetBaseNode().Merged = true
		node.GetBaseNode().Val = map[string]interface{}{v0.GetBaseNode().JsonName: v0.GetBaseNode().Val}
		node.GetBaseNode().Merged = true
		node.(*MapNode).ValType = v0.GetBaseNode().Type
		node.(*MapNode).ValNode = []Node{v0}
		
	case *ObjectNode:
		valNode := node.(*ObjectNode).ValNode
		for i, sn := range valNode {
			if !sn.IsSimple() {
				valNode[i] = fullNode(sn)
			}
		}
	}
	return node
}

func mergeNode(n1, n2 Node) Node {
	base1, base2 := n1.GetBaseNode(), n2.GetBaseNode()
	jsonName := base1.JsonName
	if base1.JsonName != base2.JsonName {
		fmt.Fprint(os.Stderr, "mergeNode n1, n2 node jsonname is not same <" + base1.JsonName + "," + base2.JsonName + ">")
		jsonName = ""
	}
	if base1.Type == base2.Type {
		if isSimpleType(base1.Type) {
			return n1
		}
		switch base1.Type {
		case SliceType:
			return mergeSliceNode(n1.(*SliceNode), n2.(*SliceNode))
		case MapType:
			return mergeMapNode(n1.(*MapNode), n2.(*MapNode))
		case StructType:
			return mergeObjectNode(n1.(*ObjectNode), n2.(*ObjectNode))
		}
		return n1
	}
	typeMap := make(map[JsonType]bool, 2)
	typeMap[base1.Type] = true
	typeMap[base2.Type] = true
	if typeMap[StructType] && typeMap[MapType] {
		return mergeMapObjectNode(n1, n2)
	}
	jt := mergeValueType(typeMap)
	if jt == InterfaceType {
		return NewNode(jsonName, base1.Val, base1.GenType, InterfaceType, base1.print)
	}
	return NewNode(jsonName, base1.Val, base1.GenType, jt, base1.print)
}

