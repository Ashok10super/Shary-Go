package main

import (
	"fmt"
)
func main()  {
	connection:= Connecttodb()


    fmt.Println(connection)

   fmt.Println(Disconnectdb(connection))

 defer Startserver()

}