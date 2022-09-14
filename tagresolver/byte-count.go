package tagresolver

import (
	"fmt"
	"math"
	"strconv"
)

func humanReadableByteCount(b int) string {
	if b <= 0 {
		return ""
	}
	var unit float64 = 1000
	bytes := float64(b)
	if bytes < unit {
		return strconv.Itoa(b) + " B"
	}
	exp := (int)(math.Log(bytes) / math.Log(unit))
	pre := string("kMGTPE"[exp-1])
	return fmt.Sprintf("%.1f %sB", bytes/math.Pow(unit, float64(exp)), pre)
}
