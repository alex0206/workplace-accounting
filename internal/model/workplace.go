package model

// WorkplaceInfo has information about a workplace
type WorkplaceInfo struct {
	ComputerName string `json:"computer_name"`
	IP           string `json:"ip"`
	Username     string `json:"username"`
}
