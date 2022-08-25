/*
 * Copyright (c) 2022. rluisr Takuya Hasegawa
 */

package datadog

import (
	"fmt"
	"net/http"
	"os"
)

// echoError expected err is always not nil
func echoError(r *http.Response) {
	fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
}
