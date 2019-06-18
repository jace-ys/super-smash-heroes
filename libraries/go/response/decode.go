package response

import (
	"encoding/json"
	"io"
	"io/ioutil"
)

func Decode(r io.Reader, v interface{}) error {
	body, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, &v)
	if err != nil {
		return err
	}
	return nil
}
