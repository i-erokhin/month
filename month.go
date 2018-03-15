package month

import (
	"encoding/json"
	"fmt"
	"time"
)


type Month struct {
	Year  int
	Month int
}

func (m *Month) String() string {
	return fmt.Sprintf("%d-%02d", m.Year, int(m.Month))
}

func (m *Month) AsISODate() string {
	return fmt.Sprintf("%d-%02d-01", m.Year, int(m.Month))
}

func (m *Month) IsCurrent() bool {
	return m.Year == time.Now().Year() && m.Month == int(time.Now().Month())
}

func (m *Month) IsPrevious() bool {
	currentMonth := int(time.Now().Month())
	if currentMonth == 1 {
		return m.Year == time.Now().Year()-1 && m.Month == 12
	} else {
		return m.Year == time.Now().Year() && m.Month == int(time.Now().Month())-1
	}
}


func (m *Month) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%q", m.String())), nil
}

func (m *Month) UnmarshalJSON(b []byte) error {
	var str string
	err := json.Unmarshal(b, &str)
	if err != nil {
		return err
	}
	per, err := String(str).Period()
	if err != nil {
		return err
	}
	m.Year = per.Year
	m.Month = per.Month
	return nil
}


func NewCurrent() *Month {
	return &Month{
		Year:  time.Now().Year(),
		Month: int(time.Now().Month()),
	}
}

func NewPrevious() *Month {
	currentMonth := int(time.Now().Month())
	if currentMonth == 1 {
		return &Month{
			Year:  time.Now().Year() - 1,
			Month: 12,
		}
	} else {
		return &Month{
			Year:  time.Now().Year(),
			Month: int(time.Now().Month()) - 1,
		}
	}
}
