package dialogue

import (
	"errors"
	"fmt"
	"strings"
)

func ExtractValue(st string, key string) (string, error) {
	pat := fmt.Sprintf("%s=", key)

	startIdx := strings.Index(st, pat)
	if startIdx < 0 {
		return "", errors.New("STRING DOESN'T CONTAIN SPECIFIED KEY")
	}

	temp := st[startIdx:]

	endIdx := strings.Index(temp, "&")
	if endIdx < 0 {
		if strings.Count(temp, "=") > 1 {
			return "", errors.New("STRING CONTAINS MULTIPLE '=' BUT NO '&' SEPARATOR")
		} else {
			return temp[(strings.Index(temp, "=") + 1):], nil
		}
	}

	value := temp[(strings.Index(temp, "=") + 1):endIdx]

	return value, nil
}

// Turn a list of key-value pairs into '&key=value&key=value' string.
func BuildKeyValueString(kvPairs []KVStruct) string {
	st := ""

	for _, kv := range kvPairs {
		st = st + "&" + kv.Key + "=" + kv.Value
	}

	return st
}
