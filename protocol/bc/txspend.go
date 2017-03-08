package bc

type Spend struct {
	body struct {
		SpentOutput Hash // the hash of an output entry
		Data        Hash
		ExtHash     Hash
	}
	ordinal int

	witness struct {
		Destination ValueDestination
		Arguments   [][]byte
	}

	// SpentOutput contains (a pointer to) the manifested entry
	// corresponding to body.SpentOutput.
	SpentOutput *Output
}

func (Spend) Type() string         { return "spend1" }
func (s *Spend) Body() interface{} { return s.body }

func (s Spend) Ordinal() int { return s.ordinal }

func (s *Spend) SpentOutputID() Hash {
	return s.body.SpentOutput
}

func (s *Spend) Data() Hash {
	return s.body.Data
}

func (s *Spend) AssetID() AssetID {
	return s.SpentOutput.AssetID()
}

func (s *Spend) ControlProgram() Program {
	return s.SpentOutput.ControlProgram()
}

func (s *Spend) Amount() uint64 {
	return s.SpentOutput.Amount()
}

func (s *Spend) Destination() ValueDestination {
	return s.witness.Destination
}

func (s *Spend) Arguments() [][]byte {
	return s.witness.Arguments
}

func (s *Spend) SetDestination(id Hash, pos uint64, e Entry) {
	s.witness.Destination = ValueDestination{
		Ref:      id,
		Position: pos,
		Entry:    e,
	}
}

func NewSpend(out *Output, data Hash, ordinal int) *Spend {
	s := new(Spend)
	s.body.SpentOutput = EntryID(out)
	s.body.Data = data
	s.ordinal = ordinal
	s.SpentOutput = out
	return s
}