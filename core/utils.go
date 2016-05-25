package core

import (
	"errors"
	"strings"
)

func SplitPreview(s string) (string, error) {
	str := strings.Split(s, "{preview}")
	if len(str) == 1 {
		return "", errors.New("No split {preview} descriptions")
	}
	return str[0], nil
}
