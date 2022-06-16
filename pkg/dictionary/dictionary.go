package dictionary

import (
	"math/rand"
	"sort"
	"time"

	"github.com/pkg/errors"
	"github.com/uladzislaudd/wordle/internal/pkg/util"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type Dictionary struct {
	Alphabet []rune
	WordSize uint8
	Words    []string
}

func LoadFromJSON(filename string) (*Dictionary, error) {
	rv := &Dictionary{}

	err := util.LoadFromJSON(filename, &rv)
	if err == nil {
		rv.Sort()
	}

	return rv, errors.Wrap(err, "util.LoadFromJSON() failed")
}

func (d *Dictionary) SaveToJSON(filename string) (err error) {
	d.Sort()
	return util.SaveToJSON(filename, d)
}

func (d *Dictionary) Sort() {
	sort.Slice(d.Alphabet, func(i, j int) bool { return d.Alphabet[i] < d.Alphabet[j] })
	sort.Strings(d.Words)
}
