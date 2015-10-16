// setmember_test.go
package om

import (
	"testing"
	"reflect"
)

func TestSetObservationCollection(t *testing.T) {
	tMember := []Member{Member{Id: "someID", OmType: "Measurement",
		ObservedProperty: Href{"http://example.com/foo"},
		Procedure:        Href{"http://example.com/bar"},
		ResultTime:       "2015-10-15T10:10:10Z",
		Result:           Result{15, "http://example.com/baz"}}}
	tSetMember := SetMember("someID", "Measurement", "http://example.com/foo",
		"http://example.com/bar", "2015-10-15T10:10:10Z", 15,
		"http://example.com/baz")
	expected := ObservationCollection{Id: "someOtherId",
		FeatureOfInterest: Href{"http://example.com/qux"},
		PhenomenonTime: Instant{"2015-10-15T10:10:10Z"},Member: tMember}
	tr := SetObservationCollection("someOtherId","http://example.com/qux",
		"2015-10-15T10:10:10Z",tSetMember)
	if reflect.DeepEqual(tr, expected) == false {
		t.Errorf("SetMember failed unit test...")
	}
}
