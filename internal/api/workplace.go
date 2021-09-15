package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/alex0206/workplace-accounting/internal/model"
)

// WorkplaceAPIClient api client to server side
type WorkplaceAPIClient struct {
	host string
}

// NewWorkplaceAPIClient create a new api client
func NewWorkplaceAPIClient(host string) *WorkplaceAPIClient {
	return &WorkplaceAPIClient{host: host}
}

// UpdateWorkplace method for updating or creating a workplace
func (c *WorkplaceAPIClient) UpdateWorkplace(info *model.WorkplaceInfo) error {
	data, err := json.Marshal(info)
	if err != nil {
		return fmt.Errorf("error marhaling workplace info: " + err.Error())
	}

	resp, err := http.Post(c.host+"/workplaces", "application/json", bytes.NewBuffer(data))
	if err != nil {
		return fmt.Errorf("error performing request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("unspecified status code %v", resp.StatusCode)
	}

	return nil
}
