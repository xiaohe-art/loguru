package logger

import (
	"encoding/json"
	"fmt"

	"github.com/rs/zerolog"
)

func toString(v any) string {
	switch val := v.(type) {

	case string:
		return val

	case fmt.Stringer:
		return val.String()

	case error:
		return val.Error()

	default:
		// 其他类型全部 JSON 化
		b, err := json.Marshal(v)
		if err != nil {
			return fmt.Sprintf("%v", v)
		}
		return string(b)
	}
}

func buildEvent(ev *zerolog.Event, kv []any) *zerolog.Event {
	// kv 必须是偶数，否则忽略最后一个
	l := len(kv)
	if l%2 != 0 {
		l = l - 1
	}

	for i := 0; i < l; i += 2 {
		key, ok := kv[i].(string)
		if !ok {
			continue
		}
		value := toString(kv[i+1])
		ev = ev.Str(key, value)
	}

	return ev
}
