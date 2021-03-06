package domains

import "github.com/gophercloud/gophercloud"

func getURL(client *gophercloud.ServiceClient, domainID string) string {
	return client.ServiceURL("domains", domainID)
}

func listURL(client *gophercloud.ServiceClient) string {
	return client.ServiceURL("domains")
}
