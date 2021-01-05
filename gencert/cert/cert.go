package cert

import (
	"fmt"
	"strings"
	"time"
)

// MaxLenCourse is max size for string Course
const MaxLenCourse int = 20

// MaxLenName is max size for string Name
const MaxLenName int = 30

//Cert struct
type Cert struct {
	Course string
	Name   string
	Date   time.Time

	LabelTitle         string
	LabelCompletion    string
	LabelPresented     string
	LabelParticipation string
	LabelDate          string
}

// Saver interface
type Saver interface {
	Save(c Cert) error
}

// New method for create a new cert item
func New(course, name, date string) (*Cert, error) {
	c, err := validateCourse(course)
	if err != nil {
		return nil, err
	}
	n, err := validateName(name)
	if err != nil {
		return nil, err
	}
	d, err := parseDate(date)
	if err != nil {
		return nil, err
	}

	cert := &Cert{
		Course:             c,
		Name:               n,
		Date:               d,
		LabelTitle:         fmt.Sprintf("%v Certificater - %v", c, n),
		LabelCompletion:    "Certificate of Completion",
		LabelPresented:     "This Certificate is Presented To",
		LabelParticipation: fmt.Sprintf("For participation in the %v", c),
		LabelDate:          fmt.Sprintf("Date: %v", d.Format("02/01/2006")),
	}
	return cert, nil
}

func parseDate(date string) (time.Time, error) {
	t, err := time.Parse("2006/01/02", date)
	if err != nil {
		return t, err
	}
	return t, nil
}

func validateName(name string) (string, error) {
	n, err := validateStr(name, MaxLenName)
	if err != nil {
		return "", err
	}
	return strings.ToTitle(n), nil
}

func validateCourse(course string) (string, error) {
	c, err := validateStr(course, MaxLenCourse)
	if err != nil {
		return "", err
	}
	if !strings.HasSuffix(c, " course") {
		c = c + " course"
	}
	return strings.ToTitle(c), nil
}

func validateStr(str string, maxLen int) (string, error) {
	s := strings.TrimSpace(str)
	if len(s) <= 0 {
		return s, fmt.Errorf("Invalid string. got='%s', len=%d", s, len(s))
	} else if len(s) >= maxLen {
		return s, fmt.Errorf("Too long String (max 19 char). got='%s', len=%d", s, len(s))
	}
	return s, nil
}
