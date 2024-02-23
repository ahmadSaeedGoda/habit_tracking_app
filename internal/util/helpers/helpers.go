package helpers

// PanicIfError is a helper to centralize error checking in case of panic is required.
// Extracted for reusability & DRY.
func PanicIfError(err error) {
	if err != nil {
		panic(err)
	}
}
