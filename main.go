package main

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/fatih/color"
	"io"
	"log"
	"math"
	"os"
)

func countLines(r io.Reader) (uint64, error) {

	var count uint64
	const lineBreak = '\n'

	buf := make([]byte, bufio.MaxScanTokenSize)

	for {
		bufferSize, err := r.Read(buf)
		if err != nil && err != io.EOF {
			return 0, err
		}

		var buffPosition int
		for {
			i := bytes.IndexByte(buf[buffPosition:], lineBreak)
			if i == -1 || bufferSize == buffPosition {
				break
			}
			buffPosition += i + 1
			count++
		}
		if err == io.EOF {
			break
		}
	}

	return count, nil
}

func processFile(fn string) {

	_, width, err := analyzeFile(fn)
	if err != nil {
		log.Printf("analyze error with %s: %s\n", fn, err.Error())
		return
	}

	err = catFile(fn, width)
	if err != nil {
		log.Printf("catfile error with %s: %s\n", fn, err.Error())
		return

	}
}

func catFile(fileName string, numWidth uint) error {

	dim := color.New(color.FgHiBlack)

	input, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer func() {
		if closeErr := input.Close(); closeErr != nil {
			log.Printf("warning: error closing file %s in catfile: %s\n", fileName, err.Error())
		}
	}()

	fileScanner := bufio.NewScanner(input)
	var lineNumber uint64 = 1
	for fileScanner.Scan() {
		_, _ = dim.Printf("%*d: ", numWidth, lineNumber)
		_, _ = fmt.Println(fileScanner.Text())
		lineNumber++
	}

	return nil
}

func analyzeFile(fn string) (lineCount uint64, width uint, err error) {

	var input *os.File
	input, err = os.Open(fn)
	if err != nil {
		return
	}
	defer func() {
		if closeErr := input.Close(); closeErr != nil {
			log.Printf("warning: error closing file %s in analyzeFile: %s\n", fn, err.Error())
		}
	}()

	lineCount, err = countLines(input)
	if err != nil {
		return
	}

	width = uint(math.Ceil(math.Log10(float64(lineCount))))

	return

}

func main() {
	for _, fileName := range os.Args[1:] {
		processFile(fileName)
	}
}
