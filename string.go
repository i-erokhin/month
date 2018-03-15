package month

import (
	"fmt"
	"regexp"
	"strconv"
)

var periodRE = regexp.MustCompile(`^(\d{4})-(\d{2})$`)

type String string

func (s String) String() string {
	return string(s)
}

func (s String) Validate() (errs []string) {
	if s == "" {
		return append(errs, "empty value")
	}
	sub := periodRE.FindStringSubmatch(s.String())
	if len(sub) == 0 {
		return append(errs, fmt.Sprintf("bad period format: %q", s))
	}
	return
}

func (s String) Period() (*Month, error) {
	sub := periodRE.FindStringSubmatch(s.String())
	if len(sub) == 0 {
		return nil, fmt.Errorf("bad period format: %q", s)
	}
	y, err := strconv.Atoi(sub[1])
	if err != nil {
		return nil, err
	}
	m, err := strconv.Atoi(sub[2])
	if err != nil {
		return nil, err
	}

	return &Month{y, m}, nil
}

func (s String) PeriodMust() *Month {
	period, err := s.Period()
	if err != nil {
		panic(err)
	}
	return period
}
