package modulegroup_management

import(
	"fmt"
	"net/http"
	//"encoding/json"
)
type Services interface {
	getAllModuleGroup(w http.ResponseWriter, r *http.Request) 
	getDemoJson(w http.ResponseWriter, r *http.Request)
}

type DemoJson struct {
	DemoName    string
	DemoArray []string
  }

  
func getAllModuleGroup(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World!")
}


func getDemoJson(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "Hello World!")
	// demoJson := DemoJson{"Name", []string{"item 1","item 2"}}
	// js, err := json.Marshal(demoJson)
	// if err != nil {
	//   http.Error(w, err.Error(), http.StatusInternalServerError)
	//   return
	// }
  
	// w.Header().Set("Content-Type", "application/json")
	// w.Write(js)
}