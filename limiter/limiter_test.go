package limiter

import (
	"bytes"
	"context"
	"errors"
	"testing"
	"time"
	"example.com/go43/quota"
	"example.com/go43/window"
)

func TestProcessCancellationAndDeduplication(t *testing.T) {
	ctx,cancel:=context.WithCancel(context.Background()); cancel(); calls:=0
	err:=Process(ctx,[]quota.QuotaGrant{{ID:"a",Priority:1}},func(quota.QuotaGrant)error{calls++;return nil})
	if !errors.Is(err,context.Canceled)||calls!=0 { t.Fatalf("cancel result err=%v calls=%d",err,calls) }
	calls=0; err=Process(context.Background(),[]quota.QuotaGrant{{ID:"a",Priority:1},{ID:"a",Priority:1}},func(quota.QuotaGrant)error{calls++;return nil})
	if err!=nil||calls!=1 { t.Fatalf("dedupe err=%v calls=%d",err,calls) }
}

func TestExportTransitionDefaultAndExpiry(t *testing.T) {
	want:=errors.New("encode"); if err:=Export(context.Background(),func()([]byte,error){return nil,want},&bytes.Buffer{}); !errors.Is(err,want){t.Fatalf("lost encoder error: %v",err)}
	v:=quota.QuotaGrant{ID:"a",State:"new"}; if err:=Transition(&v,"done"); err==nil||v.State!="new" { t.Fatalf("invalid transition changed state: %#v err=%v",v,err) }
	off:=false; if ResolveEnabled(&off,true) { t.Fatal("explicit false was replaced") }
	s:= window.New(); at:=time.Unix(100,0); s.Save(quota.QuotaGrant{ID:"a",UpdatedAt:at}); if Expire(s,at)!=0 { t.Fatal("exact cutoff expired") }
}
