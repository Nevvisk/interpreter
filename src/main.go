package main

import (
    "fmt"
    "interpreter/src/repl"
    "os"
    "os/user"
) 

func main(){
    user, err := user.Current()
    if err != nil {
        panic(err)
    }
    fmt.Printf("Hello %s this is the interpreter.", user.Username)
    fmt.Printf("Type in commands\n")
    repl.Start(os.Stdin, os.Stdout)
}
