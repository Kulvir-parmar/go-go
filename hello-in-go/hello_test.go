package main

import "testing"

func TestHello(t* testing.T) {
    got := Hello("JJ")
    want := "Hello, JJ"

    if got != want {
        t.Errorf("got %q want %q", got, want)
    }
}
