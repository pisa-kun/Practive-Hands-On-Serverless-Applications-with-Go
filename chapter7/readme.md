#### cjapter7
## CICD

- buildspec  
ルートディレクトリにbuildspecファイルを配置する。フェーズ詳細でどこで失敗しているか確認する
https://maku.blog/p/xdnu3ah/

- s3へのcpコマンドに対するアクセス権利

codebuildにs3への、putとgetに対するアクセス権限が必要

> [Container] 2021/08/12 07:36:48 Running command aws s3 cp $CODEBUILD_RESOLVED_SOURCE_VERSION.zip s3://$S3_BUCKET/
upload failed: ./5e0f47338af1e90db484aa7cae8f1fac84306899.zip to s3://hello-serverless-morichan/5e0f47338af1e90db484aa7cae8f1fac84306899.zip An error occurred (AccessDenied) when calling the PutObject operation: Access Denied

環境を編集するから iamのarnをコピーし、iamにアクセスを付与する
```
s3:putobject
s3:getobject
lambda:UpdateFunctionCode
```