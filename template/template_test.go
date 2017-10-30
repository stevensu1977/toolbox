package template

import (
	"os"
	"testing"
	"time"
)

func TestLoadTemplate(t *testing.T) {
	SetTemplatePath("../temp")

	tpl, err := LoadTemplate("login.tpl")

	tpl.Execute(os.Stdout, time.Now())
	if err != nil {
		t.Fatal(err)
	}

}
