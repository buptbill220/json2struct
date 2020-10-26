package json2struct

import (
	"fmt"
	"runtime"
)

type BaseNode struct {
	Val        interface{}
	Type       JsonType
	JsonName   string
	ScopeDepth int
	GenType    GeneratorType
	Optional    bool // node is optional or not
	Merged      bool // true for don't need merge
	print      *P // all node share the print
}

func NewNode(name string, val interface{}, genType GeneratorType, jt JsonType, print *P) Node {
	bd := &BaseNode{
		Val:      val,
		Type:     jt,
		JsonName: name,
		GenType:  genType,
		print:    print,
	}
	switch jt {
	case SliceType:
		return &SliceNode{
			BaseNode: bd,
		}
	case MapType:
		return &MapNode{
			BaseNode: bd,
		}
	case StructType:
		return &ObjectNode{
			BaseNode: bd,
		}
	}
	return bd
}

type Node interface {
	JsonToStruct(str string) (string, error)
	IsSimple() bool
	GetBaseNode() *BaseNode
	SubNode() []Node
	fillSubNode()
	copy(Node) Node
	dump()
	debug()
	isMerged() bool
}

func (p *BaseNode) JsonToStruct(str string) (ret string, err error) {
	defer func() {
		err = p.recover()
	}()
	// 1: parse str to Node
	p.parse(str)
	// 2: fill sub Node
	p.fillSubNode()
	// 3: dump
	p.dump()
	ret = p.print.Output()
	return
}

func (p *BaseNode) dump() {
	ptr := ""
	omit := ""
	if p.Optional {
		ptr = "*"
		omit = ",omitempty"
	}
	if p.JsonName == "" {
		p.print.P(p.Type.string())
	} else {
		p.print.P(fmt.Sprintf("%s %s%s `json:\"%s%s\"`\n", toProperCase(p.JsonName), ptr, p.Type.string(), p.JsonName, omit))
	}
}

func (p *BaseNode) recover() error {
	if e := recover(); e != nil {
		buf := make([]byte, 64 << 10)
		buf = buf[:runtime.Stack(buf, false)]
		return fmt.Errorf("%v", e)
	}
	return nil
}

func (p *BaseNode) IsSimple() bool {
	return isSimpleType(p.Type)
}

func (p *BaseNode) GetBaseNode() *BaseNode {
	return p
}

func (p *BaseNode) SubNode() []Node {
	return nil
}

func (p *BaseNode) parse(str string) {
	t, v := toValue(str)
	p.Val = v
	p.Type = t
}

func (p *BaseNode) fillSubNode() {}

func (p *BaseNode) isMerged() bool {
	return p.Merged
}

func (p *BaseNode) copy(n Node) Node {
	if n.isMerged() {
		return n
	}
	base := n.GetBaseNode()
	node := NewNode(base.JsonName, base.Val, base.GenType, base.Type, base.print)
	node.GetBaseNode().Optional = base.Optional
	node.GetBaseNode().ScopeDepth = base.ScopeDepth
	node.fillSubNode()
	return node
}

func (p *BaseNode) debug() {
	fmt.Printf("======debug base node: %+v\n", *p)
	fmt.Printf("======debug subnode: \n")
	for _, sn := range p.SubNode() {
		fmt.Printf("subnode: %+v\n", sn)
	}
	fmt.Println("|||||||||debug end")
}