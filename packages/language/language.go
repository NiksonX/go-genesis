// MIT License
//
// Copyright (c) 2016 GenesisCommunity
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.
package language

import (
	"encoding/json"
	"strings"
	"unicode/utf8"

	"strconv"

	"github.com/GenesisCommunity/go-genesis/packages/consts"
	"github.com/GenesisCommunity/go-genesis/packages/converter"
	"github.com/GenesisCommunity/go-genesis/packages/model"

	log "github.com/sirupsen/logrus"
)

//cacheLang is cache for language, first level map is app_id, second is lang_name, third is lang dictionary
type cacheLang struct {
	res map[int]map[string]*map[string]string
}

var (
	// LangList is the list of available languages. It stores two-bytes codes
	LangList []string
	lang     = make(map[int]*cacheLang)
)

// IsLang checks if there is a language with code name
func IsLang(code string) bool {
	if LangList == nil {
		return true
	}
	for _, val := range LangList {
		if val == code {
			return true
		}
	}
	return false
}

// DefLang returns the default language
func DefLang() string {
	if LangList == nil {
		return `en`
	}
	return LangList[0]
}

// UpdateLang updates language sources for the specified state
func UpdateLang(state, appID int, name, value string, vde bool) {
	if vde {
		state = -state
	}
	if _, ok := lang[state]; !ok {
		lang[state] = &cacheLang{make(map[int]map[string]*map[string]string)}
	}
	var ires map[string]string
	err := json.Unmarshal([]byte(value), &ires)
	if err != nil {
		log.WithFields(log.Fields{"type": consts.JSONUnmarshallError, "value": value, "error": err}).Error("Unmarshalling json")
	}
	for key, val := range ires {
		ires[strings.ToLower(key)] = val
	}
	if len(ires) > 0 {
		if _, ok := (*lang[state]).res[appID]; !ok {
			(*lang[state]).res[appID] = map[string]*map[string]string{}
		}
		(*lang[state]).res[appID][name] = &ires
	}
}

// loadLang download the language sources from database for the state
func loadLang(state int, vde bool) error {
	language := &model.Language{}
	prefix := strconv.FormatInt(int64(state), 10)
	if vde {
		prefix += `_vde`
	}
	languages, err := language.GetAll(prefix)
	if err != nil {
		log.WithFields(log.Fields{"type": consts.DBError, "error": err}).Error("Error querying all languages")
		return err
	}
	list := make([]map[string]string, 0)
	for _, l := range languages {
		list = append(list, l.ToMap())
	}
	res := make(map[int]map[string]*map[string]string)
	for _, ilist := range list {
		var ires map[string]string
		err := json.Unmarshal([]byte(ilist[`res`]), &ires)
		if err != nil {
			log.WithFields(log.Fields{"type": consts.JSONUnmarshallError, "value": ilist["res"], "error": err}).Error("Unmarshalling json")
		}
		for key, val := range ires {
			ires[strings.ToLower(key)] = val
		}
		if _, ok := res[converter.StrToInt(ilist[`app_id`])]; !ok {
			res[converter.StrToInt(ilist[`app_id`])] = map[string]*map[string]string{}
		}
		res[converter.StrToInt(ilist[`app_id`])][ilist[`name`]] = &ires
	}
	langInd := langIndex(state, vde)
	if _, ok := lang[langInd]; !ok {
		lang[langInd] = &cacheLang{}
	}
	lang[langInd].res = res
	return nil
}

func langIndex(state int, vde bool) int {
	if vde {
		return -state
	}
	return state
}

// LangText looks for the specified word through language sources and returns the meaning of the source
// if it is found. Search goes according to the languages specified in 'accept'
func LangText(in string, state, appID int, accept string, vde bool) (string, bool) {
	if strings.IndexByte(in, ' ') >= 0 || state == 0 {
		return in, false
	}
	istate := langIndex(state, vde)
	if _, ok := lang[istate]; !ok {
		if err := loadLang(state, vde); err != nil {
			return err.Error(), false
		}
	}
	langs := strings.Split(accept, `,`)
	if _, ok := (*lang[istate]).res[appID]; !ok {
		return in, false
	}
	if lres, ok := (*lang[istate]).res[appID][in]; ok {
		lng := DefLang()
		for _, val := range langs {
			val = strings.ToLower(val)
			if len(val) < 2 {
				break
			}
			if !IsLang(val[:2]) {
				continue
			}
			if len(val) >= 5 && val[2] == '-' {
				if _, ok := (*lres)[val[:5]]; ok {
					lng = val[:5]
					break
				}
			}
			if _, ok := (*lres)[val[:2]]; ok {
				lng = val[:2]
				break
			}
		}
		if len((*lres)[lng]) == 0 {
			for _, val := range *lres {
				return val, true
			}
		}
		return (*lres)[lng], true
	}
	return in, false
}

// LangMacro replaces all inclusions of $resname$ in the incoming text with the corresponding language resources,
// if they exist
func LangMacro(input string, state, appID int, accept string, vde bool) string {
	if !strings.ContainsRune(input, '$') {
		return input
	}
	syschar := '$'
	length := utf8.RuneCountInString(input)
	result := make([]rune, 0, length)
	isName := false
	name := make([]rune, 0, 128)
	clearname := func() {
		result = append(append(result, syschar), name...)
		isName = false
		name = name[:0]
	}
	for _, r := range input {
		if r != syschar {
			if isName {
				name = append(name, r)
				if len(name) > 64 || r < ' ' {
					clearname()
				}
			} else {
				result = append(result, r)
			}
			continue
		}
		if isName {
			value, ok := LangText(string(name), state, appID, accept, vde)
			if ok {
				result = append(result, []rune(value)...)
				isName = false
			} else {
				result = append(append(result, syschar), name...)
			}
			name = name[:0]
		} else {
			isName = true
		}
	}
	if isName {
		result = append(append(result, syschar), name...)
	}

	return string(result)
}

// GetLang returns the first language from accept-language
func GetLang(state int, accept string) (lng string) {
	lng = DefLang()
	for _, val := range strings.Split(accept, `,`) {
		if len(val) < 2 {
			continue
		}
		if !IsLang(val[:2]) {
			continue
		}
		lng = val[:2]
		break
	}
	return
}
