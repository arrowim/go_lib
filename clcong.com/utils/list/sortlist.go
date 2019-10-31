package list

type Comparator interface {
	// -1是小于，0是等于，1是大于
	Compare(comparator Comparator) int
}

type SortList struct {
	Data []Comparator
}

func (self *SortList) Add(comparator Comparator) {

	if len(self.Data) <= 1 {
		self.Data = append(self.Data, comparator)
		return
	}

	_, i := self.find(comparator)

	d := []Comparator{}

	if i == -100 {
		return
	}
	if i < 0 {
		d = []Comparator{comparator}
		self.Data = append(d, self.Data...)
	} else if i >= len(self.Data) {
		d = []Comparator{comparator}
		self.Data = append(self.Data, d...)
	} else {
		d = append(d, self.Data[0:i+1]...)
		c := self.Data[i+1:]
		d = append(d, comparator)

		d = append(d, c...)
		self.Data = d
	}

}

func (self *SortList) PopFront() Comparator {
	d := self.Data[0]
	self.Data = self.Data[1:]

	return d
}

func (self *SortList) PopBack() Comparator {

	if len(self.Data) == 0 {
		return nil
	}

	d := self.Data[len(self.Data)-1]

	n := []Comparator{}

	n = append(n, self.Data[0:len(self.Data)-1]...)

	self.Data = n

	return d
}

func (self *SortList) Front() Comparator {

	if len(self.Data) == 0 {
		return nil
	}

	d := self.Data[len(self.Data)-1]

	return d
}

func (self *SortList) Find(comparator Comparator) Comparator {
	c, _ := self.find(comparator)
	return c
}

func (self *SortList) find(comparator Comparator) (Comparator, int) {
	left := 0
	right := len(self.Data) - 1
	i := 0
	for left <= right {

		mid := (left + right) / 2
		d := self.Data[mid]

		c := comparator.Compare(d)
		if c == 0 {
			return d, mid
		} else if c == -1 {
			right = mid - 1
			i = mid - 1
		} else if c == 1 {
			left = mid + 1
			i = mid + 1
		} else {
			return nil, -100
		}

	}

	return nil, i
}
