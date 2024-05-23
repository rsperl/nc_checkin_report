package main

import (
	"fmt"
	"os"

	"github.com/unidoc/unioffice/common/license"
)

func GetLicense() string {
	l, err := os.ReadFile("license.txt")
	if err != nil {
		panic(err)
	}
	return string(l)
}

func init() {
	// Make sure to load your metered License API key prior to using the library.
	// If you need a key, you can sign up and create a free one at https://cloud.unidoc.io
	lk := os.Getenv(`UNIDOC_LICENSE_API_KEY`)
	if lk == "" {
		fmt.Printf("Licence key must be set in the env var UNIDOC_LICENSE_API_KEY\n")
		os.Exit(1)
	}
	err := license.SetMeteredKey(lk)
	if err != nil {
		fmt.Printf("Failed to set metered license key: %s\n", err)
		os.Exit(1)
	}

}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: nc_check_in_report <config file>")
		os.Exit(1)
	}
	configFile := os.Args[1]
	s, err := os.ReadFile(configFile)
	if err != nil {
		panic(err)
	}
	config, err := LoadConfig(s)
	if err != nil {
		panic(err)
	}
	headerRow, records := ReadExcel(config.InFile, config)
	headers := RowToMap(headerRow)
	WriteDoc(config.OutFile, records, headers, config)
}
