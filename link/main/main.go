package main

import (
    "fmt"
    "github.com/jzuhusky/gophercises/link/link"
)

func main() {
    fmt.Println("hello world")

    fnames := []string{"ex1.html", "ex2.html", "ex3.html", "signorile.html"}

    for _, fname :=  range fnames {
        filebytes := link.HtmlBytesFromFile(fname)
        links := link.ParseHtml(filebytes)
        fmt.Println(links)
    }
}
