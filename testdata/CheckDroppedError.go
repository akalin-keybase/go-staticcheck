package pkg

func fn1() error {
	var err error
	if err != nil {
		return err
	}
	if err != nil {
		if err.Error() == "foo" {
			return err
		}
		return nil
	}
	if err != nil {
		return nil // MATCH /error dropped/
	}
	return nil
}

func fn2() (int, error) {
	var err error
	if err != nil {
		return 0, err
	}
	if err != nil {
		if err.Error() == "foo" {
			return 0, err
		}
		return 4, nil
	}
	if err != nil {
		return 5, nil // MATCH /error dropped/
	}
	return 0, nil
}
