package msgraph

import (
	"reflect"
	"strings"
	"testing"
)

func testName(t *testing.T, q GraphQueryOption, expectedName string) {
	if q.Name() != expectedName {
		t.Errorf("Expected %s, got %s", expectedName, q.Name())
	}
}

func TestSelectQuery(t *testing.T) {
	var q SelectQuery
	var props []string

	testName(t, q, "$select")

	// One item in the list should not include a comma.
	t.Log("Testing single property")
	props = append(props, "prop1")
	q.AddProperty("prop1")
	if reflect.DeepEqual(q.Properties, props) == false {
		t.Errorf("Expected %v, got %v", props, q.Properties)
	}
	if q.Value() != "prop1" {
		t.Errorf("Expected \"prop1\", got \"%s\"", q.Value())
	}

	// Two items should include a comma.
	t.Log("Testing two properties")
	props = append(props, "prop2")
	q.AddProperty("prop2")
	if reflect.DeepEqual(q.Properties, props) == false {
		t.Errorf("Expected %v, got %v", props, q.Properties)
	}
	if q.Value() != "prop1,prop2" {
		t.Errorf("Expected \"prop1,prop2\", got \"%s\"", q.Value())
	}

	// Make sure two different arrays are, in fact, different.
	t.Log("Testing differing slices")
	props = append(props, "prop3")
	if reflect.DeepEqual(q.Properties, props) == true {
		t.Errorf("Expected %v != %v", props, q.Properties)
	}

	t.Log("Testing NewSelectQuery()")
	props = []string{"p1", "p2", "p3", "p4"}
	q = NewSelectQuery(props)
	if reflect.DeepEqual(q.Properties, props) == false {
		t.Errorf("Expected %v, got %v", props, q.Properties)
	}
	if q.Value() != strings.Join(props, ",") {
		t.Errorf("Expected %v, got %v", props, q.Value())
	}
}

func TestOrderByQuery(t *testing.T) {
	var q OrderByQuery

	testName(t, q, "$orderby")
}
