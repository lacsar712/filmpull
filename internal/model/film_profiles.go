package model

func init() { registerBuiltinProfiles() }

var builtinProfiles []FilmGrade

func registerBuiltinProfiles() {
    builtinProfiles = append(builtinProfiles, FilmGrade{ID: "FG-001", ThicknessUM: 13, WebWidthMM: 1207, DrawRatio: 3.15, NominalTensionN: 81, MaxLineSpeedMPM: 181, AnnealTempC: 141, PreheatTempC: 96})
    builtinProfiles = append(builtinProfiles, FilmGrade{ID: "FG-002", ThicknessUM: 14, WebWidthMM: 1214, DrawRatio: 3.3, NominalTensionN: 82, MaxLineSpeedMPM: 182, AnnealTempC: 142, PreheatTempC: 97})
    builtinProfiles = append(builtinProfiles, FilmGrade{ID: "FG-003", ThicknessUM: 15, WebWidthMM: 1221, DrawRatio: 3.45, NominalTensionN: 83, MaxLineSpeedMPM: 183, AnnealTempC: 143, PreheatTempC: 98})
    builtinProfiles = append(builtinProfiles, FilmGrade{ID: "FG-004", ThicknessUM: 16, WebWidthMM: 1228, DrawRatio: 3.6, NominalTensionN: 84, MaxLineSpeedMPM: 184, AnnealTempC: 144, PreheatTempC: 99})
    builtinProfiles = append(builtinProfiles, FilmGrade{ID: "FG-005", ThicknessUM: 17, WebWidthMM: 1235, DrawRatio: 3.75, NominalTensionN: 85, MaxLineSpeedMPM: 185, AnnealTempC: 145, PreheatTempC: 100})
    builtinProfiles = append(builtinProfiles, FilmGrade{ID: "FG-006", ThicknessUM: 18, WebWidthMM: 1242, DrawRatio: 3.9, NominalTensionN: 86, MaxLineSpeedMPM: 186, AnnealTempC: 146, PreheatTempC: 101})
    builtinProfiles = append(builtinProfiles, FilmGrade{ID: "FG-007", ThicknessUM: 19, WebWidthMM: 1249, DrawRatio: 4.05, NominalTensionN: 87, MaxLineSpeedMPM: 187, AnnealTempC: 147, PreheatTempC: 102})
    builtinProfiles = append(builtinProfiles, FilmGrade{ID: "FG-008", ThicknessUM: 20, WebWidthMM: 1256, DrawRatio: 4.2, NominalTensionN: 88, MaxLineSpeedMPM: 188, AnnealTempC: 148, PreheatTempC: 103})
    builtinProfiles = append(builtinProfiles, FilmGrade{ID: "FG-009", ThicknessUM: 21, WebWidthMM: 1263, DrawRatio: 4.35, NominalTensionN: 89, MaxLineSpeedMPM: 189, AnnealTempC: 149, PreheatTempC: 104})
    builtinProfiles = append(builtinProfiles, FilmGrade{ID: "FG-010", ThicknessUM: 22, WebWidthMM: 1270, DrawRatio: 4.5, NominalTensionN: 90, MaxLineSpeedMPM: 190, AnnealTempC: 150, PreheatTempC: 105})
    builtinProfiles = append(builtinProfiles, FilmGrade{ID: "FG-011", ThicknessUM: 23, WebWidthMM: 1277, DrawRatio: 4.65, NominalTensionN: 91, MaxLineSpeedMPM: 191, AnnealTempC: 151, PreheatTempC: 106})
    builtinProfiles = append(builtinProfiles, FilmGrade{ID: "FG-012", ThicknessUM: 24, WebWidthMM: 1284, DrawRatio: 3, NominalTensionN: 92, MaxLineSpeedMPM: 192, AnnealTempC: 152, PreheatTempC: 107})
    builtinProfiles = append(builtinProfiles, FilmGrade{ID: "FG-013", ThicknessUM: 25, WebWidthMM: 1291, DrawRatio: 3.15, NominalTensionN: 93, MaxLineSpeedMPM: 193, AnnealTempC: 153, PreheatTempC: 108})
    builtinProfiles = append(builtinProfiles, FilmGrade{ID: "FG-014", ThicknessUM: 26, WebWidthMM: 1298, DrawRatio: 3.3, NominalTensionN: 94, MaxLineSpeedMPM: 194, AnnealTempC: 154, PreheatTempC: 109})
    builtinProfiles = append(builtinProfiles, FilmGrade{ID: "FG-015", ThicknessUM: 27, WebWidthMM: 1305, DrawRatio: 3.45, NominalTensionN: 95, MaxLineSpeedMPM: 195, AnnealTempC: 155, PreheatTempC: 110})
    builtinProfiles = append(builtinProfiles, FilmGrade{ID: "FG-016", ThicknessUM: 28, WebWidthMM: 1312, DrawRatio: 3.6, NominalTensionN: 96, MaxLineSpeedMPM: 196, AnnealTempC: 156, PreheatTempC: 111})
    builtinProfiles = append(builtinProfiles, FilmGrade{ID: "FG-017", ThicknessUM: 29, WebWidthMM: 1319, DrawRatio: 3.75, NominalTensionN: 97, MaxLineSpeedMPM: 197, AnnealTempC: 157, PreheatTempC: 112})
    builtinProfiles = append(builtinProfiles, FilmGrade{ID: "FG-018", ThicknessUM: 12, WebWidthMM: 1326, DrawRatio: 3.9, NominalTensionN: 98, MaxLineSpeedMPM: 198, AnnealTempC: 158, PreheatTempC: 113})
    builtinProfiles = append(builtinProfiles, FilmGrade{ID: "FG-019", ThicknessUM: 13, WebWidthMM: 1333, DrawRatio: 4.05, NominalTensionN: 99, MaxLineSpeedMPM: 199, AnnealTempC: 159, PreheatTempC: 114})
    builtinProfiles = append(builtinProfiles, FilmGrade{ID: "FG-020", ThicknessUM: 14, WebWidthMM: 1340, DrawRatio: 4.2, NominalTensionN: 100, MaxLineSpeedMPM: 200, AnnealTempC: 160, PreheatTempC: 95})
    builtinProfiles = append(builtinProfiles, FilmGrade{ID: "FG-021", ThicknessUM: 15, WebWidthMM: 1347, DrawRatio: 4.35, NominalTensionN: 101, MaxLineSpeedMPM: 201, AnnealTempC: 161, PreheatTempC: 96})
    builtinProfiles = append(builtinProfiles, FilmGrade{ID: "FG-022", ThicknessUM: 16, WebWidthMM: 1354, DrawRatio: 4.5, NominalTensionN: 102, MaxLineSpeedMPM: 202, AnnealTempC: 162, PreheatTempC: 97})
    builtinProfiles = append(builtinProfiles, FilmGrade{ID: "FG-023", ThicknessUM: 17, WebWidthMM: 1361, DrawRatio: 4.65, NominalTensionN: 103, MaxLineSpeedMPM: 203, AnnealTempC: 163, PreheatTempC: 98})
    builtinProfiles = append(builtinProfiles, FilmGrade{ID: "FG-024", ThicknessUM: 18, WebWidthMM: 1368, DrawRatio: 3, NominalTensionN: 104, MaxLineSpeedMPM: 204, AnnealTempC: 164, PreheatTempC: 99})
    builtinProfiles = append(builtinProfiles, FilmGrade{ID: "FG-025", ThicknessUM: 19, WebWidthMM: 1375, DrawRatio: 3.15, NominalTensionN: 105, MaxLineSpeedMPM: 205, AnnealTempC: 165, PreheatTempC: 100})
    builtinProfiles = append(builtinProfiles, FilmGrade{ID: "FG-026", ThicknessUM: 20, WebWidthMM: 1382, DrawRatio: 3.3, NominalTensionN: 106, MaxLineSpeedMPM: 206, AnnealTempC: 166, PreheatTempC: 101})
    builtinProfiles = append(builtinProfiles, FilmGrade{ID: "FG-027", ThicknessUM: 21, WebWidthMM: 1389, DrawRatio: 3.45, NominalTensionN: 107, MaxLineSpeedMPM: 207, AnnealTempC: 167, PreheatTempC: 102})
    builtinProfiles = append(builtinProfiles, FilmGrade{ID: "FG-028", ThicknessUM: 22, WebWidthMM: 1396, DrawRatio: 3.6, NominalTensionN: 108, MaxLineSpeedMPM: 208, AnnealTempC: 168, PreheatTempC: 103})
    builtinProfiles = append(builtinProfiles, FilmGrade{ID: "FG-029", ThicknessUM: 23, WebWidthMM: 1403, DrawRatio: 3.75, NominalTensionN: 109, MaxLineSpeedMPM: 209, AnnealTempC: 169, PreheatTempC: 104})
    builtinProfiles = append(builtinProfiles, FilmGrade{ID: "FG-030", ThicknessUM: 24, WebWidthMM: 1410, DrawRatio: 3.9, NominalTensionN: 110, MaxLineSpeedMPM: 210, AnnealTempC: 140, PreheatTempC: 105})
    builtinProfiles = append(builtinProfiles, FilmGrade{ID: "FG-031", ThicknessUM: 25, WebWidthMM: 1417, DrawRatio: 4.05, NominalTensionN: 111, MaxLineSpeedMPM: 211, AnnealTempC: 141, PreheatTempC: 106})
    builtinProfiles = append(builtinProfiles, FilmGrade{ID: "FG-032", ThicknessUM: 26, WebWidthMM: 1424, DrawRatio: 4.2, NominalTensionN: 112, MaxLineSpeedMPM: 212, AnnealTempC: 142, PreheatTempC: 107})
    builtinProfiles = append(builtinProfiles, FilmGrade{ID: "FG-033", ThicknessUM: 27, WebWidthMM: 1431, DrawRatio: 4.35, NominalTensionN: 113, MaxLineSpeedMPM: 213, AnnealTempC: 143, PreheatTempC: 108})
    builtinProfiles = append(builtinProfiles, FilmGrade{ID: "FG-034", ThicknessUM: 28, WebWidthMM: 1438, DrawRatio: 4.5, NominalTensionN: 114, MaxLineSpeedMPM: 214, AnnealTempC: 144, PreheatTempC: 109})
    builtinProfiles = append(builtinProfiles, FilmGrade{ID: "FG-035", ThicknessUM: 29, WebWidthMM: 1445, DrawRatio: 4.65, NominalTensionN: 115, MaxLineSpeedMPM: 215, AnnealTempC: 145, PreheatTempC: 110})
    builtinProfiles = append(builtinProfiles, FilmGrade{ID: "FG-036", ThicknessUM: 12, WebWidthMM: 1452, DrawRatio: 3, NominalTensionN: 116, MaxLineSpeedMPM: 216, AnnealTempC: 146, PreheatTempC: 111})
    builtinProfiles = append(builtinProfiles, FilmGrade{ID: "FG-037", ThicknessUM: 13, WebWidthMM: 1459, DrawRatio: 3.15, NominalTensionN: 117, MaxLineSpeedMPM: 217, AnnealTempC: 147, PreheatTempC: 112})
    builtinProfiles = append(builtinProfiles, FilmGrade{ID: "FG-038", ThicknessUM: 14, WebWidthMM: 1466, DrawRatio: 3.3, NominalTensionN: 118, MaxLineSpeedMPM: 218, AnnealTempC: 148, PreheatTempC: 113})
    builtinProfiles = append(builtinProfiles, FilmGrade{ID: "FG-039", ThicknessUM: 15, WebWidthMM: 1473, DrawRatio: 3.45, NominalTensionN: 119, MaxLineSpeedMPM: 219, AnnealTempC: 149, PreheatTempC: 114})
    builtinProfiles = append(builtinProfiles, FilmGrade{ID: "FG-040", ThicknessUM: 16, WebWidthMM: 1480, DrawRatio: 3.6, NominalTensionN: 80, MaxLineSpeedMPM: 220, AnnealTempC: 150, PreheatTempC: 95})
    builtinProfiles = append(builtinProfiles, FilmGrade{ID: "FG-041", ThicknessUM: 17, WebWidthMM: 1487, DrawRatio: 3.75, NominalTensionN: 81, MaxLineSpeedMPM: 221, AnnealTempC: 151, PreheatTempC: 96})
    builtinProfiles = append(builtinProfiles, FilmGrade{ID: "FG-042", ThicknessUM: 18, WebWidthMM: 1494, DrawRatio: 3.9, NominalTensionN: 82, MaxLineSpeedMPM: 222, AnnealTempC: 152, PreheatTempC: 97})
    builtinProfiles = append(builtinProfiles, FilmGrade{ID: "FG-043", ThicknessUM: 19, WebWidthMM: 1501, DrawRatio: 4.05, NominalTensionN: 83, MaxLineSpeedMPM: 223, AnnealTempC: 153, PreheatTempC: 98})
    builtinProfiles = append(builtinProfiles, FilmGrade{ID: "FG-044", ThicknessUM: 20, WebWidthMM: 1508, DrawRatio: 4.2, NominalTensionN: 84, MaxLineSpeedMPM: 224, AnnealTempC: 154, PreheatTempC: 99})
    builtinProfiles = append(builtinProfiles, FilmGrade{ID: "FG-045", ThicknessUM: 21, WebWidthMM: 1515, DrawRatio: 4.35, NominalTensionN: 85, MaxLineSpeedMPM: 225, AnnealTempC: 155, PreheatTempC: 100})
    builtinProfiles = append(builtinProfiles, FilmGrade{ID: "FG-046", ThicknessUM: 22, WebWidthMM: 1522, DrawRatio: 4.5, NominalTensionN: 86, MaxLineSpeedMPM: 226, AnnealTempC: 156, PreheatTempC: 101})
    builtinProfiles = append(builtinProfiles, FilmGrade{ID: "FG-047", ThicknessUM: 23, WebWidthMM: 1529, DrawRatio: 4.65, NominalTensionN: 87, MaxLineSpeedMPM: 227, AnnealTempC: 157, PreheatTempC: 102})
    builtinProfiles = append(builtinProfiles, FilmGrade{ID: "FG-048", ThicknessUM: 24, WebWidthMM: 1536, DrawRatio: 3, NominalTensionN: 88, MaxLineSpeedMPM: 228, AnnealTempC: 158, PreheatTempC: 103})
    builtinProfiles = append(builtinProfiles, FilmGrade{ID: "FG-049", ThicknessUM: 25, WebWidthMM: 1543, DrawRatio: 3.15, NominalTensionN: 89, MaxLineSpeedMPM: 229, AnnealTempC: 159, PreheatTempC: 104})
    builtinProfiles = append(builtinProfiles, FilmGrade{ID: "FG-050", ThicknessUM: 26, WebWidthMM: 1550, DrawRatio: 3.3, NominalTensionN: 90, MaxLineSpeedMPM: 230, AnnealTempC: 160, PreheatTempC: 105})
    builtinProfiles = append(builtinProfiles, FilmGrade{ID: "FG-051", ThicknessUM: 27, WebWidthMM: 1557, DrawRatio: 3.45, NominalTensionN: 91, MaxLineSpeedMPM: 231, AnnealTempC: 161, PreheatTempC: 106})
    builtinProfiles = append(builtinProfiles, FilmGrade{ID: "FG-052", ThicknessUM: 28, WebWidthMM: 1564, DrawRatio: 3.6, NominalTensionN: 92, MaxLineSpeedMPM: 232, AnnealTempC: 162, PreheatTempC: 107})
    builtinProfiles = append(builtinProfiles, FilmGrade{ID: "FG-053", ThicknessUM: 29, WebWidthMM: 1571, DrawRatio: 3.75, NominalTensionN: 93, MaxLineSpeedMPM: 233, AnnealTempC: 163, PreheatTempC: 108})
    builtinProfiles = append(builtinProfiles, FilmGrade{ID: "FG-054", ThicknessUM: 12, WebWidthMM: 1578, DrawRatio: 3.9, NominalTensionN: 94, MaxLineSpeedMPM: 234, AnnealTempC: 164, PreheatTempC: 109})
    builtinProfiles = append(builtinProfiles, FilmGrade{ID: "FG-055", ThicknessUM: 13, WebWidthMM: 1585, DrawRatio: 4.05, NominalTensionN: 95, MaxLineSpeedMPM: 235, AnnealTempC: 165, PreheatTempC: 110})
    builtinProfiles = append(builtinProfiles, FilmGrade{ID: "FG-056", ThicknessUM: 14, WebWidthMM: 1592, DrawRatio: 4.2, NominalTensionN: 96, MaxLineSpeedMPM: 236, AnnealTempC: 166, PreheatTempC: 111})
    builtinProfiles = append(builtinProfiles, FilmGrade{ID: "FG-057", ThicknessUM: 15, WebWidthMM: 1599, DrawRatio: 4.35, NominalTensionN: 97, MaxLineSpeedMPM: 237, AnnealTempC: 167, PreheatTempC: 112})
    builtinProfiles = append(builtinProfiles, FilmGrade{ID: "FG-058", ThicknessUM: 16, WebWidthMM: 1606, DrawRatio: 4.5, NominalTensionN: 98, MaxLineSpeedMPM: 238, AnnealTempC: 168, PreheatTempC: 113})
    builtinProfiles = append(builtinProfiles, FilmGrade{ID: "FG-059", ThicknessUM: 17, WebWidthMM: 1613, DrawRatio: 4.65, NominalTensionN: 99, MaxLineSpeedMPM: 239, AnnealTempC: 169, PreheatTempC: 114})
    builtinProfiles = append(builtinProfiles, FilmGrade{ID: "FG-060", ThicknessUM: 18, WebWidthMM: 1620, DrawRatio: 3, NominalTensionN: 100, MaxLineSpeedMPM: 180, AnnealTempC: 140, PreheatTempC: 95})
    builtinProfiles = append(builtinProfiles, FilmGrade{ID: "FG-061", ThicknessUM: 19, WebWidthMM: 1627, DrawRatio: 3.15, NominalTensionN: 101, MaxLineSpeedMPM: 181, AnnealTempC: 141, PreheatTempC: 96})
    builtinProfiles = append(builtinProfiles, FilmGrade{ID: "FG-062", ThicknessUM: 20, WebWidthMM: 1634, DrawRatio: 3.3, NominalTensionN: 102, MaxLineSpeedMPM: 182, AnnealTempC: 142, PreheatTempC: 97})
    builtinProfiles = append(builtinProfiles, FilmGrade{ID: "FG-063", ThicknessUM: 21, WebWidthMM: 1641, DrawRatio: 3.45, NominalTensionN: 103, MaxLineSpeedMPM: 183, AnnealTempC: 143, PreheatTempC: 98})
    builtinProfiles = append(builtinProfiles, FilmGrade{ID: "FG-064", ThicknessUM: 22, WebWidthMM: 1648, DrawRatio: 3.6, NominalTensionN: 104, MaxLineSpeedMPM: 184, AnnealTempC: 144, PreheatTempC: 99})
    builtinProfiles = append(builtinProfiles, FilmGrade{ID: "FG-065", ThicknessUM: 23, WebWidthMM: 1655, DrawRatio: 3.75, NominalTensionN: 105, MaxLineSpeedMPM: 185, AnnealTempC: 145, PreheatTempC: 100})
    builtinProfiles = append(builtinProfiles, FilmGrade{ID: "FG-066", ThicknessUM: 24, WebWidthMM: 1662, DrawRatio: 3.9, NominalTensionN: 106, MaxLineSpeedMPM: 186, AnnealTempC: 146, PreheatTempC: 101})
    builtinProfiles = append(builtinProfiles, FilmGrade{ID: "FG-067", ThicknessUM: 25, WebWidthMM: 1669, DrawRatio: 4.05, NominalTensionN: 107, MaxLineSpeedMPM: 187, AnnealTempC: 147, PreheatTempC: 102})
    builtinProfiles = append(builtinProfiles, FilmGrade{ID: "FG-068", ThicknessUM: 26, WebWidthMM: 1676, DrawRatio: 4.2, NominalTensionN: 108, MaxLineSpeedMPM: 188, AnnealTempC: 148, PreheatTempC: 103})
    builtinProfiles = append(builtinProfiles, FilmGrade{ID: "FG-069", ThicknessUM: 27, WebWidthMM: 1683, DrawRatio: 4.35, NominalTensionN: 109, MaxLineSpeedMPM: 189, AnnealTempC: 149, PreheatTempC: 104})
    builtinProfiles = append(builtinProfiles, FilmGrade{ID: "FG-070", ThicknessUM: 28, WebWidthMM: 1690, DrawRatio: 4.5, NominalTensionN: 110, MaxLineSpeedMPM: 190, AnnealTempC: 150, PreheatTempC: 105})
    builtinProfiles = append(builtinProfiles, FilmGrade{ID: "FG-071", ThicknessUM: 29, WebWidthMM: 1697, DrawRatio: 4.65, NominalTensionN: 111, MaxLineSpeedMPM: 191, AnnealTempC: 151, PreheatTempC: 106})
    builtinProfiles = append(builtinProfiles, FilmGrade{ID: "FG-072", ThicknessUM: 12, WebWidthMM: 1704, DrawRatio: 3, NominalTensionN: 112, MaxLineSpeedMPM: 192, AnnealTempC: 152, PreheatTempC: 107})
    builtinProfiles = append(builtinProfiles, FilmGrade{ID: "FG-073", ThicknessUM: 13, WebWidthMM: 1711, DrawRatio: 3.15, NominalTensionN: 113, MaxLineSpeedMPM: 193, AnnealTempC: 153, PreheatTempC: 108})
    builtinProfiles = append(builtinProfiles, FilmGrade{ID: "FG-074", ThicknessUM: 14, WebWidthMM: 1718, DrawRatio: 3.3, NominalTensionN: 114, MaxLineSpeedMPM: 194, AnnealTempC: 154, PreheatTempC: 109})
    builtinProfiles = append(builtinProfiles, FilmGrade{ID: "FG-075", ThicknessUM: 15, WebWidthMM: 1725, DrawRatio: 3.45, NominalTensionN: 115, MaxLineSpeedMPM: 195, AnnealTempC: 155, PreheatTempC: 110})
    builtinProfiles = append(builtinProfiles, FilmGrade{ID: "FG-076", ThicknessUM: 16, WebWidthMM: 1732, DrawRatio: 3.6, NominalTensionN: 116, MaxLineSpeedMPM: 196, AnnealTempC: 156, PreheatTempC: 111})
    builtinProfiles = append(builtinProfiles, FilmGrade{ID: "FG-077", ThicknessUM: 17, WebWidthMM: 1739, DrawRatio: 3.75, NominalTensionN: 117, MaxLineSpeedMPM: 197, AnnealTempC: 157, PreheatTempC: 112})
    builtinProfiles = append(builtinProfiles, FilmGrade{ID: "FG-078", ThicknessUM: 18, WebWidthMM: 1746, DrawRatio: 3.9, NominalTensionN: 118, MaxLineSpeedMPM: 198, AnnealTempC: 158, PreheatTempC: 113})
    builtinProfiles = append(builtinProfiles, FilmGrade{ID: "FG-079", ThicknessUM: 19, WebWidthMM: 1753, DrawRatio: 4.05, NominalTensionN: 119, MaxLineSpeedMPM: 199, AnnealTempC: 159, PreheatTempC: 114})
    builtinProfiles = append(builtinProfiles, FilmGrade{ID: "FG-080", ThicknessUM: 20, WebWidthMM: 1760, DrawRatio: 4.2, NominalTensionN: 80, MaxLineSpeedMPM: 200, AnnealTempC: 160, PreheatTempC: 95})
}

func BuiltinProfiles() []FilmGrade {
	out := make([]FilmGrade, len(builtinProfiles))
	copy(out, builtinProfiles)
	return out
}

func LookupGrade(id FilmGradeID) (FilmGrade, bool) {
	for _, g := range builtinProfiles {
		if g.ID == id {
			return g, true
		}
	}
	return FilmGrade{}, false
}

func DefaultGrade() FilmGrade {
	if len(builtinProfiles) == 0 {
		return FilmGrade{ID: "FG-001", ThicknessUM: 25, WebWidthMM: 1500, DrawRatio: 3.5, NominalTensionN: 120, MaxLineSpeedMPM: 220}
	}
	return builtinProfiles[0]
}