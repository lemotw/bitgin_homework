package event

import (
	"BitginHomework/config"
	"BitginHomework/model"
	"log"
)

func pointDiscount(p int) int {
	return p * config.POINT_EXCHANGE_RATE
}
func discountPoint(d int) int {
	return d / config.POINT_EXCHANGE_RATE
}

// help logic Upset
func CostPoint(u model.User, record *model.TradeRecord, val float64) error {
	record.PointDiff = -record.PointDiff
	return nil
}
func CostBalance(u model.User, record *model.TradeRecord, val float64) error {
	record.BalanceDiff = -record.BalanceDiff
	return nil
}

// apply role discount
func RoleDiscount(u model.User, record *model.TradeRecord, val float64) error {
	switch u.Role {
	case config.USERROLE_VIP1:
		record.BalanceDiff = record.BalanceDiff * config.USERROLE_VIP1_DISCOUNT
	case config.USERROLE_VIP2:
		record.BalanceDiff = record.BalanceDiff * config.USERROLE_VIP2_DISCOUNT
	case config.USERROLE_VIP3:
		record.BalanceDiff = record.BalanceDiff * config.USERROLE_VIP3_DISCOUNT
	case config.USERROLE_NORMAL:
		record.BalanceDiff = record.BalanceDiff * config.USERROLE_NOR_DISCOUNT
	}

	return nil
}

// apply point discount
func PointDiscount(u model.User, record *model.TradeRecord, val float64) error {
	log.Println(record)
	availableDiscount := int(val) / 1000
	availableDiscount *= 100
	discount := availableDiscount

	pointUsed := record.PointDiff
	if record.PointDiff > record.Point {
		pointUsed = record.Point
	}

	if pointDiscount(pointUsed) > availableDiscount {
		pointUsed = discountPoint(availableDiscount)
	}

	discount = pointDiscount(pointUsed)
	record.BalanceDiff -= float64(discount)
	record.PointDiff = pointUsed

	return nil
}
