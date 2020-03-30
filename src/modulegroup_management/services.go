package modulegroup_management

import(
	"fmt"
	"net/http"
)
type Services interface {
	getAllModuleGroup(w http.ResponseWriter, r *http.Request) {
}



func getAllModuleGroup(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World!")
}