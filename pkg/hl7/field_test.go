package hl7

import "testing"

func TestParseField_SingleComponentField(t *testing.T) {
	str := "1234567FF"
	f := ParseField(str)

	comp := string(f.Components[0].GetSubComponent(1))
	if comp != str {
		t.Errorf("Exepected %s to be %s", comp, str)
	}
}

func TestParseField_MultipleComponentField(t *testing.T) {
	str := "1234567FF^^^^PRN"
	f := ParseField(str)

	providerNumber := string(f.Components[0].GetSubComponent(1))
	typeCode := string(f.Components[4].GetSubComponent(1))

	if providerNumber != "1234567FF" {
		t.Errorf("Exepected %s to be %s", providerNumber, "1234567FF")
	}

	if typeCode != "PRN" {
		t.Errorf("Exepected %s to be %s", typeCode, "PRN")
	}
}

func TestParseField_SubComponent(t *testing.T) {
	str := "SWAB&Swab&L^^^NOSE&Nose&L"
	f := ParseField(str)

	code := string(f.Components[0].GetSubComponent(1))
	description := string(f.Components[0].GetSubComponent(2))
	system := string(f.Components[0].GetSubComponent(3))

	if code != "SWAB" {
		t.Errorf("Exepected %s to be %s", code, "SWAB")
	}

	if description != "Swab" {
		t.Errorf("Exepected %s to be %s", description, "Swab")
	}

	if system != "L" {
		t.Errorf("Exepected %s to be %s", system, "L")
	}
}

func TestNewField_SingleComponent(t *testing.T) {
	f := NewField([]*Component{
		{Data: []SubComponent{"SWAB"}},
	})

	comp := string(f.Components[0].GetSubComponent(1))
	if comp != "SWAB" {
		t.Errorf("Exepected %s to be %s", comp, "SWAB")
	}
}

func TestNewField_MultipleComponent(t *testing.T) {
	f := NewField([]*Component{
		{Data: []SubComponent{"SWAB"}},
		{Data: []SubComponent{"Swab"}},
		{Data: []SubComponent{"L"}},
	})

	code := string(f.Components[0].GetSubComponent(1))
	description := string(f.Components[1].GetSubComponent(1))
	system := string(f.Components[2].GetSubComponent(1))

	if code != "SWAB" {
		t.Errorf("Exepected %s to be %s", code, "SWAB")
	}

	if description != "Swab" {
		t.Errorf("Exepected %s to be %s", description, "Swab")
	}

	if system != "L" {
		t.Errorf("Exepected %s to be %s", system, "L")
	}
}

func TestFieldFromComponents_SingleComponent(t *testing.T) {
	str := []ComponentString{
		ComponentString("1234567FF"),
	}
	f := FieldFromComponents(str)

	comp := string(f.Components[0].GetSubComponent(1))
	if comp != string(str[0]) {
		t.Errorf("Exepected %s to be %s", comp, str)
	}
}

func TestFieldFromComponents_MultipleSubcomponents(t *testing.T) {
	str := []ComponentString{
		ComponentString("SWAB&Swab"),
	}
	f := FieldFromComponents(str)

	code := string(f.Components[0].GetSubComponent(1))
	if code != "SWAB" {
		t.Errorf("Exepected %s to be %s", code, "SWAB")
	}

	description := string(f.Components[0].GetSubComponent(2))
	if description != "Swab" {
		t.Errorf("Exepected %s to be %s", description, "Swab")
	}
}

func TestFieldFromComponents_MultipleComponents(t *testing.T) {
	str := []ComponentString{
		ComponentString("SWAB"),
		ComponentString("Swab"),
	}
	f := FieldFromComponents(str)

	code := string(f.Components[0].GetSubComponent(1))
	if code != "SWAB" {
		t.Errorf("Exepected %s to be %s", code, "SWAB")
	}

	description := string(f.Components[1].GetSubComponent(1))
	if description != "Swab" {
		t.Errorf("Exepected %s to be %s", description, "Swab")
	}
}

func TestFieldToString_SingleComponent(t *testing.T) {
	f := NewField([]*Component{
		{Data: []SubComponent{"SWAB"}},
	})

	str := f.ToString(StandardDelimters())

	if str != "SWAB" {
		t.Errorf("Exepected %s to be %s", str, "SWAB")
	}
}

func TestFieldToString_MultipleSubComponents(t *testing.T) {
	f := NewField([]*Component{
		{Data: []SubComponent{"SWAB", "Swab"}},
	})

	str := f.ToString(StandardDelimters())

	if str != "SWAB&Swab" {
		t.Errorf("Exepected %s to be %s", str, "SWAB&Swab")
	}
}

func TestFieldToString_MultipleComponents(t *testing.T) {
	f := NewField([]*Component{
		{Data: []SubComponent{"SWAB"}},
		{Data: []SubComponent{"Swab"}},
	})

	str := f.ToString(StandardDelimters())

	if str != "SWAB^Swab" {
		t.Errorf("Exepected %s to be %s", str, "SWAB^Swab")
	}
}

func TestFieldGetComponentString_SingleComponent(t *testing.T) {
	f := NewField([]*Component{
		{Data: []SubComponent{"SWAB"}},
	})

	str := f.GetComponentString(1)
	if str != "SWAB" {
		t.Errorf("Exepected %s to be %s", str, "SWAB")
	}
}

func TestFieldGetComponentString_MultipleSubComponents(t *testing.T) {
	f := NewField([]*Component{
		{Data: []SubComponent{"SWAB", "Swab"}},
	})

	str := f.GetComponentString(1)
	if str != "SWAB&Swab" {
		t.Errorf("Exepected %s to be %s", str, "SWAB&Swab")
	}
}

func TestFieldGetComponentString_MultipleComponents(t *testing.T) {
	f := NewField([]*Component{
		{Data: []SubComponent{"SWAB"}},
		{Data: []SubComponent{"Swab"}},
	})

	comp1 := f.GetComponentString(1)
	if comp1 != "SWAB" {
		t.Errorf("Exepected %s to be %s", comp1, "SWAB")
	}

	comp2 := f.GetComponentString(2)
	if comp2 != "Swab" {
		t.Errorf("Exepected %s to be %s", comp2, "Swab")
	}
}

func TestFieldGetComponentString_NonExistingComponent(t *testing.T) {
	f := NewField([]*Component{
		{Data: []SubComponent{"SWAB"}},
	})

	comp2 := f.GetComponentString(2)
	if comp2 != "" {
		t.Errorf("Exepected %s to be %s", comp2, "''")
	}
}

func TestFieldGetComponent_SingleComponent(t *testing.T) {
	f := NewField([]*Component{
		{Data: []SubComponent{"SWAB"}},
	})

	comp := f.GetComponent(1)
	if comp.ToString(StandardDelimters()) != "SWAB" {
		t.Errorf("Exepected %s to be %s", comp.ToString(StandardDelimters()), "SWAB")
	}
}

func TestFieldGetComponent_MultipleSubComponents(t *testing.T) {
	f := NewField([]*Component{
		{Data: []SubComponent{"SWAB", "Swab"}},
	})

	comp := f.GetComponent(1)
	if comp.ToString(StandardDelimters()) != "SWAB&Swab" {
		t.Errorf("Exepected %s to be %s", comp.ToString(StandardDelimters()), "SWAB&Swab")
	}
}

func TestFieldGetComponent_MultipleComponents(t *testing.T) {
	f := NewField([]*Component{
		{Data: []SubComponent{"SWAB"}},
		{Data: []SubComponent{"Swab"}},
	})

	comp1 := f.GetComponent(1)
	if comp1.ToString(StandardDelimters()) != "SWAB" {
		t.Errorf("Exepected %s to be %s", comp1.ToString(StandardDelimters()), "SWAB")
	}

	comp2 := f.GetComponent(2)
	if comp2.ToString(StandardDelimters()) != "Swab" {
		t.Errorf("Exepected %s to be %s", comp2.ToString(StandardDelimters()), "Swab")
	}
}

func TestFieldGetComponent_NonExistingComponent(t *testing.T) {
	f := NewField([]*Component{
		{Data: []SubComponent{"SWAB"}},
	})

	comp2 := f.GetComponent(2)
	if comp2 != nil {
		t.Errorf("Exepected %s to be %s", comp2, "nil")
	}
}

func TestFieldGetSubComponent_SingleComponent(t *testing.T) {
	f := NewField([]*Component{
		{Data: []SubComponent{"SWAB"}},
	})

	comp := f.GetSubComponent(1, 1)
	if comp != "SWAB" {
		t.Errorf("Exepected %s to be %s", comp, "SWAB")
	}
}

func TestFieldGetSubComponent_MultipleSubComponents(t *testing.T) {
	f := NewField([]*Component{
		{Data: []SubComponent{"SWAB", "Swab"}},
	})

	comp := f.GetSubComponent(1, 2)
	if comp != "Swab" {
		t.Errorf("Exepected %s to be %s", comp, "Swab")
	}
}

func TestFieldGetSubComponent_MultipleComponents(t *testing.T) {
	f := NewField([]*Component{
		{Data: []SubComponent{"SWAB", "A"}},
		{Data: []SubComponent{"Swab", "B"}},
	})

	comp2 := f.GetSubComponent(2, 1)
	if comp2 != "Swab" {
		t.Errorf("Exepected %s to be %s", comp2, "Swab")
	}
}

func TestFieldGetSubComponent_NonExistingComponent(t *testing.T) {
	f := NewField([]*Component{
		{Data: []SubComponent{"SWAB", "A"}},
	})

	comp2 := f.GetSubComponent(2, 1)
	if comp2 != "" {
		t.Errorf("Exepected %s to be %s", comp2, "''")
	}
}

func TestFieldGetSubComponent_NonExistingSubComponent(t *testing.T) {
	f := NewField([]*Component{
		{Data: []SubComponent{"SWAB", "A"}},
	})

	comp2 := f.GetSubComponent(1, 3)
	if comp2 != "" {
		t.Errorf("Exepected %s to be %s", comp2, "''")
	}
}
