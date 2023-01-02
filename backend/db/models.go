package db

type Board struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
}

type Card struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	Status      string `json:"status" binding:"required"`
}

type Workspace struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
}

type WorkspaceStruct struct {
	Workspace
	WorkspaceId int32 `json:"workspaceId" binding:"required"`
	User        int32 `json:"user" binding:"required"`
}

type BoardStruct struct {
	Board
	BoardId   int32 `json:"boardId" binding:"required"`
	Workspace int32 `json:"workspace" binding:"required"`
}

type CardStruct struct {
	Card
	CardId int32 `json:"cardId" binding:"required"`
	Board  int32 `json:"board" binding:"required"`
}
