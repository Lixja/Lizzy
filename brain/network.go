package brain

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"time"
)

type Network struct {
	Deep [][]Neuron
}

func (n *Network) TrainF(tvalue float64, InputV, WantedResult []float64, iterations int) bool {
	write("Training Brain")
	write("FastTraining")
	var backup *Network
	changed := 0.0
	for c := 0; c < iterations; c++ {
		Results := n.Think(InputV)
		difference := 0.0
		pon := 0
		mvalue := tvalue
		for i, res := range WantedResult {
			difference += positive(res - Results[i])
			if res > Results[i] {
				pon++
			} else if res < Results[i] {
				pon--
			}
		}
		if pon < 0 {
			mvalue = -tvalue
		} else if pon == 0 {
			return true
		}

		backup = n.Backup()
		for i := 0; i < len(n.Deep); i++ {
			for i2 := 0; i2 < len(n.Deep[i]); i2++ {
				n.Deep[i][i2].modify(mvalue)
			}
		}

		Results = n.Think(InputV)
		ndifference := 0.0
		for i, res := range WantedResult {
			ndifference += positive(res - Results[i])
		}
		if ndifference > difference {
			for i := 0; i < len(n.Deep); i++ {
				for i2 := 0; i2 < len(n.Deep[i]); i2++ {
					n.Deep[i][i2].Weight = backup.Deep[i][i2].Weight
				}
			}

		}
		if math.Mod(float64(c)/(float64(iterations)/100), 10) == 0 {
			write("Training completed: " + strconv.FormatFloat((float64(c)/(float64(iterations)/100)), 'f', -1, 64) + "% - " + strconv.FormatFloat(difference, 'f', -1, 64))
			if difference == changed {
				return false
			}
			changed = difference
		}
	}
	return false
}
func (n *Network) TrainS(tvalue float64, InputV, WantedResult []float64, iterations int) bool {
	write("Training Brain")
	write("SlowTraining")
	var backupw []float64
	var dp int
	var dlp int
	again := false
	oldpon := 0
	changed := 0.0
	for c := 0; c < iterations; c++ {
		Results := n.Think(InputV)
		difference := 0.0
		pon := 0
		mvalue := tvalue
		for i, res := range WantedResult {
			difference += positive(res - Results[i])
			if res > Results[i] {
				pon++
			} else if res < Results[i] {
				pon--
			}
		}
		if pon < 0 {
			mvalue = -tvalue
			if oldpon > 0 {
				again = false
				oldpon = 0
			}
		} else if pon == 0 {
			return true
		} else {
			if oldpon < 0 {
				oldpon = 0
				again = false
			}
		}

		if !again {
			dp = rand.Intn(len(n.Deep))
			dlp = rand.Intn(len(n.Deep[dp]))
		} else {
			again = false
		}
		backupw = make([]float64, len(n.Deep[dp][dlp].Weight))
		copy(backupw, n.Deep[dp][dlp].Weight)
		n.Deep[dp][dlp].modify(mvalue)

		Results = n.Think(InputV)
		ndifference := 0.0
		for i, res := range WantedResult {
			ndifference += positive(res - Results[i])
		}
		if ndifference > difference {
			copy(n.Deep[dp][dlp].Weight, backupw)
		} else if ndifference < difference {
			again = true
			oldpon = pon
		}
		if math.Mod(float64(c)/(float64(iterations)/100), 10) == 0 {
			write("Training completed: " + strconv.FormatFloat((float64(c)/(float64(iterations)/100)), 'f', -1, 64) + "% - " + strconv.FormatFloat(difference, 'f', -1, 64))
			if difference == changed {
				return false
			}
			changed = difference
		}
	}
	return false
}

func (n *Network) TrainL(tvalue float64, InputV, WantedResult []float64, breakup int) bool {
	write("Training Brain")
	write("LoopTraining")
	var backupw []float64
	var dp int
	var dlp int
	again := false
	oldpon := 0
	changed := 0
	randomPon := false
	for i := 0; changed != breakup; i++ {
		Results := n.Think(InputV)
		difference := 0.0
		pon := 0
		mvalue := tvalue
		for i, res := range WantedResult {
			difference += positive(res - Results[i])
			if res > Results[i] {
				pon++
			} else if res < Results[i] {
				pon--
			}
		}
		if pon < 0 {
			mvalue = -tvalue
			if oldpon > 0 {
				again = false
				oldpon = 0
			}
		} else if pon == 0 {
			return true
		} else {
			if oldpon < 0 {
				oldpon = 0
				again = false
			}
		}
		if randomPon {
			pon := rand.Intn(2)
			if pon == 0 {
				mvalue = -tvalue
			} else {
				mvalue = tvalue
			}

		}
		if again {
			if oldpon < 0 {
				mvalue = -tvalue
			} else {
				mvalue = tvalue
			}
		}
		if !again {
			dp = rand.Intn(len(n.Deep))
			dlp = rand.Intn(len(n.Deep[dp]))
		} else {
			again = false
		}
		backupw = make([]float64, len(n.Deep[dp][dlp].Weight))
		copy(backupw, n.Deep[dp][dlp].Weight)
		n.Deep[dp][dlp].modify(mvalue)
		Results = n.Think(InputV)
		ndifference := 0.0
		for i, res := range WantedResult {
			ndifference += positive(res - Results[i])
		}
		if ndifference > difference {
			copy(n.Deep[dp][dlp].Weight, backupw)
			changed++
		} else if ndifference < difference {
			changed = 0
			again = true
			oldpon = pon
		}
		if math.Mod(float64(i), float64(breakup)) == 0 {
			write("Trainingr running: " + strconv.Itoa(i) + ". Loop - " + strconv.FormatFloat(difference, 'f', -1, 64))
		}
		if(changed > breakup/2){
			randomPon = true
		}
	}
	return false
}

func GenerateBrain(ainputs, aoutputs, alayers, mnpl int) *Network {
	write("Generating Brain")
	write("Creating Neurons")
	var net Network
	rand.Seed(time.Now().UTC().UnixNano())
	var layer []Neuron
	for i := 0; i < ainputs; i++ {
		layer = append(layer, Neuron{})
	}
	net.Deep = append(net.Deep, layer)
	for i := 0; i < alayers; i++ {
		layer = nil
		for i2 := rand.Intn(mnpl - 1); i2 < mnpl; i2++ {
			layer = append(layer, Neuron{})
		}
		net.Deep = append(net.Deep, layer)
	}
	layer = nil
	for i := 0; i < aoutputs; i++ {
		layer = append(layer, Neuron{})
	}
	net.Deep = append(net.Deep, layer)
	write("Created Neurons")
	write("Connecting Neurons")
	for i := 0; i < len(net.Deep); i++ {
		for i2 := 0; i2 < len(net.Deep[i]); i2++ {
			if i+1 < len(net.Deep) {
				for i3 := 0; i3 < len(net.Deep[i+1]); i3++ {
					net.Deep[i][i2].Connect(&net.Deep[i+1][i3])
				}
			}
		}
		write(strconv.Itoa(i+1) + " of " + strconv.Itoa(len(net.Deep)) + " layers ready.")
	}
	write("Connected Neurons")
	return &net
}

func (n *Network) Think(val []float64) []float64 {
	var output []float64
	for i := 0; i < len(n.Deep[0]); i++ {
		n.Deep[0][i].DoTheThing(val[i], &output)
	}
	for len(output) < len(n.Deep[len(n.Deep)-1]) {
		time.Sleep(100 * time.Microsecond)
	}
	return output
}

func write(msg string) {
	fmt.Println("Brain: " + msg)
}

func positive(i float64) float64 {
	if i >= 0 {
		return i
	} else {
		return -i
	}
}

func (n *Network) Backup() *Network {
	var backup Network
	backup.Deep = make([][]Neuron, len(n.Deep))
	for i := 0; i < len(n.Deep); i++ {
		backup.Deep[i] = make([]Neuron, len(n.Deep[i]))
		copy(backup.Deep[i], n.Deep[i])
		for i2 := 0; i2 < len(n.Deep[i]); i2++ {
			backup.Deep[i][i2].Weight = make([]float64, len(n.Deep[i][i2].Weight))
			copy(backup.Deep[i][i2].Weight, n.Deep[i][i2].Weight)
		}
	}
	return &backup
}
