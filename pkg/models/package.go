package models

import (
	"regexp"
	"strings"
)

type Pkg struct {
	Id          int64    `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Depends_on  []string `json:"depends_on"`
	Dependees   []string `json:"dependees"`
}

func NewPkg(v map[string]interface{}) *Pkg {
	var depends []string

	if d := v["Depends"]; d != nil {
		depends = getUniqDepends(d.(string))
	}

	return &Pkg{
		Name:        v["Package"].(string),
		Description: v["Description"].(string),
		Depends_on:  depends,
	}
}

func getUniqDepends(depends string) []string {
	re := regexp.MustCompile(`\s\(([^)]*)\)|\s`)
	depends = re.ReplaceAllString(depends, "")

	arr := strings.Split(depends, ",")
	uniq := make(map[string]bool)
	result := []string{}

	for _, depend := range arr {
		if !uniq[depend] {
			uniq[depend] = true
			result = append(result, strings.Split(depend, "|")...)
		}
	}

	return result
}
