package repo1

import (
	"fmt"
	"math"
)

const (
	MagicValue = math.MaxInt64
)

func AwesomeFunc() string {
	return fmt.Sprintf("awesomeFuncnewversion: %d", MagicValue)
}
