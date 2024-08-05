// resource_uptime_kuma_monitor.go
package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"fmt"
	"net/http"
	"bytes"
	"encoding/json"
)

func resourceUptimeKumaMonitor() *schema.Resource {
	return &schema.Resource{
		Create: resourceUptimeKumaMonitorCreate,
		Read:   resourceUptimeKumaMonitorRead,
		Update: resourceUptimeKumaMonitorUpdate,
		Delete: resourceUptimeKumaMonitorDelete,

		Schema: map[string]*schema.Schema{
			// Define all the parameters available in the API here
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"url": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Add other parameters as per the API documentation
		},
	}
}

func resourceUptimeKumaMonitorCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)

	data := map[string]interface{}{
		"name": d.Get("name").(string),
		"url":  d.Get("url").(string),
		// Include other parameters as per the API documentation
	}

	jsonData, _ := json.Marshal(data)
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/api/monitor", client.URL), bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", client.Token))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to create monitor, status code: %d", resp.StatusCode)
	}

	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)
	id, ok := result["id"].(string)
	if !ok {
		return fmt.Errorf("failed to get monitor id from response")
	}
	d.SetId(id)
	return resourceUptimeKumaMonitorRead(d, m)
}

func resourceUptimeKumaMonitorRead(d *schema.ResourceData, m interface{}) error {
	// Implement the read function to fetch the monitor details
	return nil
}

func resourceUptimeKumaMonitorUpdate(d *schema.ResourceData, m interface{}) error {
	// Implement the update function to update the monitor details
	return nil
}

func resourceUptimeKumaMonitorDelete(d *schema.ResourceData, m interface{}) error {
	// Implement the delete function to delete the monitor
	return nil
}
