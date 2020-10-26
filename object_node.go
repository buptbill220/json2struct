package json2struct

import "fmt"

type ObjectNode struct {
	*BaseNode
	KeyType JsonType // key type: default is KeywordStringType
	ValNode []Node  // sub node list of map value
}

func (p *ObjectNode) JsonToStruct(str string) (ret string, err error) {
	defer func() {
		err = p.recover()
	}()
	p.parse(str)
	p.fillSubNode()
	p.dump()
	ret = p.print.Output()
	return
}

func (p *ObjectNode) fillSubNode() {
	if p.Val == nil {
		return
	}
	v, ok := p.Val.(map[string]interface{})
	if !ok {
		panic("ObjectNode is not map[string]interface{}, impossible")
	}
	if len(v) == 0 {
		return
	}
	p.KeyType = KeywordStringType
	p.ValNode = make([]Node, 0, len(v))
	for sk, sv := range v {
		jt := parseValueType(sv)
		p.ValNode = append(p.ValNode, NewNode(sk, sv, p.GenType, jt, p.print))
	}
	for _, node := range p.ValNode {
		node.fillSubNode()
	}
}

func (p *ObjectNode) dump() {
	// don't change p
	fNode := fullNode(p.copy(p)).(*ObjectNode)
	if p.JsonName == "" {
		p.print.P(fNode.Type.string(), " {")
	} else {
		p.print.P(toProperCase(fNode.JsonName), " ", fNode.Type.string(), " {")
	}
	// for embedding struct type, each sub line have tabs at line head
	p.print.In()
	p.print.P("\n")
	for _, sn := range fNode.SubNode() {
		sn.dump()
	}
	p.print.Out()
	if p.JsonName == "" {
		p.print.P("\n}", fNode.JsonName)
	} else {
		p.print.P(fmt.Sprintf("\n} `json:\"%s\"`\n", fNode.JsonName))
	}
}

func (p *ObjectNode) SubNode() []Node {
	return p.ValNode
}


func (p *ObjectNode) debug() {
	fmt.Printf("======debug object node: %+v\n", *p)
	p.BaseNode.debug()
}