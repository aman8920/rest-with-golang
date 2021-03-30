package util

import (
	"database/sql/driver"
	"fmt"
	"log"
	"strconv"
	"time"
)

type SimpleDate struct {
	time.Time
}

const layoutISO = "2006-01-02"

func (sd SimpleDate) MarshalJSON() ([]byte, error) {
	log.Printf("\n ### calling marshall %v", sd)
	out := sd.Format(layoutISO)
	log.Printf("\n Date Un Formatted %v", sd)
	log.Printf("\n Date Formatted %s", out)
	return []byte(`"` + out + `"`), nil
}

func (sd SimpleDate) Value() (driver.Value, error) {
	log.Printf("\n ### calling value %v", sd)
	var zeroTime time.Time
	if sd.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return sd.Time, nil
}
func (sd *SimpleDate) Scan(v interface{}) error {
	log.Printf("\n ### calling scan %v", sd)
	value, ok := v.(time.Time)
	if ok {
		*sd = SimpleDate{Time: value}
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}

func (sd *SimpleDate) UnmarshalJSON(b []byte) error {
	log.Printf("\n ### calling Unmarshall %v", sd)
	dateStr, err := strconv.Unquote(string(b))
	if err != nil {
		return nil
	}
	log.Printf("Layout: %s", layoutISO)
	log.Printf("Date String: %s", dateStr)
	t, err := time.Parse(layoutISO, dateStr)
	if err != nil {
		return err
	}
	*sd = SimpleDate{t}
	return nil
}
