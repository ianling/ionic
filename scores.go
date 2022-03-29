package ionic

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/ion-channel/ionic/risk"
)

// GetScores takes one or more purl or other software ids, then performs a request for scores
// against the Ion API, returning a set of scores based on the ids
func (ic *IonClient) GetScores(ids []string, token string) ([]risk.Scores, error) {
	body, err := json.Marshal(ids)
	if err != nil {
		return nil, fmt.Errorf("session: failed to marshal request body: %v", err.Error())
	}

	buff := bytes.NewBuffer(body)

	b, err := ic.Post(risk.GetScoresEnpoint, token, nil, *buff, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get productidentifiers search: %v", err.Error())
	}

	var results []risk.Scores
	err = json.Unmarshal(b, &results)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal product search results: %v", err.Error())
	}

	return results, nil
}
