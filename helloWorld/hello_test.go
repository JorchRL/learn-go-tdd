package hello

import "testing"

func TestHello(t *testing.T) {
	t.Run("Hello with a name", func(t *testing.T) {
		got := Hello("Jorch", "")
		want := "Hello, Jorch!"

		assertCorrectMessage(t, got, want)
	})

	t.Run("Hello with empty string", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World!"

		assertCorrectMessage(t, got, want)
	})

	t.Run("Hello in Spanish", func(t *testing.T) {
		got := Hello("Jorch", "Spanish")
		want := "Hola, Jorch!"

		assertCorrectMessage(t, got, want)
	})

	t.Run("Hello in French", func(t *testing.T) {
		got := Hello("Jorch", "French")
		want := "Bonjour, Jorch!"

		assertCorrectMessage(t, got, want)
	})


}

	func assertCorrectMessage (t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

