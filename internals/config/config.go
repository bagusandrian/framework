package config

import (
	"fmt"
	"net"
	"os"
	"path/filepath"
	"strings"
	yaml "gopkg.in/yaml.v2"
	const "github.com/bagusandrian/framework/internals/constants"
)


func New(repoName string) (*Config, error) {
	filename := getConfigFile(repoName)
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var cfg Config
	err = yaml.NewDecoder(f).Decode(&cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}

type Config struct {
	Server         Server `yaml:"server"`
}
func getHostInfo() (name, ip string) {
	var err error
	name, err = os.Hostname()
	if err != nil {
		name = "-"
	}

	ip = "-"
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return name, ip
	}
	defer conn.Close()

	localAddr, ok := conn.LocalAddr().(*net.UDPAddr)
	if ok {
		ip = localAddr.IP.String()
	}
	return name, ip
}

// getConfigFile get  config file name
// - files/etc/repo_name/repo_name.development.yaml in dev
// - otherwise /etc/repo_name/repo_name.{TKPENV}.yaml
func getConfigFile(repoName string) string {
	var (
		SysEnv   = os.Getenv("SysEnv")
		filename = fmt.Sprintf("%s.%s.yaml", repoName, SysEnv)
	)

	// for non dev env, use config in /etc
	if SysEnv != const.DevelopmentEnv {
		return filepath.Join("/etc", repoName, string(SysEnv), filename)
	}

	// use local files in dev
	repoPath := filepath.Join(os.Getenv("GOPATH"), "src/github.com/bagusandrian", repoName)
	return filepath.Join(repoPath, "files/etc", repoName, filename)
}

func getEnv() string{
	env := os.Getenv("SysEnv")
	if env == "" {
		return "development"
	}
	return env
}
