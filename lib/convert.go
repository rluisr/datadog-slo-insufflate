/*
 * Copyright (c) 2022. rluisr Takuya Hasegawa
 */

package lib

import (
	"fmt"
)

func ConvertF64ToString(f float64) string {
	return fmt.Sprintf("%.2f", f)
}
