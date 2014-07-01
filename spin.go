package spin

// Spinner types.
var (
	Box1    = `⠋⠙⠹⠸⠼⠴⠦⠧⠇⠏`
	Box2    = `⠋⠙⠚⠞⠖⠦⠴⠲⠳⠓`
	Box3    = `⠄⠆⠇⠋⠙⠸⠰⠠⠰⠸⠙⠋⠇⠆`
	Box4    = `⠋⠙⠚⠒⠂⠂⠒⠲⠴⠦⠖⠒⠐⠐⠒⠓⠋`
	Box5    = `⠁⠉⠙⠚⠒⠂⠂⠒⠲⠴⠤⠄⠄⠤⠴⠲⠒⠂⠂⠒⠚⠙⠉⠁`
	Box6    = `⠈⠉⠋⠓⠒⠐⠐⠒⠖⠦⠤⠠⠠⠤⠦⠖⠒⠐⠐⠒⠓⠋⠉⠈`
	Box7    = `⠁⠁⠉⠙⠚⠒⠂⠂⠒⠲⠴⠤⠄⠄⠤⠠⠠⠤⠦⠖⠒⠐⠐⠒⠓⠋⠉⠈⠈`
	Spin1   = `|/-\`
	Spin2   = `◴◷◶◵`
	Spin3   = `◰◳◲◱`
	Spin4   = `◐◓◑◒`
	Default = Box1
)

// Spinner.
type Spinner struct {
	Frames []rune
	length int
	pos    int
}

// New returns a spinner initialized with Default frames.
func New() *Spinner {
	s := &Spinner{}
	s.Set(Default)
	return s
}

// Set frames to the given string which must not use spaces.
func (s *Spinner) Set(frames string) {
	s.Frames = []rune(frames)
	s.length = len(s.Frames)
}

// Next returns the next rune in the sequence.
func (s *Spinner) Next() string {
	r := s.Frames[s.pos%s.length]
	s.pos++
	return string(r)
}

// Reset the spinner to its initial frame.
func (s *Spinner) Reset() {
	s.pos = 0
}
