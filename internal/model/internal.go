package model

type ResponseChannel struct {
	Value interface{}
	Err   error
}

type SetQuery struct {
	Set interface{} `json:"set" validate:"required"`
}

// Commodity Query
type DeleteCommodityCondition struct {
	UUID string `json:"uuid" validate:"required"`
}
type DeleteCommodityQuery struct {
	Condition *DeleteCommodityCondition `json:"condition" validate:"required,dive"`
}
type UpdateCommodityCondition struct {
	UUID string `json:"uuid" validate:"required"`
}
type UpdateCommodityQuery struct {
	Condition *UpdateCommodityCondition `json:"condition" validate:"required,dive"`
	*SetQuery
}

// Size Option Query
type DeleteSizeOptCondition struct {
	Size string `json:"size" validate:"required"`
}
type DeleteSizeOptQuery struct {
	Condition *DeleteSizeOptCondition `json:"condition" validate:"required,dive"`
}
type UpdateSizeOptCondition struct {
	Size string `json:"size" validate:"required"`
}
type UpdateSizeOptQuery struct {
	Condition *UpdateSizeOptCondition `json:"condition" validate:"required,dive"`
	*SetQuery
}

// Area Option Query
type DeleteAreaOptCondition struct {
	Province string `json:"province" validate:"required"`
	City     string `json:"city" validate:"required"`
}
type DeleteAreaOptQuery struct {
	Condition *DeleteAreaOptCondition `json:"condition" validate:"required,dive"`
}
type UpdateAreaOptCondition struct {
	Province string `json:"province" validate:"required"`
	City     string `json:"city" validate:"required"`
}
type UpdateAreaOptQuery struct {
	Condition *UpdateAreaOptCondition `json:"condition" validate:"required,dive"`
	*SetQuery
}
