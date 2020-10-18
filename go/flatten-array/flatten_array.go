package flatten

func Flatten(input interface{}) []interface{} {
	if v, ok := input.([]interface{}); ok {
		ret := []interface{}{}
		for _, e := range v {
			for _, el := range Flatten(e) {
				if el != nil {
					ret = append(ret, el)
				}
			}
		}
		return ret
	} else {
		return []interface{}{input}
	}
}
