package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

type student struct {
	class  int
	grades int
}

type question struct {
	left  int
	right int
}

type cusum struct {
	l int
	s []int
}

func init() {
	sc.Split(bufio.ScanWords)
	buf := make([]byte, 0)
	sc.Buffer(buf, 131072)
}

func newcusum(sl []int) *cusum {
	c := &cusum{}
	c.l = len(sl)
	c.s = make([]int, len(sl)+1)
	for i, v := range sl {
		c.s[i+1] = c.s[i] + v
	}
	return c
}

func (c *cusum) getRange(f, t int) int {
	if f > t || f >= c.l {
		return 0
	}
	return c.s[t+1] - c.s[f]
}

func rl() int {
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	return n
}

func main() {
	s := []student{}
	n := rl()
	var sd student
	for i := 0; i < n; i++ {
		sd.class, sd.grades = rl(), rl()
		s = append(s, sd)
	}
	q := []question{}
	qn := rl()
	var qt question
	for i := 0; i < qn; i++ {
		qt.left, qt.right = rl(), rl()
		q = append(q, qt)
	}
	c1 := make([]int, n)
	c2 := make([]int, n)
	for i := 0; i < n; i++ {
		if s[i].class == 1 {
			c1[i] = s[i].grades
		} else {
			c2[i] = s[i].grades
		}
	}
	cu1 := newcusum(c1)
	cu2 := newcusum(c2)
	for i := 0; i < qn; i++ {
		l, r := q[i].left, q[i].right
		fmt.Println(cu1.getRange(l-1, r-1), " ", cu2.getRange(l-1, r-1))
	}
}
