package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"path"
	"strings"
)

var home string

func init() {
	var err error
	home, err = os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	printAbsent("5")
	printAbsent("6")
}

func printAbsent(cls string) {
	names, err := readLineToSlice(getFilePath(cls, "all.txt"))
	if err != nil {
		log.Fatal(err)
	}
	uniq := sliceToSet(names)

	signs, err := readLineToSlice(getFilePath(cls, "sign.txt"))
	if err != nil {
		log.Fatal(err)
	}

	for _, s := range signs {
		for u := range uniq {
			if strings.Contains(s, u) {
				delete(uniq, u)
				break
			}
		}
	}
	fmt.Printf("\n---------- Absent in class %s -----------\n", cls)
	fmtUniq(uniq)
}

func getFilePath(cls string, name string) string {
	return path.Join(home, ".uniq", cls, name)
}

func readLineToSlice(filePath string) (names []string, err error) {
	f, err := os.Open(filePath)
	if err != nil {
		return
	}
	defer f.Close()
	names = make([]string, 0, 64)
	bf := bufio.NewReader(f)
	for {
		var line string
		line, err = bf.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				err = nil
			}
			return
		}
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		names = append(names, line)
	}
	return
}

func sliceToSet(s []string) map[string]bool {
	m := make(map[string]bool)
	for _, it := range s {
		m[it] = true
	}
	return m
}

func fmtUniq(m map[string]bool) {
	for k := range m {
		fmt.Println(k)
	}
}
