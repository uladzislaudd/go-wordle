package main

import (
	"github.com/pborman/getopt"
	"github.com/uladzislaudd/wordle/pkg/ruleset"
)

var (
	attempts  = uint(6)
	rulesetFn = "ruleset.json"
)

func init() {
	getopt.UintVarLong(&attempts, "attemts", 'a', "attempts count", "6")
	getopt.StringVarLong(&rulesetFn, "output", 'o', "filepath to save", "/path/to/output.json")
	if err := getopt.Getopt(nil); err != nil {
		panic(err)
	}
}
func main() {
	r := ruleset.Ruleset{Attempts: attempts}
	if err := r.SaveToJSON(rulesetFn); err != nil {
		panic(err)
	}
}
