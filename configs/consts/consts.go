package consts

import (
	"fmt"
	"time"
)

const (
	ConfigEnv      = ConfigDevFile
	ConfigDevFile  = "./configs/config_dev.yaml"
	ConfigProdFile = "./configs/config_prod.yaml"
)

var (
	TokenExpired     = fmt.Errorf("Token is expired")
	TokenNotValidYet = fmt.Errorf("Token not active yet")
	TokenMalformed   = fmt.Errorf("That's not even a token")
	TokenInvalid     = fmt.Errorf("Couldn't handle this token")
)

const (
	// Short writes for common usage durations.

	D  = 24 * time.Hour
	H  = time.Hour
	M  = time.Minute
	S  = time.Second
	MS = time.Millisecond
	US = time.Microsecond
	NS = time.Nanosecond

	Jan  = 1
	Feb  = 2
	Mar  = 3
	Apr  = 4
	May  = 5
	Jun  = 6
	Jul  = 7
	Aug  = 8
	Sept = 9
	Oct  = 10
	Nov  = 11
	Dec  = 12

	January   = 1
	February  = 2
	March     = 3
	April     = 4
	June      = 6
	July      = 7
	August    = 8
	September = 9
	October   = 10
	November  = 11
	December  = 12
)

var (
	TopicOfWarn = "warn"
)
