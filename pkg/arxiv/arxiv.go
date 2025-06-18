// This package implements a backtracker using the ArXiv API
package arxiv

import (
	"fmt"

	"github.com/pbnjk/hdwgh/pkg/backtracker"
	"github.com/pbnjk/hdwgh/pkg/util"
)

type ArxivBT struct{}

func NewArxivBT() ArxivBT {
	return ArxivBT{}
}

func (bt *ArxivBT) Backtrack(article string, options backtracker.Options) (backtracker.Paper, error) {
	if util.IsValidURL(article) {
		fmt.Println("TODO")
	}

	return backtracker.Paper{}, nil
}
