package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReaderSize(os.Stdin, 1e6)

func Nextint(a ...*int) {
	l := strings.Split(Nextline(), " ")
	for i, n := range a {
		*n, _ = strconv.Atoi(l[i])
	}
}

func Nextskicei(a *[]int) {
	l := strings.Split(Nextline(), " ")
	if len(l) == 1 && l[0] == "" {
		(*a) = make([]int, 0)
		return
	}
	(*a) = make([]int, len(l))
	for i := 0; i < len(l); i++ {
		(*a)[i], _ = strconv.Atoi(l[i])
	}
}

func Nextline() string {
	buffer := make([]byte, 0)
	for {
		line, isPrefix, err := reader.ReadLine()
		if err != nil {
			panic(err)
		}
		buffer = append(buffer, line...)
		if !isPrefix {
			break
		}
	}
	return string(buffer)
}

func Sli(l, m int, def int) [][]int {
	sl := make([][]int, l)
	for i := 0; i < l; i++ {
		sl[i] = make([]int, m)
		for j := 0; j < m; j++ {
			sl[i][j] = def
		}
	}
	return sl
}

func main() {
	var h, w int
	Nextint(&h, &w)
	sl := Sli(h, w, 0)
	r := make([]int, h)
	c := make([]int, w)
	for i, _ := range sl {
		Nextskicei(&sl[i])
		for j := 0; j < len(sl[i]); j++ {
			r[i] += sl[i][j]
			c[j] += sl[i][j]
		}
	}
	//fmt.Println(h, w)
	//fmt.Println(sl)
	//fmt.Println("r:", r)
	//fmt.Println("c:", c)
	re := Sli(h, w, 0)
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			re[i][j] = r[i] + c[j] - sl[i][j]
			fmt.Print(re[i][j], " ")
		}
		fmt.Print("\n")
	}
}
