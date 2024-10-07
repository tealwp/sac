package parser

import (
	"github.com/tealwp/gofileparser"
)

func ParseGoFile(filePath string) (*gofileparser.GFPGoFile, error) {
	return gofileparser.ParseGoFile(filePath)
}
