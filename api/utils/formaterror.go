package formaterror

import (
	"errors"
	"strings"
)

func FormatError(err string) error {
	if strings.Contains(err, "email") {
		return errors.New("Email ja esta sendo usado")
	}

	return errors.New("Detalhes incorretos")
}