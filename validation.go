// Copyright 2015 Felipe A. Cavani. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// Start date:		2014-08-08
// Last modification:	2014-x

package text

import (
	"net/url"
	"regexp"
	"unicode"

	"github.com/fcavani/e"
	uni "github.com/fcavani/util/unicode"
	"github.com/golang/exp/utf8string"
)

var MinPassLen = 8
var MaxPassLen = 100
var MinEmailLen = 6
var MaxEmailLen = 60

const ErrInvLengthNumber = "wrong number of digits"
const ErrInvDigit = "caracter isn't a digit"
const ErrInvNumberChars = "invalid number of characters"
const ErrInvCharacter = "invalid character"
const ErrInvalidPassLength = "password length is invalid"
const ErrInvalidPassChar = "invalid password character"
const ErrInvEmailLength = "email length is invalid"
const ErrCantCheckEmail = "can't check the e-mail address"
const ErrInvEmailString = "invalid e-mail address"
const ErrInvalidChar = "character is invalid"

func CheckNumber(number string, min, max int) error {
	if len(number) < min || len(number) > max {
		return e.New(ErrInvLengthNumber)
	}
	for _, v := range number {
		if !unicode.IsDigit(v) {
			return e.New(ErrInvDigit)
		}
	}
	return nil
}

func CheckLetters(text string, min, max int) error {
	if len(text) < min || len(text) > max {
		return e.New(ErrInvNumberChars)
	}
	for _, v := range text {
		if !uni.IsLetter(v) {
			return e.Push(e.New(ErrInvCharacter), e.New("the character '%v' is invalid", string([]byte{byte(v)})))
		}
	}
	return nil
}

func CheckLettersNumber(text string, min, max uint64) error {
	if uint64(len(text)) < min || uint64(len(text)) > max {
		return e.New(ErrInvNumberChars)
	}
	for _, v := range text {
		if !uni.IsLetter(v) && !unicode.IsDigit(v) {
			return e.Push(e.New(ErrInvCharacter), e.New("the character '%v' is invalid", string([]byte{byte(v)})))
		}
	}
	return nil
}

func CheckText(text string, min, max int) error {
	if len(text) < min || len(text) > max {
		return e.New(ErrInvNumberChars)
	}
	for _, v := range text {
		if !uni.IsLetter(v) && !unicode.IsDigit(v) && v != '\n' && v != ' ' && v != '`' && v != '~' && v != '!' && v != '@' && v != '#' && v != '$' && v != '%' && v != '^' && v != '&' && v != '*' && v != '(' && v != ')' && v != '_' && v != '-' && v != '+' && v != '=' && v != '{' && v != '}' && v != '[' && v != ']' && v != '|' && v != '\\' && v != ':' && v != ';' && v != '"' && v != '\'' && v != '?' && v != '/' && v != ',' && v != '.' {
			return e.Push(e.New(ErrInvCharacter), e.New("the character '%v' is invalid", string([]byte{byte(v)})))
		}
	}
	return nil
}

// Check the user password. Graphics character are allowed. See unicode.IsGraphic.
func CheckPassword(pass string, min, max int) error {
	if len(pass) < min || len(pass) > max {
		return e.New(ErrInvalidPassLength)
	}
	for _, r := range pass {
		if !unicode.IsGraphic(r) {
			return e.New(ErrInvalidPassChar)
		}
	}
	return nil
}

func CheckEmail(email string) error {
	if len(email) < MinEmailLen || len(email) > MaxEmailLen {
		return e.New(ErrInvEmailLength)
	}
	r, err := regexp.Compile(`([a-zA-Z0-9]+)([.-_][a-zA-Z0-9]+)*@([a-zA-Z0-9]+)([.-_][a-zA-Z0-9]+)*`)
	if err != nil {
		return e.Push(e.New(err), ErrCantCheckEmail)
	}
	if email != r.FindString(email) {
		return e.New(ErrInvEmailString)
	}
	return nil
}

func CheckName(name string, min, max int) error {
	if len(name) < min || len(name) > max {
		return e.New(ErrInvNumberChars)
	}
	for _, v := range name {
		if !uni.IsLetter(v) && !unicode.IsDigit(v) && v != ' ' && v != '`' && v != '~' && v != '!' && v != '@' && v != '#' && v != '$' && v != '%' && v != '^' && v != '&' && v != '*' && v != '(' && v != ')' && v != '_' && v != '-' && v != '+' && v != '=' && v != '{' && v != '}' && v != '[' && v != ']' && v != '|' && v != '\\' && v != ':' && v != ';' && v != '"' && v != '\'' && v != '?' && v != '/' && v != ',' && v != '.' {
			return e.Push(e.New(ErrInvCharacter), e.New("the character '%v' is invalid", string([]byte{byte(v)})))
		}
	}
	return nil
}

func CheckNameWithoutSpecials(name string, min, max int) error {
	if len(name) < min || len(name) > max {
		return e.New(ErrInvNumberChars)
	}
	for _, v := range name {
		if !uni.IsLetter(v) && !unicode.IsDigit(v) && v != ' ' && v != '&' && v != '(' && v != ')' && v != '-' && v != ':' && v != '/' && v != ',' && v != '.' && v != '_' {
			return e.Push(e.New(ErrInvCharacter), e.New("the character '%v' is invalid", string([]byte{byte(v)})))
		}
	}
	return nil
}

func CheckFileName(nome string, min, max int) error {
	if len(nome) < min || len(nome) > max {
		return e.New(ErrInvNumberChars)
	}
	for _, v := range nome {
		if !uni.IsLetter(v) && !unicode.IsDigit(v) && v != ' ' && v != '-' && v != ':' && v != ',' && v != '.' && v != '_' {
			return e.Push(e.New(ErrInvCharacter), e.New("the character '%v' in filename is invalid", string([]byte{byte(v)})))
		}
	}
	return nil
}

func ValidateRedirect(redirect string, min, max int) error {
	utf8name := utf8string.NewString(redirect)
	len := utf8name.RuneCount()
	if len < min || len > max {
		return e.New(e.ErrInvalidLength)
	}
	for _, s := range utf8name.Slice(1, len) {
		if !uni.IsLetter(s) && !unicode.IsDigit(s) && s != '/' && s != '-' && s != '_' && s != '?' && s != '&' && s != '=' && s != '%' && s != '*' && s != '+' && s != ' ' && s != ',' {
			println("redirect:", string([]byte{byte(s)}))
			return e.Push(e.New("the character '%v' in redirect is invalid", string([]byte{byte(s)})), e.New("invalid redirect"))
		}
	}
	return nil
}

func CheckSearch(query string, min, max int) error {
	if len(query) < min || len(query) > max {
		return e.New(ErrInvNumberChars)
	}
	for _, v := range query {
		if !uni.IsLetter(v) && !unicode.IsDigit(v) && v != '@' && v != '.' && v != '-' && v != '_' && v != ' ' && v != '+' && v != '"' {
			return e.New(ErrInvCharacter)
		}
	}
	return nil
}

const ErrInvUrl = "invalid url"
const ErrNoScheme = "url without scheme"

func CheckUrl(rawurl string, min, max int) error {
	if len(rawurl) < min || len(rawurl) > max {
		return e.Push(e.New(ErrInvUrl), e.New("invalid url length"))
	}
	for _, v := range rawurl {
		if !uni.IsLetter(v) && !unicode.IsDigit(v) && v != '/' && v != ':' && v != '[' && v != ']' && v != '?' && v != '@' && v != '.' && v != '-' && v != '_' && v != ' ' && v != '+' && v != '%' && v != '#' {
			return e.Push(e.New(ErrInvUrl), e.New("the character '%v' in redirect is invalid", string([]byte{byte(v)})))
		}
	}
	return nil
}

// CleanUrl check the characteres in url and parser it with url.Parse.
// If url is ok return one string with it or if the scheme is missing
// return the url and an error.
func CleanUrl(rawurl string, min, max int) (string, error) {
	err := CheckUrl(rawurl, min, max)
	if err != nil {
		return "", e.Forward(err)
	}
	u, err := url.Parse(rawurl)
	if err != nil {
		return "", e.Push(e.New(ErrInvUrl), err)
	}
	if u.Scheme == "" {
		return u.String(), e.New(ErrNoScheme)
	}
	return u.String(), nil
}

func CheckDomain(domain string) error {
	for _, v := range domain {
		if !uni.IsLetter(v) && !unicode.IsDigit(v) && v != '.' {
			return e.Push(e.New("invalid domain name"), e.New("the character '%v' in redirect is invalid", string([]byte{byte(v)})))
		}
	}
	return nil
}
