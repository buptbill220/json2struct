package json2struct

import (
	"bytes"
	"strings"
)

type P struct {
	buf   *bytes.Buffer
	depth int
}

func NewP() *P{
	return &P{
		buf: bytes.NewBuffer(make([]byte, 0, 1024)),
	}
}

var Tabs = [8]byte {
	'\t', '\t', '\t', '\t',
	'\t', '\t', '\t', '\t',
}

func (p *P) In() {
	p.depth++
}

func (p *P) Out() {
	p.depth--
}

func (p *P) Output() string {
	return p.buf.String()
}

func (p *P) P(args ...string) {
	if len(args) == 0 {
		return
	}
	var tabs []byte
	if p.depth <= 8 {
		tabs = Tabs[:p.depth]
	} else {
		tabs = make([]byte, p.depth)
		for i := 0; i < p.depth; i += 8 {
			copy(tabs[i:], Tabs[:])
		}
	}
	for _, arg := range args {
		arg = strings.Replace(arg, "\n", "\n" + string(tabs), -1)
		p.buf.WriteString(arg)
	}
}
