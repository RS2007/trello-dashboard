package db

type WorkspaceStruct struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	WorkspaceId int32  `json:"workspaceId" binding:"required"`
	User        int32  `json:"user" binding:"required"`
}

type BoardStruct struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	BoardId     int32  `json:"boardId" binding:"required"`
	Workspace   int32  `json:"workspace" binding:"required"`
}

type CardStruct struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	Status      string `json:"status" binding:"required"`
	CardId      int32  `json:"cardId" binding:"required"`
	Board       int32  `json:"board" binding:"required"`
}
