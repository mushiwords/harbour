package worker

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type XYZData struct {
	X string
	Y string
	Z string
	A string
}
func ReadTxtFile(path string) (*[]XYZData, error) {
	var results []XYZData
	lines,err := readTxt(path)
	if err != nil {
		return nil, err
	}
	for idx,line := range lines{
		items := strings.Split(line,"\t")
		if len(items) == 4 && !strings.HasPrefix(items[0],"#"){
			results = append(results,XYZData{X:items[0],Y:items[1],Z:items[2],A:items[3]})
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