package hello

import (
	"fmt"
)

const (
 Spanish = "Spanish"
 French = "French"

 englishHelloPrefix = "Hello, "
 spanishHelloPrefix = "Hola, "
 frenchHelloPrefix = "Bonjour, "
)

func Hello(name string, lang string) string {

	if name == "" {
		name = "World"
	}

	return greetingPrexif(lang) + name + "!"

}

func greetingPrexif(lang string) (prefix string) {
	switch lang {
	case Spanish:
		prefix = spanishHelloPrefix
	case French:
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return prefix 
}

 
func main() {
	fmt.Println(Hello("Jorch", ""))
}