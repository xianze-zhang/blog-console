package service

type CountTagRequest struct {
	Name  string `json:"name" binding:"max=100"`
	State uint8  `json:"state,default=1" binding:"oneof=0 1"`
}

type TagListRequest struct {
	Name  string `json:"name" binding:"max=100"`
	State uint8  `json:"state,default=1" binding:"oneof=0 1"`
}

type CreateTagRequest struct {
	Name      string `json:"name" binding:"required,min=3,max=100"`
	CreatedBy string `json:"created_by" binding:"required,min=3,max=100"`
	State     uint8  `json:"state,default=1" binding:"oneof=0 1"`
}

type UpdateTagRequest struct {
	ID         uint32 `json:"id" binding:"required,gte=1"`
	Name       string `json:"name" binding:"min=3,max=100"`
	State      uint8  `json:"state" binding:"required,oneof=0 1"`
	ModifiedBy string `json:"modified_by" binding:"required,min=3,max=100"`
}

type DeleteTagRequest struct {
	ID uint32 `json:"id" binding:"required,gte=1"`
}
