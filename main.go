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
		for u, v := range uniq {
			if strings.Contains(s, u) {
				delete(uniq, u)
			}
			for _, it := range v {
				if strings.Contains(s, it) {
					delete(uniq, u)
					break
				}
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
	for err != io.EOF {
		var line string
		line, err = bf.ReadString('\n')
		if err != nil && err != io.EOF {
			return
		}
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		line = strings.ReplaceAll(line, " ", "")
		names = append(names, line)
	}
	err = nil
	return
}

func sliceToSet(s []string) map[string][]string {
	m := make(map[string][]string)
	for _, it := range s {
		a := strings.Split(it, "|")
		m[a[0]] = a[1:]  // set alias in value if have
	}
	return m
}

func fmtUniq(m map[string][]string) {
	for k := range m {
		fmt.Println(k)
	}
}
