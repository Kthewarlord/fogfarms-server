package models

type Module struct {
	ModuleID      int    `json:"module_id"`
	ModuleGroupID int    `json:"module_group_id"`
	Token         string `json:"token"`
	ModuleLabel   string `json:"module_label"`
	Fogger        []bool `json:"fogger"`
	LED           []bool `json:"led"`
	Mixer         []bool `json:"mixer"`
	SolenoidValve []bool `json:"solenoid_valve"`
}
