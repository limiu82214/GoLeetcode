package Q20ValidParenthesesPkg

func Solve(s string) bool {
	buff := []rune{}
	for _, c := range s {
		switch c {
		case '(', '{', '[':
			buff = append(buff, c)
		case ')', '}', ']':
			ls := len(buff)
			if ls == 0 {
				return false
			}
			pop := buff[ls-1]
			switch {
			case c == '}' && pop == '{':
			case c == ')' && pop == '(':
			case c == ']' && pop == '[':
			default:
				return false
			}
			buff = buff[:ls-1]
		}
	}

	return len(buff) == 0
}
