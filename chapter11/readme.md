## chapter11 Monitoring and TroubleShooting

#### memo
- handler(input string)
handler(event Event) でlambdaでは入力値にjsonを渡す
[MyEventの参考](https://dev.classmethod.jp/articles/aws-lambda-supports-go/)

- x-rayのtrace
1. lambda - 設定- モニタリングとおよび運用ツール
2. AWS X-Ray アクティブトレースをONにして保存
3. モニタリング-トレースでサービスマップ確認
4. tracingフォルダのfindAllコードをlambdaに差し替え
5. x-rayで再度トレースを確認すると、findAll関数をより詳細にトレースできる
![画像](.\images\x-ray_1.png)

![画像](.\images\x-ray_2.png)
[参考](https://qiita.com/smith-30/items/225e27e6d9a110bce725)