package params

import (
	"fmt"
	"regexp"
	"strconv"
)

type Tag interface {
	GetName() string
	Validate(v string) error
}

type httpTag struct{}

var Http httpTag

func (mail httpTag) GetName() string { return "http" }
func (mail httpTag) Validate(v string) error {
	return nil
}

type mailTag struct{}

var Mail mailTag

func (mail mailTag) GetName() string { return "mail" }
func (mail mailTag) Validate(v string) error {
	matched, err := regexp.MatchString("^[a-zA-Z0-9_+-]+(.[a-zA-Z0-9_+-]+)*@([a-zA-Z0-9][a-zA-Z0-9-]*[a-zA-Z0-9]*\\.)+[a-zA-Z]{2,}$", v)
	if err != nil {
		return fmt.Errorf("mail: %w", err)
	}
	if !matched {
		return fmt.Errorf("mail: invalid format")
	}
	return nil
}

type positiveTag struct{}

var Positive positiveTag

func (pos positiveTag) GetName() string { return "positive" }
func (pos positiveTag) Validate(v string) error {
	num, err := strconv.ParseInt(v, 10, 64)
	if err != nil {
		return fmt.Errorf("positive: %w", err)
	}
	if num < 0 {
		return fmt.Errorf("positive: not positive %d", num)
	}
	return nil
}

var tags []Tag = []Tag{Mail, Positive, Http}
