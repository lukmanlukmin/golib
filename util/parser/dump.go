package parser

import (
	"bytes"
	"encoding/json"

	"github.com/lukmanlukmin/golib/log"
)

func DumpToString(v interface{}) string {
	str, ok := v.(string)
	if !ok {
		buff := &bytes.Buffer{}
		if err := json.NewEncoder(buff).Encode(v); err != nil {
			log.WithError(err).Errorln("[dumper] failed to json encode value")
			return ""
		}

		return buff.String()
	}
	return str
}
