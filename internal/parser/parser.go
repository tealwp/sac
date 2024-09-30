package parser

import (
	"github.com/tealwp/gofileparser"
)

func ParseGoFile(filePath string) (*gofileparser.GFP_GoFile, error) {
	return gofileparser.ParseGoFile(filePath)
}
