package vps_actions

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/g1eng/savac/pkg/vps"
)

type VpsActionGenerator struct {
	ApiClient  *vps.SavaClient
	OutputType int
	NoHeader   bool
}

func isRegexPatternExpr(pat string) bool {
	if len(pat) == 0 {
		return false
	}
	return len(strings.TrimLeft(pat, "/")) == len(pat)-1 && len(strings.TrimRight(pat, "/")) == len(pat)-1
}

func toRegex(pat string) (*regexp.Regexp, error) {
	if !isRegexPatternExpr(pat) {
		return nil, fmt.Errorf("%s is not a regular expression argument", pat)
	}
	pat = strings.TrimLeft(strings.TrimRight(pat, "/"), "/")
	return regexp.CompilePOSIX(pat)
}
