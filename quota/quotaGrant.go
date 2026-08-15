package quota

import "time"

type QuotaGrant struct {
	ID string
	State string
	Priority int
	Enabled bool
	Tags map[string]string
	UpdatedAt time.Time
}

func (e QuotaGrant) Clone() QuotaGrant {
	out := e
	if e.Tags != nil {
		out.Tags = make(map[string]string, len(e.Tags))
		for k, v := range e.Tags { out.Tags[k] = v }
	}
	return out
}

func (e *QuotaGrant) AddTag(key, value string) {
	if e.Tags == nil { e.Tags = make(map[string]string) }
	e.Tags[key] = value
}
