package debug

import (
	"fmt"
	"time"
)


func Log(msg string, args ...interface{}) {
    ts := time.Now().Format("15:04:05.00")
    fmt.Printf("[%s] DEBUG: %s\n", ts, fmt.Sprintf(msg, args...))

}
