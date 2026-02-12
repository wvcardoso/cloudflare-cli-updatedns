package cmd

import (
	"io"
	"net/http"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

//	STRUCTS

type Config struct {
	Cloudflare struct {
		ZoneID   string `yaml:"zone_id"`
		APIToken string `yaml:"api_token"`
		APIURL   string `yaml:"api_url"`
	} `yaml:"cloudflare"`

	IP struct {
		CheckURL string `yaml:"check_url"`
	} `yaml:"ip"`
}

type RecordDNS struct {
	ID      string `json:"id,omitempty"` // omitempty pois no create/update o ID não vai no body
	Name    string `json:"name"`
	Type    string `json:"type"`
	Content string `json:"content"`
}

type CloudflareResponse struct {
	Success bool        `json:"success"`
	Errors  []any       `json:"errors"`
	Result  []RecordDNS `json:"result"`
}

type UpdateDNSRequest struct {
	Type    string `json:"type"`
	Name    string `json:"name"`
	Content string `json:"content"`
	TTL     int    `json:"ttl"`
	Proxied bool   `json:"proxied"`
}

//	FUNÇÕES

func getPublicIP() (string, error) {
	resp, err := http.Get(AppConfig.IP.CheckURL)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	// Remove quebra de linha
	ip := strings.TrimSpace(string(body))

	return ip, nil
}

var AppConfig Config

func LoadConfig(path string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	return yaml.Unmarshal(data, &AppConfig)
}
