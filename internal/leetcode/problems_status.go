package leetcode

var problemStatus = map[int]bool{
	1:    true,
	2:    true,
	3:    true,
	4:    true,
	5:    true,
	6:    true,
	7:    true,
	8:    true,
	9:    true,
	10:   true,
	11:   true,
	12:   true,
	13:   true,
	14:   true,
	15:   true,
	16:   true,
	17:   true,
	18:   true,
	20:   true,
	21:   true,
	22:   true,
	24:   true,
	26:   true,
	27:   true,
	28:   true,
	35:   true,
	36:   true,
	38:   true,
	48:   true,
	50:   true,
	53:   true,
	58:   true,
	65:   true,
	66:   true,
	67:   true,
	69:   true,
	70:   true,
	71:   true,
	83:   true,
	88:   true,
	100:  true,
	101:  true,
	104:  true,
	108:  true,
	110:  true,
	112:  true,
	118:  true,
	119:  true,
	121:  true,
	122:  true,
	125:  true,
	136:  true,
	137:  true,
	139:  true,
	141:  true,
	151:  true,
	155:  true,
	160:  true,
	165:  true,
	167:  true,
	168:  true,
	169:  true,
	171:  true,
	172:  true,
	182:  true,
	189:  true,
	190:  true,
	191:  true,
	192:  true,
	193:  true,
	194:  true,
	195:  true,
	198:  true,
	202:  true,
	204:  true,
	206:  true,
	215:  true,
	217:  true,
	219:  true,
	226:  true,
	231:  true,
	233:  true,
	234:  true,
	237:  true,
	242:  true,
	258:  true,
	263:  true,
	264:  true,
	268:  true,
	283:  true,
	290:  true,
	292:  true,
	303:  true,
	319:  true,
	326:  true,
	342:  true,
	343:  true,
	344:  true,
	345:  true,
	350:  true,
	367:  true,
	371:  true,
	383:  true,
	387:  true,
	393:  true,
	412:  true,
	413:  true,
	414:  true,
	415:  true,
	434:  true,
	441:  true,
	443:  true,
	445:  true,
	448:  true,
	453:  true,
	455:  true,
	459:  true,
	461:  true,
	468:  true,
	485:  true,
	504:  true,
	509:  true,
	520:  true,
	521:  true,
	532:  true,
	541:  true,
	551:  true,
	557:  true,
	561:  true,
	565:  true,
	566:  true,
	581:  true,
	595:  true,
	605:  true,
	606:  true,
	617:  true,
	620:  true,
	626:  true,
	627:  true,
	628:  true,
	643:  true,
	657:  true,
	661:  true,
	665:  true,
	674:  true,
	679:  true,
	680:  true,
	686:  true,
	696:  true,
	697:  true,
	705:  true,
	706:  true,
	709:  true,
	717:  true,
	724:  true,
	728:  true,
	744:  true,
	746:  true,
	747:  true,
	766:  true,
	771:  true,
	788:  true,
	793:  true,
	796:  true,
	804:  true,
	819:  true,
	824:  true,
	830:  true,
	832:  true,
	840:  true,
	849:  true,
	859:  true,
	867:  true,
	877:  true,
	888:  true,
	893:  true,
	896:  true,
	905:  true,
	908:  true,
	912:  true,
	914:  true,
	917:  true,
	922:  true,
	925:  true,
	929:  true,
	933:  true,
	937:  true,
	938:  true,
	941:  true,
	942:  true,
	944:  true,
	949:  true,
	953:  true,
	961:  true,
	965:  true,
	970:  true,
	973:  true,
	976:  true,
	977:  true,
	984:  true,
	985:  true,
	989:  true,
	993:  true,
	994:  true,
	997:  true,
	999:  true,
	1000: true,
	1002: true,
	1005: true,
	1009: true,
	1010: true,
	1013: true,
	1018: true,
	1021: true,
	1022: true,
	1023: true,
	1025: true,
	1029: true,
	1037: true,
	1041: true,
	1042: true,
	1046: true,
	1047: true,
	1051: true,
	1071: true,
	1108: true,
	1154: true,
	1163: true,
	1170: true,
	1185: true,
	1189: true,
	1221: true,
}

// IsSolved - leetcode.IsSolved
func IsSolved(id int) bool {
	return problemStatus[id]
}
