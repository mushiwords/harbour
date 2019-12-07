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
    command := exec.Command("git", "clone", url)  
    fmt.Println(command.Args)  
  
    var bys bytes.Buffer  
    command.Stdout = &bys  
  
    fmt.Printf("=====>开始下载：%s<========\n", url)  
  
    err := command.Start()  
    if err != nil {  
        panic(err)  
    }  
  
    command.Wait()  
  
    fmt.Println(bys.String())  
  
    fmt.Printf("=====>下载完毕：%s<========\n", url)  
  
}
