package libhost_test

import (
	"fmt"
	"strings"
	"testing"
	"time"

	. "github.com/ipcjk/mlxsh/libhost"
)

var hostYaml = `
- Hostname: decix-router
  Username: decixUser
  Password: decixPassword
  EnablePassword: enableDecix
  KeyFile: id_rsa_decix
  SSHPort: 2242
  SpeedMode: False
  ScriptFile: scripts/bgp_summary
  Labels:
    location: frankfurt
    environment: production
    type: netiron
- Hostname: amsix-router
  Username: amsix-user
  Password: amxis-password
  EnablePassword: enableAmsix
  SSHKeyFile: id_rsa_amsix
  SSHPort: 22
  StrictHostCheck: False
  SpeedMode: False
  ScriptFile: scripts/bgp_summary
  Labels:
    location: amsterdam
    environment: production
    type: netiron
- Hostname: foo-router
  Username: amsix-user
  Password: amsix-password
  EnablePassword: enableAmsix
  KeyFile: id_rsa_amsix
  SSHPort: 22
  SSHIP: 172.15.4.1
  StrictHostCheck: False
  SpeedMode: False
  ScriptFile: scripts/bgp_summary
  Labels:
    location: berlin
    environment: stage
    type: netiron`

func TestLoadFromYaml(t *testing.T) {

	r := strings.NewReader(hostYaml)

	hostsConfig, err := LoadAllFromYAML(r)
	if err != nil {
		t.Error("Cant parse YAML")
	}

	if hostsConfig[0].Labels["location"] != "frankfurt" {
		t.Error("Cant find city of frankfurt in list member 0")
	}

	if hostsConfig[1].Labels["location"] != "amsterdam" {
		t.Error("Cant find city of amsterdam in list member 1")
	}

	if hostsConfig[2].SSHIP != "172.15.4.1" {
		t.Error("Cant find ssh ip for amsterdam router")
	}

	if hostsConfig[2].SSHPort != 22 {
		t.Error("Cant find ssh port for amsterdam router")
	}

	if hostsConfig[2].KeyFile != "id_rsa_amsix" {
		t.Error("Cant find ssh key file for amsterdam router")
	}

}

func TestLoadHostNameFromYaml(t *testing.T) {
	r := strings.NewReader(hostYaml)
	cliLabel := ""
	cliHostname := "amsix-router"

	selectedHosts, _, err := LoadMatchesFromYAML(r, cliLabel, cliHostname)
	if err != nil {
		t.Error(err)
	}

	if len(selectedHosts) != 1 {
		t.Errorf("Too many or less hosts found: %d", len(selectedHosts))
	}

	if len(selectedHosts) > 1 && selectedHosts[0].Hostname != "amsix-router" {
		t.Error("Router not found")
	}

}

func TestMatchLabels(t *testing.T) {
	r := strings.NewReader(hostYaml)
	hostsConfig, err := LoadAllFromYAML(r)
	if err != nil {
		t.Error("Cant parse YAML")
	}

	if hostsConfig[0].MatchLabels("environment=production,type=netiron,location=frankfurt") != true {
		t.Error("Test did not match expected labels")
	}

	var count int
	for _, hosts := range hostsConfig {
		if hosts.MatchLabels("environment=production") {
			count++
		}
	}

	if count != 2 {
		fmt.Print("Not enough hosts matched for label environment")
	}

}

func TestAppCli(t *testing.T) {
	r := strings.NewReader(hostYaml)
	hostsConfig, err := LoadAllFromYAML(r)
	if err != nil {
		t.Error("Cant parse YAML")
	}

	for x := range hostsConfig {
		hostsConfig[x].ApplyCliSettings("script", "config", time.Second*10, time.Second*5, false, "", "")
	}

	for x := range hostsConfig {
		if hostsConfig[x].ScriptFile != "script" || hostsConfig[x].ConfigFile != "config" {
			t.Error("ApplyCliSettings did not set parameters right ")
		}
	}

}
