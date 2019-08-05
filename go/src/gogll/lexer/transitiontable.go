// Code generated by gocc; DO NOT EDIT.

package lexer

/*
Let s be the current state
Let r be the current input rune
transitionTable[s](r) returns the next state.
*/
type TransitionTable [NumStates]func(rune) int

var TransTab = TransitionTable{
	// S0
	func(r rune) int {
		switch {
		case r == 9: // ['\t','\t']
			return 1
		case r == 10: // ['\n','\n']
			return 1
		case r == 13: // ['\r','\r']
			return 1
		case r == 32: // [' ',' ']
			return 1
		case r == 34: // ['"','"']
			return 2
		case r == 39: // [''',''']
			return 3
		case r == 47: // ['/','/']
			return 4
		case r == 58: // [':',':']
			return 5
		case r == 59: // [';',';']
			return 6
		case 65 <= r && r <= 90: // ['A','Z']
			return 7
		case r == 97: // ['a','a']
			return 8
		case r == 101: // ['e','e']
			return 9
		case r == 108: // ['l','l']
			return 10
		case r == 110: // ['n','n']
			return 11
		case r == 112: // ['p','p']
			return 12
		case r == 115: // ['s','s']
			return 13
		case r == 117: // ['u','u']
			return 14
		case r == 124: // ['|','|']
			return 15
		}
		return NoState
	},
	// S1
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S2
	func(r rune) int {
		switch {
		case r == 92: // ['\','\']
			return 16
		default:
			return 17
		}
	},
	// S3
	func(r rune) int {
		switch {
		case r == 92: // ['\','\']
			return 18
		default:
			return 19
		}
	},
	// S4
	func(r rune) int {
		switch {
		case r == 42: // ['*','*']
			return 20
		case r == 47: // ['/','/']
			return 21
		}
		return NoState
	},
	// S5
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S6
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S7
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 22
		case 65 <= r && r <= 90: // ['A','Z']
			return 23
		case r == 95: // ['_','_']
			return 24
		case 97 <= r && r <= 122: // ['a','z']
			return 25
		}
		return NoState
	},
	// S8
	func(r rune) int {
		switch {
		case r == 110: // ['n','n']
			return 26
		}
		return NoState
	},
	// S9
	func(r rune) int {
		switch {
		case r == 109: // ['m','m']
			return 27
		}
		return NoState
	},
	// S10
	func(r rune) int {
		switch {
		case r == 101: // ['e','e']
			return 28
		case r == 111: // ['o','o']
			return 29
		}
		return NoState
	},
	// S11
	func(r rune) int {
		switch {
		case r == 111: // ['o','o']
			return 30
		case r == 117: // ['u','u']
			return 31
		}
		return NoState
	},
	// S12
	func(r rune) int {
		switch {
		case r == 97: // ['a','a']
			return 32
		}
		return NoState
	},
	// S13
	func(r rune) int {
		switch {
		case r == 112: // ['p','p']
			return 33
		}
		return NoState
	},
	// S14
	func(r rune) int {
		switch {
		case r == 112: // ['p','p']
			return 34
		}
		return NoState
	},
	// S15
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S16
	func(r rune) int {
		switch {
		case r == 34: // ['"','"']
			return 35
		case r == 39: // [''',''']
			return 35
		case r == 110: // ['n','n']
			return 35
		case r == 114: // ['r','r']
			return 35
		case r == 116: // ['t','t']
			return 35
		}
		return NoState
	},
	// S17
	func(r rune) int {
		switch {
		case r == 34: // ['"','"']
			return 36
		case r == 92: // ['\','\']
			return 37
		default:
			return 17
		}
	},
	// S18
	func(r rune) int {
		switch {
		case r == 34: // ['"','"']
			return 38
		case r == 39: // [''',''']
			return 38
		case r == 110: // ['n','n']
			return 38
		case r == 114: // ['r','r']
			return 38
		case r == 116: // ['t','t']
			return 38
		}
		return NoState
	},
	// S19
	func(r rune) int {
		switch {
		case r == 39: // [''',''']
			return 39
		}
		return NoState
	},
	// S20
	func(r rune) int {
		switch {
		case r == 42: // ['*','*']
			return 40
		default:
			return 20
		}
	},
	// S21
	func(r rune) int {
		switch {
		case r == 10: // ['\n','\n']
			return 41
		default:
			return 21
		}
	},
	// S22
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 22
		case 65 <= r && r <= 90: // ['A','Z']
			return 23
		case r == 95: // ['_','_']
			return 24
		case 97 <= r && r <= 122: // ['a','z']
			return 25
		}
		return NoState
	},
	// S23
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 22
		case 65 <= r && r <= 90: // ['A','Z']
			return 23
		case r == 95: // ['_','_']
			return 24
		case 97 <= r && r <= 122: // ['a','z']
			return 25
		}
		return NoState
	},
	// S24
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 22
		case 65 <= r && r <= 90: // ['A','Z']
			return 23
		case r == 95: // ['_','_']
			return 24
		case 97 <= r && r <= 122: // ['a','z']
			return 25
		}
		return NoState
	},
	// S25
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 22
		case 65 <= r && r <= 90: // ['A','Z']
			return 23
		case r == 95: // ['_','_']
			return 24
		case 97 <= r && r <= 122: // ['a','z']
			return 25
		}
		return NoState
	},
	// S26
	func(r rune) int {
		switch {
		case r == 121: // ['y','y']
			return 42
		}
		return NoState
	},
	// S27
	func(r rune) int {
		switch {
		case r == 112: // ['p','p']
			return 43
		}
		return NoState
	},
	// S28
	func(r rune) int {
		switch {
		case r == 116: // ['t','t']
			return 44
		}
		return NoState
	},
	// S29
	func(r rune) int {
		switch {
		case r == 119: // ['w','w']
			return 45
		}
		return NoState
	},
	// S30
	func(r rune) int {
		switch {
		case r == 116: // ['t','t']
			return 46
		}
		return NoState
	},
	// S31
	func(r rune) int {
		switch {
		case r == 109: // ['m','m']
			return 47
		}
		return NoState
	},
	// S32
	func(r rune) int {
		switch {
		case r == 99: // ['c','c']
			return 48
		}
		return NoState
	},
	// S33
	func(r rune) int {
		switch {
		case r == 97: // ['a','a']
			return 49
		}
		return NoState
	},
	// S34
	func(r rune) int {
		switch {
		case r == 99: // ['c','c']
			return 50
		}
		return NoState
	},
	// S35
	func(r rune) int {
		switch {
		case r == 34: // ['"','"']
			return 36
		case r == 92: // ['\','\']
			return 37
		default:
			return 17
		}
	},
	// S36
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S37
	func(r rune) int {
		switch {
		case r == 34: // ['"','"']
			return 35
		case r == 39: // [''',''']
			return 35
		case r == 110: // ['n','n']
			return 35
		case r == 114: // ['r','r']
			return 35
		case r == 116: // ['t','t']
			return 35
		}
		return NoState
	},
	// S38
	func(r rune) int {
		switch {
		case r == 39: // [''',''']
			return 39
		}
		return NoState
	},
	// S39
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S40
	func(r rune) int {
		switch {
		case r == 42: // ['*','*']
			return 40
		case r == 47: // ['/','/']
			return 51
		default:
			return 20
		}
	},
	// S41
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S42
	func(r rune) int {
		switch {
		case r == 111: // ['o','o']
			return 52
		}
		return NoState
	},
	// S43
	func(r rune) int {
		switch {
		case r == 116: // ['t','t']
			return 53
		}
		return NoState
	},
	// S44
	func(r rune) int {
		switch {
		case r == 116: // ['t','t']
			return 54
		}
		return NoState
	},
	// S45
	func(r rune) int {
		switch {
		case r == 99: // ['c','c']
			return 55
		}
		return NoState
	},
	// S46
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S47
	func(r rune) int {
		switch {
		case r == 98: // ['b','b']
			return 56
		}
		return NoState
	},
	// S48
	func(r rune) int {
		switch {
		case r == 107: // ['k','k']
			return 57
		}
		return NoState
	},
	// S49
	func(r rune) int {
		switch {
		case r == 99: // ['c','c']
			return 58
		}
		return NoState
	},
	// S50
	func(r rune) int {
		switch {
		case r == 97: // ['a','a']
			return 59
		}
		return NoState
	},
	// S51
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S52
	func(r rune) int {
		switch {
		case r == 102: // ['f','f']
			return 60
		}
		return NoState
	},
	// S53
	func(r rune) int {
		switch {
		case r == 121: // ['y','y']
			return 61
		}
		return NoState
	},
	// S54
	func(r rune) int {
		switch {
		case r == 101: // ['e','e']
			return 62
		}
		return NoState
	},
	// S55
	func(r rune) int {
		switch {
		case r == 97: // ['a','a']
			return 63
		}
		return NoState
	},
	// S56
	func(r rune) int {
		switch {
		case r == 101: // ['e','e']
			return 64
		}
		return NoState
	},
	// S57
	func(r rune) int {
		switch {
		case r == 97: // ['a','a']
			return 65
		}
		return NoState
	},
	// S58
	func(r rune) int {
		switch {
		case r == 101: // ['e','e']
			return 66
		}
		return NoState
	},
	// S59
	func(r rune) int {
		switch {
		case r == 115: // ['s','s']
			return 67
		}
		return NoState
	},
	// S60
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S61
	func(r rune) int {
		switch {
		case r == 65: // ['A','A']
			return 68
		}
		return NoState
	},
	// S62
	func(r rune) int {
		switch {
		case r == 114: // ['r','r']
			return 69
		}
		return NoState
	},
	// S63
	func(r rune) int {
		switch {
		case r == 115: // ['s','s']
			return 70
		}
		return NoState
	},
	// S64
	func(r rune) int {
		switch {
		case r == 114: // ['r','r']
			return 71
		}
		return NoState
	},
	// S65
	func(r rune) int {
		switch {
		case r == 103: // ['g','g']
			return 72
		}
		return NoState
	},
	// S66
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S67
	func(r rune) int {
		switch {
		case r == 101: // ['e','e']
			return 73
		}
		return NoState
	},
	// S68
	func(r rune) int {
		switch {
		case r == 108: // ['l','l']
			return 74
		}
		return NoState
	},
	// S69
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S70
	func(r rune) int {
		switch {
		case r == 101: // ['e','e']
			return 75
		}
		return NoState
	},
	// S71
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S72
	func(r rune) int {
		switch {
		case r == 101: // ['e','e']
			return 76
		}
		return NoState
	},
	// S73
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S74
	func(r rune) int {
		switch {
		case r == 116: // ['t','t']
			return 77
		}
		return NoState
	},
	// S75
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S76
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S77
	func(r rune) int {
		switch {
		}
		return NoState
	},
}
