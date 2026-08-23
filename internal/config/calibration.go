package config

import (
	"github.com/lacsar712/filmpull/internal/model"
)

// Calibration holds per-zone tension sensor scaling factors loaded from plant engineering tables.
type Calibration struct {
	ZoneFactors map[string]float64
	NipMaxKPa   map[string]float64
}

func DefaultCalibration(zoneCount, nipCount int) Calibration {
	zf := make(map[string]float64, zoneCount)
	for i := 1; i <= zoneCount; i++ {
		zf[model.ZoneID("zone-"+itoa(i)).String()] = 1.0 + float64(i%5)*0.02
	}
	nm := make(map[string]float64, nipCount)
	for i := 1; i <= nipCount; i++ {
		nm["nip-"+itoa(i)] = 450 + float64(i%3)*25
	}
	return Calibration{ZoneFactors: zf, NipMaxKPa: nm}
}

func (c Calibration) ScaleZone(zone model.ZoneID, raw float64) float64 {
	f, ok := c.ZoneFactors[zone.String()]
	if !ok || f <= 0 {
		return raw
	}
	return raw * f
}

func (c Calibration) MaxNipPressure(id model.NipID) float64 {
	if v, ok := c.NipMaxKPa[string(id)]; ok {
		return v
	}
	return 500
}

func itoa(n int) string {
	if n == 0 {
		return "0"
	}
	var buf [12]byte
	i := len(buf)
	for n > 0 {
		i--
		buf[i] = byte('0' + n%10)
		n /= 10
	}
	return string(buf[i:])
}
