package main

import (
	"bufio"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

var rownum = flag.Int("n", 0, "row number")
var maxc = flag.Int("c", 1, "max column number")
var path = flag.String("p", "data.txt", "file path")

func init() {
	rand.Seed(time.Now().UnixNano())
	flag.Parse()
}

func main() {
	f, err := os.Create(*path)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	w := bufio.NewWriter(f)
	var cnt uint64
	for r := 0; r < *rownum; r++ {
		for i, l := 0, rand.Intn(*maxc)+1; i < l; i++ {
			s := strconv.Itoa(rand.Intn(2))
			w.WriteString(s)
			if s == "1" {
				cnt++
			}
		}
		w.WriteByte(0x0a)
		if err := w.Flush(); err != nil {
			panic(err)
		}
	}
	fmt.Println(cnt)
}
