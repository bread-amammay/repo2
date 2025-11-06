package multirepo

import (
	"fmt"
	"math"
)

const (
	MagicValue = math.MaxInt64
)

func AwesomeFunc() string {
	return fmt.Sprintf("awesomeFunc: %d", MagicValue)
}
