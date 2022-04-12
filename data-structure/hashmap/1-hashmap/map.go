package __hashmap

import "errors"

type dictionary map[string]string

func Search(dict map[string]string, key string) (string, error) {
	value, ok := dict[key]
	if !ok {
		return "", errors.New("value not found")
	}
	return value, nil
}

func (d dictionary) Search(key string) (string, error) {
	return Search(d, key)
}

func (d *dictionary) Add(key, value string) {
	(*d)[key] = value
}

func (d *dictionary) Update(key, value string) error {
	if _, ok := (*d)[key]; !ok {
		return errors.New("key not found")
	}
	(*d)[key] = value
	return nil
}
