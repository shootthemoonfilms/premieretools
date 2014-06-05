package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

var (
	extIn  = flag.String("extIn", "mov", "Input file extension")
	extOut = flag.String("extOut", "mpg", "Output file extension")
)

func main() {
	flag.Parse()

	args := flag.Args()

	// Determine if we're using local directory or list of provided ones.
	var dirs []string
	if len(args) < 1 {
		log.Print("Using current working directory to scan")
		dirs = make([]string, 1)
		cwd, err := os.Getwd()
		if err != nil {
			panic(err)
		}
		dirs[0] = cwd
		log.Print("Found cwd " + cwd)
	} else {
		dirs = args[:]
	}

	// Iterate through directories
	for idx := range dirs {
		scanDir(dirs[idx])
	}
}

func scanDir(dirName string) {
	files, _ := ioutil.ReadDir(dirName)
	for _, f := range files {
		//log.Print(f.Name())
		fullPath := dirName + string(os.PathSeparator) + f.Name()
		if FileExists(fullPath) && (strings.HasSuffix(f.Name(), "."+*extIn) || strings.HasSuffix(f.Name(), "."+strings.ToUpper(*extIn))) {
			log.Print("Processing " + f.Name())
			processFile(dirName, f.Name())
		}
	}
}

func processFile(pathName, fileName string) {
	// Figure base filename without suffix
	baseFileName := ""
	if strings.HasSuffix(fileName, "."+*extIn) {
		baseFileName = strings.TrimSuffix(fileName, "."+*extIn)
	} else if strings.HasSuffix(fileName, "."+strings.ToUpper(*extIn)) {
		baseFileName = strings.TrimSuffix(fileName, "."+strings.ToUpper(*extIn))
	}

	log.Print("Processing " + fileName + " in '" + pathName + "'")
	inPath := pathName + string(os.PathSeparator) + fileName
	outPath := pathName + string(os.PathSeparator) + baseFileName + "." + *extOut

	os.Rename(inPath, outPath)

	log.Print("Successfully processed")
}

// FileExists reports whether the named file exists.
func FileExists(name string) bool {
	st, err := os.Stat(name)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	if st.IsDir() {
		return false
	}
	return true
}
