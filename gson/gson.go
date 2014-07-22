package gson

import (
	"encoding/json"
	"errors"
)

// JSON contains JSON-format data.
type JSON struct {
	Data interface{}
}

// NewJSON returns the pointer to a new JSON
// unmarshalled with the body bytes.
func NewJSON(body []byte) (*JSON, error) {
	// To decode JSON data we use the Unmarshal function.
	// func Unmarshal(data []byte, v interface{}) error
	// create a place where the decoded data will be stored
	var s JSON
	err := json.Unmarshal(body, &s.Data)

	if err != nil {
		return nil, err
	}
	return &s, nil
}

// MarshalJSON implements the json.Marshaler interface.
func (j *JSON) MarshalJSON() ([]byte, error) {
	return json.Marshal(&j.Data)
}

// Encode marshals its data into []byte.
func (j *JSON) Encode() ([]byte, error) {
	return j.MarshalJSON()
}

// Map type asserts JSON to map[string]interface{}.
func (j *JSON) Map() (map[string]interface{}, error) {
	if m, ok := j.Data.(map[string]interface{}); ok {
		return m, nil
	}
	return nil, errors.New("Can't type assert with map[string]interface{}")
}

// SetMap mapifies the JSON
// and map key to value.
func (j *JSON) SetMap(key string, val interface{}) {
	m, ok := j.Map()
	if ok != nil {
		return
	}
	m[key] = val
}

// Get find the data by its key
// and returns a pointer to a new JSON object.
// Useful for chaining the nested JSON.
// js.Get("top_level").Get("dict").Get("value").Int()
func (j *JSON) Get(key string) *JSON {
	m, err := j.Map()
	if err == nil {
		if val, ok := m[key]; ok {
			return &JSON{val}
		}
	}
	return &JSON{nil}
}

// GetCheck find the data by its key
// returning the boolean value of success or failure.
func (j *JSON) GetCheck(key string) (*JSON, bool) {
	m, err := j.Map()
	if err == nil {
		if val, ok := m[key]; ok {
			return &JSON{val}, ok
		}
	}
	return nil, false
}

// GetBranch returns the JSON in the branch.
func (j *JSON) GetBranch(items ...string) *JSON {
	for i := range items {
		m, err := j.Map()
		if err != nil {
			return &JSON{nil}
		}
		if val, ok := m[items[i]]; ok {
			j = &JSON{val}
		} else {
			return &JSON{nil}
		}
	}
	return j
}

// GetByIndex returns a pointer to the JSON
// finding by its index.
// js.Get("top_level").Get("array").GetByIndex(1).Get("key").Int()
func (j *JSON) GetByIndex(idx int) *JSON {
	slice, err := j.Slice()
	if err == nil {
		if idx < len(slice) {
			return &JSON{slice[idx]}
		}
	}
	return &JSON{nil}
}

// Bool type asserts JSON to bool.
func (j *JSON) Bool() (bool, error) {
	if s, ok := j.Data.(bool); ok {
		return s, nil
	}
	return false, errors.New("Can't type assert with bool")
}

// String type asserts JSON to string.
func (j *JSON) String() (string, error) {
	if s, ok := j.Data.(string); ok {
		return s, nil
	}
	return "", errors.New("Can't type assert with string")
}

// Byte type asserts JSON to []byte.
func (j *JSON) Byte() ([]byte, error) {
	// type assert to string first
	if s, ok := j.Data.(string); ok {
		return []byte(s), nil
	}
	return nil, errors.New("Can't type assert with []byte")
}

// Int type asserts JSON to int.
func (j *JSON) Int() (int, error) {
	if f, ok := j.Data.(int); ok {
		return int(f), nil
	}
	return -1, errors.New("Can't type assert with int")
}

// Int64 type asserts JSON to int.
func (j *JSON) Int64() (int64, error) {
	if f, ok := j.Data.(int64); ok {
		return int64(f), nil
	}
	return -1, errors.New("Can't type assert with int64")
}

// Float64 type asserts JSON to float64.
func (j *JSON) Float64() (float64, error) {
	if f, ok := j.Data.(float64); ok {
		return float64(f), nil
	}
	return -1, errors.New("Can't type assert with float64")
}

// Slice type asserts JSON to []interface{}.
func (j *JSON) Slice() ([]interface{}, error) {
	if a, ok := j.Data.([]interface{}); ok {
		return a, nil
	}
	return nil, errors.New("Can't type assert with []interface{}")
}

// StringSlice type asserts JSON to []string.
func (j *JSON) StringSlice() ([]string, error) {
	sl, err := j.Slice()
	if err != nil {
		return nil, err
	}
	result := []string{}
	for _, v := range sl {
		s, ok := v.(string)
		if !ok {
			return nil, err
		}
		result = append(result, s)
	}
	return result, nil
}

// Float64Slice type asserts JSON to []float64.
func (j *JSON) Float64Slice() ([]float64, error) {
	sl, err := j.Slice()
	if err != nil {
		return nil, err
	}
	result := []float64{}
	for _, v := range sl {
		s, ok := v.(float64)
		if !ok {
			return nil, err
		}
		result = append(result, s)
	}
	return result, nil
}
