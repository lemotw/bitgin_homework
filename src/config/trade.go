package config

const (
	USERROLE_NORMAL        = "N"
	USERROLE_VIP1          = "V1"
	USERROLE_VIP2          = "V2"
	USERROLE_VIP3          = "V3"
	USERROLE_NOR_DISCOUNT  = 1
	USERROLE_VIP1_DISCOUNT = 0.95
	USERROLE_VIP2_DISCOUNT = 0.9
	USERROLE_VIP3_DISCOUNT = 0.85

	POINT_EXCHANGE_RATE = 1
)

func RoleLegal(role string) bool {
	roleList := []string{USERROLE_NORMAL, USERROLE_VIP1, USERROLE_VIP2, USERROLE_VIP3}
	for _, role_str := range roleList {
		if role == role_str {
			return true
		}
	}
	return false
}
