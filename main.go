package main

import (
        "fmt"
        "github.com/MarinX/keylogger"
)

func main() {
        keyboard := keylogger.FindKeyboardDevice()
        if len(keyboard) <= 0 {
                fmt.Println("No keyboard found")
                return
        }
        fmt.Println("Found a keyboard at", keyboard)
        k, err := keylogger.New(keyboard)
        if err != nil {
                fmt.Println("Error")
                return
        }
        defer k.Close()

        events := k.Read()

        for e:= range events {
                switch e.Type {
                case keylogger.EvKey:
                        if e.KeyPress() {
                                fmt.Println("Key pressed ", e.KeyString())
                        }
                        break
                }
        }
}
