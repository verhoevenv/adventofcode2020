package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
	"unsafe"
)

type operation string
type argument int
type instruction struct {
	op  operation
	arg argument
}
type program []instruction

func (p *program) accumulateUntilLoopDetected() int {
	//doing some hard work just to prove a point(er)
	var execution *instruction = &(*p)[0] //Go is noisy C
	const instrSize uintptr = unsafe.Sizeof(*execution)

	var accumulator int = 0
	visited := make(map[uintptr]bool)
	for {
		// note: using a uintptr over multiple statements is completely
		//  invalid Go code according to the unsafe docs
		// this could be made safe by inlining the ep variable but
		//  it seems more readable to only to the type wrangling once
		// I assume that as long as Go doesn't decide to garbage collect,
		//  this should work, #YOLO
		// (dear colleagues, I would never do this in production code)
		ep := uintptr(unsafe.Pointer(execution))

		if visited[ep] {
			break
		}
		visited[ep] = true

		switch execution.op {
		case "nop":
			ep += instrSize
		case "acc":
			accumulator += int(execution.arg)
			ep += instrSize
		case "jmp":
			ep += uintptr(execution.arg) * instrSize
		default:
			panic(fmt.Sprintf("unknown instruction %v", execution.arg))
		}

		execution = (*instruction)(unsafe.Pointer(ep))
	}

	return accumulator
}

var instructionRE = regexp.MustCompile(`(.+) ([+-]\d+)`)

func parse(in string) program {
	result := make(program, 0)

	for _, line := range strings.Split(in, "\n") {
		matches := instructionRE.FindStringSubmatch(line)
		if len(matches) != 3 {
			panic(fmt.Sprintf("Line %v doesn't match", line))
		}

		op := operation(matches[1])
		num, err := strconv.Atoi(matches[2])
		if err != nil {
			panic(err)
		}
		arg := argument(num)

		result = append(result, instruction{op, arg})
	}

	return result
}

func main() {
	input, _ := ioutil.ReadAll(os.Stdin)

	program := parse(string(input))

	fmt.Println(program.accumulateUntilLoopDetected())
}
