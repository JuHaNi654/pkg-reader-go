package system

import (
	"bufio"
	"os"
	"strings"

	"github.com/JuHaNi654/pkg-reader/pkg/models"
	"golang.org/x/exp/slices"
)

func GetPkgs() ([]*models.Pkg, error) {
	pkgs := []*models.Pkg{}
	file, err := os.Open("./mock/status.real")
	if err != nil {
		return nil, err
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)
	tmp := map[string]interface{}{}
	prev := ""

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			pkgs = append(pkgs, models.NewPkg(tmp))
			tmp = map[string]interface{}{}
			continue
		}

		if strings.HasPrefix(line, " ") {
			v, _ := tmp[prev].(string)
			tmp[prev] = v + line
			continue
		}

		property := strings.SplitN(line, ":", 2)
		prev = property[0]
		tmp[prev] = strings.TrimSpace(property[1])
	}

	// Append last pkg
	pkgs = append(pkgs, models.NewPkg(tmp))

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	for _, i := range pkgs {
		for _, j := range pkgs {
			if slices.Contains(j.Depends_on, i.Name) {
				i.Dependees = append(i.Dependees, j.Name)
			}
		}
	}

	return pkgs, nil
}
