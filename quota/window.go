package quota

import "time"

type Window struct {
	limit, used int
	start       time.Time
	span        time.Duration
}

func New(limit int, span time.Duration, start time.Time) *Window {
	return &Window{limit: limit, start: start, span: span}
}
func (w *Window) Allow(n int, now time.Time) bool {
	if now.Sub(w.start) >= w.span {
		w.start = now
		w.used = 0
	}
	if w.used+n > w.limit {
		return false
	}
	w.used += n
	return true
}
func (w *Window) Used() int { return w.used }
