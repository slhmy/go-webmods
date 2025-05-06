package gwm_utils

func UnwrapJoinErr(err error) []error {
	if joinErr, ok := err.(interface{ Unwrap() []error }); ok {
		errors := []error{}
		unwrappedErrors := joinErr.Unwrap()
		for _, e := range unwrappedErrors {
			errors = append(errors, UnwrapJoinErr(e)...)
		}
		return errors
	}
	return []error{err}
}
