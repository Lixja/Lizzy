package main

import (
	"LizzyAI/brain"
	"LizzyAI/utils"
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	output := ""
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	var liz *brain.Network
	for output != text {
		output, liz = startLizzy(utils.GetAsciiI(text))
	}
	fmt.Println(*liz)

}

func startLizzy(textI []int) (string, *brain.Network) {
	liz := brain.GenerateBrain(len(textI), len(textI), 30, 30)
	var Inputs []float64
	InputInt := textI
	for _, runI := range textI {
		Inputs = append(Inputs, float64(runI))
	}
	var Converted []int
	Outputs := liz.Think(Inputs)
	Converted = convert(Outputs)
	print(InputInt, Converted)
	write(utils.GetAsciiS(Converted))
	liz.TrainF(0.001, Inputs, Inputs, 5000)
	Outputs = nil
	Outputs = liz.Think(Inputs)
	Converted = convert(Outputs)
	print(InputInt, Converted)
	write(utils.GetAsciiS(Converted))
	liz.TrainS(0.001, Inputs, Inputs, 100000)
	Outputs = nil
	Outputs = liz.Think(Inputs)
	Converted = convert(Outputs)
	print(InputInt, Converted)
	liz.TrainL(0.0001, Inputs, Inputs, 20000)
	Outputs = nil
	Outputs = liz.Think(Inputs)
	Converted = convert(Outputs)
	print(InputInt, Converted)
	hello := utils.GetAsciiS(Converted)
	write(hello)
	return hello, liz
}

func convert(f []float64) []int {
	var res []int
	for _, ff := range f {
		res = append(res, int(ff))
	}
	return res
}

func positive(i int) int {
	if i >= 0 {
		return i
	} else {
		return -i
	}
}

func print(r []int, o []int) {
	for i, sol := range r {
		write(strconv.Itoa(sol) + " - " + strconv.Itoa(o[i]))
	}

}

func write(msg string) {
	fmt.Print("Lizzy: " + msg + "\n")
}
