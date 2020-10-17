package sublist

type Relation string

func Sublist(list1, list2 []int) Relation {

	if len(list1) < len(list2) && isSublist(list1, list2) {
		return "sublist"
	} else if len(list1) > len(list2) && isSublist(list2, list1) {
		return "superlist"
	} else if len(list1) == len(list2) && isSublist(list1, list2) {
		return "equal"
	} else {
		return "unequal"
	}
}

func isSublist(list1, list2 []int) bool {
	for i := 0; i <= len(list2)-len(list1); i++ {
		j := 0
		for j = 0; j < len(list1) && list1[j] == list2[i+j]; j++ {
		}

		if j == len(list1) {
			return true
		}
	}
	return false
}
