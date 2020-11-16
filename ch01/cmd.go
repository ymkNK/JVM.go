package main
import "flag"
import "os"
import "fmt"
type Cmd struct{
    helpFlag    bool
    versionFlag bool
    cpOption    string
    class       string
    args        []string
}