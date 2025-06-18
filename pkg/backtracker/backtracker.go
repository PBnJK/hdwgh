package backtracker

import "github.com/pbnjk/hdwgh/internal/article"

type Backtracker struct {
	article article.Article
	options Options
}

// Struct defining options for the backtracking process
type Options struct {
	depth int // Maximum paper depth (<= 0 means no limit, though not recommended)
	count int // Maximum sources to search, per paper (<= 0 means no limit)
}

// Struct representing the paper info obtained
type Paper struct {
	title    string   // Title of the paper
	abstract string   // Abstract of the paper
	children []*Paper // Referenced papers
}

func DefaultOptions() Options {
	return Options{
		depth: 10,
		count: -1,
	}
}

func New(article article.Article, options Options) Backtracker {
	return Backtracker{
		article,
		options,
	}
}

func (bt Backtracker) Backtrack() (Paper, error) {
	return bt.backtrack(bt.article)
}

func (bt Backtracker) backtrack(article article.Article) (Paper, error) {
	return Paper{}, nil
}
