// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"time"
)

type Item struct {
	// アイテムID
	ID string `json:"id"`
	// タイトル
	Title string `json:"title"`
	// カテゴリーID
	CategoryID int `json:"categoryID"`
	// 日付
	Date    time.Time `json:"date"`
	Like    bool      `json:"like"`
	Dislike bool      `json:"dislike"`
	// 作成日時
	CreatedAt time.Time `json:"createdAt"`
	// 更新日時
	UpdatedAt time.Time `json:"updatedAt"`
}

type NewItem struct {
	// タイトル
	Title string `json:"title"`
	// カテゴリーID
	CategoryID int `json:"categoryID"`
	// 日付
	Date    time.Time `json:"date"`
	Like    bool      `json:"like"`
	Dislike bool      `json:"dislike"`
}

type NewUser struct {
	// ユーザーID
	ID string `json:"id"`
}

type User struct {
	// ユーザーID
	ID string `json:"id"`
	// 作成日時
	CreatedAt time.Time `json:"createdAt"`
	// 更新日時
	UpdatedAt time.Time `json:"updatedAt"`
}
