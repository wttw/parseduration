package parseduration

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const (
	nanosecond  float64 = 1
	microsecond         = 1000 * nanosecond
	millisecond         = 1000 * microsecond
	second              = 1000 * millisecond
	minute              = 60 * second
	hour                = 60 * minute
	day                 = 24 * hour
	week                = 7 * day
	month               = 30 * day
	year                = 365 * day
)

var alt8601a = regexp.MustCompile(`^P(\d{4})-(\d{2})-(\d{2})(?:T(\d{2}):(\d{2}):(\d{2}(?:[.]\d+)))?`)
var alt8601b = regexp.MustCompile(`^P(\d{4})(\d{2})(\d{2})(?:T(\d{2})(\d{2})(\d{2}(?:[.]\d+)))?`)

var days8601 = regexp.MustCompile(`^P([+-]?[0-9.]+[YMWD])*(T.*)?$`)
var seconds8601 = regexp.MustCompile(`^T([+-]?[0-9.]+[HMS])*`)

// Parse8601 parses an ISO 8601-ish format duration string
func Parse8601(s string) (time.Duration, error) {
	fmt.Printf("Parsing '%s'\n", s)
	if !strings.HasPrefix(s, "P") {
		return 0, errors.New("Invalid ISO8601 duration")
	}
	s = strings.Replace(s, ",", ".", -1)

	alta := alt8601a.FindStringSubmatch(s)
	if alta != nil {
		return alt8601(alta)
	}
	altb := alt8601b.FindStringSubmatch(s)
	if altb != nil {
		return alt8601(altb)
	}

	parts := strings.Split(strings.TrimPrefix(s, "P"), "T")

	if len(parts) < 1 || len(parts) > 2 {
		return 0, errors.New("Invalid ISO8601 duration")
	}

	ret, err := parse8601Date(parts[0])
	if err != nil {
		return 0, err
	}
	if len(parts) == 2 {
		r, err := parse8601Hour(parts[1])
		if err != nil {
			return 0, err
		}
		ret = ret + r
	}
  return time.Duration(ret), nil
}

func parse8601Date(s string) (int64, error) {

}

func parse8601Hour(s string) (int64, error) {
  
}
	t := strings.Index()

	daypart := days8601.FindStringSubmatch(s)
	if daypart == nil {
		return 0, errors.New("Invalid ISO8601 duration")
	}

	var ret int64

	for _, part := range daypart[1:] {
		if part == "" {
			continue
		}
		if part[:1] == "T" {
			fmt.Printf("tpart='%s'\n", part)
			secondspart := seconds8601.FindStringSubmatch(part)
			if secondspart == nil {
				return 0, errors.New("Invalid ISO8601 duration")
			}
			fmt.Printf("secondspart=%+v\n", secondspart)
			for _, spart := range secondspart[1:] {
				fmt.Printf("spart='%s'\n", spart)
				sz := len(spart)
				num := spart[:sz-1]
				switch spart[sz-1:] {
				case "H":
					hours, err := strconv.ParseFloat(num, 64)
					if err != nil {
						return time.Duration(0), err
					}
					ret = ret + int64(hours*hour)
				case "M":
					minutes, err := strconv.ParseFloat(num, 64)
					if err != nil {
						return time.Duration(0), err
					}
					ret = ret + int64(minutes*minute)
				case "S":
					seconds, err := strconv.ParseFloat(num, 64)
					if err != nil {
						return time.Duration(0), err
					}
					ret = ret + int64(seconds*second)
				}
			}
		} else {
			sz := len(part)
			num := part[:sz-1]
			switch part[sz-1:] {
			case "Y":
				years, err := strconv.ParseFloat(num, 64)
				if err != nil {
					return time.Duration(0), err
				}
				ret = ret + int64(years*year)
			case "M":
				months, err := strconv.ParseFloat(num, 64)
				if err != nil {
					return time.Duration(0), err
				}
				ret = ret + int64(months*month)
			case "W":
				weeks, err := strconv.ParseFloat(num, 64)
				if err != nil {
					return time.Duration(0), err
				}
				ret = ret + int64(weeks*week)
			case "D":
				days, err := strconv.ParseFloat(num, 64)
				if err != nil {
					return time.Duration(0), err
				}
				ret = ret + int64(days*day)
			}
		}
	}
	return time.Duration(ret), nil
}

func alt8601(parts []string) (time.Duration, error) {
	var ret int64
	if len(parts) >= 4 {
		years, err := strconv.ParseInt(parts[1], 10, 64)
		if err != nil {
			return 0, err
		}
		ret = ret + years*int64(year)

		months, err := strconv.ParseInt(parts[2], 10, 64)
		if err != nil {
			return 0, err
		}
		ret = ret + months*int64(month)

		days, err := strconv.ParseInt(parts[3], 10, 64)
		if err != nil {
			return 0, err
		}
		ret = ret + days*int64(day)
	}
	if len(parts) >= 7 {
		hours, err := strconv.ParseInt(parts[4], 10, 64)
		if err != nil {
			return 0, err
		}
		ret = ret + hours*int64(hour)

		minutes, err := strconv.ParseInt(parts[5], 10, 64)
		if err != nil {
			return 0, err
		}
		ret = ret + minutes*int64(minute)

		seconds, err := strconv.ParseFloat(parts[6], 64)
		if err != nil {
			return 0, err
		}
		ret = ret + int64(seconds*second)

	}
	return time.Duration(ret), nil
}
