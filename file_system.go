package main

import (
	l4g "github.com/alecthomas/log4go"
	"path/filepath"
)

const (
	FILE_TYPE = "*.xml"
)

func ListFiles(globPattern string) (matches []string, err error) {
	return filepath.Glob(globPattern)
}

func GetDirectoriesList(options map[string]interface{}) (output []string) {
	baseDir := options["<path_location>"].(string)
	rfcList, _ := getPatternList(baseDir)

	for _, dir := range rfcList {
		output = append(output, filepath.Join(dir, FILE_TYPE))
		l4g.Debug(filepath.Join(dir, FILE_TYPE))
	}
	return
}

func getPatternList(baseDir string) (matches []string, err error) {
	return filepath.Glob(filepath.Join(baseDir))
}
