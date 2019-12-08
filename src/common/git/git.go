package git  
  
import (  
    "os/exec"  
    "fmt"  
    "bytes"  
)  
  
type GitConfig struct {
    GitAddress string `json:"GitAddress"`
    WebRoot string `json:"WebRoot"`
} 

func (g *GitConfig)GitClone() {  
    // todo existed
   go func(){
        command := exec.Command("git", "clone", url, g.GitAddress)  
  
        var bys bytes.Buffer  
        command.Stdout = &bys  
  
        err := command.Start()  
        if err != nil {  
            panic(err)  
        }  
  
        command.Wait()  
  
        fmt.Println(bys.String())  
    }()
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
