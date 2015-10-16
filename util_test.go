// Copyright 2015 Felipe A. Cavani. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// Start date:		2015-01-07
// Last modification:	2015-x

package text

import (
	"testing"
)

type test struct {
	data   string
	length int
	result string
}

var tests []test = []test{
	{"qqqq", 1, "qqqq"},
	{"qqqq", 2, "qqqq"},
	{"qqqqqq", 2, "qq..."},
	{"qqqq", 10, "qqqq"},
	{"qq qq", 4, "qq qq"},
	{"ab cdefg", 4, "ab..."},
	{"Lorem status quo sit amet lapsus sustos.", 300, "Lorem status quo sit amet lapsus sustos."},
	{"lorem sit: vchj", 10, "lorem sit..."},
	{"lorem sit... vchj", 12, "lorem sit..."},
	{"Lorem status quo sit amet lapsus sustos.", 7, "Lorem..."},
}

func TestReticence(t *testing.T) {
	for i, test := range tests {
		str := Reticence(test.data, test.length)
		//t.Log(i, test.data, test.length, test.result, str)
		if str != test.result {
			t.Fatal(i, test.data, test.length, test.result, str)
		}
	}
}

type testTailStruct struct {
	msg  string
	head string
	mid  string
	tail string
	fail bool
}

var testTail []testTailStruct = []testTailStruct{
	{"a|b|c", "a", "b", "c", false},
	{"aaaa|bbbb|cccc", "aaaa", "bbbb", "cccc", false},
	{"aaaa|bbbb|", "aaaa", "bbbb", "", true},
	{"aaaa|bbbb", "aaaa", "bbbb", "", true},
	{"aaaa", "aaaa", "", "", true},
	{"a|a|aa|bbbb|cccc", "a|a|aa", "bbbb", "cccc", false},
}

func TestTail(t *testing.T) {
	for i, test := range testTail {
		headmid, tail_ := HeadTail(test.msg, "|")
		head, mid := HeadTail(headmid, "|")
		t.Logf("\"%v\", \"%v\", \"%v\"", string(head), string(mid), string(tail_))
		if (string(head) == "" || string(mid) == "" || string(tail_) == "") && !test.fail {
			t.Fatal("Failed:", i, string(head), string(mid), string(tail_))
		}
		if (string(head) != test.head || string(mid) != test.mid || string(tail_) != test.tail) && !test.fail {
			t.Fatal("Not spected:", i, string(head), string(mid), string(tail_))
		}
	}
}
