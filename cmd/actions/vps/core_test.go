package vps_actions

var (
	validPatterns = []string{
		"/^thisis-okay$/",
		"/[thisis].okay/",
		"/[t-x].0kay$/",
	}
	invalidRegex    = "/them(+/"
	inValidPatterns = []string{
		"them+",
		invalidRegex,
		"/them**",
		"them**/",
	}
)

func (v *VpsActionSuite) Test_isRegexPatternString() {
	for _, p := range validPatterns {
		if isRegexPatternExpr(p) == false {
			v.Failf("fail", "patterm %s is invalid", p)
		}
	}
	for _, p := range inValidPatterns {
		if isRegexPatternExpr(p) == true {
			if invalidRegex == p {
				continue
			}
			v.Failf("fail", "patterm %s is valid", p)
		}
	}
}

func (v *VpsActionSuite) Test_toRegex() {
	for _, p := range validPatterns {
		if _, err := toRegex(p); err != nil {
			v.Failf("fail", "patterm %s is invalid", p)
		}
	}
	for _, p := range inValidPatterns {
		if _, err := toRegex(p); err == nil {
			v.Failf("fail", "patterm %s is valid", p)
		}
	}
}
