package model

import "time"

type Dates struct {
	Created  *time.Time `bson:"created,omitempty"`
	Modified *time.Time `bson:"modified,omitempty"`
	Deleted  *time.Time `bson:"deleted,omitempty"`
}
