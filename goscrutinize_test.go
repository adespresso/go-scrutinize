package main

import (
	"os"
	"testing"
)

func TestSelfScrutinize(t *testing.T) {
	main()

	if _, err := os.Stat("coverage.xml"); os.IsNotExist(err) {
		t.Error("coverage.xml not generated")
	}
	if _, err := os.Stat("checkstyle_report.xml"); os.IsNotExist(err) {
		t.Error("checkstyle_report.xml not generated")
	}
}
