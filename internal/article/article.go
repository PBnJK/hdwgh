package article

import (
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
	"path"
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

	return Article{}, nil
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

	fmt.Printf("%s\n", response.Status)

	_, err = io.Copy(file, response.Body)
	if err != nil {
		return err
	}

	return nil
}
