package model

type ResultCheckpoint struct {
	Iterations int32
	Successes  int32
}

func (c ResultCheckpoint) GetSuccessRate() float32 {
	return float32(c.Successes) / float32(c.Iterations) * 100.0
}
