package modulegroup_management

import(
	"fmt"
	"net/http"
	"encoding/json"
)
type Services interface {
	getAllModuleGroup(w http.ResponseWriter, r *http.Request) 
	getDemoJson(w http.ResponseWriter, r *http.Request)
}

func getAllModuleGroup(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World!")
}

type DemoJson struct {
	DempName    string
	DemoArray []string
  }

func getDemoJson(w http.ResponseWriter, r *http.Request){
	demoJson := DemoJson{"Name", []string{"item 1","item 2"}}
	js, err := json.Marshal(demoJson)
	if err != nil {
	  http.Error(w, err.Error(), http.StatusInternalServerError)
	  return
	}
  
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}