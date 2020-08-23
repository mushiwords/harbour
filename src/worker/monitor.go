package worker

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type XYZData struct {
	X int
	Y int
	Z int
	//A string
}
func ReadTxtFile(path string) (*[]XYZData, error) {
	var results []XYZData
	lines,err := readTxt(path)
	if err != nil {
		return nil, err
	}
	for idx,line := range lines{
		items := strings.Split(line,"\t")
		if len(items) >= 3 && !strings.HasPrefix(items[0],"#"){
			x,_ := strconv.Atoi(strings.Split(items[0],".")[0])
			y,_ := strconv.Atoi(strings.Split(items[1],".")[0])
			z,_ := strconv.Atoi(strings.Split(items[2],".")[0])
			results = append(results,XYZData{X:x,Y:y,Z:z})
		}
		if idx == 10 {
			fmt.Printf("read success: %s , %s, %s",items[0],items[1],items[2])
		}
	}

	return &results ,nil
}

func readTxt(path string) ([]string, error)  {
	var res []string

	fi, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer fi.Close()

	br := bufio.NewReader(fi)
	for {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		res = append(res,string(a))
	}
	return res,nil
}