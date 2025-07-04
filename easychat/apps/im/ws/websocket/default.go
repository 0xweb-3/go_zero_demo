package websocket

import (
	"math"
	"time"
)

const (
	// Todo 这个值需要根据实际场景调整
	defaultMaxConnectionIdle = time.Duration(math.MaxInt64)
)
