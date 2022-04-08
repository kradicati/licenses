package model

type License struct {
	Id             string   `firestore:"id" json:"id,omitempty"`
	Creator        string   `firestore:"creator" json:"creator,omitempty"`
	Created        int64    `firestore:"created" json:"created,omitempty"`
	Expires        int64    `firestore:"expires" json:"expires,omitempty"`
	Product        string   `firestore:"product" json:"product,omitempty"`
	WhitelistedIps []string `firestore:"whitelisted_ips" json:"whitelisted_ips,omitempty"`
	BlacklistedIps []string `firestore:"blacklisted_ips" json:"blacklisted_ips,omitempty"`
	AdditionalInfo string   `firestore:"additional_info" json:"additional_info,omitempty"`
	IpLog          []string `firestore:"ip_log" json:"ip_log,omitempty"`
}
