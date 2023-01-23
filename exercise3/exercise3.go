package exercise3

func WhichOperation(S string, T string) string {
	if S == T {
		return "NOTHING"
	}

	isAdd, solution := add(S, T)
	if isAdd {
		return "ADD " + solution
	}

	isJoin, solution := join(S, T)
	if isJoin {
		return "JOIN " + solution
	}

	isSwap, solution := swap(S, T)
	if isSwap {
		return "SWAP " + solution
	}

	return "IMPOSSIBLE"
}

func add(S string, T string) (bool, string) {

	for i := 0; i < len(T); i++ {
		if S+string(T[i]) == T {
			return true, string(T[i])
		}
	}
	return false, ""
}

func join(S string, T string) (bool, string) {
	for i := 0; i < len(S)-1; i++ {
		if S[i] == S[i+1] {
			//see if the word without that one duplicated letter is the same as T
			newWord := deleleteDuplicatedChar(S, i)
			if newWord == T {
				return true, string(S[i])
			}
		}
	}
	return false, ""
}

func deleleteDuplicatedChar(S string, index int) string { //https://golangbyexample.com/golang-delete-index-string/
	runeS := []rune(S)
	result := append(runeS[0:index], runeS[index+1:]...)
	return string(result)
}

func swap(S string, T string) (bool, string) { //https://golangbyexample.com/swap-characters-string-golang/
	lenS := len(S)
	for i := 0; i < lenS-1; i++ {

		for j := i + 1; j < lenS; j++ {
			runeS := []rune(S)

			valueAux := runeS[i]
			runeS[i] = runeS[j]
			runeS[j] = valueAux

			if string(runeS) == T {
				return true, string(runeS[j]) + " " + string(runeS[i])
			}
		}
	}

	return false, ""
}
