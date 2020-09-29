## GET

##### GET: `/v1/todos`

全てのtodo情報を取得する．

```json
[
	{
		"id": 1,
		"content": "砂糖を買う",
		"completed": false
	},
	{
		"id": 2,
		"content": "課題を終わらせる",
		"completed": false
	},
	{
		"id": 3,
		"content": "歯医者に行く",
		"completed": false
	}
]
```

##### GET: `/v1/todos/${id}`

任意のtodo情報を取得する．

```json
{
	"id": 1,
	"content": "砂糖を買う",
	"completed": false
}
```

## POST

##### POST: `/v1/todos`

todo情報を送信する．

```json
{
	"id": 1,
	"content": "hoge",
	"completed": false
}
```

## PATCH

##### PATCH: `/v1/todos/${id}/toggle-completed`

任意のtodoの完了・未完了をtoggleする．

```json
{
	"completed": !completed //書き方がよくわからなかった
}
```

## DELETE

##### DELETE `/v1/todos/${id}`

任意のtodoを削除する