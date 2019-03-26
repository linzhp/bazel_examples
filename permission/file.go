package permission

import "io/ioutil"

func WriteX(dir string) error {
	f, err := ioutil.TempFile(dir, "x")
	if err != nil {
		return err
	}
	_, err = f.WriteString("Hello, Go!")
	return err
}
