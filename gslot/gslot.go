//  Copyright 2019 Marius Ackerman
//
//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the License.
//  You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.

/*
Package gslot implements grammar slots
*/
package gslot

import (
	"bytes"
	"fmt"
	"gogll/ast"
	"gogll/frstflw"
	"gogll/symbols"
	"sort"
)

type Label struct {
	Head      string
	Alternate int
	Pos       int
}

var slots = make(map[Label]symbols.Symbols)

type Slots []Label

func GetSlots() Slots {
	if len(slots) == 0 {
		genSlots()
	}
	res := make(Slots, 0, len(slots))
	for l, _ := range slots {
		res = append(res, l)
	}
	sort.Sort(res)
	return res
}

func (s Label) Label() string {
	return fmt.Sprintf("%s%dR%d", s.Head, s.Alternate, s.Pos)
}

func (s Label) IsEoR() bool {
	symbols := slots[s]
	return s.Pos >= len(symbols)
}

func (s Label) IsFiR() bool {
	symbols := slots[s]
	if s.Pos > 1 || len(symbols) <= 1 {
		return false
	}
	if frstflw.FirstOfSymbol(symbols[0]).Contain(ast.Empty) &&
		symbols[0] == symbols[1] {
		return false
	}
	return true
}

func (s Label) String() string {
	symbols := slots[s]
	buf := new(bytes.Buffer)
	fmt.Fprintf(buf, "%s : ", s.Head)
	for i, sym := range symbols {
		if i == s.Pos {
			fmt.Fprintf(buf, "∙")
		}
		fmt.Fprintf(buf, "%s ", sym)
	}
	// fmt.Printf("slotLabel.String(): %s pos=%d len(symbols)=%d\n", s.Head, s.Pos, len(symbols))
	if s.Pos >= len(symbols) {
		fmt.Fprintf(buf, "∙")
	}
	// fmt.Println("  ", buf.String())
	return buf.String()
}

func (s Label) Symbols() symbols.Symbols {
	return slots[s]
}

func (ss Slots) Labels() (labels []string) {
	for _, s := range ss {
		labels = append(labels, s.Label())
	}
	return
}

func (ss Slots) Len() int {
	return len(ss)
}

func (ss Slots) Less(i, j int) bool {
	if ss[i].Head < ss[j].Head {
		return true
	}

	if ss[i].Head == ss[j].Head {
		if ss[i].Alternate < ss[j].Alternate {
			return true
		}
		if ss[i].Alternate == ss[j].Alternate {
			if ss[i].Pos < ss[j].Pos {
				return true
			}
		}
	}
	return false
}

func (ss Slots) Swap(i, j int) {
	iTmp := ss[i]
	ss[i] = ss[j]
	ss[j] = iTmp
}

func genSlots() {
	for _, rule := range ast.GetRules() {
		genSlotsOfRule(rule)
	}
}

func genSlotsOfRule(r *ast.Rule) {
	for i, a := range r.Alternates {
		genSlotsOfAlternate(r.Head.StringValue(), i, a.Symbols()...)
	}
}

func genSlotsOfAlternate(nt string, altI int, symbols ...string) {
	if symbols[0] == ast.Empty {
		genSlot(nt, altI, 0, []string{}...)
	} else {
		for pos := 0; pos <= len(symbols); pos++ {
			genSlot(nt, altI, pos, symbols...)
		}
	}
}

func genSlot(nt string, altI, pos int, symbols ...string) {
	slot := Label{
		Head:      nt,
		Alternate: altI,
		Pos:       pos,
	}
	slots[slot] = symbols
}