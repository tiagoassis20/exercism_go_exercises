package flatten

func Flatten(input interface{}) []interface{} {
	return flatten([]interface{}{}, input)
}

func flatten(base []interface{}, input interface{}) (ret []interface{}) {
	ret = base
	if v, ok := input.([]interface{}); ok {
		for _, e := range v {
			ret = flatten(ret, e)
		}
	} else {
		if input != nil {
			ret = append(ret, input)
		}
	}
	return
}
