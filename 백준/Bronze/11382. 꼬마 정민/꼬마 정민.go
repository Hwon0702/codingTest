package main
import (
    "fmt"
    "bufio"
    "os"
)

func main(){
    reader := bufio.NewReader(os.Stdin)
    writer := bufio.NewWriter(os.Stdout)
    defer writer.Flush()
    var a, b, c int
    fmt.Fscanf(reader, "%d %d %d\n", &a, &b, &c)
    fmt.Fprintf(writer, "%d", a+b+c)
}