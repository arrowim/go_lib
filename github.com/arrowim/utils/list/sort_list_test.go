package list

import (
	"fmt"
	"testing"
)

type IntegerNode struct{
	Data int
}

func createIntegerNode (d int) *IntegerNode{
	s := IntegerNode{}
	s.Data = d
	return &s
}

func (self *IntegerNode) Compare(comparator Comparator) int{
	if v,ok:= comparator.(*IntegerNode);ok{

		if v.Data > self.Data{
			return -1;
		}else if v.Data < self.Data{
			return 1;
		}else{
			return 0;
		}
	}

	return -2;
}

func print (l *SortList){
	for _,v1 := range l.Data{
		if v,ok:= v1.(*IntegerNode);ok{
			fmt.Println(v.Data)
		}
	}
	fmt.Println("")
}

func TestSortList_Add(t *testing.T) {
	var l =&SortList{}

	l.Add(createIntegerNode(1))
	//print(l)

	l.Add(createIntegerNode(3))
	//print(l)

	l.Add(createIntegerNode(2))
	//print(l)

	l.Add(createIntegerNode(5))

	l.Add(createIntegerNode(5))
	l.Add(createIntegerNode(5))
	l.Add(createIntegerNode(5))
	l.Add(createIntegerNode(101))


	print(l)



}
