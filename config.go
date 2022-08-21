package echochamber

import (
	"bytes"
	"fmt"
)

/***
function to parse single string config to map[string]string for example
a=b c=d e=f
it will become
map[a:b c:d e:f]
***/
func ParseConfig(input string) (map[string]string, error) {
	var b bytes.Buffer
	escapeSeparator := false
	keyValList := []string{}
	res := map[string]string{}
	for _, v := range input {
		if v == '"' || v == '\'' {
			if escapeSeparator {
				escapeSeparator = false
			} else {
				escapeSeparator = true
			}
		} else if v == ' ' {
			if escapeSeparator {
				b.WriteRune(v)
			} else {
				keyValList = append(keyValList, b.String())
				b.Reset()
			}
		} else if v == '=' {
			if escapeSeparator {
				return nil, fmt.Errorf(`illegal character = after "`)
			} else {
				keyValList = append(keyValList, b.String())
				b.Reset()
			}
		} else {
			b.WriteRune(v)
		}
	}
	if escapeSeparator {
		return nil, fmt.Errorf(`need to close "`)
	}
	if b.Len() > 0 || len(keyValList) > 0 {
		keyValList = append(keyValList, b.String())
	}
	if len(keyValList) != 0 && len(keyValList)%2 != 0 {
		return nil, fmt.Errorf("invalid number of key and value, should be odd")
	}

	key := ""
	for k, v := range keyValList {
		if k%2 == 0 {
			key = v
		} else {
			res[key] = v
		}
	}
	return res, nil
}
