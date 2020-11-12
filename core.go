package sirius

import (
	"fmt"
	"strconv"
	"strings"
)

type jsonNode map[string]interface{}

func getJSONPathValue(path string, node jsonNode) interface{} {
	var currentNode = node
	var value interface{}
	keys := strings.Split(path, ".")
	for _, key := range keys {
		tempNode := getNodeValue(key, -1, currentNode)
		if tempNode == nil {
			value = currentNode[key]
			continue
		}
		currentNode = tempNode
		value = currentNode
	}
	return value
}

func getNodeValue(key string, arrayIndex int, node jsonNode) jsonNode {
	tempNode := getNode(node[key], arrayIndex)
	if tempNode == nil {
		tempKey, index, err := formatKey(key)
		if err != nil {
			return nil
		}
		return getNodeValue(tempKey, index, node)
	}
	return tempNode
}

func getNode(v interface{}, i int) jsonNode {
	switch node := v.(type) {
	case map[string]interface{}:
		configNode := jsonNode(node)
		return configNode
	case []interface{}:
		if len(node) <= i {
			return nil
		}
		if i == -1 {
			return nil
		}
		configNode := jsonNode(getNode(node[i], 0))
		return configNode
	default:
		return nil
	}
}

func formatKey(key string) (string, int, error) {
	i := strings.Index(key, "[")
	k := strings.Index(key, "]")
	if i == -1 {
		return "", 0, fmt.Errorf("Not format key for index")
	}
	formatKey := key[:i]
	if formatKey == "" {
		return "", 0, fmt.Errorf("Not format key")
	}
	index, err := strconv.Atoi(key[i+1 : k])
	if err != nil {
		return "", 0, err
	}
	return formatKey, index, nil
}
