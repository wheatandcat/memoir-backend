// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"time"
)

type DeleteItem struct {
	// アイテムID
	ID string `json:"id"`
}

type InputItemsInPeriod struct {
	After     *string   `json:"after"`
	First     int       `json:"first"`
	StartDate time.Time `json:"startDate"`
	EndDate   time.Time `json:"endDate"`
}

type InputRelationshipRequests struct {
	After *string `json:"after"`
	First int     `json:"first"`
}

type InputRelationships struct {
	After *string `json:"after"`
	First int     `json:"first"`
}

type Invite struct {
	// ユーザーID
	UserID string `json:"userID"`
	// 招待コード
	Code string `json:"code"`
	// 作成日時
	CreatedAt time.Time `json:"createdAt"`
	// 更新日時
	UpdatedAt time.Time `json:"updatedAt"`
}

type Item struct {
	// アイテムID
	ID string `json:"id"`
	// ユーザーID
	UserID string `json:"userID"`
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

type ItemsInPeriod struct {
	PageInfo *PageInfo            `json:"pageInfo"`
	Edges    []*ItemsInPeriodEdge `json:"edges"`
}

type ItemsInPeriodEdge struct {
	Node   *Item  `json:"node"`
	Cursor string `json:"cursor"`
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

type NewPushToken struct {
	// Push通知トークン
	Token string `json:"token"`
	// デバイスID
	DeviceID string `json:"deviceID"`
}

type NewRelationshipRequest struct {
	// 招待コード
	Code string `json:"code"`
}

type NewUser struct {
	// ユーザーID
	ID string `json:"id"`
}

type PageInfo struct {
	EndCursor   string `json:"endCursor"`
	HasNextPage bool   `json:"hasNextPage"`
}

type PushToken struct {
	// ユーザーID
	UserID string `json:"userID"`
	// Push通知トークン
	Token string `json:"token"`
	// デバイスID
	DeviceID string `json:"deviceID"`
	// 作成日時
	CreatedAt time.Time `json:"createdAt"`
	// 更新日時
	UpdatedAt time.Time `json:"updatedAt"`
}

type Relationship struct {
	// ID
	ID string `json:"id"`
	// フォローしたユーザーID
	FollowerID string `json:"followerId"`
	// フォローされたユーザーID
	FollowedID string `json:"followedId"`
	// 作成日時
	CreatedAt time.Time `json:"createdAt"`
	// 更新日時
	UpdatedAt time.Time `json:"updatedAt"`
	// ユーザー情報
	User *User `json:"user"`
}

type RelationshipEdge struct {
	Node   *Relationship `json:"node"`
	Cursor string        `json:"cursor"`
}

type RelationshipRequest struct {
	// ID
	ID string `json:"id"`
	// フォローしたユーザーID
	FollowerID string `json:"followerId"`
	// フォローされたユーザーID
	FollowedID string `json:"followedId"`
	// 1:申請中、2:拒否、3: 承認
	Status int `json:"status"`
	// 作成日時
	CreatedAt time.Time `json:"createdAt"`
	// 更新日時
	UpdatedAt time.Time `json:"updatedAt"`
	// ユーザー情報
	User *User `json:"user"`
}

type RelationshipRequestEdge struct {
	Node   *RelationshipRequest `json:"node"`
	Cursor string               `json:"cursor"`
}

type RelationshipRequests struct {
	PageInfo *PageInfo                  `json:"pageInfo"`
	Edges    []*RelationshipRequestEdge `json:"edges"`
}

type Relationships struct {
	PageInfo *PageInfo           `json:"pageInfo"`
	Edges    []*RelationshipEdge `json:"edges"`
}

type UpdateItem struct {
	// アイテムID
	ID string `json:"id"`
	// タイトル
	Title *string `json:"title"`
	// カテゴリーID
	CategoryID *int `json:"categoryID"`
	// 日付
	Date    *time.Time `json:"date"`
	Like    *bool      `json:"like"`
	Dislike *bool      `json:"dislike"`
}

type UpdateUser struct {
	// 表示名
	DisplayName string `json:"displayName"`
	// 画像URL
	Image string `json:"image"`
}

type User struct {
	// ユーザーID
	ID string `json:"id"`
	// 表示名
	DisplayName string `json:"displayName"`
	// 画像URL
	Image string `json:"image"`
	// 作成日時
	CreatedAt time.Time `json:"createdAt"`
	// 更新日時
	UpdatedAt time.Time `json:"updatedAt"`
}
