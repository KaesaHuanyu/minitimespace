package models

const (
	ORDER_BY_CREATE      = "created_at"
	ORDER_BY_ID          = "id"
	ORDER_BY_CREATE_DESC = "created_at desc"
	ORDER_BY_ID_DESC     = "id desc"

	//status
	ActivityCreated = iota
	ActivityCanceled
	ActivityStarted
	ActivitySuccessed
)
