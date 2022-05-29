package main
import(
"fmt"
"bufio"
"os"
)

func main(){
    reader:=bufio.NewReader(os.Stdin)
    writer:=bufio.NewWriter(os.Stdout)
    defer writer.Flush()
    
    var a,b int
    fmt.Fscanf(reader, "%d\n%d", &a,&b)
    fmt.Fprintf(writer, "%d", a+b)
}