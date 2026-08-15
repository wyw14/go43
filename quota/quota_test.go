package quota

import "testing"

func TestCloneAndTags(t *testing.T) {
	original := QuotaGrant{ID:"a", Tags:map[string]string{"source":"api"}}
	cloned := original.Clone(); original.Tags["source"] = "changed"
	if cloned.Tags["source"] != "api" { t.Fatal("clone retained caller map") }
	var fresh QuotaGrant; fresh.AddTag("owner","ops")
	if fresh.Tags["owner"] != "ops" { t.Fatal("tag was not recorded") }
}

func TestValidateAllowsZeroPriority(t *testing.T) {
	if err := Validate(QuotaGrant{ID:"a", Priority:0}); err != nil { t.Fatalf("zero priority should be valid: %v",err) }
}
