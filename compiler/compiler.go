package compiler

import (
	"bufio"
	"fmt"
	"os"
)

var (
	commands = map[byte]func(*Compiler){
		'+': (*Compiler).plus,
		'-': (*Compiler).minus,
		'>': (*Compiler).next,
		'<': (*Compiler).back,
		'.': (*Compiler).print,
		',': (*Compiler).input,
		'[': (*Compiler).cycleStart,
		']': (*Compiler).cycleStop,
	}
)

func New() *Compiler {
	return &Compiler{
		i:     0,
		arr:   [30000]rune{},
		index: 0,
		gotos: []int{},
	}
}

type Compiler struct {
	i     int
	arr   [30000]rune
	index int
	gotos []int
}

func (c *Compiler) Compile(inp []byte) {
	for ; c.i < len(inp); c.i++ {
		f, ok := commands[inp[c.i]]
		if !ok {
			continue
		}

		f(c)
	}
}

func (c *Compiler) plus() {
	c.arr[c.index]++
}

func (c *Compiler) minus() {
	c.arr[c.index]--
}

func (c *Compiler) next() {
	c.index++
}

func (c *Compiler) back() {
	c.index--
}

func (c *Compiler) print() {
	fmt.Printf("%c", (c.arr[c.index]))
}

func (c *Compiler) input() {
	r, _, err := bufio.NewReader(os.Stdin).ReadRune()
	if err != nil {
		panic(err)
	}

	c.arr[c.index] = r
}

func (c *Compiler) cycleStart() {
	c.gotos = append(c.gotos, c.i)
}

func (c *Compiler) cycleStop() {
	if len(c.gotos) == 0 {
		panic("Unmatched ]")
	}

	if c.arr[c.index] == 0 {
		c.gotos = c.gotos[:len(c.gotos)-1]

		return
	}

	c.i = c.gotos[len(c.gotos)-1]
}
