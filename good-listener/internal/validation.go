package internal

import (
	"errors"
	"fmt"
	"strings"
	"unicode"
)

func CheckPortCorrect(port int) error {
	if port < 0 || port > 65535 {
		return fmt.Errorf("port must be integer number in range 0 <= p <= 65535, but %d was provided", port)
	}
	return nil
}

func hasSpacesUnicode(s string) bool {
	for _, r := range s {
		if unicode.IsSpace(r) {
			return true
		}
	}
	return false
}

func CheckEndpointsCorrect(eps []string) error {
	var errs []error
	for _, ep := range eps {
		err := CheckEndpointCorrect(ep)
		if err != nil {
			errs = append(errs, err)
		}
	}
	return errors.Join(errs...)
}

func CheckEndpointCorrect(ep string) error {
	var errs []error

	if hasSpacesUnicode(ep) {
		errs = append(errs, fmt.Errorf("incorrect endpoint '%s': endpoint can't contain empty spaces", ep))
	}

	if !strings.HasPrefix(ep, "/") {
		errs = append(errs, fmt.Errorf("incorrect endpoint '%s': endpoint must start with '/'", ep))
	}

	return errors.Join(errs...)
}
