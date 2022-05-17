package ionic

import (
	"bytes"
	"encoding/json"
	"fmt"
)

const GraphQLQueryEndpoint = "v1/query"

// GraphQLQuery sends an arbitrary GraphQL query to the API, returning the result.
func (ic *IonClient) GraphQLQuery(query string, token string) (json.RawMessage, error) {
	body, err := json.Marshal(query)
	if err != nil {
		return nil, fmt.Errorf("GraphQLQuery: failed to marshal query: %v", err.Error())
	}

	result, err := ic.Post(GraphQLQueryEndpoint, token, nil, *bytes.NewBuffer(body), nil)
	if err != nil {
		return nil, fmt.Errorf("graphql query failed: %v", err.Error())
	}

	return result, nil
}
