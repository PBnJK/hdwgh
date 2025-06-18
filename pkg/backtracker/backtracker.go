package backtracker

// Interface defining a backtracker, responsible for walking back through the
// history of sources in an article
type Backtracker interface {
	Backtrack(article string, options Options) (Paper, error)
}

// Struct defining options for the backtracking process
type Options struct {
	depth int // Maximum paper depth
	count int // Maximum sources to search, per paper
}

// Struct representing the paper info obtained
type Paper struct {
	title    string   // Title of the paper
	abstract string   // Abstract of the paper
	children []*Paper // Referenced papers
}
