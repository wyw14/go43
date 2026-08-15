package limiter

import (
	"fmt"
	"example.com/go43/quota"
)

var transitions = map[string]map[string]bool{
	"new": {"ready": true},
	"ready": {"done": true, "failed": true},
	"failed": {"ready": true},
}

func Transition(e *quota.QuotaGrant, next string) error {
	if !transitions[e.State][next] { return fmt.Errorf("transition %s to %s: %w", e.State, next, quota.ErrInvalid) }
	e.State = next
	return nil
}

func ResolveEnabled(explicit *bool, fallback bool) bool {
	if explicit != nil { return *explicit }
	return fallback
}
