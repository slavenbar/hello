package main

import "fmt"

func main() {
    quit := make(chan bool)
    go func(){
        for{
            select{
            case<-quit:
                return
            default:
                fmt.Println("Quit")
            }
        }
    }()
    fmt.Println("Stop")
quit<-true        
}
