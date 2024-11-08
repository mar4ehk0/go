package file

import (
	"bufio"
	"os"

	"github.com/mar4ehk0/go/hw12_log_util/internal/analyzer"
)

func ReadFile(pathToFile string, out chan string) error {

	f, err := os.OpenFile(pathToFile, os.O_RDONLY, os.ModePerm)
	if err != nil {
		return err
	}
	defer f.Close()

	sc := bufio.NewScanner(f)

	for sc.Scan() {
		value := sc.Text()
		out <- value
	}

	close(out)

	return nil
}

func CreateFile(output string, stat analyzer.Stat) error {
	f, err := os.OpenFile(output, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	_, err = f.Write([]byte(stat.String()))
	if err != nil {
		return err
	}

	f.Close()

	return nil
}
