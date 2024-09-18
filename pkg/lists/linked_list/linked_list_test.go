package linked_list

import (
	"testing"
)

func TestLinkedListPush(t *testing.T) {
	ll := New()

	tests := []struct {
		input            int
		expectedData     interface{}
		expectedPrevData interface{}
	}{
		{10, 10, nil},
		{20, 20, 10},
		{30, 30, 20},
	}

	for _, tt := range tests {
		ll.Add(tt.input)
		if ll.dataLast.data != tt.expectedData {
			t.Errorf("data is not %T. got=%d", tt.expectedData, ll.dataLast.data)
		}

		if ll.dataLast.prev != nil && ll.dataLast.prev.data != tt.expectedPrevData {
			t.Errorf("previous node data is not %T. got=%d", tt.expectedPrevData, ll.dataLast.prev.data)
		}
	}
}

func TestLinkedListRemove(t *testing.T) {
	ll := New()

	ll.Add(10)
	ll.Add(20)
	ll.Add(30)

	tests := []struct {
		removeIndex    int
		expectedData   []interface{}
		expectedLength int
	}{
		{1, []interface{}{10, 30}, 2},
		{0, []interface{}{30}, 1},
		{0, []interface{}{}, 0},
	}

	for _, tt := range tests {
		ll.Remove(tt.removeIndex)

		if ll.Size() != tt.expectedLength {
			t.Errorf("length is not %d. got=%d", tt.expectedLength, ll.Size())
		}

		for i, expectedData := range tt.expectedData {
			data, _ := ll.Get(i)
			if data != expectedData {
				t.Errorf("at index %d, expected data is %d. got=%d", i, expectedData, data)
			}
		}
	}
}

func TestLinkedListGet(t *testing.T) {
	ll := New()

	tests := []struct {
		input        int
		expectedData interface{}
	}{
		{0, 10},
		{1, 20},
		{2, 30},
		{-1, nil},
		{3, nil},
	}

	ll.Add(10)
	ll.Add(20)
	ll.Add(30)

	for _, tt := range tests {
		data, _ := ll.Get(tt.input)

		if data != tt.expectedData {
			t.Errorf("data is not %d. got=%d", tt.expectedData, data)
		}
	}
}

func TestLinkedListContains(t *testing.T) {
	ll := New()

	ll.Add(10)
	ll.Add(20)
	ll.Add(30)

	tests := []struct {
		input          interface{}
		expectedResult bool
	}{
		{10, true},
		{20, true},
		{30, true},
		{40, false},
		{nil, false},
	}

	for _, tt := range tests {
		result := ll.Contains(tt.input)
		if result != tt.expectedResult {
			t.Errorf("expected Contains(%v) to be %v, but got %v", tt.input, tt.expectedResult, result)
		}
	}
}

func TestLinkedListSwap(t *testing.T) {
	ll := New()

	ll.Add(10)
	ll.Add(20)
	ll.Add(30)
	ll.Add(40)

	tests := []struct {
		index1       int
		index2       int
		expectedData []interface{}
	}{
		{0, 1, []interface{}{20, 10, 30, 40}},  // Swap first two elements
		{1, 3, []interface{}{20, 40, 30, 10}},  // Swap second and last element
		{2, 2, []interface{}{20, 40, 30, 10}},  // Swap same element (no changes)
		{0, 4, []interface{}{20, 40, 30, 10}},  // Index out of bounds (no changes)
		{-1, 2, []interface{}{20, 40, 30, 10}}, // Negative index (no changes)
	}

	for _, tt := range tests {
		ll.Swap(tt.index1, tt.index2)

		for i, expectedData := range tt.expectedData {
			data, _ := ll.Get(i)
			if data != expectedData {
				t.Errorf("at index %d, expected data is %v, but got %v", i, expectedData, data)
			}
		}
	}
}

func TestLinkedListInsert(t *testing.T) {
	ll := New()

	ll.Add(10)
	ll.Add(20)
	ll.Add(30)

	tests := []struct {
		index        int
		data         interface{}
		expectedData []interface{}
	}{
		{0, 5, []interface{}{5, 10, 20, 30}},           // Insert at the beginning
		{2, 15, []interface{}{5, 10, 15, 20, 30}},      // Insert in the middle
		{5, 35, []interface{}{5, 10, 15, 20, 30, 35}},  // Insert at the end
		{7, 40, []interface{}{5, 10, 15, 20, 30, 35}},  // Insert out of bounds
		{-1, 50, []interface{}{5, 10, 15, 20, 30, 35}}, // Negative index
	}

	for _, tt := range tests {
		ll.Insert(tt.index, tt.data)

		// Verify the data in the linked list if the insertion was successful
		for i, expectedData := range tt.expectedData {
			data, _ := ll.Get(i)
			if data != expectedData {
				t.Errorf("at index %d, expected data is %v, but got %v", i, expectedData, data)
			}
		}

	}
}
