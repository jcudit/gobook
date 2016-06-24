// Implements an ordered slice
package my_oslice

type OrderedSlice struct {
	oslice []interface{}
	less   func(interface{}, interface{}) bool
}

// New creates an ordered slice of any type and requires
// a less than function to be passed in
func New(less func(interface{}, interface{}) bool) OrderedSlice {
	return OrderedSlice{oslice: []interface{}{}, less: less}
}

// NewIntSlice creates an ordered slice of strings
func NewStringSlice() OrderedSlice {
	return OrderedSlice{
		oslice: []interface{}{},
		less: func(a interface{}, b interface{}) bool {
			return a.(string) < b.(string)
		},
	}
}

// NewIntSlice creates an ordered slice of integers
func NewIntSlice() OrderedSlice {
	return OrderedSlice{
		oslice: []interface{}{},
		less: func(a interface{}, b interface{}) bool {
			return a.(int) < b.(int)
		},
	}
}

// Len returns the number of items in the slice
func (s *OrderedSlice) Len() int {
	return len(s.oslice)
}

// Add inserts an item into the slice while maintaining order
func (s *OrderedSlice) Add(item interface{}) {
	if len(s.oslice) == 0 {
		// Add to empty slice
		s.oslice = append(s.oslice, item)
	} else if len(s.oslice) == 1 {
		// Add to a uni-element slice
		if s.less(item, s.oslice[0]) {
			// prepend when new item is smaller
			temp := []interface{}{}
			temp = append(temp, item)
			temp = append(temp, s.oslice...)
			s.oslice = temp
		} else {
			// append when new item is greater
			s.oslice = append(s.oslice, item)
		}
	} else {
		// Add to multi-element slice
		for i := range s.oslice {
			if s.less(item, s.oslice[i]) {
				if i == 0 {
					// prepend to entire slice when introducing smallest element
					temp := []interface{}{}
					temp = append(temp, item)
					temp = append(temp, s.oslice...)
					s.oslice = temp
					break
				} else {
					// insert within slice when introducing intermediate element
					temp := []interface{}{}
					temp = append(temp, s.oslice[:i]...)
					temp = append(temp, item)
					temp = append(temp, s.oslice[i:]...)
					s.oslice = temp
					break
				}
			}
			if i+1 == len(s.oslice) {
				// append to entire slice when introducing largest element
				s.oslice = append(s.oslice, item)
				break
			}
		}
	}
}

// Clear empties the slice
func (s *OrderedSlice) Clear() {
	s.oslice = []interface{}{}
}

// Removes the first occurrence of the specified item and
// returns whether or not the removal was successful
func (s *OrderedSlice) Remove(item interface{}) bool {
	for i := range s.oslice {
		if s.oslice[i] == item {
			if len(s.oslice) == 1 {
				s.Clear()
			} else if i+1 == len(s.oslice) {
				s.oslice = s.oslice[:i-1]
			} else {
				temp := []interface{}{}
				temp = append(temp, s.oslice[:i]...)
				temp = append(temp, s.oslice[i+1:]...)
				s.oslice = temp
			}
			return true
		}
	}
	return false
}

// Index reports the position of the first occurrence of the
// specified item or -1 if not found
func (s *OrderedSlice) Index(item interface{}) int {
	for i := range s.oslice {
		if s.oslice[i] == item {
			return i
		}
	}
	return -1
}

// At returns the item at the given position
func (s *OrderedSlice) At(index int) interface{} {
	if 0 <= index && index <= len(s.oslice)-1 {
		return s.oslice[index]
	} else {
		return nil
	}
}
