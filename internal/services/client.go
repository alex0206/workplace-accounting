package services

import (
	"fmt"
	"net"
	"os"
	"os/user"

	"github.com/alex0206/workplace-accounting/internal/model"
)

// APIClient server api client
type APIClient interface {
	UpdateWorkplace(info *model.WorkplaceInfo) error
}

// ClientService describe workplace client side
type ClientService struct {
	apiClient APIClient
}

// NewClientService create a new client service
func NewClientService(client APIClient) *ClientService {
	return &ClientService{apiClient: client}
}

// UpdateWorkplace performing workplace creating or updating
func (s *ClientService) UpdateWorkplace() error {
	ip, err := s.ip()
	if err != nil {
		return fmt.Errorf("error getting ip address: %v", err)
	}

	hostName, err := os.Hostname()
	if err != nil {
		return fmt.Errorf("error getting hostname: %v", err)
	}

	currentUser, err := user.Current()
	if err != nil {
		return fmt.Errorf("error getting user: %v", err)
	}

	wp := model.WorkplaceInfo{
		ComputerName: hostName,
		IP:           ip,
		Username:     currentUser.Username,
	}

	return s.apiClient.UpdateWorkplace(&wp)
}

func (s *ClientService) ip() (string, error) {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return "", err
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP.String(), nil
}
