package main

import (
	"flag"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type MacosConfigData struct {
	Label       string
	ProgramPath string
	ConfigPath  string
	StdOutPath  string
	StdErrPath  string
	KeepAlive   bool
	RunAtLoad   bool
}

var macosConfig MacosConfigData = MacosConfigData{
	Label:       "com.guiln.promgateway",
	ProgramPath: "/usr/local/bin/promgateway",
	ConfigPath:  "/etc/promgateway.conf.json",
	StdOutPath:  "/tmp",
	StdErrPath:  "/tmp",
	RunAtLoad:   true,
	KeepAlive:   true,
}

var lggr *log.Logger = log.New(os.Stdout, "[INSTALLER] ", 0)

// TODO: system.d config format
// for time being this installer only supports macos daemon (plist format)
// configuration and folder structure
func main() {
	if !doesUserGivesPermission() {
		return
	}

	filePaths := getCliEntry()

	lggr.Println("Generating plist file...")
	if err := generatePlistFile(); err != nil {
		lggr.Fatal(err)
		os.Exit(1)
	}

	lggr.Println("Copying binary files...")
	if err := putFile(filePaths.binaryPath, macosConfig.ProgramPath); err != nil {
		lggr.Fatal(err)
		os.Exit(1)
	}

	lggr.Println("Copying config files...")
	if err := putFile(filePaths.configPath, macosConfig.ConfigPath); err != nil {
		lggr.Fatal(err)
		os.Exit(1)
	}
	lggr.Println("Installation succeeded!")
}

func doesUserGivesPermission() bool {
	fmt.Printf("Do you like to proceed with the installation of the %s?\n(Y/n)\n", macosConfig.Label)
	var usersAnswer string
	fmt.Scanln(&usersAnswer)
	if strings.TrimSpace(strings.ToUpper(usersAnswer)) == "Y" {
		return true
	}

	return false
}

func getCliEntry() struct {
	binaryPath string
	configPath string
} {
	binaryPath := flag.String("binary-file", "", "binary file path of promgateway daemon.")
	configPath := flag.String("config-file", "", "config file path of promgateway daemon.")

	flag.Parse()

	return struct {
		binaryPath string
		configPath string
	}{
		binaryPath: *binaryPath,
		configPath: *configPath,
	}
}

func generatePlistFile() error {
	plistUserInstallationPath := fmt.Sprintf("%s/Library/LaunchAgents/%s.plist", os.Getenv("HOME"), macosConfig.Label)
	f, err := os.Create(plistUserInstallationPath)
	defer f.Close()
	if err != nil {
		lggr.Fatalf("Template file creation failed: %s", err)
		return err
	}

	t := template.Must(template.New("launchdConfig").Parse(configurationPlistTemplate))
	err = t.Execute(f, macosConfig)
	if err != nil {
		lggr.Fatalf("Template generation failed: %s", err)
		return err
	}

	return nil
}

// putFile writes file from source to the destination.
func putFile(source, destination string) error {
	input, err := ioutil.ReadFile(source)
	if err != nil {
		return err
	}

	const permissionMask = 0644
	err = ioutil.WriteFile(destination, input, permissionMask)
	if err != nil {
		return err
	}
	return nil
}

// Configuration templates.

// Macos configuration plist template.
const configurationPlistTemplate string = `
<?xml version='1.0' encoding='UTF-8'?>
 <!DOCTYPE plist PUBLIC \"-//Apple Computer//DTD PLIST 1.0//EN\" \"http://www.apple.com/DTDs/PropertyList-1.0.dtd\" >
 <plist version='1.0'>
   <dict>
     <key>Label</key><string>{{.Label}}</string>
     <key>ProgramArguments</key>
        <array>
          <string>{{.ProgramPath}}</string>
          <string>-config-file</string>
          <string>{{.ConfigPath}}</string>
        </array>
	 <key>StandardOutPath</key><string>{{.StdOutPath}}/{{.Label}}.out.lggr</string>
     <key>StandardErrorPath</key><string>{{.StdErrPath}}/{{.Label}}.err.lggr</string>
     <key>KeepAlive</key><{{.KeepAlive}}/>
     <key>RunAtLoad</key><{{.RunAtLoad}}/>
   </dict>
</plist>
	`
