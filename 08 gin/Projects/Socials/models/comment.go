package models

type Comment struct {
    ID      int    `json:"id"`
    Text    string `json:"text"`
    PostID  int    `json:"post_id"`
    UserID  int    `json:"user_id"`
}


