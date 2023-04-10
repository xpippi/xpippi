package drugs

type Pill int

const (
	Placebo Pill = iota
	Amphetamine
	Heroine
	Cannabis
	Methamphetamine = Heroine
)
