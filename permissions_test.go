package ionic

import "testing"

func TestAllPermissionsHaveObjectTypes(t *testing.T) {
	for _, permission := range AllPermission {
		if objectTypeByPermission[permission] == "" {
			t.Errorf("objectTypeByPermission map missing entry for %s", permission)
		}
	}
}
