package rotation

import (
	"errors"
)

type Rotation struct {
	StartIdx int
	People   []string
	Roles    []string
}

type RoatationResult struct {
	Role       string
	Supervisor string
}

func (m *Rotation) Rotate() []RoatationResult {

	result := make([]RoatationResult, 0, len(m.Roles))
	for _, v := range m.Roles {
		err, supervisor := m.nextSupervisor()
		if err != nil {
			break
		}
		result = append(result, RoatationResult{Role: v, Supervisor: supervisor})
	}
	return result
}

func (m *Rotation) nextSupervisor() (error, string) {
	if len(m.People) == 0 || len(m.Roles) == 0 {
		return errors.New("no one people"), ""
	}

	val := m.People[m.StartIdx]
	m.idxIncreament()
	return nil, val
}
func (m *Rotation) idxIncreament() {
	m.StartIdx += 1
	for {
		if len(m.People) <= m.StartIdx {
			m.StartIdx -= len(m.People)
		} else {
			break
		}
	}
}
