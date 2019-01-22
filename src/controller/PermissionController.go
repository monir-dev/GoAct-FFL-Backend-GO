package controller

import (
	"Structure/src/model"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"strings"
)

func GetPermissions(w http.ResponseWriter, r *http.Request)  {
	roleId := mux.Vars(r)["role_id"]

	// All permission
	var permissions []struct{
		ID string `json:"id"`
		Name string `json:"name"`
		Module string `json:"module"`
	}
	DB.Table("permissions").Scan(&permissions)

	var modules []struct{
		Module string `json:"module"`
		Count int `json:"count"`
	}
	DB.Table("permissions").Select("module, count(module) as count").Group("module").Scan(&modules)

	// Specified roles permission
	var thisRolePermissions string
	row := DB.Raw("select GROUP_CONCAT(permissions.id) as thisRolePermissions from permissions join permission_role ON permissions.id = permission_role.`permission_id` join roles on roles.id = permission_role.`role_id` where roles.id = ?", roleId).Row()
	row.Scan(&thisRolePermissions)

	// response
	var data = map[string]interface{} {
		"permissions" : permissions,
		"modules": modules,
		"thisRolePermissions": thisRolePermissions,
	}
	SendSuccessResponse(w, data)
}

func AssignPermissionsToRole(w http.ResponseWriter, r *http.Request)  {
	roleId, _ := strconv.Atoi(mux.Vars(r)["role_id"])
	var permissionRole []model.PermissionRole
	DB.Table("permission_role").Where("role_id =?", roleId).Delete(&permissionRole)

	var body map[string]string
	json.NewDecoder(r.Body).Decode(&body)
	permissionIdArray := strings.Split(body["ids"], ",")

	for pos := range permissionIdArray {
		pId, _ := strconv.Atoi(permissionIdArray[pos])

		var temp model.PermissionRole
		temp.RoleId = roleId
		temp.PermissionId = pId

		//permissionRole = append(permissionRole, temp)
		DB.Table("permission_role").Create(&temp)
	}

	// response
	var data = map[string]string{}
	SendSuccessResponse(w, data)
}