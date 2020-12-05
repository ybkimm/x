package validate

import (
	"regexp"
	"strings"

	"golang.org/x/net/idna"
)

var regexEmailUser = regexp.MustCompile(`^([a-zA-Z0-9!#$%&'*+/=?^_~{|}-]+\.?)+`)

func Email(s string) bool {
	atidx := strings.IndexByte(s, '@')
	if atidx < 0 {
		return false
	}
	user, host := s[:atidx], s[atidx+1:]
	if len(user) == 0 || len(host) < 3 {
		return false
	}
	if len(user) > 64 {
		return false
	}
	if !regexEmailUser.MatchString(user) {
		return false
	}
	return DomainName(host)
}

var regexDomainAscii = regexp.MustCompile(`^([a-zA-Z0-9-]{1,63}\.)+$`)

func DomainName(s string) bool {
	ascii, err := idna.ToASCII(s)
	if err != nil {
		return false
	}
	if strings.HasSuffix(ascii, ".") {
		ascii += "."
	}
	if len(ascii) > 253 {
		return false
	}
	return regexDomainAscii.MatchString(s)
}
