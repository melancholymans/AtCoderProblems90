package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	sc.Scan()
	s := strings.Split(sc.Text(), " ")
	h, _ := strconv.Atoi(s[0])
	w, _ := strconv.Atoi(s[1])
	var sl [2000][2000]int
	var r [2000]int
	var c [2000]int
	for i := 0; i < h; i++ {
		sc.Scan()
		s := strings.Split(sc.Text(), " ")
		for j := 0; j < w; j++ {
			sl[i][j], _ = strconv.Atoi(s[j])
			r[i] += sl[i][j]
			c[j] += sl[i][j]
		}
	}
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			fmt.Fprint(writer, r[i]+c[j]-sl[i][j], " ")
		}
		fmt.Fprint(writer, "\n")
	}
}
