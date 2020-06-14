package lib

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
	"time"
)

func FindLastDuration(haystack, pattern string) (time.Duration, error) {
	reg, err := regexp.Compile(strings.ReplaceAll(pattern, "?", "(\\d+:\\d+:[\\d\\.]+)"))
	if err != nil {
		return 0, err
	}
	if !reg.MatchString(haystack) {
		return 0, errors.New("not found")
	}


	results := reg.FindAllStringSubmatch(haystack, -1)
	lastMatch := results[len(results)-1][1]
	parts := strings.Split(lastMatch, ":")
	duration, err := time.ParseDuration(fmt.Sprintf("%sh%sm%ss", parts[0], parts[1], parts[2]))
	if err != nil {
		return 0, err
	}
	return duration, nil
}
