package matcher

import "regexp"

func New() map[string]*regexp.Regexp {
	matchers := make(map[string]*regexp.Regexp)

	matchers[":date"] = regexp.MustCompile("((Mon|Tue|Wed|Thu|Fri|Sat|Sun), (0\\d|1\\d|2\\d|3[01]) (Jan|Feb|Mar|Apr|May|Jun|Jul|Aug|Sep|Oct|Nov|Dec) 2\\d{3} (0\\d|1\\d|2[0-3]):(0\\d|1\\d|2\\d|3\\d|4\\d|5\\d):(0\\d|1\\d|2\\d|3\\d|4\\d|5\\d) (A|M|N|Y|Z|UT|GMT|[A-Z]{3}|[+-](0\\d|1[012])))")
	matchers[":b62:22"] = regexp.MustCompile("([0-9A-Za-z]{22})")
	matchers[":uuid"] = regexp.MustCompile("([[:xdigit:]]{8}-[[:xdigit:]]{4}-[[:xdigit:]]{4}-[[:xdigit:]]{4}-[[:xdigit:]]{12})")

	return matchers
}
