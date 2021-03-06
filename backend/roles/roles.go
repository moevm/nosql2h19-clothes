package roles

import (
	"encoding/json"
	"fmt"
	"github.com/mikespook/gorbac"
	"log"
	"os"
)

var Permissions = make(gorbac.Permissions)
var Rbac = gorbac.New()

const AdminRole = "admin"
const UserRole = "user"

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
		fmt.Println("The admin has been granted permissions")
	} else {
		fmt.Println("The admin has NOT been granted permissions")
	}

	if Rbac.IsGranted(UserRole, Permissions["create-worker"], nil) {
		fmt.Println("The user has been granted permissions")
	} else {
		fmt.Println("The user has NOT been granted permissions")
	}

	return true
}
