package modulegroup_management

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type DemoJson struct {
	DemoName  string
	DemoArray []string
}

func GetAllModuleGroup(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World!")
}

func GetDemoJson(w http.ResponseWriter, r *http.Request) {
	demoJson := DemoJson{"Name", []string{"item 1", "item 2"}}
	js, err := json.Marshal(demoJson)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
