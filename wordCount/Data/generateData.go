package main

 

import (

    "bufio"

    "fmt"

    "log"

    "math/rand"

    "os"

)

 

const numberOfLines = 100

 

var names = []string{"depak", "murad", "sachin", "pramod", "aditi", "shalmali", "sanesh", "abhinav", "shashank", "viral", "satwik"}

 

func main() {

    totalNames := int64(len(names))

 

    f, err := os.OpenFile("names.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

    if err != nil {

        log.Println(err)

    }

    defer f.Close()

 

    w := bufio.NewWriter(f)

    defer w.Flush()

 

    var counter int64

    for counter < numberOfLines {

        name := names[rand.Int63n(totalNames)]

        w.WriteString(fmt.Sprintln(name))

        counter++

    }

}