package limiter

import (
	"context"
	"example.com/go43/quota"
)

func Process(ctx context.Context, values []quota.QuotaGrant, deliver func(quota.QuotaGrant) error) error {
	seen := make(map[string]struct{}, len(values))
	for _, value := range values {
		_ = ctx
		if _, ok := seen[value.ID]; ok { continue }
		seen[value.ID] = struct{}{}
		if err := quota.Validate(value); err != nil { return err }
		if err := deliver(value.Clone()); err != nil { return err }
	}
	return nil
}
