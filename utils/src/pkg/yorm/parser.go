package yorm

import (
	"fmt"
	"github.com/timchine/pole/pkg/yorm/parse"
	"github.com/timchine/pole/pkg/yorm/parse/util"
	"gopkg.in/yaml.v2"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)



func (db *DB) LoadSqlYaml(path string) error {
	var (
		file *os.File
		err error
		info os.FileInfo
	)
	file, err = os.Open(path)
	if err != nil {
		return err
	}
	defer Close(file)
	info, err = file.Stat()
	if err != nil {
		return err
	}
	if info.IsDir() {
		err = db.loadDir(nil, path)
		if err != nil {
			return err
		}
	} else {
		db.mp, err = db.parse(path)
		if err != nil {
			return err
		}
	}
	return err
}

func (db *DB) loadDir(prefix []string, path string) (err error) {
	var (
		infos []os.FileInfo
		mp *util.ModelPathMap
	)
	if infos, err = ioutil.ReadDir(path); err != nil {
		return
	}
	for _, fileInfo := range infos {
		if fileInfo.IsDir() {
			prefix = append(prefix, fileInfo.Name())
			err = db.loadDir(prefix, filepath.Join(path, fileInfo.Name()))
			if err != nil {
				return err
			}
			prefix = nil
		} else {
			name := fileInfo.Name()
			ext := filepath.Ext(name)
			mp, err = db.parse(filepath.Join(path, name))
			if err != nil {
				return err
			}

			if mp != nil {
				pfx := append(prefix, name[0:len(name)-len(ext)])
				db.makeModelPathMap(strings.Join(pfx, "."), mp)
			}
		}

	}
	return
}

type Content struct {
	Header Info
	Models map[string]interface{}
}

func (c Content) isYorm() bool {
	if c.Header.Lang == YamlLang {
		return true
	}
	return false
}


type Info struct {
	Version string
	Lang string
}

const YamlLang = "YORM"

func (db *DB) parse(path string) (*util.ModelPathMap, error) {
	var (
		file *os.File
		decoder *yaml.Decoder
		content Content
		err error
	)
	file, err = os.Open(path)
	if err != nil {
		return nil, err
	}
	defer Close(file)
	decoder = yaml.NewDecoder(file)
	err = decoder.Decode(&content)
	if err != nil {
		return nil, fmt.Errorf("path %s error: %s", path, err)
	}
	if !content.isYorm() {
		return nil, fmt.Errorf("path %s error: is not a sql yaml", path)
	}
	parser := parse.GetParser(content.Header.Version, db.dialect)

	return parser.Parse(content.Models)
}


func Close(close io.Closer) {
	err := close.Close()
	if err != nil {
		panic(err)
	}
}
