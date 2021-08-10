## chapter5
#### dynamodb

#### memo
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