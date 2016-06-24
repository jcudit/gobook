package my_oslice

import (
	"testing"
)

func TestNew(t *testing.T) {
	var _ OrderedSlice = New(func(a interface{}, b interface{}) bool {
		return a.(float64) < b.(float64)
	})
}

func TestNewStringSlice(t *testing.T) {
	var _ OrderedSlice = NewStringSlice()
}

func TestNewIntSlice(t *testing.T) {
	var _ OrderedSlice = NewIntSlice()
}

func TestLen(t *testing.T) {
	var o OrderedSlice = New(func(a interface{}, b interface{}) bool {
		return a.(float64) < b.(float64)
	})
	o.Add(float64(99.999))

	var s OrderedSlice = NewStringSlice()
	s.Add("hello")

	var i OrderedSlice = NewIntSlice()
	i.Add(9)

	if o.Len() != 1 || i.Len() != 1 || s.Len() != 1 {
		t.Fatalf("Length value incorrectly reported: %d %d %d", o.Len(), i.Len(), s.Len())
	}
}

func TestAdd(t *testing.T) {
	o := New(func(a interface{}, b interface{}) bool {
		return a.(float64) < b.(float64)
	})

	// Adding a generic element to an empty slice
	o.Add(float64(99.999))
	if o.Len() != 1 {
		t.Fatal("Adding generic element to ordered slice failed")
	}

	// Adding a string element to an empty slice
	var s OrderedSlice = NewStringSlice()
	s.Add("hello")
	if s.Len() != 1 {
		t.Fatalf("Adding string element to ordered slice failed")
	}

	// Adding an integer element to an empty slice
	var i OrderedSlice = NewIntSlice()
	i.Add(9)
	if i.Len() != 1 {
		t.Fatal("Adding integer element to ordered slice failed")
	}

	// Adding to uni-element slice
	i.Add(8)
	if i.Len() != 2 {
		t.Fatal("Adding integer element to the front of uni-element ordered slice failed")
	}
	i.Clear()
	i.Add(9)
	i.Add(10)
	if i.Len() != 2 {
		t.Fatal("Adding integer element to the back uni-element ordered slice failed")
	}

	// Adding to a multi-element slice
	i.Add(8)
	if i.Len() != 3 {
		t.Fatal("Adding integer element to the front of a multi-element ordered slice failed")
	}
	i.Clear()
	i.Add(9)
	i.Add(10)
	i.Add(11)
	if i.Len() != 3 {
		t.Fatal("Adding integer element to the back of a multi-element ordered slice failed")
	}
	i.Clear()
	i.Add(9)
	i.Add(11)
	i.Add(10)
	if i.Len() != 3 {
		t.Fatal("Adding integer element to the middle of a multi-element ordered slice failed")
	}

	// Adding repeat values to the slice
	i.Add(9)
	i.Add(10)
	i.Add(11)
	if i.Len() != 6 {
		t.Fatal("Adding repeat values to ordered slice failed")
	}

}

func TestClear(t *testing.T) {
	var o OrderedSlice = New(func(a interface{}, b interface{}) bool {
		return a.(float64) < b.(float64)
	})
	o.Clear()
	if o.Len() != 0 {
		t.Fatal("Clearing ordered generic-type slice failed")
	}

	var s OrderedSlice = NewStringSlice()
	s.Clear()
	if s.Len() != 0 {
		t.Fatal("Clearing ordered string slice failed")
	}

	var i OrderedSlice = NewIntSlice()
	i.Clear()
	if i.Len() != 0 {
		t.Fatal("Clearing ordered integer slice failed")
	}
}

func TestRemove(t *testing.T) {
	o := New(func(a interface{}, b interface{}) bool {
		return a.(float64) < b.(float64)
	})

	// Removing an intermediate element
	o.Add(float64(3.999))
	o.Add(float64(2.999))
	o.Add(float64(1.999))
	o.Len()
	if o.Remove(float64(2.999)) != true && o.Len() != 2 {
		t.Fatal("Failed to remove existing element")
	}

	// Removing a non-existent element
	if o.Remove(float64(88.888)) != false && o.Len() != 3 {
		t.Fatal("Failed to fail removing a non-existing element")
	}

	// Removing the only element
	o.Clear()
	o.Add(float64(99.999))
	if o.Remove(float64(99.999)) != true && o.Len() != 0 {
		t.Fatal("Failed to remove the only element")
	}

	// Removing the first element
	o.Clear()
	o.Add(float64(99.999))
	o.Add(float64(97.999))
	o.Add(float64(96.999))
	if o.Remove(float64(96.999)) != true && o.Len() != 2 {
		t.Fatal("Failed to remove the first element")
	}

	// Removing the last element
	o.Clear()
	o.Add(float64(99.999))
	o.Add(float64(97.999))
	o.Add(float64(96.999))
	if o.Remove(float64(99.999)) != true && o.Len() != 2 {
		t.Fatal("Failed to remove the first element")
	}
}

func TestIndex(t *testing.T) {
	o := New(func(a interface{}, b interface{}) bool {
		return a.(float64) < b.(float64)
	})
	o.Add(float64(99.999))
	o.Add(float64(97.999))
	o.Add(float64(96.999))
	if o.Index(float64(96.999)) != 0 && o.Index(float64(97.999)) != 1 &&
		o.Index(float64(99.999)) != -1 && o.Index(float64(99.999)) != 2 {
		t.Fatal("Failed to correctly detect indices of all elements")
	}
}

func TestAt(t *testing.T) {
	o := New(func(a interface{}, b interface{}) bool {
		return a.(float64) < b.(float64)
	})
	o.Add(float64(99.999))
	o.Add(float64(97.999))
	o.Add(float64(96.999))
	if o.At(0) != float64(96.999) && o.At(1) != float64(97.999) &&
		o.At(2) != float64(99.999) && o.At(99) != nil {
		t.Fatal("Failed to return valid elements at all indices")
	}
}

func TestIntLess(t *testing.T) {
	var i OrderedSlice = NewIntSlice()
	i.Add(9)
	if i.Len() != 1 {
		t.Fatal("Adding integer element to ordered slice failed")
	}
}
