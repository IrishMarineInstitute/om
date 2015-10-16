// setmember_test.go
package om

import (
	"testing"
)

func TestSetMember(t *testing.T) {
	expected := Member{Id: "someID", OmType: "Measurement",
		ObservedProperty: Href{"http://example.com/foo"},
		Procedure:        Href{"http://example.com/bar"},
		ResultTime:       "2015-10-15T10:10:10Z",
		Result:           Result{15, "http://example.com/baz"}}
	tr := SetMember("someID", "Measurement", "http://example.com/foo",
		"http://example.com/bar", "2015-10-15T10:10:10Z", 15,
		"http://example.com/baz")
	if expected != tr {
		t.Errorf("SetMember failed unit test...")
	}
}
