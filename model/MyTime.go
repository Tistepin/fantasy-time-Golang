package model

import "time"

/**
* User:徐国纪
* Create_Time:上午 10:46
**/

type MyTime struct {
	time.Time
}

func (m *MyTime) UnmarshalJSON(data []byte) error {
	// Ignore null, like in the main JSON package.
	if string(data) == "null" || string(data) == `""` {
		return nil
	}
	// Fractional seconds are handled implicitly by Parse.
	tt, err := time.Parse(`"`+time.DateTime+`"`, string(data))
	*m = MyTime{tt}
	return err
}
