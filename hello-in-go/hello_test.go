package main

import "testing"

func assertCorrectMessage(t testing.TB, got, want string) {
    t.Helper()
    if got != want {
        t.Errorf("got %q want %q", got, want)
    }
}

func TestHello(t* testing.T) {
    t.Run("saying hello to JJ in spanish", func(t* testing.T) {
        got := Hello("JJ", "spanish")
        want := "Hola, JJ"
        assertCorrectMessage(t, got, want)
    })

    t.Run("saying hello to JJ in french", func(t* testing.T) {
        got := Hello("JJ", "french")
        want := "Bonjour, JJ"
        assertCorrectMessage(t, got, want)
    })

    t.Run("saying hello to JJ in deafult lang", func(t* testing.T) {
        got := Hello("JJ", "")
        want := "Hello, JJ"
        assertCorrectMessage(t, got, want)
    })

    t.Run("saying generic hello to all", func(t* testing.T) {
        got := Hello("", "")
        want := "Hello, world"
        assertCorrectMessage(t, got, want)
    })

}
