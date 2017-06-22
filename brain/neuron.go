package brain

import "math/rand"

type Neuron struct {
	Weight       []float64
	Inputs       []float64
	ConnectionsI []*Neuron
	ConnectionsO []*Neuron
}

func (in *Neuron) Connect(out *Neuron) {
	in.ConnectionsO = append(in.ConnectionsO, out)
	in.Weight = append(in.Weight, rand.Float64())
	out.ConnectionsI = append(out.ConnectionsI, in)
}

func (n *Neuron) DoTheThing(v float64, output *[]float64) {
	n.Inputs = append(n.Inputs, v)
	if len(n.Inputs) == len(n.ConnectionsI) || (len(n.ConnectionsI) == 0 && len(n.Inputs) == 1) {
		var question float64
		for i := 0; i < len(n.Inputs); i++ {
			question += n.Inputs[i]
		}
		n.Inputs = nil
		if len(n.ConnectionsO) == 0 {
			*output = append(*output, question)
		} else {
			for i := 0; i < len(n.ConnectionsO); i++ {
				n.ConnectionsO[i].DoTheThing(question*n.Weight[i], output)
			}
		}
	}
}

func (n *Neuron) modify(mval float64) {
	for i := 0; i < len(n.ConnectionsO); i++ {
		n.Weight[i] += rand.Float64() * mval
	}
}
