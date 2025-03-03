package greet

import "fmt"

func Hello(name string) string {
	return fmt.Sprintf("Hello from package2, %s", name)
}
