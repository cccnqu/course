package main

import (
    "fmt"
    "io/ioutil"
)

func main() {
    b, err := ioutil.ReadFile("file.txt") // just pass the file name
    if err != nil {
        fmt.Print(err)
    }

    fmt.Println(b) // print the content as 'bytes'

    str := string(b) // convert content to a 'string'

    fmt.Println(str) // print the content as a 'string'
}