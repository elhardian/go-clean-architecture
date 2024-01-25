package errorHelper

import "fmt"

func ErrorDataNotExist(object, identifier string) error {
	return fmt.Errorf("%s with identifier %s doesn't exists", object, identifier)
}
