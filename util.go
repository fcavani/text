// Copyright 2015 Felipe A. Cavani. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// Start date:		2014-12-04
// Last modification:	2014-x

package text

import (
	"strings"
	"unicode"
)

func FirstCaps(str string) string {
	ret := ""
	if len(str) >= 1 {
		ret = string(unicode.ToUpper(rune(str[0])))
	}
	if len(str) >= 2 {
		ret += str[1:]
	}
	return ret
}

// Reticence trucate the string in the space or on pontuation mark and put
// reticences in the resulting string.
func Reticence(str string, length int) string {
	if length > len(str) {
		return str
	}
	var i int
F:
	for i = len(str) - 1; i >= 0; i-- {
		switch str[i] {
		case ' ', ',', '?', ';', ':', '\'', '"', '!':
			if i <= length {
				break F
			}
		case '.':
			if i-2 >= 0 {
				s := str[i-2 : i]
				if s == ".." {
					i = i - 2
					if i <= length {
						break F
					}
				}
			}
			if i <= length {
				break F
			}
		}
	}
	if i-1 > 0 {
		switch str[i-1] {
		case ' ', ',', '?', ';', ':', '\'', '"', '!':
			i--
		case '.':
			if i-2 > 0 && str[i-2:i] == ".." {
				i -= 3
			}
		}
	}
	if i >= 2 {
		if i+3 >= len(str) {
			return str
		}
		return str[:i] + "..."
	}
	if length >= 2 && length < len(str) {
		if length+3 >= len(str) {
			return str
		}
		return str[:length] + "..."
	}
	return str
}

func HeadTail(in, sep string) (head, tail string) {
	sepi := strings.LastIndex(in, sep)
	if sepi == -1 || sepi >= len(in)-1 {
		head = in
		tail = ""
		return
	}
	tail = in[sepi+1:]
	head = in[:sepi]
	return
}
