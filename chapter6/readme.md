## chapter6 deploying your serverless application
#### 

#### memo
- invoke command

aws cli2からjsonファイルがbase64デフォになるとかで、`--cli-binary-format raw-in-base64-out`をつける必要がある

>aws lambda invoke --function InsertIdole --payload file://input.json result.json --cli-binary-format raw-in-base64-out

```json
//body
{
    "body":"{\"ID\":\"12\",\"name\":\"nanakusa haduki\"}"
}
```

※invoke commandは呼び出し ぐらいの意味

- update / version

関数コードのアップデート
> aws lambda update-function-code --function-name FindAllIdols --zip-file fileb://./FindAll.zip

関数のpublish  
>aws lambda publish-version --function-name FindAllIdols --description 1.1.0
