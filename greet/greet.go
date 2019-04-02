package greet

import "fmt"

func Greet(name string, lang string) string {
	if name == lang {
		fmt.Println("name == lang")
	} else {
		fmt.Println("name != lang")
	}

	fmt.Println("dododo")
	switch lang {
	case "en":
		return fmt.Sprintf("Hello, %s.", name)
	case "ja":
		return fmt.Sprintf("こんにちは, %s.", name)
	}
	return "..."
}
