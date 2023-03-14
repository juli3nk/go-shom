package shom

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/juli3nk/stack/client"
)

const URL = "https://services.data.shom.fr/b2q8lrcdl4s04cbabsj4nhcb/hdm/spm"

type Config struct {
	HarborName string
	Date       string
	Utc        string
}

func New(harborName, date, utc string) (*Config, error) {
	config := Config{
		HarborName: harborName,
		Date:       date,
		Utc:        utc,
	}

	return &config, nil
}

// https://services.data.shom.fr/x13f1b4faeszdyinv9zqxmx1/wfs?service=WFS&version=1.0.0&srsName=EPSG:3857&request=GetFeature&typeName=SPM_PORTS_WFS:liste_ports_spm_h2m&outputFormat=application/json
func Wfs() (*Wfs, error) {
	url := fmt.Sprintf("%s/wfs", URL)

	u, err := client.ParseUrl(url)
	if err != nil {
		return nil, err
	}

	cc := &client.Config{
		Scheme: u.Scheme,
		Host:   u.Host,
		Port:   u.Port,
		Path:   u.Path,
	}

	req, err := client.New(cc)
	if err != nil {
		return nil, err
	}

	req.HeaderAdd("Referer", "https://maree.shom.fr/")
	req.HeaderAdd("Accept", "application/json")

	req.ValueAdd("service", "WFS")
	req.ValueAdd("version", "1.0.0")
	req.ValueAdd("srsName", "EPSG:3857")
	req.ValueAdd("request", "GetFeature")
	req.ValueAdd("typeName", "SPM_PORTS_WFS:liste_ports_spm_h2m")
	req.ValueAdd("outputFormat", "application/json")

	result := req.Get()

	if result.Response.StatusCode != 200 {
		return nil, result.Error
	}

	response := new(Wfs)

	if err := json.Unmarshal(resp, response); err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Config) get(path string, queryValues map[string][]string) ([]byte, error) {
	u, err := client.ParseUrl(URL)
	if err != nil {
		return nil, err
	}

	urlPath := fmt.Sprintf("%s%s", u.Path, path)

	cc := &client.Config{
		Scheme: u.Scheme,
		Host:   u.Host,
		Port:   u.Port,
		Path:   urlPath,
	}

	req, err := client.New(cc)
	if err != nil {
		return nil, err
	}

	req.HeaderAdd("Referer", "https://maree.shom.fr/")
	req.HeaderAdd("Accept", "application/json")

	req.ValueAdd("harborName", c.HarborName)
	req.ValueAdd("date", c.Date)
	req.ValueAdd("utc", c.Utc)

	for key, values := range queryValues {
		for _, v := range values {
			req.ValueAdd(key, v)
		}
	}

	result := req.Get()

	if result.Response.StatusCode != 200 {
		return nil, result.Error
	}

	return result.Body, nil
}

// https://services.data.shom.fr/b2q8lrcdl4s04cbabsj4nhcb/hdm/spm/hlt?harborName=LE_LEGUE_BOUEE&duration=7&date=2022-03-22&utc=standard&correlation=1
func (c *Config) Hlt(duration, correlation int) (*AnnuaireMarees, error) {
	urlPath := "/hlt"

	duration1 := []string{}
	if duration > 0 {
		duration1 = []string{strconv.Itoa(duration)}
	}
	correlation1 := []string{}
	if correlation > 0 {
		correlation1 = []string{strconv.Itoa(correlation)}
	}

	values := map[string][]string{
		"duration":    duration1,
		"correlation": correlation1,
	}

	resp, err := c.get(urlPath, values)
	if err != nil {
		return nil, err
	}

	response := new(AnnuaireMarees)

	if err := json.Unmarshal(resp, response); err != nil {
		return nil, err
	}

	return response, nil
}

// https://services.data.shom.fr/b2q8lrcdl4s04cbabsj4nhcb/hdm/spm/wl?harborName=LE_LEGUE_BOUEE&duration=1&date=2022-03-22&utc=standard&nbWaterLevels=288
func (c *Config) Wl(waterLevels int) (*HauteurEauParHeure, error) {
	urlPath := "/wl"

	waterLevels1 := []string{}
	if waterLevels > 0 {
		waterLevels1 = []string{strconv.Itoa(waterLevels)}
	}

	values := map[string][]string{
		"waterLevels": waterLevels1,
	}

	resp, err := c.get(urlPath, values)
	if err != nil {
		return nil, err
	}

	response := new(HauteurEauParHeure)

	if err := json.Unmarshal(resp, response); err != nil {
		return nil, err
	}

	return response, nil
}

// https://services.data.shom.fr/b2q8lrcdl4s04cbabsj4nhcb/hdm/spm/coeff?harborName=LE_LEGUE_BOUEE&duration=365&date=2022-03-01&utc=1&correlation=1
func (c *Config) Coeff(duration, correlation int) (*Coeff, error) {
	urlPath := "/coeff"

	duration1 := []string{}
	if duration > 0 {
		duration1 = []string{strconv.Itoa(duration)}
	}
	correlation1 := []string{}
	if correlation > 0 {
		correlation1 = []string{strconv.Itoa(correlation)}
	}

	values := map[string][]string{
		"duration":    duration1,
		"correlation": correlation1,
	}

	resp, err := c.get(urlPath, values)
	if err != nil {
		return nil, err
	}

	response := new(Coeff)

	if err := json.Unmarshal(resp, response); err != nil {
		return nil, err
	}

	return response, nil
}
