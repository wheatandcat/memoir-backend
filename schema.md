yarn run v1.22.10
$ /Users/iinoyouhei/go/src/github.com/wheatandcat/memoir-backend/node_modules/.bin/graphql-markdown http://localhost:8080/query

# Schema Types

<details>
  <summary><strong>Table of Contents</strong></summary>

- [Query](#query)
- [Mutation](#mutation)
- [Objects](#objects)
  - [AuthUser](#authuser)
  - [ExistAuthUser](#existauthuser)
  - [Invite](#invite)
  - [Item](#item)
  - [ItemsInPeriod](#itemsinperiod)
  - [ItemsInPeriodEdge](#itemsinperiodedge)
  - [PageInfo](#pageinfo)
  - [PushToken](#pushtoken)
  - [Relationship](#relationship)
  - [RelationshipEdge](#relationshipedge)
  - [RelationshipRequest](#relationshiprequest)
  - [RelationshipRequestEdge](#relationshiprequestedge)
  - [RelationshipRequests](#relationshiprequests)
  - [Relationships](#relationships)
  - [User](#user)
- [Inputs](#inputs)
  - [DeleteItem](#deleteitem)
  - [InputItemsInPeriod](#inputitemsinperiod)
  - [InputRelationshipRequests](#inputrelationshiprequests)
  - [InputRelationships](#inputrelationships)
  - [NewAuthUser](#newauthuser)
  - [NewItem](#newitem)
  - [NewPushToken](#newpushtoken)
  - [NewRelationshipRequest](#newrelationshiprequest)
  - [NewUser](#newuser)
  - [UpdateItem](#updateitem)
  - [UpdateUser](#updateuser)
- [Scalars](#scalars)
  - [Boolean](#boolean)
  - [Float](#float)
  - [ID](#id)
  - [Int](#int)
  - [String](#string)
  - [Time](#time)

</details>

## Query

<table>
<thead>
<tr>
<th align="left">Field</th>
<th align="right">Argument</th>
<th align="left">Type</th>
<th align="left">Description</th>
</tr>
</thead>
<tbody>
<tr>
<td colspan="2" valign="top"><strong>user</strong></td>
<td valign="top"><a href="#user">User</a>!</td>
<td>

ユーザーを取得する

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>existAuthUser</strong></td>
<td valign="top"><a href="#existauthuser">ExistAuthUser</a>!</td>
<td>

認証ユーザーが存在するか判定する

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>item</strong></td>
<td valign="top"><a href="#item">Item</a></td>
<td>

アイテムを取得する

</td>
</tr>
<tr>
<td colspan="2" align="right" valign="top">id</td>
<td valign="top"><a href="#id">ID</a>!</td>
<td></td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>itemsByDate</strong></td>
<td valign="top">[<a href="#item">Item</a>]</td>
<td>

アイテムを日付で取得する

</td>
</tr>
<tr>
<td colspan="2" align="right" valign="top">date</td>
<td valign="top"><a href="#time">Time</a>!</td>
<td></td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>itemsInDate</strong></td>
<td valign="top">[<a href="#item">Item</a>]</td>
<td>

アイテムを日付で取得する

</td>
</tr>
<tr>
<td colspan="2" align="right" valign="top">date</td>
<td valign="top"><a href="#time">Time</a>!</td>
<td></td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>itemsInPeriod</strong></td>
<td valign="top"><a href="#itemsinperiod">ItemsInPeriod</a>!</td>
<td>

期間でアイテムを取得する

</td>
</tr>
<tr>
<td colspan="2" align="right" valign="top">input</td>
<td valign="top"><a href="#inputitemsinperiod">InputItemsInPeriod</a>!</td>
<td></td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>invite</strong></td>
<td valign="top"><a href="#invite">Invite</a>!</td>
<td>

招待コードを取得する

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>inviteByCode</strong></td>
<td valign="top"><a href="#user">User</a>!</td>
<td>

招待コードからユーザーを取得する

</td>
</tr>
<tr>
<td colspan="2" align="right" valign="top">code</td>
<td valign="top"><a href="#string">String</a>!</td>
<td></td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>relationshipRequests</strong></td>
<td valign="top"><a href="#relationshiprequests">RelationshipRequests</a>!</td>
<td>

招待の申請リクエストを取得する

</td>
</tr>
<tr>
<td colspan="2" align="right" valign="top">input</td>
<td valign="top"><a href="#inputrelationshiprequests">InputRelationshipRequests</a>!</td>
<td></td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>relationships</strong></td>
<td valign="top"><a href="#relationships">Relationships</a>!</td>
<td>

共有ユーザーを取得する

</td>
</tr>
<tr>
<td colspan="2" align="right" valign="top">input</td>
<td valign="top"><a href="#inputrelationships">InputRelationships</a>!</td>
<td></td>
</tr>
</tbody>
</table>

## Mutation

<table>
<thead>
<tr>
<th align="left">Field</th>
<th align="right">Argument</th>
<th align="left">Type</th>
<th align="left">Description</th>
</tr>
</thead>
<tbody>
<tr>
<td colspan="2" valign="top"><strong>createUser</strong></td>
<td valign="top"><a href="#user">User</a>!</td>
<td>

ユーザーを作成する

</td>
</tr>
<tr>
<td colspan="2" align="right" valign="top">input</td>
<td valign="top"><a href="#newuser">NewUser</a>!</td>
<td></td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>createAuthUser</strong></td>
<td valign="top"><a href="#authuser">AuthUser</a>!</td>
<td>

認証ユーザーを作成する

</td>
</tr>
<tr>
<td colspan="2" align="right" valign="top">input</td>
<td valign="top"><a href="#newauthuser">NewAuthUser</a>!</td>
<td></td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>updateUser</strong></td>
<td valign="top"><a href="#user">User</a>!</td>
<td>

ユーザーを更新する

</td>
</tr>
<tr>
<td colspan="2" align="right" valign="top">input</td>
<td valign="top"><a href="#updateuser">UpdateUser</a>!</td>
<td></td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>deleteUser</strong></td>
<td valign="top"><a href="#user">User</a>!</td>
<td>

ユーザーを削除する

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>createItem</strong></td>
<td valign="top"><a href="#item">Item</a>!</td>
<td>

アイテムを作成する

</td>
</tr>
<tr>
<td colspan="2" align="right" valign="top">input</td>
<td valign="top"><a href="#newitem">NewItem</a>!</td>
<td></td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>updateItem</strong></td>
<td valign="top"><a href="#item">Item</a>!</td>
<td>

アイテムを更新する

</td>
</tr>
<tr>
<td colspan="2" align="right" valign="top">input</td>
<td valign="top"><a href="#updateitem">UpdateItem</a>!</td>
<td></td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>deleteItem</strong></td>
<td valign="top"><a href="#item">Item</a>!</td>
<td>

アイテムを削除する

</td>
</tr>
<tr>
<td colspan="2" align="right" valign="top">input</td>
<td valign="top"><a href="#deleteitem">DeleteItem</a>!</td>
<td></td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>createInvite</strong></td>
<td valign="top"><a href="#invite">Invite</a>!</td>
<td>

招待コード作成する

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>updateInvite</strong></td>
<td valign="top"><a href="#invite">Invite</a>!</td>
<td>

招待コード更新する

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>createRelationshipRequest</strong></td>
<td valign="top"><a href="#relationshiprequest">RelationshipRequest</a>!</td>
<td>

招待をリクエストする

</td>
</tr>
<tr>
<td colspan="2" align="right" valign="top">input</td>
<td valign="top"><a href="#newrelationshiprequest">NewRelationshipRequest</a>!</td>
<td></td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>acceptRelationshipRequest</strong></td>
<td valign="top"><a href="#relationshiprequest">RelationshipRequest</a>!</td>
<td>

招待リクエストを承諾する

</td>
</tr>
<tr>
<td colspan="2" align="right" valign="top">followedID</td>
<td valign="top"><a href="#string">String</a>!</td>
<td></td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>ngRelationshipRequest</strong></td>
<td valign="top"><a href="#relationshiprequest">RelationshipRequest</a>!</td>
<td>

招待リクエストを拒否する

</td>
</tr>
<tr>
<td colspan="2" align="right" valign="top">followedID</td>
<td valign="top"><a href="#string">String</a>!</td>
<td></td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>deleteRelationship</strong></td>
<td valign="top"><a href="#relationship">Relationship</a>!</td>
<td>

共有メンバーを解除する

</td>
</tr>
<tr>
<td colspan="2" align="right" valign="top">followedID</td>
<td valign="top"><a href="#string">String</a>!</td>
<td></td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>createPushToken</strong></td>
<td valign="top"><a href="#pushtoken">PushToken</a>!</td>
<td>

Push 通知のトークンを作成する

</td>
</tr>
<tr>
<td colspan="2" align="right" valign="top">input</td>
<td valign="top"><a href="#newpushtoken">NewPushToken</a>!</td>
<td></td>
</tr>
</tbody>
</table>

## Objects

### AuthUser

<table>
<thead>
<tr>
<th align="left">Field</th>
<th align="right">Argument</th>
<th align="left">Type</th>
<th align="left">Description</th>
</tr>
</thead>
<tbody>
<tr>
<td colspan="2" valign="top"><strong>id</strong></td>
<td valign="top"><a href="#id">ID</a>!</td>
<td>

ユーザー ID

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>displayName</strong></td>
<td valign="top"><a href="#string">String</a>!</td>
<td>

表示名

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>image</strong></td>
<td valign="top"><a href="#string">String</a>!</td>
<td>

画像 URL

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>new</strong></td>
<td valign="top"><a href="#boolean">Boolean</a>!</td>
<td>

新規作成

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>createdAt</strong></td>
<td valign="top"><a href="#time">Time</a>!</td>
<td>

作成日時

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>updatedAt</strong></td>
<td valign="top"><a href="#time">Time</a>!</td>
<td>

更新日時

</td>
</tr>
</tbody>
</table>

### ExistAuthUser

<table>
<thead>
<tr>
<th align="left">Field</th>
<th align="right">Argument</th>
<th align="left">Type</th>
<th align="left">Description</th>
</tr>
</thead>
<tbody>
<tr>
<td colspan="2" valign="top"><strong>exist</strong></td>
<td valign="top"><a href="#boolean">Boolean</a>!</td>
<td></td>
</tr>
</tbody>
</table>

### Invite

<table>
<thead>
<tr>
<th align="left">Field</th>
<th align="right">Argument</th>
<th align="left">Type</th>
<th align="left">Description</th>
</tr>
</thead>
<tbody>
<tr>
<td colspan="2" valign="top"><strong>userID</strong></td>
<td valign="top"><a href="#id">ID</a>!</td>
<td>

ユーザー ID

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>code</strong></td>
<td valign="top"><a href="#string">String</a>!</td>
<td>

招待コード

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>createdAt</strong></td>
<td valign="top"><a href="#time">Time</a>!</td>
<td>

作成日時

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>updatedAt</strong></td>
<td valign="top"><a href="#time">Time</a>!</td>
<td>

更新日時

</td>
</tr>
</tbody>
</table>

### Item

<table>
<thead>
<tr>
<th align="left">Field</th>
<th align="right">Argument</th>
<th align="left">Type</th>
<th align="left">Description</th>
</tr>
</thead>
<tbody>
<tr>
<td colspan="2" valign="top"><strong>id</strong></td>
<td valign="top"><a href="#id">ID</a>!</td>
<td>

アイテム ID

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>userID</strong></td>
<td valign="top"><a href="#string">String</a>!</td>
<td>

ユーザー ID

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>title</strong></td>
<td valign="top"><a href="#string">String</a>!</td>
<td>

タイトル

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>categoryID</strong></td>
<td valign="top"><a href="#int">Int</a>!</td>
<td>

カテゴリー ID

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>date</strong></td>
<td valign="top"><a href="#time">Time</a>!</td>
<td>

日付

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>like</strong></td>
<td valign="top"><a href="#boolean">Boolean</a>!</td>
<td></td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>dislike</strong></td>
<td valign="top"><a href="#boolean">Boolean</a>!</td>
<td></td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>createdAt</strong></td>
<td valign="top"><a href="#time">Time</a>!</td>
<td>

作成日時

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>updatedAt</strong></td>
<td valign="top"><a href="#time">Time</a>!</td>
<td>

更新日時

</td>
</tr>
</tbody>
</table>

### ItemsInPeriod

<table>
<thead>
<tr>
<th align="left">Field</th>
<th align="right">Argument</th>
<th align="left">Type</th>
<th align="left">Description</th>
</tr>
</thead>
<tbody>
<tr>
<td colspan="2" valign="top"><strong>pageInfo</strong></td>
<td valign="top"><a href="#pageinfo">PageInfo</a>!</td>
<td></td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>edges</strong></td>
<td valign="top">[<a href="#itemsinperiodedge">ItemsInPeriodEdge</a>!]!</td>
<td></td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>totalCount</strong></td>
<td valign="top"><a href="#int">Int</a>!</td>
<td></td>
</tr>
</tbody>
</table>

### ItemsInPeriodEdge

<table>
<thead>
<tr>
<th align="left">Field</th>
<th align="right">Argument</th>
<th align="left">Type</th>
<th align="left">Description</th>
</tr>
</thead>
<tbody>
<tr>
<td colspan="2" valign="top"><strong>node</strong></td>
<td valign="top"><a href="#item">Item</a></td>
<td></td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>cursor</strong></td>
<td valign="top"><a href="#string">String</a>!</td>
<td></td>
</tr>
</tbody>
</table>

### PageInfo

<table>
<thead>
<tr>
<th align="left">Field</th>
<th align="right">Argument</th>
<th align="left">Type</th>
<th align="left">Description</th>
</tr>
</thead>
<tbody>
<tr>
<td colspan="2" valign="top"><strong>endCursor</strong></td>
<td valign="top"><a href="#string">String</a>!</td>
<td></td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>hasNextPage</strong></td>
<td valign="top"><a href="#boolean">Boolean</a>!</td>
<td></td>
</tr>
</tbody>
</table>

### PushToken

<table>
<thead>
<tr>
<th align="left">Field</th>
<th align="right">Argument</th>
<th align="left">Type</th>
<th align="left">Description</th>
</tr>
</thead>
<tbody>
<tr>
<td colspan="2" valign="top"><strong>userID</strong></td>
<td valign="top"><a href="#id">ID</a>!</td>
<td>

ユーザー ID

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>token</strong></td>
<td valign="top"><a href="#string">String</a>!</td>
<td>

Push 通知トークン

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>deviceID</strong></td>
<td valign="top"><a href="#string">String</a>!</td>
<td>

デバイス ID

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>createdAt</strong></td>
<td valign="top"><a href="#time">Time</a>!</td>
<td>

作成日時

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>updatedAt</strong></td>
<td valign="top"><a href="#time">Time</a>!</td>
<td>

更新日時

</td>
</tr>
</tbody>
</table>

### Relationship

<table>
<thead>
<tr>
<th align="left">Field</th>
<th align="right">Argument</th>
<th align="left">Type</th>
<th align="left">Description</th>
</tr>
</thead>
<tbody>
<tr>
<td colspan="2" valign="top"><strong>id</strong></td>
<td valign="top"><a href="#id">ID</a>!</td>
<td>

ID

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>followerId</strong></td>
<td valign="top"><a href="#string">String</a>!</td>
<td>

フォローしたユーザー ID

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>followedId</strong></td>
<td valign="top"><a href="#string">String</a>!</td>
<td>

フォローされたユーザー ID

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>createdAt</strong></td>
<td valign="top"><a href="#time">Time</a>!</td>
<td>

作成日時

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>updatedAt</strong></td>
<td valign="top"><a href="#time">Time</a>!</td>
<td>

更新日時

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>user</strong></td>
<td valign="top"><a href="#user">User</a></td>
<td>

ユーザー情報

</td>
</tr>
<tr>
<td colspan="2" align="right" valign="top">skip</td>
<td valign="top"><a href="#boolean">Boolean</a></td>
<td></td>
</tr>
</tbody>
</table>

### RelationshipEdge

<table>
<thead>
<tr>
<th align="left">Field</th>
<th align="right">Argument</th>
<th align="left">Type</th>
<th align="left">Description</th>
</tr>
</thead>
<tbody>
<tr>
<td colspan="2" valign="top"><strong>node</strong></td>
<td valign="top"><a href="#relationship">Relationship</a></td>
<td></td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>cursor</strong></td>
<td valign="top"><a href="#string">String</a>!</td>
<td></td>
</tr>
</tbody>
</table>

### RelationshipRequest

<table>
<thead>
<tr>
<th align="left">Field</th>
<th align="right">Argument</th>
<th align="left">Type</th>
<th align="left">Description</th>
</tr>
</thead>
<tbody>
<tr>
<td colspan="2" valign="top"><strong>id</strong></td>
<td valign="top"><a href="#id">ID</a>!</td>
<td>

ID

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>followerId</strong></td>
<td valign="top"><a href="#string">String</a>!</td>
<td>

フォローしたユーザー ID

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>followedId</strong></td>
<td valign="top"><a href="#string">String</a>!</td>
<td>

フォローされたユーザー ID

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>status</strong></td>
<td valign="top"><a href="#int">Int</a>!</td>
<td>

1:申請中、2:拒否、3: 承認

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>createdAt</strong></td>
<td valign="top"><a href="#time">Time</a>!</td>
<td>

作成日時

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>updatedAt</strong></td>
<td valign="top"><a href="#time">Time</a>!</td>
<td>

更新日時

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>user</strong></td>
<td valign="top"><a href="#user">User</a></td>
<td>

ユーザー情報

</td>
</tr>
<tr>
<td colspan="2" align="right" valign="top">skip</td>
<td valign="top"><a href="#boolean">Boolean</a></td>
<td></td>
</tr>
</tbody>
</table>

### RelationshipRequestEdge

<table>
<thead>
<tr>
<th align="left">Field</th>
<th align="right">Argument</th>
<th align="left">Type</th>
<th align="left">Description</th>
</tr>
</thead>
<tbody>
<tr>
<td colspan="2" valign="top"><strong>node</strong></td>
<td valign="top"><a href="#relationshiprequest">RelationshipRequest</a></td>
<td></td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>cursor</strong></td>
<td valign="top"><a href="#string">String</a>!</td>
<td></td>
</tr>
</tbody>
</table>

### RelationshipRequests

<table>
<thead>
<tr>
<th align="left">Field</th>
<th align="right">Argument</th>
<th align="left">Type</th>
<th align="left">Description</th>
</tr>
</thead>
<tbody>
<tr>
<td colspan="2" valign="top"><strong>pageInfo</strong></td>
<td valign="top"><a href="#pageinfo">PageInfo</a>!</td>
<td></td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>edges</strong></td>
<td valign="top">[<a href="#relationshiprequestedge">RelationshipRequestEdge</a>!]!</td>
<td></td>
</tr>
</tbody>
</table>

### Relationships

<table>
<thead>
<tr>
<th align="left">Field</th>
<th align="right">Argument</th>
<th align="left">Type</th>
<th align="left">Description</th>
</tr>
</thead>
<tbody>
<tr>
<td colspan="2" valign="top"><strong>pageInfo</strong></td>
<td valign="top"><a href="#pageinfo">PageInfo</a>!</td>
<td></td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>edges</strong></td>
<td valign="top">[<a href="#relationshipedge">RelationshipEdge</a>!]!</td>
<td></td>
</tr>
</tbody>
</table>

### User

<table>
<thead>
<tr>
<th align="left">Field</th>
<th align="right">Argument</th>
<th align="left">Type</th>
<th align="left">Description</th>
</tr>
</thead>
<tbody>
<tr>
<td colspan="2" valign="top"><strong>id</strong></td>
<td valign="top"><a href="#id">ID</a>!</td>
<td>

ユーザー ID

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>displayName</strong></td>
<td valign="top"><a href="#string">String</a>!</td>
<td>

表示名

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>image</strong></td>
<td valign="top"><a href="#string">String</a>!</td>
<td>

画像 URL

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>createdAt</strong></td>
<td valign="top"><a href="#time">Time</a>!</td>
<td>

作成日時

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>updatedAt</strong></td>
<td valign="top"><a href="#time">Time</a>!</td>
<td>

更新日時

</td>
</tr>
</tbody>
</table>

## Inputs

### DeleteItem

<table>
<thead>
<tr>
<th colspan="2" align="left">Field</th>
<th align="left">Type</th>
<th align="left">Description</th>
</tr>
</thead>
<tbody>
<tr>
<td colspan="2" valign="top"><strong>id</strong></td>
<td valign="top"><a href="#id">ID</a>!</td>
<td>

アイテム ID

</td>
</tr>
</tbody>
</table>

### InputItemsInPeriod

<table>
<thead>
<tr>
<th colspan="2" align="left">Field</th>
<th align="left">Type</th>
<th align="left">Description</th>
</tr>
</thead>
<tbody>
<tr>
<td colspan="2" valign="top"><strong>after</strong></td>
<td valign="top"><a href="#string">String</a></td>
<td></td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>first</strong></td>
<td valign="top"><a href="#int">Int</a>!</td>
<td></td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>startDate</strong></td>
<td valign="top"><a href="#time">Time</a>!</td>
<td>

開始日

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>endDate</strong></td>
<td valign="top"><a href="#time">Time</a>!</td>
<td>

終了日

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>userIDList</strong></td>
<td valign="top">[<a href="#string">String</a>]</td>
<td>

ユーザー ID

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>categoryID</strong></td>
<td valign="top"><a href="#int">Int</a></td>
<td>

カテゴリー ID

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>like</strong></td>
<td valign="top"><a href="#boolean">Boolean</a></td>
<td>

Good

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>dislike</strong></td>
<td valign="top"><a href="#boolean">Boolean</a></td>
<td>

Bad

</td>
</tr>
</tbody>
</table>

### InputRelationshipRequests

<table>
<thead>
<tr>
<th colspan="2" align="left">Field</th>
<th align="left">Type</th>
<th align="left">Description</th>
</tr>
</thead>
<tbody>
<tr>
<td colspan="2" valign="top"><strong>after</strong></td>
<td valign="top"><a href="#string">String</a></td>
<td></td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>first</strong></td>
<td valign="top"><a href="#int">Int</a>!</td>
<td></td>
</tr>
</tbody>
</table>

### InputRelationships

<table>
<thead>
<tr>
<th colspan="2" align="left">Field</th>
<th align="left">Type</th>
<th align="left">Description</th>
</tr>
</thead>
<tbody>
<tr>
<td colspan="2" valign="top"><strong>after</strong></td>
<td valign="top"><a href="#string">String</a></td>
<td></td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>first</strong></td>
<td valign="top"><a href="#int">Int</a>!</td>
<td></td>
</tr>
</tbody>
</table>

### NewAuthUser

<table>
<thead>
<tr>
<th colspan="2" align="left">Field</th>
<th align="left">Type</th>
<th align="left">Description</th>
</tr>
</thead>
<tbody>
<tr>
<td colspan="2" valign="top"><strong>id</strong></td>
<td valign="top"><a href="#id">ID</a>!</td>
<td>

ユーザー ID

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>isNewUser</strong></td>
<td valign="top"><a href="#boolean">Boolean</a>!</td>
<td>

true: ユーザー作成を行う

</td>
</tr>
</tbody>
</table>

### NewItem

<table>
<thead>
<tr>
<th colspan="2" align="left">Field</th>
<th align="left">Type</th>
<th align="left">Description</th>
</tr>
</thead>
<tbody>
<tr>
<td colspan="2" valign="top"><strong>title</strong></td>
<td valign="top"><a href="#string">String</a>!</td>
<td>

タイトル

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>categoryID</strong></td>
<td valign="top"><a href="#int">Int</a>!</td>
<td>

カテゴリー ID

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>date</strong></td>
<td valign="top"><a href="#time">Time</a>!</td>
<td>

日付

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>like</strong></td>
<td valign="top"><a href="#boolean">Boolean</a>!</td>
<td></td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>dislike</strong></td>
<td valign="top"><a href="#boolean">Boolean</a>!</td>
<td></td>
</tr>
</tbody>
</table>

### NewPushToken

<table>
<thead>
<tr>
<th colspan="2" align="left">Field</th>
<th align="left">Type</th>
<th align="left">Description</th>
</tr>
</thead>
<tbody>
<tr>
<td colspan="2" valign="top"><strong>token</strong></td>
<td valign="top"><a href="#string">String</a>!</td>
<td>

Push 通知トークン

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>deviceID</strong></td>
<td valign="top"><a href="#string">String</a>!</td>
<td>

デバイス ID

</td>
</tr>
</tbody>
</table>

### NewRelationshipRequest

<table>
<thead>
<tr>
<th colspan="2" align="left">Field</th>
<th align="left">Type</th>
<th align="left">Description</th>
</tr>
</thead>
<tbody>
<tr>
<td colspan="2" valign="top"><strong>code</strong></td>
<td valign="top"><a href="#string">String</a>!</td>
<td>

招待コード

</td>
</tr>
</tbody>
</table>

### NewUser

<table>
<thead>
<tr>
<th colspan="2" align="left">Field</th>
<th align="left">Type</th>
<th align="left">Description</th>
</tr>
</thead>
<tbody>
<tr>
<td colspan="2" valign="top"><strong>id</strong></td>
<td valign="top"><a href="#id">ID</a>!</td>
<td>

ユーザー ID

</td>
</tr>
</tbody>
</table>

### UpdateItem

<table>
<thead>
<tr>
<th colspan="2" align="left">Field</th>
<th align="left">Type</th>
<th align="left">Description</th>
</tr>
</thead>
<tbody>
<tr>
<td colspan="2" valign="top"><strong>id</strong></td>
<td valign="top"><a href="#id">ID</a>!</td>
<td>

アイテム ID

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>title</strong></td>
<td valign="top"><a href="#string">String</a></td>
<td>

タイトル

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>categoryID</strong></td>
<td valign="top"><a href="#int">Int</a></td>
<td>

カテゴリー ID

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>date</strong></td>
<td valign="top"><a href="#time">Time</a></td>
<td>

日付

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>like</strong></td>
<td valign="top"><a href="#boolean">Boolean</a></td>
<td></td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>dislike</strong></td>
<td valign="top"><a href="#boolean">Boolean</a></td>
<td></td>
</tr>
</tbody>
</table>

### UpdateUser

<table>
<thead>
<tr>
<th colspan="2" align="left">Field</th>
<th align="left">Type</th>
<th align="left">Description</th>
</tr>
</thead>
<tbody>
<tr>
<td colspan="2" valign="top"><strong>displayName</strong></td>
<td valign="top"><a href="#string">String</a>!</td>
<td>

表示名

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>image</strong></td>
<td valign="top"><a href="#string">String</a>!</td>
<td>

画像 URL

</td>
</tr>
</tbody>
</table>

## Scalars

### Boolean

The `Boolean` scalar type represents `true` or `false`.

### Float

The `Float` scalar type represents signed double-precision fractional values as specified by [IEEE 754](http://en.wikipedia.org/wiki/IEEE_floating_point).

### ID

The `ID` scalar type represents a unique identifier, often used to refetch an object or as key for a cache. The ID type appears in a JSON response as a String; however, it is not intended to be human-readable. When expected as an input type, any string (such as "4") or integer (such as 4) input value will be accepted as an ID.

### Int

The `Int` scalar type represents non-fractional signed whole numeric values. Int can represent values between -(2^31) and 2^31 - 1.

### String

The `String`scalar type represents textual data, represented as UTF-8 character sequences. The String type is most often used by GraphQL to represent free-form human-readable text.

### Time
