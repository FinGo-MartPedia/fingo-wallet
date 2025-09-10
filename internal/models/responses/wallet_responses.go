package responses

type WalletTransactionResponse struct {
	Balance   float64 `json:"balance"`
	Type      string  `json:"type"`
	Reference string  `json:"reference"`
}

type BalanceResponse struct {
	Balance float64 `json:"balance"`
}

type PaginatedResult[T any] struct {
	Items      []T `json:"items"`
	Total      int `json:"total"`
	Page       int `json:"page"`
	PageSize   int `json:"page_size"`
	TotalPages int `json:"total_pages"`
}
