//  Copyright 2020 Marius Ackerman
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

package parser

import (
	"bytes"
	"path"
	"text/template"

	"github.com/goccmack/gogll/cfg"
	"github.com/goccmack/goutil/ioutil"
)

func genErrors(pkg string) {
	tmpl, err := template.New("parser errors").Parse(errorsSrc)
	if err != nil {
		panic(err)
	}
	wr := new(bytes.Buffer)
	tmpl.Execute(wr, path.Join(pkg, "token"))
	if err := ioutil.WriteFile(path.Join(cfg.BaseDir, "src", "errors", "mod.rs"), wr.Bytes()); err != nil {
		panic(err)
	}
}

const errorsSrc = `//! Generated by GoGLL. Do not edit.

use crate::ast;
use crate::token;

pub struct Error struct {
	pub err             String,
	pub error_token     Box<token::Token>,
	pub error_symbols   Vec<ast::Node>,
	pub expected_tokens Vec<String>,
}

impl fmt::Display for Error {
    fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
		write!(f, "Error {}\n", self.err);
		write!(f, "Token: type={}, literal={}", 
			self.error_token.typ, 
			self.error_token.literal_string())
		let (ln, col) = self.error_token.get_line_column();
		write!(f, "line {} column {}\n", ln, col);
		write!(f, "Expected one of:\n");
		for tok in self.expected_tokens.iter() {
			write!(f, "{}", tok);
		}
	}

}

// (E *Error) String() string {
// 	w := new(bytes.Buffer)
// 	fmt.Fprintf(w, "Error")
// 	if E.Err != nil {
// 		fmt.Fprintf(w, " %s\n", E.Err)
// 	} else {
// 		fmt.Fprintf(w, "\n")
// 	}
// 	fmt.Fprintf(w, "Token: type=%d, lit=%s\n", E.ErrorToken.Type, E.ErrorToken.Literal())
// 	ln, col := E.ErrorToken.GetLineColumn()
// 	fmt.Fprintf(w, "Pos: offset=%d, line=%d, column=%d\n", E.ErrorToken.Lext(), ln, col)
// 	fmt.Fprintf(w, "Expected one of: ")
// 	for _, sym := range E.ExpectedTokens {
// 		fmt.Fprintf(w, "%s ", sym)
// 	}
// 	fmt.Fprintf(w, "ErrorSymbol:\n")
// 	for _, sym := range E.ErrorSymbols {
// 		fmt.Fprintf(w, "%v\n", sym)
// 	}
// 	return w.String()
// }
`