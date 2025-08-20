package utility

import (
	"strconv"
)

func String2Int64(s string) (i int64) {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		panic(err)
	}
	return
}
