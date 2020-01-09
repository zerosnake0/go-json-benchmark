package json_benchmark

import (
	"encoding/json"

	"github.com/mreiferson/go-ujson"
)

type ujsonObjectStore struct {
}

var defaultUjsonObjectStore ujson.ObjectStore = ujsonObjectStore{}

func (ujsonObjectStore) NewObject() (interface{}, error) {
	return make(map[string]interface{}), nil
}

func (ujsonObjectStore) NewArray() (interface{}, error) {
	a := make([]interface{}, 0)
	return &a, nil
}

func (ujsonObjectStore) ObjectAddKey(mi interface{}, k interface{}, v interface{}) error {
	ks := k.(string)
	m := mi.(map[string]interface{})
	m[ks] = v
	return nil
}

func (ujsonObjectStore) ArrayAddItem(ai interface{}, v interface{}) error {
	a := ai.(*[]interface{})
	*a = append(*a, v)
	return nil
}

func (ujsonObjectStore) NewString(b []byte) (interface{}, error) {
	quoted := append([]byte{'"'}, b...)
	quoted = append(quoted, '"')
	var s string
	err := json.Unmarshal(quoted, &s)
	return s, err
}

func (ujsonObjectStore) NewNumeric(b []byte) (interface{}, error) {
	var f float64
	err := json.Unmarshal(b, &f)
	return f, err
}

func (ujsonObjectStore) NewTrue() (interface{}, error) {
	return true, nil
}

func (ujsonObjectStore) NewFalse() (interface{}, error) {
	return false, nil
}

func (ujsonObjectStore) NewNull() (interface{}, error) {
	return nil, nil
}
