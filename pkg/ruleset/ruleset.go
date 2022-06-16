package ruleset

import "github.com/uladzislaudd/wordle/internal/pkg/util"

type Ruleset struct {
	Attempts uint
}

func LoadFromJSON(filename string) (rv *Ruleset, err error) {
	rv = &Ruleset{}
	err = util.LoadFromJSON(filename, &rv)
	return
}

func (r *Ruleset) SaveToJSON(filename string) error {
	return util.SaveToJSON(filename, r)
}
