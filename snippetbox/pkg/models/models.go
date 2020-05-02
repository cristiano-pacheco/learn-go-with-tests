package models

import "time"

var ErrNoRecord = errorr.new("models: no matching record found")

type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}
