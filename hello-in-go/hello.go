package main

import "fmt"

const (
    spanish = "Hola, "
    frenchka14 = "Bonjour, "
)

// defining the output variable creates a varibale in function scope for us.
func getLanguageHello(language string) (prefix string) {
    if language == "spanish" {
        prefix = spanish
    } else if language == "french" {
        prefix = frenchka14
    } else {
        prefix = "Hello, "
    }

    return
}

func Hello(name string, language string) string{
    if name == "" {
        name = "world"
    }
    return getLanguageHello(language) + name
}

func main() {
    fmt.Println(Hello("JJ", "french"));
}
