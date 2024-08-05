// main.go
package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: func() *schema.Provider {
			return &schema.Provider{
				Schema: map[string]*schema.Schema{
					"url": {
						Type:        schema.TypeString,
						Optional:    true,
						DefaultFunc: schema.EnvDefaultFunc("UPTIME_KUMA_URL", nil),
					},
					"username": {
						Type:        schema.TypeString,
						Optional:    true,
						DefaultFunc: schema.EnvDefaultFunc("UPTIME_KUMA_USERNAME", nil),
					},
					"password": {
						Type:        schema.TypeString,
						Optional:    true,
						Sensitive:   true,
						DefaultFunc: schema.EnvDefaultFunc("UPTIME_KUMA_PASSWORD", nil),
					},
					"api_key": {
						Type:        schema.TypeString,
						Optional:    true,
						Sensitive:   true,
						DefaultFunc: schema.EnvDefaultFunc("UPTIME_KUMA_API_KEY", nil),
					},
				},
				ResourcesMap: map[string]*schema.Resource{
					"uptime_kuma_monitor": resourceUptimeKumaMonitor(),
				},
				ConfigureFunc: providerConfigure,
			}
		},
	})
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	config := Config{
		URL:      d.Get("url").(string),
		Username: d.Get("username").(string),
		Password: d.Get("password").(string),
		APIKey:   d.Get("api_key").(string),
	}
	return config.Client()
}

type Config struct {
	URL      string
	Username string
	Password string
	APIKey   string
}

func (c *Config) Client() (*Client, error) {
	client := &Client{
		URL:      c.URL,
		Username: c.Username,
		Password: c.Password,
		APIKey:   c.APIKey,
	}
	err := client.Authenticate()
	if err != nil {
		return nil, err
	}
	return client, nil
}
