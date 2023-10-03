package model

type Feedback struct {
	Description TextareaField `json:"description"`
	NameInput   TextField     `json:"name_input"`
	EmailInput  TextField     `json:"email_input"`
}

// FuncFeedback is a helper function that binds appropriate properties to the model on instantiation
func (p Page) FuncFeedback() Feedback {
	return Feedback{
		Description: TextareaField{
			Input: Input{
				Autocomplete: "off",
				DataAttributes: []DataAttribute{
					{
						Key: "value-missing",
						Value: Localisation{
							LocaleKey: "ImproveThisPageInvalid",
							Plural:    1,
						},
					},
				},
				ID:         "description-field",
				IsRequired: true,
				Label: Localisation{
					LocaleKey: "ImproveThisPage",
					Plural:    1,
				},
				Language: p.Language,
				Name:     "description",
			},
		},
		NameInput: TextField{
			Input: Input{
				Autocomplete: "name",
				ID:           "name-field",
				Label: Localisation{
					LocaleKey: "NameOpt",
					Plural:    1,
				},
				Name:     "name",
				Type:     Text,
				Language: p.Language,
			},
		},
		EmailInput: TextField{
			Input: Input{
				Autocomplete: "email",
				DataAttributes: []DataAttribute{
					{
						Key: "type-mismatch",
						Value: Localisation{
							LocaleKey: "EmailInvalid",
							Plural:    1,
						},
					},
				},
				ID: "email-field",
				Label: Localisation{
					LocaleKey: "EmailOpt",
					Plural:    1,
				},
				Name:     "email",
				Type:     Email,
				Language: p.Language,
			},
		},
	}
}
