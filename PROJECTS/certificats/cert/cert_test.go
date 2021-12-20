package cert

import "testing"

func TestValidCertData(t *testing.T) {
	c, err := New("Golang", "bob", "2018-05-31")
	if err != nil {
		t.Errorf("Cert data should be valid: %v\n", err)
	}
	if c == nil {
		t.Errorf("Cert should be a valid reference: got=nil")
	}
	// if c.Course != "GOLANG COURSE" {
	// 	t.Errorf("Course name is not valid: Expected='GOLANG COURSE', got=%v", c.Course)
	// }
}

func TestCourseEmptyValue(t *testing.T) {
	_, err := New("", "bob", "2018-05-31")
	if err == nil {
		t.Errorf("Error should be return on empty course name")
	}
}

func TestCourseNameTooLong(t *testing.T) {
	_, err := New("qsffqgshsfghqdfgsrgrgssthqegsghgqssfsfgbsfgsdfgsdfgsdfgsddfgsddgfs", "bob", "2018-05-31")
	if err == nil {
		t.Errorf("Error should be return on a too long course name")
	}
}

func TestNameEmptyValue(t *testing.T) {
	_, err := New("Golang", "qsdqsdqsdqsdqsdqsdqsdqsdqsdqsdqd", "2018-05-31")
	if err == nil {
		t.Errorf("Error should be return on empty student name")
	}
}

func TestNameTooLong(t *testing.T) {
	_, err := New("golang", "bqsdqsdqsdqsdqsdqsdqsdqsdqsdqsdqsdqsdqsob", "2018-05-31")
	if err == nil {
		t.Errorf("Error should be return on a too long student name")
	}
}
