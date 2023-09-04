package main

import (
	"leetcode/common"
)

// 2 = abc
// 3 = def
// 4 = ghi
// 5 = jkl
// 6 = mno
// 7 = pqrs
// 8 = tuv
// 9 = wxyz

var numMap = map[string][]string{
	"2": {"a", "b", "c"},
	"3": {"d", "e", "f"},
	"4": {"g", "h", "i"},
	"5": {"j", "k", "l"},
	"6": {"m", "n", "o"},
	"7": {"p", "q", "r", "s"},
	"8": {"t", "u", "v"},
	"9": {"w", "x", "y", "z"},
}

func letterCombinations(digits string) []string {
	target := []string{}

	for _, letter := range digits {
		num := string(letter)
		target = addLetter(target, numMap[num])
	}

	return target
}

func addLetter(sources []string, toAdd []string) []string {
	var targets []string

	if sources == nil || len(sources) == 0 {
		return toAdd
	}

	for _, s := range sources {
		for _, a := range toAdd {
			targets = append(targets, s+a)
		}
	}

	return targets

}

func main() {
	common.PrintArr(letterCombinations(""))
}
