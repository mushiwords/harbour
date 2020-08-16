package worker

import (
	"bufio"
	"io"
	"os"
	"strings"
)

type XYZData struct {
	X string
	Y string
	Z string
	Amplitude  string
}
func ReadTxtFile(path string) (*[]XYZData, error) {
	var results []XYZData
	lines,err := readTxt(path)
	if err != nil {
		return nil, err
	}
	for _,line := range lines{
		items := strings.Split(line,"\t")
		if len(items) == 4 && !strings.HasPrefix(items[0],"#"){
			results = append(results,XYZData{X:items[0],Y:items[1],Z:items[2],Amplitude:items[3]})
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