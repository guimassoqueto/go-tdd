package main

type Dict map[string]string
type DictionaryError string


const (
	ErrNotFound = DictionaryError("could not find the word you were looking for")
	ErrKeyAlreadyExists = DictionaryError("key already exists")
	ErrKeyDoesNotExist = DictionaryError("key already exists")
)

func (e DictionaryError) Error() string {
	return string(e)
}


func (d Dict) Search(key string) (string, error) {
	value, ok := d[key]
	if !ok {
		return "", ErrNotFound
	}
	return value, nil
}

func (d Dict) Add(key string, value string) error {
	_, err := d.Search(key)

	switch err{
	case ErrNotFound:
		d[key] = value
	case nil:
		return ErrKeyAlreadyExists
	default:
		return err
	}
	return nil
}

func (d Dict) Update(key string, value string) error {
	_, err := d.Search(key)

	switch err {
	case ErrNotFound:
		return ErrKeyDoesNotExist
	case nil:
		d[key] = value
	default:
		return err
	}
	return nil
}

func (d Dict) Delete(key string) {
	delete(d, key)
}