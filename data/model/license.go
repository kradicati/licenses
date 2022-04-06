package model

type License struct {
	Id             string   `firestore:"id"`
	Creator        string   `firestore:"creator"`
	Product        string   `firestore:"product"`
	ConcurrentIps  int      `firestore:"concurrent_ips"`
	ClientInfo     string   `firestore:"client_info"`
	AdditionalInfo string   `firestore:"additional_info"`
	IpLog          []string `firestore:"ip_log"`
}
