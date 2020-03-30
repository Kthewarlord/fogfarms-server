package modulegroup_management

import(
	"fmt"
	"net/http"
	"github.com/KitaPDev/fogfarms-server/tree/demoReturn"
	"encoding/json"
)
type Services interface {
	getAllModuleGroup(w http.ResponseWriter, r *http.Request) {
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
}