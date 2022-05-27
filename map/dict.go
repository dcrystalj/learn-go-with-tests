package main

type Dictionary map[string]string

const (
	ErrNotFound   = DictionaryErr("could not find the word you were looking for")
	ErrWordExists = DictionaryErr("cannot add word because it already exists")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

func Search(dictionary map[string]string, word string) string {
	return dictionary[word]
}

func (d Dictionary) SearchD(word string) (string, error) {
	res, found := d[word]
	if !found {
		return "", ErrNotFound
	}
	return res, nil
}

func (d Dictionary) Add(key, val string) error {
	_, err := d.SearchD(key)
	switch err {
	case ErrNotFound:
		d[key] = val
	case nil:
		return ErrWordExists
	default:
		return err
	}
	return nil
}
