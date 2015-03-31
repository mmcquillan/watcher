package checks

import "fmt"
import "os"
import "bufio"
import "time"

func LogWatch(f string) {
    
    file, err := os.Open(f)
    if err != nil {
        fmt.Println(err)
    }
    defer file.Close()

    r := bufio.NewReader(file)
    
    for {
        p, _ := r.Peek(1)
        if len(p) > 0 {
            ln, _, _ := r.ReadLine()
            fmt.Println(string(ln))
        } else {
            time.Sleep(3000)
        }
    }

}
