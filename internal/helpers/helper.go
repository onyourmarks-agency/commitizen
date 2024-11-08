package helpers

import "github.com/onyourmarks-agency/commitizen/internal/errors"

func GetValueFromYAML[T any](data map[string]interface{}, key string) (T, error) {
	var (
		res T
		ok  bool
		v   interface{}
	)

	v, ok = data[key]
	if !ok {
		return res, errors.NewMissingErr(key)
	}
	res, ok = v.(T)
	if !ok {
		return res, errors.ErrType
	}
	return res, nil
}
