package engine

import (
	"github.com/mumax/3/cuda"
	"github.com/mumax/3/data"
)

var globalmesh_ data.Mesh // mesh for m and everything that has the same size
var blacklist = []int{ 131,  137,  139,  149,  151,  157,  163,  167,  173,  179,  181,
	191,  193,  197,  199,  211,  223,  227,  229,  233,  239,  241,
	251,  257,  262,  263,  269,  271,  274,  277,  278,  281,  283,
	293,  298,  302,  307,  311,  313,  314,  317,  326,  331,  334,
	337,  346,  347,  349,  353,  358,  359,  362,  367,  373,  379,
	382,  383,  386,  389,  393,  394,  397,  398,  401,  409,  411,
	417,  419,  421,  422,  431,  433,  439,  443,  446,  447,  449,
	453,  454,  457,  458,  461,  463,  466,  467,  471,  478,  479,
	482,  487,  489,  491,  499,  501,  502,  503,  509,  514,  519,
	521,  523,  524,  526,  537,  538,  541,  542,  543,  547,  548,
	554,  556,  557,  562,  563,  566,  569,  571,  573,  577,  579,
	586,  587,  591,  593,  596,  597,  599,  601,  604,  607,  613,
	614,  617,  619,  622,  626,  628,  631,  633,  634,  641,  643,
	647,  652,  653,  655,  659,  661,  662,  668,  669,  673,  674,
	677,  681,  683,  685,  687,  691,  692,  694,  695,  698,  699,
	701,  706,  709,  716,  717,  718,  719,  723,  724,  727,  733,
	734,  739,  743,  745,  746,  751,  753,  755,  757,  758,  761,
	764,  766,  769,  771,  772,  773,  778,  785,  786,  787,  788,
	789,  794,  796,  797,  802,  807,  809,  811,  813,  815,  818,
	821,  822,  823,  827,  829,  831,  834,  835,  838,  839,  842,
	843,  844,  849,  853,  857,  859,  862,  863,  865,  866,  877,
	878,  879,  881,  883,  886,  887,  892,  894,  895,  898,  905,
	906,  907,  908,  911,  914,  916,  917,  919,  921,  922,  926,
	929,  932,  933,  934,  937,  939,  941,  942,  947,  951,  953,
	955,  956,  958,  959,  964,  965,  967,  971,  973,  974,  977,
	978,  982,  983,  985,  991,  993,  995,  997,  998, 1002, 1004,
   1006, 1009, 1011, 1013, 1018, 1019, 1021, 1028, 1031, 1033, 1038,
   1039, 1041, 1042, 1043, 1046, 1047, 1048, 1049, 1051, 1052, 1055,
   1057, 1059, 1061, 1063, 1069, 1074, 1076, 1077, 1082, 1084, 1086,
   1087, 1091, 1093, 1094, 1096, 1097, 1099, 1101, 1103, 1108, 1109,
   1112, 1114, 1115, 1117, 1119, 1123, 1124, 1126, 1129, 1132, 1135,
   1137, 1138, 1141, 1142, 1145, 1146, 1149, 1151, 1153, 1154, 1158,
   1163, 1165, 1167, 1169, 1171, 1172, 1174, 1179, 1181, 1182, 1186,
   1187, 1191, 1192, 1193, 1194, 1195, 1198, 1201, 1202, 1203, 1205,
   1208, 1211, 1213, 1214, 1217, 1223, 1226, 1227, 1228, 1229, 1231,
   1233, 1234, 1237, 1238, 1244, 1249, 1251, 1252, 1253, 1255, 1256,
   1257, 1259, 1262, 1263, 1266, 1267, 1268, 1277, 1279, 1282, 1283,
   1285, 1286, 1289, 1291, 1293, 1294, 1297, 1299, 1301, 1303, 1304,
   1306, 1307, 1310, 1315, 1317, 1318, 1319, 1321, 1322, 1324, 1327,
   1329, 1336, 1337, 1338, 1341, 1345, 1346, 1347, 1348, 1351, 1354,
   1355, 1359, 1361, 1362, 1366, 1367, 1370, 1371, 1373, 1374, 1379,
   1381, 1382, 1383, 1384, 1385, 1388, 1389, 1390, 1393, 1396, 1398,
   1399, 1401, 1402, 1405, 1409, 1412, 1413, 1415, 1418, 1423, 1427,
   1429, 1432, 1433, 1434, 1436, 1437, 1438, 1439, 1441, 1446, 1447,
   1448, 1451, 1453, 1454, 1459, 1461, 1465, 1466, 1467, 1468, 1471,
   1473, 1477, 1478, 1481, 1483, 1486, 1487, 1489, 1490, 1492, 1493,
   1497, 1499, 1502, 1503, 1506, 1507, 1509, 1510, 1511, 1514, 1516,
   1522, 1523, 1527, 1528, 1529, 1531, 1532, 1535, 1538, 1542, 1543,
   1544, 1546, 1549, 1553, 1555, 1556, 1557, 1559, 1561, 1563, 1565,
   1567, 1569, 1570, 1571, 1572, 1574, 1576, 1578, 1579, 1583, 1585,
   1588, 1589, 1592, 1594, 1597, 1601, 1603, 1604, 1607, 1609, 1611,
   1613, 1614, 1618, 1619, 1621, 1622, 1623, 1626, 1627, 1629, 1630,
   1631, 1636, 1637, 1639, 1641, 1642, 1644, 1646, 1654, 1655, 1657,
   1658, 1661, 1662, 1663, 1667, 1668, 1669, 1670, 1671, 1673, 1676,
   1678, 1684, 1685, 1686, 1687, 1688, 1689, 1693, 1697, 1698, 1699,
   1703, 1706, 1707, 1709, 1713, 1714, 1718, 1719, 1721, 1723, 1724,
   1726, 1727, 1730, 1731, 1732, 1733, 1735, 1737, 1741, 1745, 1747,
   1753, 1754, 1756, 1757, 1758, 1759, 1761, 1762, 1765, 1766, 1772,
   1773, 1774, 1777, 1779, 1781, 1783, 1784, 1787, 1788, 1789, 1790,
   1791, 1793, 1795, 1796, 1797, 1799, 1801, 1803, 1807, 1810, 1811,
   1812, 1814, 1816, 1821, 1822, 1823, 1828, 1831, 1832, 1834, 1835,
   1837, 1838, 1839, 1841, 1842, 1844, 1847, 1851, 1852, 1857, 1858,
   1861, 1864, 1865, 1866, 1867, 1868, 1871, 1873, 1874, 1877, 1878,
   1879, 1882, 1883, 1884, 1889, 1893, 1894, 1895, 1897, 1899, 1901,
   1902, 1903, 1906, 1907, 1910, 1912, 1913, 1915, 1916, 1918, 1923,
   1928, 1929, 1930, 1931, 1933, 1934, 1937, 1939, 1941, 1942, 1945,
   1946, 1948, 1949, 1951, 1954, 1956, 1959, 1963, 1964, 1965, 1966,
   1967, 1969, 1970, 1973, 1977, 1979, 1981, 1982, 1983, 1985, 1986,
   1987, 1990, 1991, 1993, 1994, 1996, 1997, 1999}

func init() {
	DeclFunc("SetGridSize", SetGridSize, `Sets the number of cells for X,Y,Z`)
	DeclFunc("SetCellSize", SetCellSize, `Sets the X,Y,Z cell size in meters`)
	DeclFunc("SetMesh", SetMesh, `Sets GridSize, CellSize and PBC at the same time`)
	DeclFunc("SetPBC", SetPBC, "Sets the number of repetitions in X,Y,Z to create periodic boundary "+
		"conditions. The number of repetitions determines the cutoff range for the demagnetization.")
}

func Mesh() *data.Mesh {
	checkMesh()
	return &globalmesh_
}

func arg(msg string, test bool) {
	if !test {
		panic(UserErr(msg + ": illegal arugment"))
	}
}

// Set the simulation mesh to Nx x Ny x Nz cells of given size.
// Can be set only once at the beginning of the simulation.
// TODO: dedup arguments from globals
func SetMesh(Nx, Ny, Nz int, cellSizeX, cellSizeY, cellSizeZ float64, pbcx, pbcy, pbcz int) {
	SetBusy(true)
	defer SetBusy(false)

	arg("GridSize", Nx > 0 && Ny > 0 && Nz > 0)
	arg("CellSize", cellSizeX > 0 && cellSizeY > 0 && cellSizeZ > 0)
	arg("PBC", pbcx >= 0 && pbcy >= 0 && pbcz >= 0)

	prevSize := globalmesh_.Size()
	pbc := []int{pbcx, pbcy, pbcz}

	if globalmesh_.Size() == [3]int{0, 0, 0} {
		// first time mesh is set
		for _,b := range blacklist {
			if Nx == b {
				Nx += 1
			}  
			if Ny == b {
				Ny += 1
			}  
			if Nz == b {
				Nz += 1
			}  
		}
		globalmesh_ = *data.NewMesh(Nx, Ny, Nz, cellSizeX, cellSizeY, cellSizeZ, pbc...)
		M.alloc()
		regions.alloc()
	} else {
		// here be dragons
		LogOut("resizing...")

		// free everything
		conv_.Free()
		conv_ = nil
		mfmconv_.Free()
		mfmconv_ = nil
		cuda.FreeBuffers()

		// resize everything
		globalmesh_ = *data.NewMesh(Nx, Ny, Nz, cellSizeX, cellSizeY, cellSizeZ, pbc...)
		M.resize()
		regions.resize()
		geometry.buffer.Free()
		geometry.buffer = data.NilSlice(1, Mesh().Size())
		geometry.setGeom(geometry.shape)

		// remove excitation extra terms if they don't fit anymore
		// up to the user to add them again
		if Mesh().Size() != prevSize {
			B_ext.RemoveExtraTerms()
			J.RemoveExtraTerms()
		}

		if Mesh().Size() != prevSize {
			B_therm.noise.Free()
			B_therm.noise = nil
		}
	}
	lazy_gridsize = []int{Nx, Ny, Nz}
	lazy_cellsize = []float64{cellSizeX, cellSizeY, cellSizeZ}
	lazy_pbc = []int{pbcx, pbcy, pbcz}
}

func printf(f float64) float32 {
	return float32(f)
}

// for lazy setmesh: set gridsize and cellsize in separate calls
var (
	lazy_gridsize []int
	lazy_cellsize []float64
	lazy_pbc      = []int{0, 0, 0}
)

func SetGridSize(Nx, Ny, Nz int) {
	lazy_gridsize = []int{Nx, Ny, Nz}
	if lazy_cellsize != nil {
		SetMesh(Nx, Ny, Nz, lazy_cellsize[X], lazy_cellsize[Y], lazy_cellsize[Z], lazy_pbc[X], lazy_pbc[Y], lazy_pbc[Z])
	}
}

func SetCellSize(cx, cy, cz float64) {
	lazy_cellsize = []float64{cx, cy, cz}
	if lazy_gridsize != nil {
		SetMesh(lazy_gridsize[X], lazy_gridsize[Y], lazy_gridsize[Z], cx, cy, cz, lazy_pbc[X], lazy_pbc[Y], lazy_pbc[Z])
	}
}

func SetPBC(nx, ny, nz int) {
	lazy_pbc = []int{nx, ny, nz}
	if lazy_gridsize != nil && lazy_cellsize != nil {
		SetMesh(lazy_gridsize[X], lazy_gridsize[Y], lazy_gridsize[Z],
			lazy_cellsize[X], lazy_cellsize[Y], lazy_cellsize[Z],
			lazy_pbc[X], lazy_pbc[Y], lazy_pbc[Z])
	}
}

// check if mesh is set
func checkMesh() {
	if globalmesh_.Size() == [3]int{0, 0, 0} {
		panic("need to set mesh first")
	}
}