package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

var UpdateDNS = &cobra.Command{
	Use:   "update-dns",
	Short: "Atualiza os DNS",
	Run: func(cmd *cobra.Command, args []string) {

		records, err := RunGetDNS()
		if err != nil {
			log.Fatal(err)
		}

		ip, err := getPublicIP()
		if err != nil {
			panic(err)
		}
		println()
		fmt.Printf("Meu IP agora: %s\n", ip)
		println()

		if len(records) >= 1 {

			for _, record := range records {
				RunUpdateDNS(
					AppConfig.Cloudflare.ZoneID,
					record.ID,
					AppConfig.Cloudflare.APIToken,
					record.Name,
					ip,
				)
			}
		} else {
			fmt.Println("Nenhum registro de DNS encontrado")
		}

	},
}

func init() {
	rootCmd.AddCommand(UpdateDNS)
}

func RunUpdateDNS(zoneID, recordID, apiToken, name, ip string) error {

	payload := UpdateDNSRequest{
		Type:    "A",
		Name:    name,
		Content: ip,
		TTL:     120,
		Proxied: false,
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("PUT", "https://"+AppConfig.Cloudflare.APIURL+"/"+zoneID+"/dns_records/"+recordID, bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", "Bearer "+apiToken)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("erro ao atualizar DNS: %s", resp.Status)
	} else {
		fmt.Printf("DNS atualizado: %s -> %s\n", ip, name)

	}

	return nil
}

func RunGetDNS() ([]RecordDNS, error) {

	req, err := http.NewRequest("GET", "https://"+AppConfig.Cloudflare.APIURL+"/"+AppConfig.Cloudflare.ZoneID+"/dns_records", nil)
	if err != nil {
		log.Fatal(err)
	}

	// Exemplo de Headers (ajuste com seu Token)
	req.Header.Set("Authorization", "Bearer "+AppConfig.Cloudflare.APIToken)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("erro na API: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var data CloudflareResponse

	if err := json.Unmarshal(body, &data); err != nil {
		return nil, err
	}

	var recordDNS []RecordDNS

	fmt.Printf("%-35s | %-20s\n", "ID", "NAME")
	fmt.Println("------------------------------------------------------------------")
	for _, record := range data.Result {
		fmt.Printf("%-35s | %-20s\n", record.ID, record.Name)
		recordDNS = append(recordDNS, record)
	}
	fmt.Println("------------------------------------------------------------------")

	fmt.Println("Total registros:", len(recordDNS))

	return recordDNS, nil
}
