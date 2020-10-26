package json2struct

import "fmt"

type SliceNode struct {
	*BaseNode
	ValType JsonType // type of slice value
	ValNode []Node  // sub node list of slice value
}

func (p *SliceNode) JsonToStruct(str string) (ret string, err error) {
	defer func() {
		err = p.recover()
	}()
	p.parse(str)
	p.fillSubNode()
	p.dump()
	ret = p.print.Output()
	return
}

func (p *SliceNode) SubNode() []Node {
	return p.ValNode
}

func (p *SliceNode) fillSubNode() {
	if p.Val == nil {
		return
	}
	v, ok := p.Val.([]interface{})
	if !ok {
		panic("SliceNode is not []interface{}, impossible")
	}
	p.ValType = parseSliceType(v)
	if len(v) == 0 {
		return
	}
	if isSimpleType(p.ValType) {
		return
	}
	p.ValNode = make([]Node, 0, len(v))
	if p.ValType == SliceType {
		for _, sv := range v {
			p.ValNode = append(p.ValNode, NewNode("", sv, p.GenType, SliceType, p.print))
		}
	} else {
		for _, sv := range v {
			if sv == nil {
				p.ValNode = append(p.ValNode, NewNode("", sv, p.GenType, MapType, p.print))
				continue
			}
			isMap, isObj, _ := parseMapType(sv.(map[string]interface{}))
			if isMap {
				p.ValNode = append(p.ValNode, NewNode("", sv, p.GenType, MapType, p.print))
			}
			if isObj {
				p.ValNode = append(p.ValNode, NewNode("", sv, p.GenType, StructType, p.print))
			}
		}
	}
	for _, node := range p.ValNode {
		node.fillSubNode()
	}
}

func (p *SliceNode) dump() {
	// don't change p
	fNode := fullNode(p.copy(p)).(*SliceNode)
	if p.JsonName == "" {
		p.print.P(fNode.Type.string())
	} else {
		p.print.P(toProperCase(fNode.JsonName), " ", fNode.Type.string())
	}
	for _, sn := range fNode.SubNode() {
		sn.dump()
	}
	if isSimpleType(fNode.ValType) {
		p.print.P(fNode.ValType.string())
	}
	if p.JsonName != "" {
		p.print.P(fmt.Sprintf(" `json:\"%s\"`\n", fNode.JsonName))
	}
}

func (p *SliceNode) debug() {
	fmt.Printf("======debug slice node: %+v\n", *p)
	p.BaseNode.debug()
}