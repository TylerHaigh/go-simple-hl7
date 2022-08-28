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
