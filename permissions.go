package ionic

var objectTypeByPermission = map[Permission]ObjectType{
	PermissionAllPermissions:         ObjectTypeOrganization,
	PermissionOrganizationModify:     ObjectTypeOrganization,
	PermissionOrganizationView:       ObjectTypeOrganization,
	PermissionOrganizationUserCreate: ObjectTypeOrganization,
	PermissionOrganizationUserModify: ObjectTypeOrganization,
	PermissionOrganizationUserRemove: ObjectTypeOrganization,
	PermissionSoftwareListModify:     ObjectTypeSoftwareList,
	PermissionSoftwareListView:       ObjectTypeSoftwareList,
}

func (permission *Permission) ObjectType() ObjectType {
	return objectTypeByPermission[*permission]
}
