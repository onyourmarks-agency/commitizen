package validators

import "github.com/onyourmarks-agency/commitizen/internal/errors"

func MultiRequired(name string) func([]string) error {
	return func(vals []string) error {
		if len(vals) == 0 {
			return errors.NewRequiredErr(name)
		}
		return nil
	}
}
