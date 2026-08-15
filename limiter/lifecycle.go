package limiter

import (
	"time"
	"example.com/go43/window"
)

func Expire(store *window.Store, cutoff time.Time) int { return store.DeleteBefore(cutoff) }
