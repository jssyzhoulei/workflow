package parse

import (
	"github.com/timchine/pole/pkg/yorm/parse/util"
	parse_v1 "github.com/timchine/pole/pkg/yorm/parse/v1"
	"strings"
)

type Parser interface {
	Parse(mp map[string]interface{}) (*util.ModelPathMap, error)
	Version() string
	ParseModelComplex(c interface{}, query util.Query) (*strings.Builder, []interface{})
}

func GetParser(version string, b util.Binder) Parser {
	switch version {
	case "v1.0":
		return parse_v1.NewParseV1(b)
		break
	}
	return nil
}


