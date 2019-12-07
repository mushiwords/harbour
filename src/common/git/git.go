package git  
  
import (  
    "os/exec"  
    "fmt"  
    "bytes"  
)  
  
func GitClone(gitUrl string ) {  
    go worker(gitUrl)  
}  
  
func worker(url string) {  
    command := exec.Command("git", "clone", url, "/home/captain/github/auto-harbour")  
  
    var bys bytes.Buffer  
    command.Stdout = &bys  
  
  
    err := command.Start()  
    if err != nil {  
        panic(err)  
    }  
  
    command.Wait()  
  
    fmt.Println(bys.String())  
}
