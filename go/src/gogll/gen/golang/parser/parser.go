package parser

import (
	"bytes"
	"fmt"
	"gogll/ast"
	"gogll/gslot"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"sort"
	"strings"
	"text/template"
)

const FilePerm = 0731

var g *ast.Grammar

func Gen(parserDir string, grammar *ast.Grammar) {
	g = grammar
	buf := new(bytes.Buffer)
	tmpl, err := template.New("Parser Template").Parse(src)
	if err != nil {
		failError(err)
	}
	data := getData(parserDir)
	if err = tmpl.Execute(buf, data); err != nil {
		failError(err)
	}
	fname := path.Join(parserDir, "parser.go")
	if err := ioutil.WriteFile(fname, buf.Bytes(), FilePerm); err != nil {
		failError(err)
	}
	genLabels(filepath.Join(parserDir, "labels"), data.Labels)
}

type Data struct {
	Package     string
	Imports     []string
	StartSymbol string
	Labels      []string
	Rules       []*RuleCode
}

type RuleCode struct {
	Label string
	Code  string
}

func getData(baseDir string) *Data {
	data := &Data{
		Package:     getPackage(baseDir),
		Imports:     getImports(baseDir),
		StartSymbol: ast.GetStartSymbol(),
		Rules:       getRules(),
	}
	data.Labels = getLabels(data.Rules)
	return data
}

func getImports(baseDir string) []string {
	return []string{
		"io/ioutil",
	}
}

func getLabels(ruleCode []*RuleCode) (labels []string) {
	for _, r := range ruleCode {
		labels = append(labels, r.Label)
	}
	for _, s := range gslot.GetSlots() {
		labels = append(labels, s.Label())
	}
	sort.Strings(labels)
	return
}

func getRules() (rules []*RuleCode) {
	for _, nt := range ast.GetNonTerminals() {
		rules = append(rules, getRuleCode(nt)...)
	}
	return
}

func getRuleCode(nt string) (rules []*RuleCode) {
	rule := ast.GetRule(nt)
	rules = append(rules, getRuleTestCode(rule))
	for i := range rule.Alternates {
		rules = append(rules, getAlternateCode(rule, i))
	}
	return
}

func getAlternateCode(rule *ast.Rule, altI int) *RuleCode {
	rc := &RuleCode{
		Label: getAlternateLabel(rule.Head.Value(), altI),
		Code:  codeAlt(rule, altI),
	}
	return rc
}

func getAlternateLabel(nt string, i int) string {
	return fmt.Sprintf("L_%s%d", nt, i)
}

func getRuleTestCode(rule *ast.Rule) *RuleCode {
	r := &RuleCode{
		Label: "J_" + rule.Head.Value(),
		Code:  getTestSelectsForRule(rule),
	}
	return r
}

func getTestSelectsForRule(rule *ast.Rule) (code string) {
	buf := new(bytes.Buffer)
	for i, a := range rule.Alternates {
		altCode := getTestSelectForAlternate(rule.Head.Value(), i, a.Symbols()...)
		fmt.Fprintf(buf, "%s\n", altCode)
	}
	fmt.Fprint(buf, "            L = labels.L0\n")
	return buf.String()
}

func getPackage(baseDir string) string {
	if ast.GetPackage() == "" {
		pl := strings.Split(filepath.Clean(filepath.ToSlash(baseDir)), "/")
		pkg := pl[len(pl)-1]
		return pkg
	}
	return ast.GetPackage()
}

func failError(err error) {
	fmt.Printf("Error generating parser: %s\n", err)
	panic("fix me")
	os.Exit(1)
}

const src = `
/* 
Package parser is generated by gogll. Do not edit.
*/
package parser

import(
	"fmt"
	"os"
	"unicode/utf8"

	"{{.Package}}/parser/labels"
	"{{.Package}}/parser/sppf"
	{{range $i, $import := .Imports}}
	"{{$import}}" {{end}}
)

const Dollar = ""

var dummy = sppf.Dummy

func ParseFile(fname string) {
	buf, err := ioutil.ReadFile(fname)
	if err != nil {
		failError(err)
	}
	Parse(buf)
}

func Parse(input []byte) {
	m := len(input)
	u0 := Node{L: labels.L0, I: 0}
	cU, cN, cI := u0, dummy, 0

	L := labels.J_{{.StartSymbol}}
	for done := false; !done; {
		next, runeSize := decodeRune(input[cI:])
		fmt.Printf("L:%s, cI=%d, next=%s, size=%d\n", labels.String(L), cI, next, runeSize)
		switch L {
		case labels.L0:
			if !R.empty() {
				L, cU, cI, cN = R.remove()
			} else {
				if sppf.ExistSymNode("{{.StartSymbol}}", 0, m) {
					return
				} else {
					fail()
				}
			}

		{{range $i, $r := .Rules}}case labels.{{$r.Label}}:
			{{$r.Code}}

		{{end}}default:
			panic("This must not happen")
		}
	}
}

/*** descriptors ***/

var (
	R *descriptors = &descriptors{}
	U *descriptors = &descriptors{}
)

type descriptors struct {
	set []*descriptor
}

func (ds *descriptors) contain(d *descriptor) bool {
	for _, d1 := range ds.set {
		if d1 == d {
			return true
		}
	}
	return false
}

type descriptor struct {
	L int
	u Node
	i int
	w sppf.Node
}

func add(L int, u Node, i int, w sppf.Node) {
	d := &descriptor{L, u, i, w}
	if !U.contain(d) {
		R.set = append(R.set, d)
		U.set = append(U.set, d)
	}
}

func (d *descriptors) empty() bool {
	return len(d.set) == 0
}

func (ds *descriptors) remove() (L int, u Node, i int, w sppf.Node) {
	d := ds.set[len(ds.set)-1]
	ds.set = ds.set[:len(ds.set)-1]
	return d.L, d.u, d.i, d.w
}

/*** Rune decoding ***/
func decodeRune(str []byte) (string, int) {
	if len(str) == 0 {
		return Dollar, 0
	}
	r, sz := utf8.DecodeRune(str)
	if r == utf8.RuneError {
		panic(fmt.Sprintf("Rune error: %s", str))
	}
	chr := runeToString(r)
	return chr, sz
}

func runeToString(r rune) string {
	buf := make([]byte, utf8.RuneLen(r))
	utf8.EncodeRune(buf, r)
	return string(buf)
}

/*** GSS ***/

type GSS struct {
	nodes  map[Node]bool
	edges  map[Edge]bool
	popped map[Node]*poppedNode
}

type Edge struct {
	from Node
	to   Node
	w    sppf.Node
}

type Node struct {
	// slot label
	L int
	// Input position
	I int
}

type poppedNode struct {
	gss  Node
	sppf sppf.Node
}

var gss = &GSS{
	nodes:  make(map[Node]bool),
	edges:  make(map[Edge]bool),
	popped: make(map[Node]*poppedNode),
}

func create(L int, u Node, i int, w sppf.Node) Node {
	v := Node{L: L, I: i}
	if _, exist := gss.nodes[v]; !exist {
		gss.nodes[v] = true
	}
	e := Edge{from: v, to: u, w: w}
	if _, exist := gss.edges[e]; !exist {
		gss.edges[e] = true
		for _, p := range getPoppedNodes(v) {
			y := sppf.GetNode(L, w, p.sppf)
			add(L, v, p.sppf.GetRightExtent(), y)
		}
	}
	return v
}

func getGSSEdges(u Node) (edges []Edge) {
	for e, _ := range gss.edges {
		if e.from == u {
			edges = append(edges, e)
		}
	}
	return
}

func getPoppedNodes(v Node) (pns []*poppedNode) {
	for _, p := range gss.popped {
		if p.gss == v {
			pns = append(pns, p)
		}
	}
	return
}

func pop(u Node, i int, z sppf.Node) {
	gss.popped[u] = &poppedNode{u, z}
	for _, e := range getGSSEdges(u) {
		y := sppf.GetNode(u.L, e.w, z)
		add(u.L, e.to, i, y)
	}
}


/*** Errors ***/

func fail() {
	panic("implement me")
}

func failError(err error) {
	fmt.Printf("Error: %s\n", err)
	os.Exit(1)
}
`

const labelsSrc = `
/* 
Package labels is generated by gogll. Do not edit. 
*/
package labels

// Rule and grammar slot labels
const (
	L0 int = iota{{range $i, $l := .}}
	{{$l}}{{end}}
)

func String(label int) string {
	switch label {
	case L0: return "L0" {{range $i, $l := .}}
	case {{$l}}: 
		return "{{$l}}"{{end}}
	}
	panic("impossible")
}

`
func genLabels(labelsDir string, labels []string) {
	if err := os.MkdirAll(labelsDir, 0731); err != nil {
		panic(err)
	}
	tmpl, err := template.New("labels").Parse(labelsSrc)
	if err != nil {
		panic(err)
	}
	buf := new(bytes.Buffer)
	if err := tmpl.Execute(buf, labels); err != nil {
		panic(err)
	}
	if err := ioutil.WriteFile(filepath.Join(labelsDir, "labels.go"), buf.Bytes(), 0731); err != nil {
		panic(err)
	}
}