package netironDevice

import (
	"bufio"
	"fmt"
	"golang.org/x/crypto/ssh"
	"io"
	"log"
	"strings"
	"time"
)

type netironDevice struct {
	port                                        int
	enable, hostname, model, password, username string

	debug bool

	promptModes  map[string]string
	promptMode   string
	readTimeout  time.Duration
	speedMode    bool
	writeTimeout time.Duration
	w io.Writer
	ExitCode int

	sshConfigPrompt, sshEnabledPrompt, sshUnprivilegedPrompt string

	sshConfig          *ssh.ClientConfig
	sshConfigPromptPre string
	sshConnection      *ssh.Client
	sshSession         *ssh.Session
	sshStdinPipe       io.WriteCloser
	sshStdoutPipe      io.Reader
	sshStdErrPipe      io.Reader
}

/*
NetironDevice returns a new
netironDevice object
*/
func NetironDevice(model string, hostname string, port int, enable, username, password string, readTimeout time.Duration,
	writeTimeout time.Duration, debug bool, speedMode bool, w io.Writer) *netironDevice {

	return &netironDevice{model: model, port: port, hostname: hostname, enable: enable, readTimeout: readTimeout,
		speedMode:                 speedMode, writeTimeout: writeTimeout, debug: debug, w: w, promptModes: make(map[string]string),
		sshConfig:                 &ssh.ClientConfig{User: username, Auth: []ssh.AuthMethod{ssh.Password(password)}}}
}

func (b *netironDevice) ConnectPrivilegedMode() (err error) {
	b.sshConnection, err = ssh.Dial("tcp", fmt.Sprintf("%s:%d", b.hostname, b.port), b.sshConfig)
	if err != nil {
		return err
	}

	b.sshSession, err = b.sshConnection.NewSession()
	if err != nil {
		return err
	}

	b.sshStdoutPipe, err = b.sshSession.StdoutPipe()
	if err != nil {
		return err
	}

	b.sshStdinPipe, err = b.sshSession.StdinPipe()
	if err != nil {
		return err
	}

	b.sshStdErrPipe, err = b.sshSession.StderrPipe()
	if err != nil {
		return err
	}

	err = b.sshSession.Shell()
	if err != nil {
		return err
	}

	b.sshUnprivilegedPrompt, err = b.readTill([]string{">"})
	if err != nil {
		return err
	}

	b.sshEnabledPrompt = strings.Replace(b.sshUnprivilegedPrompt, ">", "#", 1)
	b.sshConfigPrompt = strings.Replace(b.sshUnprivilegedPrompt, ">", "(config)#", 1)
	b.sshConfigPromptPre = strings.Replace(b.sshUnprivilegedPrompt, ">", "(config", 1)

	b.promptModes["sshEnabled"] = b.sshEnabledPrompt
	b.promptModes["sshConfig"] = b.sshConfigPrompt
	b.promptModes["sshConfigPre"] = b.sshConfigPromptPre
	b.promptModes["sshNotEnabled"] = b.sshUnprivilegedPrompt

	if b.debug {
		fmt.Fprintf(b.w,"Enabled:(%s)\n", b.sshEnabledPrompt)
		fmt.Fprintf(b.w,"Not-Enabled:(%s)\n", b.sshUnprivilegedPrompt)
		fmt.Fprintf(b.w,"Config:(%s)\n", b.sshConfigPrompt)
		fmt.Fprintf(b.w,"ConfigSection:(%s)\n", b.sshConfigPromptPre)
	}

	if b.loginDialog() && b.debug {
		log.Println("Logged in")
	}
	return
}

func (b *netironDevice) loginDialog() bool {
	b.write("enable\n")
	_, err := b.readTill([]string{"Password:"})
	if err != nil {
		log.Fatal(err)
	}

	b.write(b.enable + "\n")
	_, err = b.readTillEnabledPrompt()
	if err != nil {
		log.Fatal(err)
	}

	b.promptMode = "sshEnabled"

	return true
}

func (b *netironDevice) write(command string) {
	_, err := b.sshStdinPipe.Write([]byte(command))
	if err != nil {
		log.Fatal(err)
	}

	if b.debug {
		fmt.Fprintf(b.w,"Send command: %s", command)
	}
	time.Sleep(b.writeTimeout)
}

func (b *netironDevice) readTill(search []string) (string, error) {
	shortBuf := make([]byte, 1)
	lineBuffer := make([]byte, 0, 32)
	foundToken := make(chan struct{}, 0)
	defer close(foundToken)

WaitInput:
	for {
		/* Reset the timer, when we received at least 1 byte */
		go func() {
			select {
			case <-(time.After(b.readTimeout)):
				if b.debug {
					log.Println(string(lineBuffer[:]))
				}
				b.sshSession.Close()
				b.sshConnection.Close()
			case <-foundToken:
				return
			}
		}()
		var err error
		if _, err = io.ReadAtLeast(b.sshStdoutPipe, shortBuf, 1); err != nil {
			if err != io.EOF {
				return "", err
			}
		}
		foundToken <- struct{}{}
		lineBuffer = append(lineBuffer, shortBuf[0])
		for x := range search {
			if strings.Contains(string(lineBuffer[:]), search[x]) {
				break WaitInput
			}
		}

	}

	return string(lineBuffer[:]), nil
}

func (b *netironDevice) ConfigureTerminalMode() {
	b.write("conf t\n")
	_, err := b.readTill([]string{"(config)#"})
	if err != nil {
		log.Fatal(err)
	}

	if b.debug {
		log.Println("Configuration mode on")
	}
}

func (b *netironDevice) ExecPrivilegedMode(command string) {
	if err := b.SwitchMode("sshEnabled"); err != nil {
		log.Fatal("Cant switch to enabled mode")
	}

	b.write(command + "\n")
	_, err := b.readTillEnabledPrompt()
	if err != nil {
		log.Fatal(err)
	}
}

func (b *netironDevice) SkipPageDisplayMode() (string, error) {
	if err := b.SwitchMode("sshEnabled"); err != nil {
		return "", fmt.Errorf("Cant switch to enabled mode to execute skip-page-display")
	}

	b.write("skip-page-display\n")
	return b.readTill([]string{b.sshEnabledPrompt})
}

func (b *netironDevice) readTillEnabledPrompt() (string, error) {
	return b.readTill([]string{b.sshEnabledPrompt})
}

func (b *netironDevice) readTillConfigPrompt() (string, error) {
	return b.readTill([]string{b.sshConfigPrompt})
}

func (b *netironDevice) readTillConfigPromptSection() (string, error) {
	return b.readTill([]string{b.sshConfigPromptPre})
}

func (b *netironDevice) SwitchMode(targetMode string) error {
	if b.promptMode == targetMode {
		return nil
	}

	switch b.promptMode {
	case "sshEnabled":
		if targetMode == "sshConfig" {
			b.ConfigureTerminalMode()
		} else {
			b.write("exit\n")
		}
	case "sshConfig":
		if targetMode == "sshEnabled" {
			b.write("end\n")
		} else {
			b.write("end\n")
			b.write("exit\n")
		}
	case "sshNotEnabled":
		if targetMode == "sshEnabled" {
			fmt.Println("LOGIN")
		} else {
			fmt.Println("LOGIN & CONF Mode")
		}

	}

	return nil
}

func (b *netironDevice) GetPromptMode() error {
	b.write("\n")

	mode, err := b.readTill([]string{b.sshConfigPrompt, b.sshEnabledPrompt, b.sshUnprivilegedPrompt})
	if err != nil {
		return fmt.Errorf("Cant find command line mode: %s", err)
	}

	mode = strings.TrimSpace(mode)

	switch mode {
	case b.promptModes["sshEnabled"]:
		b.promptMode = "sshEnabled"
	case b.promptModes["sshConfig"]:
		b.promptMode = "sshConfig"
	case b.promptModes["sshNotEnabled"]:
		b.promptMode = "sshNotEnabled"
	default:
		b.promptMode = "unknown"
	}

	return nil

}

func (b *netironDevice) WriteConfiguration() (err error) {
	if err = b.SwitchMode("sshEnabled"); err != nil {
		return err
	}

	b.write("write memory\n")
	_, err = b.readTill([]string{"(config)#", "Write startup-config done."})
	if err != nil {
		return err
	}

	if b.debug {
		fmt.Fprint(b.w, "Write startup-config done")
	}
	return
}

func (b *netironDevice) CloseConnection() {
	b.sshConnection.Close()
}

func (b *netironDevice) PasteConfiguration(configuration io.Reader) (err error) {
	if err = b.SwitchMode("sshConfig"); err != nil {
		return err
	}

	scanner := bufio.NewScanner(configuration)
	for scanner.Scan() {
		b.write(scanner.Text() + "\n")
		/* Wait till config prompt returns or not ? */
		if !b.speedMode {
			val, err := b.readTillConfigPromptSection()
			if err != nil {
				return err
			}
			if b.debug {
				fmt.Fprintf(b.w, "Captured %s\n", val)
			}
		}
		fmt.Fprint(b.w, "+")
	}
	fmt.Fprint(b.w, "\n")
	return
}

func (b *netironDevice) RunCommandsFromReader(commands io.Reader) (err error) {
	if err = b.SwitchMode("sshEnabled"); err != nil {
		return fmt.Errorf("Cant switch to privileged mode: %s", err)
	}

	scanner := bufio.NewScanner(commands)
	for scanner.Scan() {
		b.write(scanner.Text() + "\n")
		val, err := b.readTillEnabledPrompt()
		if err != nil && err != io.EOF {
			return err
		}
		fmt.Fprintf(b.w, "%s\n", val)
	}
	return err
}
