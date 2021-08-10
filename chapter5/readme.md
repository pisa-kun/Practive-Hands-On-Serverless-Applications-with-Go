## chapter5
#### dynamodb

#### memo
- credentialsfileを使ったセッションとdynamoDBの初期化
```go
sess := session.New(&aws.Config{Region: aws.String("ap-northeast-1")})
	db := dynamodb.New(sess)
	table_name := "Idols"
```

- "github.com/aws/aws-sdk-go-v2/aws/external" の代わりにconfigモジュールを使用する  
https://shogo82148.github.io/blog/2020/10/24/aws-sdk-go-v2-broken/

- goのver互換で使えないモジュールがあるので以下のコードを使う
```go
	accessKey := ""
	secretKey := ""
	//bucket := "test"

	creds := credentials.NewStaticCredentials(accessKey, secretKey, "")
	sess, _ := session.NewSession(&aws.Config{
		Credentials: creds,
		Region:      aws.String("ap-northeast-1")},
	)
	ddb := dynamodb.New(sess)
	tableName := ""
```
[putitemの使い方](https://gammalab.net/blog/8jpanck5k5t9u/s)

- json構造体の数値フィールドは文字列、数字どちらともとれてしまう
> json: cannot unmarshal number into Go struct field

age int 'json:"fieldname,string"'のようにintで取得する

https://qiita.com/kmagai/items/b0bac178e69f59557504

- simpleなscan
https://qiita.com/kenjiskywalker/items/de3c5c0afbc4170ecd3d

- lambdaの環境変数設定

https://docs.aws.amazon.com/ja_jp/lambda/latest/dg/configuration-envvars.html
>aws lambda update-function-configuration --function-name FindAllIdols \ --environment Variables={TABLE_NAME=Idols}

- curlでbodyのjsonを整形表示  
1. jqをdownload https://zenn.dev/unsoluble_sugar/articles/e47b37b04dd1153d5b29
2. windowsはjqをリネームして環境変数に
3. curl -sX GET https://i3w8zvfye4.execute-api.ap-northeast-1.amazonaws.com/staging/idols | .jq

- simpleなgetItem
```go
input := &dynamodb.GetItemInput{
        TableName: aws.String("UserName"),
        Key: map[string]*dynamodb.AttributeValue{
            "UserID": {
                N: aws.String("1"),
            },
            "Date": {
                S: aws.String("20171215"),
            },
        },
    }

    result, err := svc.GetItem(input)
    if err != nil {
        fmt.Println("[GetItem Error]", err)
        return
    }
```
https://qiita.com/sakayuka/items/4af7fead94d589716f4d

- jqをつかってGetItemで返ってきたjsonを整形する  
Nameが欲しいとき
> curl -sX GET https://i3w8zvfye4.execute-api.ap-northeast-1.amazonaws.com/staging/idols/6 | jq ".[]|.Name|.S"

- lambda 環境変数設定
```cmd
### 1行で打ち込むこと
aws lambda update-function-configuration --function-name InsertIdole --environment Variables={TABLE_NAME=Idols}
```

- lambda zipコードのアップデート
> aws lambda update-function-code --function-name InsertIdole --zip-file fileb://./insertupdateitem.zip

- json unmarshal で 文字エラーが起きるケース(curl でpostできない)
### かなり手こずるポイント
> invalid character '\'' looking for beginning of value

https://qiita.com/Syuparn/items/233c5b38164b5ea2fdf6

windows環境はダブルクオートにエスケープをつける必要があるみたい
> {\\"id\\":\\"8\\",\\"name\\":\\"tukioka kogane\\"}

#### ダブルクォーテーションのエスケープ処理
#### 全体をダブルクォーテーションで括る
[この2点はWindows環境のコマンドプロンプトからCurl接続する場合は注意しておきたい点](http://dim5.net/windows-server/curl-jsondata-postmethod.html)