func isValid(s string) bool {
    stack := make([]rune, 0, utf8.RuneCountInString(s))
    for _, c := range s {
        if string(c) == "(" || string(c) == "[" || string(c) == "{" {
            stack = append(stack, c)
        } else {
            if len(stack) == 0 {
                return false
            }
            top := stack[len(stack)-1]
            if !(string(top) == "(" &&  string(c) == ")" || string(top) == "[" &&  string(c) == "]" || string(top) == "{" && string(c) == "}") {
                return false
            }
            stack = stack[:len(stack)-1]
        }
    }
    return len(stack) == 0
}