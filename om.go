// Package om handles data in ISO19156 Observations &
// Measurements, with specific reference to the OM-JSON encoding
package om

// Observations and Measurements - JSON encoding types

// Href contains a URL link
type Href struct {
	Href string `json:"href"`
}

// Result contains a Result value and a URL to the definition of the unit-
//      of measure for the Result Value
type Result struct {
	Value float32 `json:"value"`
	Uom   string  `json:"uom"`
}

// Member describes an individual Member of an Observation Collection
type Member struct {
	ID               string `json:"id"`
	OmType           string `json:"type"`
	ObservedProperty Href   `json:"observedProperty"`
	Procedure        Href   `json:"procedure"`
	ResultTime       string `json:"resultTime"`
	Result           Result `json:"result"`
}

// Instant is for use within a Phenomenon Time, to constrain Phenomen Time to a
// single time instant
type Instant struct {
	Instant string `json:"instant"`
}

// ObservationCollection is for grouping a number of Observations against the
// same phenomenon time and feature of interest
type ObservationCollection struct {
	ID                string   `json:"id"`
	FeatureOfInterest Href     `json:"featureOfInterest"`
	PhenomenonTime    Instant  `json:"phenomenonTime"`
	Member            []Member `json:"member"`
}

// SetMember creates an individual Member for inclusion in an Observataion
// Collection. It takes the following inputs:
//
//		id			A string giving an identifier to the Member
//		omType		A string defining the Observation Type of the Member
//		obsProp		A string giving a URL defining the Observable Property
//					this Member refers to
//		procedure	A string giving a URL defining the Procedure used to
//					generate this Member Observation
//		resultTime	A string giving the reult time
//		result		A float32 number giving the value of the result
//		uom			A string giving a URL defining the unit of measure of the
//					result
//
//	TODO: 	1.	Error handling on omType
//			2.	Checking that URLs have expected form
//			3.  Checking of result time
//
func SetMember(id string, omType string, obsProp string, procedure string,
	resultTime string, result float32, uom string) Member {
	return Member{id, omType, Href{obsProp}, Href{procedure}, resultTime,
		Result{result, uom}}
}

// SetObservationCollection creates an obserbvation colletion from an id, a 
// feature of interest, a phenomenon time and any number of member observations. 
// It takes the followinginputs:
//
//		id			A string giving an identifier to the ObservationCollection
//		foi			A dtring giving a URL to a definition of the Feature of
//					Interest
//		phenomTime	A string giving the phenomenon time
//		member[s]	Any number of om.Member objects
//
//	TODO: 	1.	Checking that URLs have expected form
//			3.  Checking of phenomenon time
//
func SetObservationCollection(id string, foi string, phenomTime string,
	members ...Member) ObservationCollection {
	return ObservationCollection{id, Href{foi}, Instant{phenomTime}, members}
}
