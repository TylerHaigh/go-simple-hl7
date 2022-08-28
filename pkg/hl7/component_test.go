package hl7

import (
	"testing"
)

func TestNewComponent_SingleComponent(t *testing.T) {
	c := NewComponent([]SubComponent{"S"})

	expected := SubComponent("S")
	actual := c.Data[0]

	if actual != expected {
		t.Errorf("got %s but wanted %s", actual, expected)
	}
}

func TestNewComponent_MultipleSubComponents(t *testing.T) {
	subComponents := []SubComponent{"S", "C"}
	c := NewComponent(subComponents)

	if len(c.Data) != len(subComponents) {
		t.Errorf("Length of Component is not %d", len(subComponents))
	}

	for i, expectedSc := range subComponents {

		actualSc := c.Data[i]

		if actualSc != expectedSc {
			t.Errorf("got %s but wanted %s", actualSc, expectedSc)
		}
	}
}

func TestParseComponent_SingleComponent(t *testing.T) {
	s := "M"
	c := ParseComponent(s)

	expected := SubComponent("M")
	actual := c.Data[0]

	if actual != expected {
		t.Errorf("Expected %s but got %s", expected, actual)
	}
}

func TestParseComponent_MultipleComponents(t *testing.T) {
	s := "M&A&B"
	c := ParseComponent(s)

	components := []SubComponent{"M", "A", "B"}
	for i, expectedSc := range components {

		actualSc := c.Data[i]

		if actualSc != expectedSc {
			t.Errorf("Expected %s but got %s", expectedSc, actualSc)
		}
	}
}

func TestToString_SingleComponent(t *testing.T) {
	c := NewComponent([]SubComponent{"S"})
	expected := "S"
	actual := c.ToString(StandardDelimters())

	if actual != expected {
		t.Errorf("Expected %s to be %s", actual, expected)
	}
}

func TestToString_MultipleComponents(t *testing.T) {
	c := NewComponent([]SubComponent{"S", "A", "B"})
	expected := "S&A&B"
	actual := c.ToString(StandardDelimters())

	if actual != expected {
		t.Errorf("Expected %s to be %s", actual, expected)
	}
}

func TestGetSubComponent_SingleComponent(t *testing.T) {
	c := NewComponent([]SubComponent{"S"})
	expected := SubComponent("S")
	actual := c.GetSubComponent(1)

	if actual != expected {
		t.Errorf("Expected %s to be %s", actual, expected)
	}
}

func TestGetSubComponent_MultipleComponents(t *testing.T) {
	c := NewComponent([]SubComponent{"S", "A"})
	expected := SubComponent("A")
	actual := c.GetSubComponent(2)

	if actual != expected {
		t.Errorf("Expected %s to be %s", actual, expected)
	}
}

func TestGetSubComponent_OutOfBounds(t *testing.T) {
	c := NewComponent([]SubComponent{"S"})
	expected := SubComponent("")
	actual := c.GetSubComponent(10)

	if actual != expected {
		t.Errorf("Expected %s to be %s", actual, expected)
	}
}

func TestSetFromString(t *testing.T) {
	c := NewComponent([]SubComponent{})
	c.SetFromString("A&B&C")

	components := []SubComponent{"A", "B", "C"}
	for i, expectedSc := range components {

		actualSc := c.Data[i]

		if actualSc != expectedSc {
			t.Errorf("Expected %s but got %s", expectedSc, actualSc)
		}
	}
}

func TestSet(t *testing.T) {
	components := []SubComponent{"A", "B", "C"}
	c := NewComponent([]SubComponent{})
	c.Set(components)

	for i, expectedSc := range components {

		actualSc := c.Data[i]

		if actualSc != expectedSc {
			t.Errorf("Expected %s but got %s", expectedSc, actualSc)
		}
	}
}
