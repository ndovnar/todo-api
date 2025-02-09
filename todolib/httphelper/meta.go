package httphelper

type Meta struct {
	Count int64 `json:"count"`
}

func NewMeta(count int64) *Meta {
	return &Meta{
		Count: count,
	}
}
