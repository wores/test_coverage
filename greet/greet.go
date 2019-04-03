package greet

import "fmt"

func Greet(name string, lang string) string {
	if name == lang {
		fmt.Println("name == lang")
	} else if name == "techi" {
		fmt.Println("i am techi")
	} else {
		fmt.Println("name != lang")
	}

	// if name == "techi" {
	// 	fmt.Println("i am techi")
	// }

	fmt.Println("dododo popo")
	switch lang {
	case "en":
		return fmt.Sprintf("Hello, %s.", name)
	case "ja":
		return fmt.Sprintf("こんにちは, %s.", name)
	}
	return "f..."
}
