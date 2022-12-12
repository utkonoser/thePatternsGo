package singleton

import "testing"

func TestGetInstance(t *testing.T) {
	counter1 := GetInstance()

	if counter1 == nil {
		t.Error("expected pointer to Singleton after " +
			"calling GetInstance(), not nil\n")
	}
	expectedCounter := counter1

	currentCount := counter1.AddOne()
	if currentCount != 1 {
		t.Errorf("after AddOne() the count must be"+
			" 1 but it is %d\n", currentCount)
	}

	counter2 := GetInstance()
	if counter2 != expectedCounter {
		t.Error("expected same instance in counter2 but" +
			" it got a different instance\n")
	}

	currentCount = counter2.AddOne()
	if currentCount != 2 {
		t.Errorf("after AddOne() the count must be"+
			" 2 but it is %d\n", currentCount)
	}
}
