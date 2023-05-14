package client

import (
	"encoding/xml"
	"fmt"
	"neb/src/utils/cnsl"
	"os/exec"
	"strings"
)

type String struct {
	Name  string `xml:"name,attr"`
	Value string `xml:",chardata"`
}

type Strings struct {
	Strings []String `xml:"string"`
}

func GetSharedPrefsToken() *string {
	cmd := exec.Command("su", "-c", "cat /data/data/software.simplicial.nebulous/shared_prefs/application.MainActivity.xml")
	output, err := cmd.Output()
	if err != nil {
		cnsl.Error("error getting account token. you either have no root or you are running this script on windows.")
		return nil
	}

	decoder := xml.NewDecoder(strings.NewReader(string(output)))

	var strings Strings
	err = decoder.Decode(&strings)
	if err != nil {
		cnsl.Error(fmt.Sprintf("unknown error getting account token -> %s", err))
		return nil
	}

	var loginTicket *string
	for _, str := range strings.Strings {
		if str.Name == "loginTicket" {
			loginTicket = &str.Value
			break
		}
	}
	if loginTicket == nil {
		nada := "nada"
		loginTicket = &nada
	}

	return loginTicket
}
