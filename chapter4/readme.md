## chapter4
#### APIGateway

#### memo
- lambdaのハンドルは文字列完全一致になるので小文字・大文字も意識すること
> FindAllIdols.exe の場合
> × findallidols

- curl コマンド
> curl -X POST -H "Content-Type: application/json" -d '{"id":7, "name":"mukimuki"}' https://i3w8zvfye4.execute-api.ap-northeast-1.amazonaws.com/staging/idols