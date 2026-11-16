package doc

import (
	"github.com/yuin/goldmark/ast"
)

// Block nodes that front ends build for constructs Markdown has no node for.

var KindCalloutNode = ast.NewNodeKind("Callout")

type CalloutKind uint8

const (
	CalloutNote CalloutKind = iota
	CalloutTip
	CalloutImportant
	CalloutWarning
	CalloutCaution
)

// Callout is a titled block with a coloured bar down its left edge, such as
// a GitHub alert.
type Callout struct {
	ast.BaseBlock
	CalloutKind CalloutKind
}

func (n *Callout) Kind() ast.NodeKind         { return KindCalloutNode }
func (n *Callout) Dump(src []byte, level int) { ast.DumpHelper(n, src, level, nil, nil) }

type calloutStyle struct {
	icon, title string
	color       func(*Theme) Color
}

var calloutStyles = map[CalloutKind]calloutStyle{
	CalloutNote:      {"ⓘ", "Note", func(t *Theme) Color { return t.Note }},
	CalloutTip:       {"✦", "Tip", func(t *Theme) Color { return t.Tip }},
	CalloutImportant: {"❖", "Important", func(t *Theme) Color { return t.Important }},
	CalloutWarning:   {"▲", "Warning", func(t *Theme) Color { return t.Warning }},
	CalloutCaution:   {"◆", "Caution", func(t *Theme) Color { return t.Caution }},
}

func (l *layouter) callout(n *Callout) {
	l.startBlock()
	cs := calloutStyles[n.CalloutKind]
	color := cs.color(l.o.Theme)
	bar := []Span{{Text: "▎ ", Style: Style{FG: color}}}
	l.push(bar, bar)
	l.emit([]Span{{Text: cs.icon + " " + cs.title, Style: Style{FG: color, Bold: true}}})
	l.blocks(n)
	l.pop()
	l.endBlock()
}
