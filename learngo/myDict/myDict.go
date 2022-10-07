package myDict

import "errors"

// Dictionary type
type Dictionary map[string]string

var (
	errNotfound   = errors.New("not found")
	errCantUpdate = errors.New("can't update non-exitsting word")
	errWordExist  = errors.New("that word already exists")
)

// Search for a word
func (d Dictionary) Search(word string) (string, error) {
	value, exist := d[word]
	if exist {
		return value, nil
	} else {
		return "", errNotfound
	}
}

// Add a word
func (d Dictionary) Add(word string, def string) error {
	_, err := d.Search(word)
	switch err {
	case errNotfound:
		d[word] = def
	case nil:
		return errWordExist
	}
	return nil
}

// Update Dictionary
func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)
	switch err {
	case nil:
		d[word] = definition
	case errNotfound:
		return errCantUpdate
	}
	return nil
}

// Delete a word
func (d Dictionary) Delete(word string) {
	delete(d, word)
}
