package sirius

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

// SJson doc ...
type SJson struct {
	Buf  []byte
	node jsonNode
}

// Open doc ...
func Open(name string) (*SJson, error) {
	jsonFile, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()

	buffer, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}
	var node jsonNode
	err = json.Unmarshal([]byte(buffer), &node)
	if err != nil {
		return nil, err
	}
	sjson := SJson{
		Buf:  buffer,
		node: node,
	}
	return &sjson, nil
}

// Get doc ..
func (g SJson) Get(key string) interface{} {
	return getJSONPathValue(key, g.node)
}

// GetString doc ..
func (g SJson) GetString(key string) string {
	v := getJSONPathValue(key, g.node)
	str, ok := v.(string)
	if !ok {
		return ""
	}
	return str
}

// GetInt doc ..
func (g SJson) GetInt(key string) int {
	v := getJSONPathValue(key, g.node)
	str := fmt.Sprintf("%v", v)
	valueInt, err := strconv.Atoi(str)
	if err != nil {
		return 0
	}
	return valueInt
}

// GetInt64 doc ..
func (g SJson) GetInt64(key string) int64 {
	v := getJSONPathValue(key, g.node)
	str := fmt.Sprintf("%v", v)
	valueInt64, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return 0
	}
	return valueInt64
}

// GetFloat64 doc ..
func (g SJson) GetFloat64(key string) float64 {
	v := getJSONPathValue(key, g.node)
	valueFloat64, ok := v.(float64)
	if !ok {
		return 0.0
	}
	return valueFloat64
}

// GetBool doc ..
func (g SJson) GetBool(key string) bool {
	v := getJSONPathValue(key, g.node)
	valueBool, ok := v.(bool)
	if !ok {
		return false
	}
	return valueBool
}

// GetArrayMaps doc ..
func (g SJson) GetArrayMaps(key string) []map[string]interface{} {
	v := getJSONPathValue(key, g.node)
	value, ok := v.([]interface{})
	if !ok {
		return nil
	}
	var a []map[string]interface{}
	for _, m := range value {
		nm, ok := m.(map[string]interface{})
		if !ok {
			return nil
		}
		a = append(a, nm)
	}
	return a
}
