package model

const (
	USERROLE_NORMAL = "N"
	USERROLE_VIP1   = "V1"
	USERROLE_VIP2   = "V2"
	USERROLE_VIP3   = "V3"
)

var USERROLE_DISCOUNT map[string]float64

func init() {
	USERROLE_DISCOUNT = make(map[string]float64)
	USERROLE_DISCOUNT[USERROLE_NORMAL] = 1
	USERROLE_DISCOUNT[USERROLE_VIP1] = 0.95
	USERROLE_DISCOUNT[USERROLE_VIP2] = 0.9
	USERROLE_DISCOUNT[USERROLE_VIP3] = 0.85
}
