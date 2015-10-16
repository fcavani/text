// Copyright 2015 Felipe A. Cavani. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// Start date:		2013-07-04
// Last modification:	2013-x

package text

import (
	"net/url"
	"strings"
)

//EscapeCommaSeparated escapes the args and make a comma separeted list with it.
func EscapeCommaSeparated(in ...string) string {
	var out string
	for i, str := range in {
		escaped := strings.Replace(url.QueryEscape(str), "%2F", "%252F", -1)
		escaped = strings.Replace(escaped, "\"", "%22", -1)
		escaped = strings.Replace(escaped, " ", "%20", -1)
		out += escaped
		if i < len(in)-1 {
			out += ","
		}
	}
	return out
}

func Escape(str string) string {
	escaped := strings.Replace(url.QueryEscape(str), "%2F", "%252F", -1)
	escaped = strings.Replace(escaped, "\"", "%22", -1)
	escaped = strings.Replace(escaped, " ", "%20", -1)
	return escaped
}
