yarn run v1.22.4
$ /Users/iinoyouhei/go/src/github.com/wheatandcat/memoir-backend/node_modules/.bin/graphql-markdown http://localhost:8080/query
# Schema Types

<details>
  <summary><strong>Table of Contents</strong></summary>

  * [Query](#query)
  * [Mutation](#mutation)
  * [Objects](#objects)
    * [Item](#item)
    * [User](#user)
  * [Inputs](#inputs)
    * [DeleteItem](#deleteitem)
    * [NewItem](#newitem)
    * [NewUser](#newuser)
    * [UpdateItem](#updateitem)
  * [Scalars](#scalars)
    * [Boolean](#boolean)
    * [Float](#float)
    * [ID](#id)
    * [Int](#int)
    * [String](#string)
    * [Time](#time)

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
</tbody>
</table>

## Objects

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

アイテムID

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

カテゴリーID

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

ユーザーID

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

アイテムID

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

カテゴリーID

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

ユーザーID

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

アイテムID

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

カテゴリーID

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

Done in 2.63s.
