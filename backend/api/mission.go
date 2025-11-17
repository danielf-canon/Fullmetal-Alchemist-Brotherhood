package api

type MissionRequestDto struct {
    Title       string `json:"title"`
    Description string `json:"description"`
    AssignedTo  uint   `json:"assigned_to"` 
}

type MissionResponseDto struct {
    ID          uint                 `json:"mission_id"`
    Title       string               `json:"title"`
    Description string               `json:"description"`
    Status      string               `json:"status"`
    AssignedTo  uint                 `json:"assigned_to"`
    Alchemist   *AlquimistaResponseDto `json:"alchemist,omitempty"`
    CreatedAt   string               `json:"created_at"`
}
