package MyMaps

import "errors"

var errNotFound = errors.New("word not found")
var errAlreadyExists = errors.New("word alread exists")

type dictionary map[string]string

func (d dictionary) Search(wd string) (string, error) {
	defintion, ok := d[wd]

	if !ok {
		return "", errNotFound
	}
	return defintion, nil
}

func (d dictionary) Add(wd string, definittio string) error {
	_, ok := d[wd]

	if ok {
		return errAlreadyExists
	}
	d[wd] = definittio
	return nil
}
