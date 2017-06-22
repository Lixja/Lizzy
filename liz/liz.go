package main

import "Lizzy/brain"
import "fmt"
import "strconv"

func main() {
	liz := brain.GenerateBrain(5, 5, 2, 7)
	Inputs := []float64{72, 65, 76, 76, 79}
	Converted := []int{0, 0, 0, 0, 0}
	Outputs := liz.Think(Inputs)
	Converted = []int{int(Outputs[0]), int(Outputs[1]), int(Outputs[2]), int(Outputs[3]), int(Outputs[4])}
	write("72 - " + strconv.Itoa(Converted[0]))
	write("65 - " + strconv.Itoa(Converted[1]))
	write("76 - " + strconv.Itoa(Converted[2]))
	write("76 - " + strconv.Itoa(Converted[3]))
	write("79 - " + strconv.Itoa(Converted[4]))
	liz.Train(0.001, Inputs, Inputs, 100000, true)
	Outputs = nil
	Outputs = liz.Think(Inputs)
	Converted = []int{int(Outputs[0]), int(Outputs[1]), int(Outputs[2]), int(Outputs[3]), int(Outputs[4])}
	write("")
	write("72 - " + strconv.Itoa(Converted[0]))
	write("65 - " + strconv.Itoa(Converted[1]))
	write("76 - " + strconv.Itoa(Converted[2]))
	write("76 - " + strconv.Itoa(Converted[3]))
	write("79 - " + strconv.Itoa(Converted[4]))
	liz.Train(0.0001, Inputs, Inputs, 1000000, false)
	Outputs = nil
	Outputs = liz.Think(Inputs)
	Converted = []int{int(Outputs[0]), int(Outputs[1]), int(Outputs[2]), int(Outputs[3]), int(Outputs[4])}
	write("")
	write("72 - " + strconv.Itoa(Converted[0]))
	write("65 - " + strconv.Itoa(Converted[1]))
	write("76 - " + strconv.Itoa(Converted[2]))
	write("76 - " + strconv.Itoa(Converted[3]))
	write("79 - " + strconv.Itoa(Converted[4]))
}

func positive(i int) int {
	if i >= 0 {
		return i
	} else {
		return -i
	}
}

func write(msg string) {
	fmt.Print("Lizzy: " + msg + "\n")
}
