// Code generated by "stringer -type=Suit"; DO NOT EDIT.

package deck

import "strconv"

const _Suit_name = "SpadesDiamondsClubsHearts"

var _Suit_index = [...]uint8{0, 6, 14, 19, 25}

func (i Suit) String() string {
	if i < 0 || i >= Suit(len(_Suit_index)-1) {
		return "Suit(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Suit_name[_Suit_index[i]:_Suit_index[i+1]]
}
