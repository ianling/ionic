package ionic

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/ion-channel/ionic/languages"
)

//GetLanguages takes a text input and returns any matching languages
func (ic *IonClient) GetLanguages(text string, token string) ([]languages.Language, error) {
	b, err := ic.Post(languages.LanguagesGetLanguages, token, nil, *bytes.NewBuffer([]byte(text)), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get languages: %v", err.Error())
	}

	var s []languages.Language
	err = json.Unmarshal(b, &s)
	if err != nil {
		return nil, fmt.Errorf("failed to read response from get languages: %v", err.Error())
	}

	return s, nil
}
