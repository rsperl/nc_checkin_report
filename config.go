package main

import (
	"strings"

	"gopkg.in/yaml.v2"
)

type Result struct {
	StartsOnColumn string `yaml:"startsOnColumn"`
	Count          int    `yaml:"count"`
}

type Config struct {
	InFile               string `yaml:"inFile"`
	OutFile              string `yaml:"outFile"`
	Title                string `yaml:"title"`
	HeaderStartsOnRow    int    `yaml:"headerStartsOnRow"`
	PointsPossible       int    `yaml:"pointsPossible"`
	FirstName            string `yaml:"firstName"`
	LastName             string `yaml:"lastName"`
	NumberCorrect        string `yaml:"numberCorrect"`
	PercentCorrect       string `yaml:"percentCorrect"`
	NumberItemsAttempted string `yaml:"numberItemsAttempted"`
	Results              Result `yaml:"results"`
}

func LoadConfig(data []byte) (Config, error) {
	var c Config
	err := yaml.Unmarshal(data, &c)
	if err != nil {
		return c, err
	}
	if c.OutFile == "" {
		c.OutFile = strings.ReplaceAll(c.Title, ":", "") + ".docx"
	}
	return c, nil
}
