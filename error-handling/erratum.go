package erratum

// Use tries to open a resource with the given ResourceOpener.
// Once successfully opened, the Frob function of the resource will be called with the given input.
func Use(o ResourceOpener, input string) (err error) {
	resource, err := o()
	if err != nil {
		if _, ok := err.(TransientError); ok {
			return Use(o, input)
		}
		return err
	}
	defer func() {
		if r := recover(); r != nil {
			if frobErr, ok := r.(FrobError); ok {
				resource.Defrob(frobErr.defrobTag)
			}
			err = r.(error)
		}
		_ = resource.Close()
	}()
	resource.Frob(input)
	return nil
}
