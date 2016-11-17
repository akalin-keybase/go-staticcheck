package pkg

func fn1() error {
	var err error
	if err != nil {
		return err
	}
	if err != nil {
		return nil // MATCH /error dropped/
	}
	return nil
}
