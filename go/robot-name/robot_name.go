package robotname

import (
	"fmt"
	"math/rand"
	"time"
)

const max = 26 * 26 * 10 * 10 * 10

var isNamesInit = false

var namesList = make([]string, max)

type Robot struct {
	name string
}

func (r *Robot) Name() (string, error) {
	var err error = nil
	if r.name == "" {
		r.name, err = generateName()
	}
	return r.name, err
}

func (r *Robot) Reset() {
	//namesList.PushBack(r.name)
	r.name = ""

}

func generateName() (string, error) {
	names := getListOfNames()
	rand.Seed(time.Now().UnixNano())
	if len(names) > 1 {
		return getName(rand.Intn(len(names) - 1))
	} else {
		return getName(0)
	}
}

func getListOfNames() []string {
	if !isNamesInit {
		for i := 0; i < max; i++ {
			namesList[i] = fmt.Sprintf("%s%03d", []byte{'A' + byte((i/26000)%26), 'A' + byte((i/1000)%26)}, i%1000)
		}

		isNamesInit = true
	}
	return namesList
}
func getName(n int) (string, error) {

	if n >= len(namesList) || n < 0 {
		return "", fmt.Errorf("Paremeter %d is Out Of Bounds", n)
	}
	if len(namesList) == 0 {
		return "", fmt.Errorf("All possible names are in use")
	}
	name := namesList[n]
	namesList = append(namesList[:n], namesList[n+1:]...)

	return name, nil

}
