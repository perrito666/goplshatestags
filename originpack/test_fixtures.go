// +build unit

package originpack

func CreateSomethingFromOrigin() *SomethingFromOrigin {
	return &SomethingFromOrigin{
		MeaninglessField: "really meaningless",
	}
}
