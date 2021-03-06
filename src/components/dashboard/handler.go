package dashboard

import (
	"net/http"

	"github.com/gorilla/mux"
)

func MakeHTTPHandler() http.Handler {
	router := mux.NewRouter()
	router.HandleFunc("/dashboard", populateDashboard).
		Methods("POST").
		Schemes("http")
	router.HandleFunc("/dashboard/module_group", getModuleGroupInfo).
		Methods("POST").
		Schemes("http")

	router.HandleFunc("/dashboard/toggle_auto", toggleAuto).
		Methods("POST").
		Schemes("http")

	router.HandleFunc("/dashboard/set_env_param", setEnvironmentParameters).
		Methods("POST").
		Schemes("http")

	router.HandleFunc("/dashboard/reset_timer", resetTimer).
		Methods("POST").
		Schemes("http")

	router.HandleFunc("/dashboard/update_device_status", updateDeviceStatus).
		Methods("POST").
		Schemes("http")

	router.HandleFunc("/dashboard/history", history).
		Methods("POST").
		Schemes("http")

	return router
}

func populateDashboard(w http.ResponseWriter, r *http.Request) {
	PopulateDashboard(w, r)
}

func toggleAuto(w http.ResponseWriter, r *http.Request) {
	ToggleAuto(w, r)
}

func setEnvironmentParameters(w http.ResponseWriter, r *http.Request) {
	SetEnvironmentParameters(w, r)
}

func resetTimer(w http.ResponseWriter, r *http.Request) {
	ResetTimer(w, r)
}

func updateDeviceStatus(w http.ResponseWriter, r *http.Request) {
	UpdateDeviceStatus(w, r)
}

func history(w http.ResponseWriter, r *http.Request) {
	GetSensorDataHistory(w, r)
}

func getModuleGroupInfo(w http.ResponseWriter, r *http.Request) {
	GetModuleGroupInfo(w, r)
}
