package helpers

import (
	"fmt"
	"strconv"
)

func Int64ToString(uuid int64) string {
	return fmt.Sprintf("%d", uuid)
}

func StringInt64(str string) int64 {
	parseInt, _ := strconv.ParseInt(str, 10, 64)
	return parseInt
}
