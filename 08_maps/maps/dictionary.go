package maps

// DictionaryErr defines dictionary error
type DictionaryErr string

const (
	// ErrWordExists defines word exists error
	ErrWordExists = DictionaryErr("cannot add word because it already exists")

	// ErrNotFound defines error not found
	ErrNotFound = DictionaryErr("could not find the word you were looking for")

	// ErrWordDoesNotExist defines word doesnot exists error
	ErrWordDoesNotExist = DictionaryErr("word does not exists")
)

func (e DictionaryErr) Error() string {
	return string(e)
}

// Dictionary map[string]string
type Dictionary map[string]string

// Search searches for a string
func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]

	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}

// Add adds a new entry
func (d Dictionary) Add(key, value string) error {

	_, err := d.Search(key)

	switch err {
	case ErrNotFound:
		d[key] = value
	case nil:
		return ErrWordExists
	default:
		return nil
	}

	d[key] = value

	return nil
}

// Update updates a definition
func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)
	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[word] = definition
	default:
		return nil
	}

	return nil
}

// Delete deletes a word
func (d Dictionary) Delete(word string) {
	delete(d, word)
}
