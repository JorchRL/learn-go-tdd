package maps

type Dictionary map[string]string

var (
	ErrNotFound         = DictionaryError("cound not find the word")
	ErrWordExists       = DictionaryError("word already exists in dictionary")
	ErrWordDoesNotExist = DictionaryError("word does not exist in dictionary")
)

type DictionaryError string

func (e DictionaryError) Error() string {
	return string(e)
}

func (d *Dictionary) Search(word string) (string, error) {
	definition, ok := (*d)[word]
	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}

func (d Dictionary) Add(word string, definition string) error {
	_, err := d.Search(word)
	switch err {
	case ErrNotFound:
		d[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}
	d[word] = definition
	return nil
}

func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)
	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[word] = definition
	default:
		return err
	}
	d[word] = definition
	return nil
}

func (d Dictionary) Delete(word string) {

	delete(d, word)
}
