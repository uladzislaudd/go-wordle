package dictionary

var (
	d *Dictionary
)

func init() {
	var err error
	d, err = LoadFromJSON("../../assets/wordle/dictionary.json")
	if err != nil {
		panic(err)
	}
}
