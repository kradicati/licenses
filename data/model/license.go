package model

type License struct {
	Id             string   `firestore:"id" json:"id"`
	Creator        string   `firestore:"creator" json:"creator"`
	Created        int64    `firestore:"created" json:"created"`
	Expires        int64    `firestore:"expires" json:"expires"`
	Product        string   `firestore:"product" json:"product"`
	WhitelistedIps []string `firestore:"whitelisted_ips" json:"whitelisted_ips"`
	BlacklistedIps []string `firestore:"blacklisted_ips" json:"blacklisted_ips"`
	AdditionalInfo string   `firestore:"additional_info" json:"additional_info"`
	IpLog          []IpLog  `firestore:"ip_log" json:"ip_log"`
}

type IpLog struct {
	Ip     string `firestore:"ip" json:"ip"`
	Time   int64  `firestore:"time" json:"time"`
	Status bool   `firestore:"status" json:"status"`
	Reason string `firestore:"reason" json:"reason"`
}
