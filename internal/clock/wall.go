package clock

import "time"

type Wall struct{}

func NewWall() *Wall { return &Wall{} }

func (w *Wall) Now() time.Time { return time.Now() }

func (w *Wall) After(d time.Duration) <-chan time.Time {
	return time.After(d)
}