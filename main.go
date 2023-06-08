package main

import (

    "gitlab.com/krink/logstream/golang/logstream"
)


func main() {

    err := logstream.Stream()

    if err != nil {
        panic(err)
    }

}
