package model

import "time"

type Dates struct {
	Created  *time.Time
	Modified *time.Time
	Deleted  *time.Time
}
