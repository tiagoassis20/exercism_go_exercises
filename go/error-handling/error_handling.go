package erratum

func Use(o ResourceOpener, input string) (err error) {
	resource, err := o()

	for err != nil {
		if _, ok := err.(TransientError); !ok {
			return
		}
		resource, err = o()
	}

	defer resource.Close()

	defer func() {
		if r := recover(); r != nil {
			switch r := r.(type) {
			case FrobError:
				{
					resource.Defrob(r.defrobTag)
					err = r
				}
			case error:
				err = r
			}
		}
	}()

	resource.Frob(input)

	return

}
