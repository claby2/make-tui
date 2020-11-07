package main

import (
    "log"
    "os"
)

func main() {
    if len(os.Args[1:]) == 0 {
        log.Fatal("no input file")
        os.Exit(1)
    }
    file_path := os.Args[1]

    content := NewParsedContent(file_path)
    content.Parse()

    Render(content)
}
