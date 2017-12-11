package main

import (
	"fmt"
	l4g "github.com/alecthomas/log4go"
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
  napoleon cfdi <path_location>
  napoleon conceptos <path_location>
  napoleon -h | --help
  napoleon -v | --version

Options:
  -h --help     Show this screen.
  -v --version  Show version.`

	options, _ := docopt.Parse(usage, nil, true, "1.0.2", false)
	l4g.Debug(options)

	c := make(<-chan int)

	if options["cfdi"].(bool) {
		c = consumer(producer(options), "cfdi")
	}
	if options["conceptos"].(bool) {
		c = consumer(producer(options), "conceptos")
	}
	<-c

	time.Sleep(time.Millisecond)
	l4g.Info("napoleon stopped")
}

func producer(options map[string]interface{}) <-chan string {
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

func consumer(in <-chan string, outputType string) <-chan int {
	l4g.Debug("Output type %s", outputType)
	out := make(chan int)
	fmt.Println(EncodeHeaders(outputType))
	go func() {
		for file := range in {
			for _, row := range EncodeAsRows(file, outputType) {
				fmt.Println(row)
			}
		}
		out <- 1
		close(out)
	}()

	return out
}
