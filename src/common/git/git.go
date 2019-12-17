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

func (g *GitConfig)GitClone( ) {  
   go func(){
        command := exec.Command("git", "clone", g.GitAddress, g.WebRoot)  
  
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
