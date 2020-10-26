package json2struct

import "fmt"

type MapNode struct {
	*BaseNode
	KeyType JsonType // type of map key: must be in string/float/int/bool
	ValType JsonType // type of map value
	ValNode []Node // sub node list of map value
}


func (p *MapNode) JsonToStruct(str string) (ret string, err error) {
	defer func() {
		err = p.recover()
	}()
	p.parse(str)
	p.fillSubNode()
	p.dump()
	ret = p.print.Output()
	return
}

func (p *MapNode) fillSubNode() {
	if p.Val == nil {
		return
	}
	v, ok := p.Val.(map[string]interface{})
	if !ok {
		panic("MapNode is not map[string]interface{}, impossible")
	}
	if len(v) == 0 {
		return
	}
	_, _, jt := parseMapType(v)
	p.KeyType = jt
	typeMap := make(map[JsonType]bool, 2)
	p.ValNode = make([]Node, 0, len(v))
	for _, sv := range v {
		jt = parseValueType(sv)
		typeMap[jt] = true
		p.ValNode = append(p.ValNode, NewNode("", sv, p.GenType, jt, p.print))
	}
	p.ValType = mergeValueType(typeMap)
	
	for _, node := range p.ValNode {
		node.fillSubNode()
	}
	return
}

func (p *MapNode) dump() {
	// don't change p
	node := fullNode(p.copy(p))
	fNode := node.(*MapNode)
	if p.JsonName == "" {
		p.print.P(fNode.Type.string(), "[", fNode.KeyType.string(), "]")
	} else {
		p.print.P(toProperCase(fNode.JsonName), " ", fNode.Type.string(), "[", fNode.KeyType.string(), "]")
	}
	if isSimpleType(fNode.ValType) {
		p.print.P(fNode.ValType.string())
	} else {
		for _, sn := range fNode.SubNode() {
			sn.dump()
		}
	}
	if p.JsonName != "" {
		p.print.P(fmt.Sprintf(" `json:\"%s\"`\n", fNode.JsonName))
	}
}

func (p *MapNode) SubNode() []Node {
	return p.ValNode
}

func (p *MapNode) debug() {
	fmt.Printf("======debug map node: %+v\n", *p)
	p.BaseNode.debug()
}