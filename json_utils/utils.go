// Author: Paweł Konopko
// License: MIT

package json_utils

import "encoding/json"

func Encode(s interface{}) string {
	b, err := json.Marshal(s)
	if err != nil {
		return "!ERR{}"
	}
	return string(b)
}

func EncodePretty(s interface{}) string {
	b, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "!ERR{}"
	}
	return string(b)
}