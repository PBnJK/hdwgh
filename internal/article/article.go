package article

import (
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
	"path"
	"regexp"

	"github.com/ledongthuc/pdf"
)

type Article struct {
	sources []string
}

const (
	TMP_ID_SIZE    = 8
	TMP_ID_CHARSET = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_-"
)

func New(input string) (Article, error) {
	file, err := createTemporaryFile()
	if err != nil {
		return Article{}, err
	}
	defer file.Close()

	fetchPDF(file, input)
	return createArticle(file)
}

func createTemporaryFile() (*os.File, error) {
	name := path.Join(os.TempDir(), generateRandomFileID(TMP_ID_SIZE))

	file, err := os.Create(name)
	if err != nil {
		return nil, err
	}

	return file, nil
}

func generateRandomFileID(size int) string {
	id := make([]byte, size)
	for i := range id {
		id[i] = TMP_ID_CHARSET[rand.Intn(len(TMP_ID_CHARSET))]
	}

	return string(id)
}

func fetchPDF(file *os.File, input string) error {
	// FIXME: for now we just assume everything is an URL
	return downloadFromURL(file, input)
}

func downloadFromURL(file *os.File, url string) error {
	response, err := http.Get(url)
	if err != nil {
		return err
	}

	defer response.Body.Close()

	_, err = io.Copy(file, response.Body)
	if err != nil {
		return err
	}

	return nil
}

func createArticle(file *os.File) (Article, error) {
	fi, err := file.Stat()
	if err != nil {
		return Article{}, err
	}

	fmt.Printf("Creating PDF reader... ")
	r, err := pdf.NewReader(file, fi.Size())
	if err != nil {
		fmt.Println("ERROR")
		return Article{}, err
	}
	fmt.Println("OK!")

	article := Article{
		sources: make([]string, 0),
	}

	pageNumber := r.NumPage()

	// Loop back through pages, since references/bibliographies are generally at
	// the end of an article
	for pageNumber > 0 {
		fmt.Printf("Reading page %d... ", pageNumber)
		page := r.Page(pageNumber)
		if page.V.IsNull() || page.V.Key("Contents").Kind() == pdf.Null {
			fmt.Println("no content! Skipping...")

			pageNumber--
			continue
		}
		fmt.Println("OK!")

		fmt.Printf("Finding sources... ")
		sources, err := getPageSources(page)
		if err != nil {
			fmt.Println("ERROR")
			return Article{}, nil
		}
		fmt.Println("OK!")

		// No sources found, assume we're done
		if len(sources) == 0 {
			break
		}

		article.sources = append(article.sources, sources...)

		pageNumber--
		break
	}

	return article, nil
}

func getPageSources(page pdf.Page) ([]string, error) {
	t, err := page.GetPlainText(nil)
	if err != nil {
		return nil, err
	}

	sources := make([]string, 0)

	sources = append(sources, findDOI(t)...)
	sources = append(sources, findArxiv(t)...)

	return sources, nil
}

func findDOI(text string) []string {
	sources := make([]string, 0)

	// TODO:
	// Find identifiers in the form: doi:10.nnnn/nnnn...

	// doiRegex := regexp.MustCompile(`10\..+?/.+`)

	return sources
}

func findArxiv(text string) []string {
	sources := make([]string, 0)

	// TODO:
	// Find identifiers in the form: arXiv:nnnn.nnnnn

	arxivRegex := regexp.MustCompile(`arXiv:(\d{4}\.\d+)`)
	for _, match := range arxivRegex.FindAllString(text, -1) {
		fmt.Println(match)
	}

	return sources
}

func findURLs(text string) []string {
	sources := make([]string, 0)

	// TODO:
	// Find and disambiguate URLs in the forms:
	// - arxiv.org/abs/nnnn.nnnnn, arxiv.org/pdf/nnnn.nnnnn as ArXiv articles
	// - doi.org/10.nnnn/nnnn... as DOI articles

	return sources
}
