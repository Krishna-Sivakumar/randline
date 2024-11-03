package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"math/rand"
	"os"
)

type FileNotFound struct {
	filename string
}

func (f FileNotFound) Error() string {
	return fmt.Sprintf("%s does not exist.", f.filename)
}

func readFromFile(filename string) (string, error) {
	contents, err := os.ReadFile(filename)
	if err != nil {
		return "", FileNotFound{filename: filename}
	}
	result, err := readFromByteBuffer(contents)
	if err != nil {
		return "", err
	}
	return result, nil
}

func readFromStdin() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	buf := make([]byte, 1024) // read a page at a time
	bytes := make([]byte, 0, 4*1024)
	totbytes := 0
	for {
		nbytes, err := reader.Read(buf)
		if nbytes == 0 {
			if err == nil {
				continue
			}
			if err == io.EOF {
				break
			}
			return "", err
		}
		totbytes += nbytes
		bytes = append(bytes, buf[:nbytes]...)
	}
	fmt.Println(totbytes)
	return readFromByteBuffer(bytes)
}

func readFromByteBuffer(contents []byte) (string, error) {
	offset := make([]int, 0)
	offset = append(offset, 0)
	lineCount := 0
	for index, byte := range contents {
		if byte == '\n' {
			lineCount += 1
			offset = append(offset, index+1)
		}
	}
	chosenLine := rand.Intn(lineCount + 1)
	if chosenLine == lineCount {
		return string(contents[offset[chosenLine]:]), nil
	} else {
		return string(contents[offset[chosenLine]:offset[chosenLine+1]]), nil
	}

}

func main() {
	filename := flag.String("filename", "", "File to be read.")
	flag.Parse()

	if *filename == "" {
		if len(os.Args) > 1 {
			// check if filename is present in the first argument
			if result, err := readFromFile(os.Args[1]); err != nil {
				fmt.Println(err.Error())
				return
			} else {
				fmt.Print(result)
				return
			}
		} else {
			// read from stdin
			if result, err := readFromStdin(); err != nil {
				fmt.Println(err.Error())
				return
			} else {
				fmt.Print(result)
				return
			}
		}
	} else {
		if result, err := readFromFile(os.Args[1]); err != nil {
			fmt.Println(err.Error())
			return
		} else {
			fmt.Print(result)
			return
		}
	}
}
