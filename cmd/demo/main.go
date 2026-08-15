package main

import (
	"context"
	"fmt"
	"time"
	"example.com/go43/quota"
	"example.com/go43/limiter"
)

func main() {
	v := quota.QuotaGrant{ID:"demo", State:"new", UpdatedAt:time.Now()}
	err := limiter.Process(context.Background(), []quota.QuotaGrant{v}, func(got quota.QuotaGrant) error { fmt.Println(got.ID); return nil })
	if err != nil { panic(err) }
}
