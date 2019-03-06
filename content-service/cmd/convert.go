package main

func GetContentFromItems(items []item) []userContentResponse {
	content := []userContentResponse{}
	for _, item := range items {
		status := calculateStatus(item)
		c := userContentResponse{
			ID:     item.ID,
			Name:   item.Name,
			Status: status,
		}
		content = append(content, c)
	}

	return content
}

func calculateStatus(item item) int {
	return 2 // TODO: Calculate it later.
}
