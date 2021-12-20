package cert

import (
	"fmt"
	"strings"
	"time"
)

const MaxLengthCourse int = 20
const MaxLengthName int = 30

type Cert struct {
	Course             string
	Name               string
	Date               time.Time
	LabelTitle         string
	LabelCompletion    string
	LabelPresented     string
	LabelParticipation string
	LabelDate          string
}

type Saver interface {
	Save(c Cert) error
}

func ValidateCourse(course string) (string, error) {
	const SUFFIX string = " course"
	c, err := ValidateStr(course, MaxLengthCourse)
	if err != nil {
		return c, err
	}
	if !strings.HasSuffix(c, SUFFIX) {
		course += SUFFIX
	}
	return strings.ToUpper(course), nil
}

func New(course, name, date string) (*Cert, error) {
	c, err := ValidateCourse(course)
	if err != nil {
		return nil, err
	}
	n, err := ValidateCourse(name)
	if err != nil {
		return nil, err
	}
	d, err := ParseDate(date)
	if err != nil {
		return nil, err
	}
	cert := &Cert{
		Course:             c,
		Name:               n,
		LabelTitle:         fmt.Sprintf("%v Certificate - %v", course, name),
		LabelCompletion:    "Certificate of completion",
		LabelPresented:     "This certificate is presented To",
		LabelParticipation: fmt.Sprintf("For participation  in the %v", course),
		LabelDate:          fmt.Sprintf("Date: %v", d.Format("02/01/2006")),
		Date:               d,
	}
	return cert, nil
}

func ValidateStr(str string, maxLen int) (string, error) {
	c := strings.TrimSpace(str)
	if len(c) <= 0 || len(c) > maxLen {
		return c, fmt.Errorf("invalid string. got='%s', len='%d'", c, len(c))
	}
	return c, nil
}

func ValidateName(name string) (string, error) {
	c, err := ValidateStr(name, MaxLengthName)
	if err != nil {
		return c, err
	}
	return strings.ToUpper(name), nil
}

func ParseDate(date string) (time.Time, error) {
	d, err := time.Parse("2006-01-02", date)
	if err != nil {
		return d, err
	}
	return d, nil
}
