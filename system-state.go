package molu

type SystemState struct {
	Ok         bool         `json:"ok"`
	Subsystems []*subsystem `json:"subsystems"`
}

func (systemState *SystemState) update(s *subsystem) {
	found, index := findSubSystemIndex(s.Name, systemState.Subsystems)
	if found {
		systemState.Subsystems[index] = s
	} else {
		systemState.Subsystems = append(systemState.Subsystems, s)
	}
	systemState.check()
}

func findSubSystemIndex(name string, subsystems []*subsystem) (found bool, index int) {
	var item *subsystem
	for index, item = range subsystems {
		if item.Name == name {
			found = true
			break
		}
	}
	return
}

func (systemState *SystemState) check() {
	for _, s := range systemState.Subsystems {
		if !s.Ok {
			systemState.Ok = false
			return
		}
	}
	systemState.Ok = true
}

type subsystem struct {
	Name string `json:"name"`
	Ok   bool   `json:"ok"`
	Msg  string `json:"msg"`
}
