package model

type License struct {
	Id             string
	Creator        string
	Product        string
	ConcurrentIps  int
	ClientInfo     string
	AdditionalInfo string
	IpLog          []string
}
