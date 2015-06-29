package main

import (
	l4g "code.google.com/p/log4go"
	"fmt"
	"github.com/docopt/docopt-go"
	"os"
	"time"
)

const LOG_CONFIGURATION_FILE = "logging-conf.xml"

func init() {
	l4g.LoadConfiguration(LOG_CONFIGURATION_FILE)
}

func main() {
	l4g.Info("Process ID: %d", os.Getpid())

	usage := `
Usage:
  napoleon run <path_location>
  napoleon -h | --help
  napoleon -v | --version

Options:
  -h --help     Show this screen.
  -v --version  Show version.`

	options, _ := docopt.Parse(usage, nil, true, "0.0.1", false)
	l4g.Debug(options)

	c := make(<-chan int)

	if options["run"].(bool) {
		c = Consumer(Producer(options))
	}
	l4g.Info("napoleon stopped")
	<-c
	time.Sleep(time.Millisecond)
}

func Producer(options map[string]interface{}) <-chan string {
	out := make(chan string)
	go func() {
		directoriesList := GetDirectoriesList(options)
		l4g.Info("Directorios encontrados: %d", len(directoriesList))

		for _, directory := range directoriesList {
			files, _ := ListFiles(directory)
			l4g.Info("%d archivos en directorio %s", len(files), directory)

			for _, filePath := range files {
				out <- filePath
			}
		}
		close(out)
	}()

	return out
}

func Consumer(in <-chan string) <-chan int {
	out := make(chan int)
	fmt.Println(EncodeHeaders())
	go func() {
		for file := range in {
			for _, row := range EncodeAsRows(file) {
				fmt.Println(row)
			}
		}
		out <- 1
		close(out)
	}()

	return out
}
