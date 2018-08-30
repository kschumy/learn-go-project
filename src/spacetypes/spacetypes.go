package spacetypes

type SpaceType int

const (
	SpecialWithoutPass SpaceType = iota
	RegularWithoutPass
	SpecialWithPass
	RegularyWithoutPass
)


func GetType(needSpecial, needPass bool) SpaceType {
	switch {
	case needSpecial && !needPass:
		return SpecialWithoutPass
	case !needSpecial && !needPass:
		return RegularWithoutPass
	case needSpecial && needPass:
		return SpecialWithPass
	default:
		return RegularyWithoutPass
	}
}

//type spaceTypes struct {
//	special bool
//	pass bool
//}
//
//func IsSpecial(currType *SpaceType) bool {
//	return currType == 1 || code == 3
//}
//
//func IsPass(code int) bool {
//	return code == 1 || code == 2
//}

