package util

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"io/fs"
	"io/ioutil"
	"strings"

	"github.com/pkg/errors"
)

func LoadStructFromFile(filename string, str interface{}) error {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return errors.Wrap(err, "ioutil.ReadFile() failed")
	}

	err = gob.NewDecoder(bytes.NewBuffer(b)).Decode(str)
	return errors.Wrap(err, "gob.NewDecoder().Decode() failed")
}

func SaveStructToFile(filename string, str interface{}) error {
	var buf bytes.Buffer
	err := gob.NewEncoder(&buf).Encode(str)
	if err != nil {
		return errors.Wrap(err, "gob.NewDecoder().Decode() failed")
	}

	err = ioutil.WriteFile(filename, buf.Bytes(), fs.ModePerm)
	return errors.Wrap(err, "ioutil.WriteFile() failed")
}

func LoadFromJSON(filename string, rv any) error {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return errors.Wrap(err, "ioutil.ReadFile() failed")
	}

	err = json.Unmarshal(b, rv)
	return errors.Wrap(err, "json.Unmarshal() failed")
}

func SaveToJSON(filename string, a any) error {
	bytes, err := json.Marshal(a)
	if err != nil {
		return errors.Wrap(err, "json.Marshal() failed")
	}

	if !strings.HasSuffix(filename, ".json") {
		filename += ".json"
	}

	err = ioutil.WriteFile(filename, bytes, fs.ModePerm)
	return errors.Wrap(err, "ioutil.WriteFile() failed")
}
