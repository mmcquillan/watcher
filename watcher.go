package main

import "fmt"
import "watcher/checks"

func main() {

    fmt.Printf("Watcher Program...\n")
    go checks.LogWatch("/var/log/alternatives.log")
    go checks.FileChangeWatch("/var/log/alternatives.log")

    var input string
    fmt.Scanln(&input)

}

