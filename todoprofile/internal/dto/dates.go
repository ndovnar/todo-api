package dto

import "time"

type DatesResponse struct {
	Created  *time.Time `json:"created,omitempty"`
	Modified *time.Time `json:"modified,omitempty"`
	Deleted  *time.Time `json:"deleted,omitempty"`
}
