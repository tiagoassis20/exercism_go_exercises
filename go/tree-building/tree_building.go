package tree

import (
	"fmt"
	"sort"
)

type Record struct {
	ID, Parent int
}

// ByParent implements sort.Interface for []Record based on
// the Parent field.
type ByParent []Record

func (a ByParent) Len() int      { return len(a) }
func (a ByParent) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByParent) Less(i, j int) bool {
	return a[i].Parent < a[j].Parent || (a[i].Parent == a[j].Parent && a[i].ID < a[j].ID)
}

type Node struct {
	ID       int
	Children []*Node
}

func Build(records []Record) (*Node, error) {

	if len(records) == 0 {
		return nil, nil
	}
	recordsMap, err := parseRecords(records)
	if err != nil {
		return nil, err
	}
	return buildNode(recordsMap, 0), nil
}

func parseRecords(records []Record) (map[int][]int, error) {
	sort.Sort(ByParent(records))

	recordsMap := make(map[int][]int)
	s := 0
	for _, rec := range records {
		if rec.ID <= rec.Parent && rec.ID != 0 {
			return nil, fmt.Errorf("Parent greater then child")
		}
		recordsMap[rec.Parent] = append(recordsMap[rec.Parent], rec.ID)
		s += rec.ID
	}
	if s != (len(records)-1)*len(records)/2 {
		return nil, fmt.Errorf("error on the entries")
	}
	if list, ok := recordsMap[0]; !ok || list[0] != 0 {
		return nil, fmt.Errorf("No root element")
	}
	return recordsMap, nil
}

func buildNode(recordsMap map[int][]int, parent int) *Node {
	node := new(Node)
	node.ID = parent
	if children, ok := recordsMap[parent]; ok {
		for _, id := range children {
			if id != 0 {
				node.Children = append(node.Children, buildNode(recordsMap, id))
			}
		}
	}
	return node
}
