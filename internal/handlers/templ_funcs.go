package handlers

import (
	"fmt"
	"strings"
)

func FormatCodeBlock(codeblock string) string {
	splitBlock := strings.Split(codeblock, "\n")
	for i := 0; i < len(splitBlock); i++ {
		if i == 0 || i == len(splitBlock) - 1 {
			continue
		}
		splitBlock[i] = fmt.Sprintf("\t%s",splitBlock[i])
	}
	return strings.Join(splitBlock, "\n")
}