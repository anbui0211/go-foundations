package unittest

import (
	"fmt"
	"regexp"
	"strings"
)

func GetLocation(address, ward, district, province string) string {
	return fmt.Sprintf("%s, %s, %s, %s",
		replaceLocationAddress(address),
		replaceLocationWard(ward),
		replaceLocationDistrict(district),
		replaceLocationProvince(province),
	)
}

// ReplaceLocationProvince ...
func replaceLocationProvince(p string) string {
	return strings.ReplaceAll(TrimSpaceAll(p), "Thành phố ", "")
}

// ReplaceLocationDistrict ...
func replaceLocationDistrict(d string) string {
	d = TrimSpaceAll(d)
	d = strings.ReplaceAll(d, "Quận ", "Q.")
	d = strings.ReplaceAll(d, "Huyện ", "H.")
	d = strings.ReplaceAll(d, "Thị xã ", "TX.")
	d = strings.ReplaceAll(d, "Thành phố ", "TP.")
	return d
}

// ReplaceLocationWard ...
func replaceLocationWard(w string) string {
	w = TrimSpaceAll(w)
	w = strings.ReplaceAll(w, "Phường ", "P.")
	w = strings.ReplaceAll(w, "Xã ", "X.")
	w = strings.ReplaceAll(w, "Thị trấn", "TT.")
	return w
}

// ReplaceLocationAddress ...
func replaceLocationAddress(a string) string {
	a = TrimSpaceAll(a)
	a = strings.ReplaceAll(a, ",", "")
	return a
}

// TrimSpaceAll ...
func TrimSpaceAll(s string) string {
	return strings.TrimSpace(ReplaceMultipleSpace(s))
}

// ReplaceMultipleSpace ...
func ReplaceMultipleSpace(value string) (result string) {
	re := regexp.MustCompile(`\s+`)
	out := re.ReplaceAllString(value, " ")
	return out
}
