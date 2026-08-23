package nip

import "math"

func RampPressure(current, target, rate float64) float64 {
	if math.Abs(target-current) <= rate {
		return target
	}
	if target > current {
		return current + rate
	}
	return current - rate
}

func SeatPressure(nominal, seatFactor float64) float64 {
	if seatFactor <= 0 {
		seatFactor = 0.85
	}
	return nominal * seatFactor
}

func ReleasePressure(current, rate float64) float64 {
	if current <= 0 {
		return 0
	}
	return math.Max(0, current-rate)
}