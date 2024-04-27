package secret

import (
	"github.com/charmbracelet/huh"
	"github.com/shipengqi/commitizen/internal/parameter/validators"

	"github.com/shipengqi/commitizen/internal/parameter/str"
)

type Param struct {
	str.Param `mapstructure:",squash"`
}

func (p Param) Render() huh.Field {
	param := p.Param.RenderInput()
	param.Password(true)

	// reset validators of the secret
	var group []validators.Validator[string]
	if p.Required {
		group = append(group, validators.Required(p.Name, p.Trim))
	}
	// if the value is not required and no value has been given, min length validator should be ignored.
	if p.Required && p.MinLength != nil {
		group = append(group, validators.MinLength(*p.MinLength))
	}
	if p.MaxLength != nil {
		group = append(group, validators.MaxLength(*p.MaxLength))
	}
	if p.Regex != "" {
		group = append(group, validators.RegexValidator(p.Regex, p.RegexMessage))
	}
	if len(group) > 0 {
		param.Validate(validators.Group(group...))
	}

	return param
}
