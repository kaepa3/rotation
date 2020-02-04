package rotation

import (
	"testing"
)

type RotationTest struct {
	Rotation Rotation
	Expected []RoatationResult
}

func Rotation_Zero(t *testing.T) {
	for _, v := range []RotationTest{
		RotationTest{Rotation: Rotation{}, Expected: []RoatationResult{}},
		RotationTest{Rotation: Rotation{Roles: []string{"jo", "hoge"}}, Expected: []RoatationResult{}},
		RotationTest{Rotation: Rotation{People: []string{"jo", "hoge"}}, Expected: []RoatationResult{}},
		RotationTest{Rotation: Rotation{
			Roles: []string{"jo", "hoge"}, People: []string{"one", "tow"}},
			Expected: []RoatationResult{
				RoatationResult{Role: "jo", Supervisor: "one"},
				RoatationResult{Role: "hoge", Supervisor: "two"},
			},
		},
		RotationTest{Rotation: Rotation{
			Roles: []string{"jo", "hoge", "sore"}, People: []string{"one", "tow"}},
			Expected: []RoatationResult{
				RoatationResult{Role: "jo", Supervisor: "one"},
				RoatationResult{Role: "hoge", Supervisor: "two"},
				RoatationResult{Role: "sore", Supervisor: "one"},
			},
		},
		RotationTest{Rotation: Rotation{
			Roles: []string{"jo", "hoge"}, People: []string{"one", "tow", "three"}},
			Expected: []RoatationResult{
				RoatationResult{Role: "jo", Supervisor: "one"},
				RoatationResult{Role: "hoge", Supervisor: "two"},
			},
		},
	} {
		result := v.Rotation.Rotate()
		if len(result) != len(v.Expected) {
			t.Error()
		}
		for idx, exp := range v.Expected {
			if exp.Role != result[idx].Role || exp.Supervisor != result[idx].Supervisor {
				t.Error(result[idx])
			}
		}
	}
}
