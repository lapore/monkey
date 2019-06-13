package main

import "fmt"
import "os"
import "os/user"
import "monkey/repl"

func main() {
    user, err := user.Current()
    if err != nil {
        panic(err)
    }

    fmt.Printf("Heelp %s! This is the Monkey programming language!\n", user.Username)

    fmt.Printf("Feel free to type in commands\n")

//    repl.Start(os.Stdin, os.Stdout)
    repl.Start(os.Stdin)
}
