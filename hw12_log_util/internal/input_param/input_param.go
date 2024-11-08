package input_param

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/pflag"
)

func GetInputParam() (string, string, string) {
	var fileToPath string
	var method string
	var output string

	fileToPathEnv, methodEnv, outputEnv := readEnv()
	fileToPathFlag, methodFlag, outputFlag := readFlags()

	if len(fileToPathFlag) == 0 {
		fileToPath = fileToPathEnv
	} else {
		fileToPath = fileToPathFlag
	}
	if len(methodFlag) == 0 {
		method = methodEnv
	} else {
		method = methodFlag
	}
	if len(outputFlag) == 0 {
		output = outputEnv
	} else {
		output = outputFlag
	}

	return fileToPath, method, output
}

func readEnv() (string, string, string) {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	fileToPathEnv := os.Getenv("LOG_ANALYZER_FILE")
	methodEnv := os.Getenv("LOG_ANALYZER_METHOD")
	outputEnv := os.Getenv("LOG_ANALYZER_OUTPUT")

	return fileToPathEnv, methodEnv, outputEnv
}

func readFlags() (string, string, string) {
	fileToPathFlag := pflag.StringP("file", "f", "", "path to file")
	methodFlag := pflag.StringP("method", "m", "", "HTTP method which will analyze")
	outputFlag := pflag.StringP("output", "o", "", "Output to file, it empty then will be stdout")

	pflag.Parse()

	return *fileToPathFlag, *methodFlag, *outputFlag
}
