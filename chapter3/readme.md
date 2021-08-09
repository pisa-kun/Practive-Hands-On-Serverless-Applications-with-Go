## chapter3 developing a serverless functiom with lambda
#### s3のzipをlambdaで呼び出す

1. set linux build, go build
> echo binary build  
> C:\WINDOWS\system32>set GOOS=linux  
> set GOOS=linux  
> go env GOOS  
> go build -o main main.go  

2. バイナリをzip化する

3. copy zip to s3
> aws s3 cp deployment.zip s3://hello-serverless-morichan
```cmd
C:\Develop\Go\ServerlessApplicationWithGo\chapter3> aws s3 cp deployment.zip s3://hello-serverless-morichan
upload: .\deployment.zip to s3://hello-serverless-morichan/deployment.zip
```

4. awsコンソールでlambda関数作成  
作成後、s3のzipを参照させる。  
ランタイム設定の編集　で ハンドラを`main`に変更


---
#### memo
- set GOOS で環境変数を変更した場合は、go envで環境が切り替わったことを確認。できれば、変更したコンソールでビルドコマンドしたほうが良い
- s3にアップロードしたzipを変更した場合、再度lambdaコンソールでs3のurlを指定しなおす
  - s3にアップロードしていても実行毎にurl先のオブジェクト参照するわけではなく、パス指定したタイミングで中のbinaryを抱え込むタイプ

- ハンドラ関数が見つからない場合は以下のエラーログになる
```
START RequestId: 0d4eae65-a614-41d5-9a74-f52a0f9b9769 Version: $LATEST
fork/exec /var/task/deployment: no such file or directory: PathError
```
- buildのosが異なる場合
```
### lambdaは linux osであることに注意
fork/exec /var/task/main: exec format error: PathError
```