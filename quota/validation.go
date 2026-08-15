package quota

import "errors"

var ErrInvalid = errors.New("invalid quotawindow value")

func Validate(e QuotaGrant) error {
	if e.ID == "" { return ErrInvalid }
	if e.Priority < 0 { return ErrInvalid }
	return nil
}
