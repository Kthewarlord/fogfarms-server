module github.com/ddfsdd/fogfarms-server

go 1.14

replace github.com/KitaPDev/fogfarms-server/modulegroup_management => github.com/ddfsdd/fogfarms-server/tree/master/src/modulegroup_management v0.0.4

require (
	github.com/KitaPDev/fogfarms-server v0.0.4
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/julienschmidt/httprouter v1.3.0
	github.com/labstack/gommon v0.3.0
	github.com/lib/pq v1.3.0
	golang.org/x/crypto v0.0.0-20200323165209-0ec3e9974c59
)
