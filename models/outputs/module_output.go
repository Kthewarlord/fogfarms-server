package outputs

import (
	"github.com/KitaPDev/fogfarms-server/models"
)

type ModuleOutput struct {
	ModuleID         int    `json:"module_id"`
	ModuleGroupID    int    `json:"module_group_id"`
	ModuleLabel      string `json:"module_label"`
	ModuleGroupLabel string `json:"module_group_label"`
}

type Dashboardoutput struct {
	Sensordata     models.SensorData `json:"sensor_module"`
	Device         DashBoardModule   `json:"controller"`
	NutrientAmount int               `json:"nutrient_amount"`
}

type DashBoardModule struct {
	Fogger        []bool `json:"fogger"`
	LED           []bool `json:"led"`
	Mixer         []bool `json:"mixer"`
	SolenoidValve []bool `json:"solenoid_valve"`
}
