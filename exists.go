package validator

const (
	// Exists represents the rule name which will be used to find the default error message.
	Exists = "exists"
	// ExistsMsg is default error message format for records that not exists.
	ExistsMsg = "% not exists"
)

// Exists checks if given value is exists in database.
func (v *Validator) Exists(value any, table, column, field, msg string) *Validator {
	v.Check(v.repo.Exists(value, table, column), field, msg)

	return v
}