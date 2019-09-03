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

package stringset

import (
	"bytes"
	"fmt"
	"sort"
)

type StringSet struct {
	smap map[string]bool
}

func New(elements ...string) *StringSet {
	ss := &StringSet{
		smap: make(map[string]bool),
	}
	for _, s := range elements {
		ss.Add(s)
	}
	return ss
}

func (set *StringSet) Add(ss ...string) {
	for _, s := range ss {
		set.smap[s] = true
	}
}

func (set *StringSet) AddSet(set1 *StringSet) {
	set.Add(set1.Elements()...)
}

func (ss *StringSet) Contain(s string) bool {
	_, exist := ss.smap[s]
	return exist
}

func (ss *StringSet) Elements() []string {
	elements := make([]string, 0, len(ss.smap))
	for s, _ := range ss.smap {
		elements = append(elements, s)
	}
	sort.Strings(elements)
	return elements
}

func (ss *StringSet) Equal(ss1 *StringSet) bool {
	if len(ss.smap) != len(ss1.smap) {
		return false
	}
	for k, _ := range ss.smap {
		if _, exist := ss1.smap[k]; !exist {
			return false
		}
	}
	return true
}

func (ss *StringSet) Len() int {
	return len(ss.smap)
}

func (ss *StringSet) Remove(s string) {
	delete(ss.smap, s)
}

func (ss *StringSet) String() string {
	if len(ss.smap) == 0 {
		return "{}"
	}
	w := new(bytes.Buffer)
	fmt.Fprint(w, "{")
	for i, s := range ss.Elements() {
		if i > 0 {
			fmt.Fprint(w, ",")
		}
		fmt.Fprint(w, s)
	}
	fmt.Fprint(w, "}")
	return w.String()
}
