package roles

import (
	"encoding/json"
	"log"
	"os"

	"fmt"
	"github.com/mikespook/gorbac"
)

var Permissions = make(gorbac.Permissions)
var Rbac = gorbac.New()

const AdminRole = "admin"
const ManagerRole = "manager"
const AccountantRole = "accountant"

func LoadJson(filename string, v interface{}) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	return json.NewDecoder(f).Decode(v)
}

func SaveJson(filename string, v interface{}) error {
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()
	return json.NewEncoder(f).Encode(v)
}

func Init() bool {
	// map[RoleId]PermissionIds
	var jsonRoles map[string][]string
	// map[RoleId]ParentIds
	var jsonInher map[string][]string
	// Load roles information
	if err := LoadJson("vendor/roles/roles.json", &jsonRoles); err != nil {
		log.Fatal(err)
		return false
	}
	// Load inheritance information
	if err := LoadJson("vendor/roles/inher.json", &jsonInher); err != nil {
		log.Fatal(err)
		return false
	}

	// Build roles and add them to goRBAC instance
	for rid, pids := range jsonRoles {
		role := gorbac.NewStdRole(rid)
		for _, pid := range pids {
			_, ok := Permissions[pid]
			if !ok {
				Permissions[pid] = gorbac.NewStdPermission(pid)
			}
			role.Assign(Permissions[pid])
		}
		Rbac.Add(role)
	}
	// Assign the inheritance relationship
	for rid, parents := range jsonInher {
		if err := Rbac.SetParents(rid, parents); err != nil {
			log.Fatal(err)
		}
	}

	if Rbac.IsGranted(AdminRole, Permissions["create-worker"], nil) {
		fmt.Println("The admin has been granted permis create-worker")
	} else {
		fmt.Println("The admin has NOT been granted permis cUser")
	}

	if Rbac.IsGranted(ManagerRole, Permissions["create-user"], nil) {
		fmt.Println("The manager has been granted permis create-user")
	} else {
		fmt.Println("The manager has NOT been granted permis create-user")
	}

	if Rbac.IsGranted(AccountantRole, Permissions["create-worker"], nil) {
		fmt.Println("The accountant has been granted permis create-worker")
	} else {
		fmt.Println("The accountant has NOT been granted permis create-worker")
	}

	return true
}
