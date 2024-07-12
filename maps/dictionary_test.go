package maps

import "testing"

// Dictionary tests
func TestSearchWord(t *testing.T) {
	dictionary := Dictionary{"test": "this is a test"}
	t.Run("word in dictionary", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is a test"
		assertStrings(t, got, want)
	})
	t.Run("word not in dictionary", func(t *testing.T) {
		_, err := dictionary.Search("unknown")
		assertError(t, err, ErrNotFound)
	})
}

func TestAddWord(t *testing.T) {

	t.Run("new word", func(t *testing.T) {

		dict := Dictionary{}
		dict.Add("test", "this is a test")

		want := "this is a test"
		got, err := dict.Search("test")

		assertNotError(t, err)
		assertStrings(t, got, want)

	})
	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is a test"
		dict := Dictionary{word: definition}
		err := dict.Add(word, "new definition")

		assertError(t, err, ErrWordExists)
		got, _ := dict.Search(word)
		assertStrings(t, got, definition)
	})
}

func TestUpdateWord(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {

		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}
		newDefinition := "new definition"

		err := dictionary.Update(word, newDefinition)

		assertNotError(t, err)
		assertDefinition(t, dictionary, word, newDefinition)
	})

	t.Run("new word retu", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{}

		err := dictionary.Update(word, definition)

		assertError(t, err, ErrWordDoesNotExist)
	})
}

func TestDeleteWord(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		word := "test"
		dictionary := Dictionary{word: "test definition"}
		dictionary.Delete(word)
		_, err := dictionary.Search(word)

		if err != ErrNotFound {
			t.Errorf("Expected %q to be deleted", word)
		}
	})

}

// Assertion helpers

func assertDefinition(t testing.TB, dictonary Dictionary, word, definition string) {
	t.Helper()

	got, err := dictonary.Search(word)
	if err != nil {
		t.Fatal("should find added word:", err)
	}
	assertStrings(t, got, definition)
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q given %q", got, want, "test")
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}
}

func assertNotError(t testing.TB, err error) {
	t.Helper()
	if err != nil {
		t.Fatal("Should not get an error: ", err)
	}
}
