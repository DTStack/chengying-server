package util

import (
	"encoding/json"
)

func ToMap(object interface{}) (map[string]interface{}, error) {
	bts, err := json.Marshal(object)
	if err != nil {
		return nil, err
	}
	m := map[string]interface{}{}
	if err = json.Unmarshal(bts, &m); err != nil {
		return nil, err
	}
	return m, nil
}

func ToObject(date interface{}, object interface{}) error {
	bts, err := json.Marshal(date)
	if err != nil {
		return err
	}
	if err = json.Unmarshal(bts, object); err != nil {
		return err
	}
	return nil
}
