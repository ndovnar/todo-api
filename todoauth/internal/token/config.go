package token

import "time"

type Config struct {
	AccessTokenDuration  time.Duration `required:"true"`
	RefreshTokenDuration time.Duration `required:"true"`
}
