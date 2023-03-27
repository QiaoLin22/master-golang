package main

func main() {
	err := Foo()
	if err != nil {
		return
	}

	y, err := Bar()
	if err != nil {
		return
	}

	_ = y

	if err := bar(); err != nil {
		// handler
	}
}

func Foo() error {
	return nil
}

func Bar() (int, error) {
	return 666, nil
}

func Baz() (int, error) {
	return 666, nil
}

func bar() error {
	if err := Foo(); err != nil {
		return err
	}
	return nil
}
