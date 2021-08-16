## chapter9 building the frontend with s3

#### memo

^ npmのインストール先

グローバルインストールとローカルインストールがある
> npm install -g bootstrap@4.0.0-alpha.6

通常はグローバルインストール(-g)をつけるみたい

- .angular-cli.jsonがない
V1.7系から angular.jsonに変更になったらしい
```
Angular CLI v1.7からv6.0へのアップデートのうち、最大の変更は.angular-cli.jsonファイルからangular.jsonファイルへの移行です。 これまでAngular CLIの各種設定を記述していた.angular-cli.jsonファイルは、名前だけでなく内部のJSON構造も互換性のない新しいangular.jsonに変わります。
```

- ng generate component

ng new で作成したルートディレクトリで上記のコマンドを打つと、src/app直下にコンポーネントを作成する

```
ng generate component components/navbar
CREATE src/app/components/navbar/navbar.component.html (21 bytes)
CREATE src/app/components/navbar/navbar.component.spec.ts (626 bytes)
CREATE src/app/components/navbar/navbar.component.ts (276 bytes)
CREATE src/app/components/navbar/navbar.component.scss (0 bytes)
```

- component追加時のお作法みたいなもの

tsファイルのselectorの名前を変えて、app.component.htmlに追加する
```ts
@Component({
  selector: 'idol-item',
  templateUrl: './idol-item.component.html',
  styleUrls: ['./idol-item.component.scss']
})
```

- ng new

cssを基本的に選択する
```
 ng new frontend
? Would you like to add Angular routing? Yes
? Which stylesheet format would you like to use? (Use arrow keys)
> CSS
  SCSS   [ https://sass-lang.com/documentation/syntax#scss                ]
  Sass   [ https://sass-lang.com/documentation/syntax#the-indented-syntax ]
  Less   [ http://lesscss.org                                             ]
```

[参考ページ](https://www.techiediaries.com/angular-bootstrap/)

```json
            "styles": [
              "./node_modules/bootstrap/dist/css/bootstrap.css",
              "src/styles.css"
              ],
            "scripts": [
              "./node_modules/jquery/dist/jquery.js",
              "./node_modules/bootstrap/dist/js/bootstrap.js"
            ] 
```

- movie-item.components.ts でのidolフィールド

フィールド変数を初期化しないとコンパイルできないので、とりあえず空文字で初期化しておく
```ts
export class IdolItemComponent implements OnInit {
  @Input()
  
  public idol:Idol;
  constructor() {
    this.idol = new Idol("","","");
  }
```

- angular/http パッケージ

下記に変更されてる
```ts
// REST クライアント実装ののためのサービスを import ( Angular 5.0.0 以降はこちらを使う )
import { HttpClient, HttpHeaders } from '@angular/common/http';
```

[参考](https://qiita.com/ksh-fthr/items/840ae54472892a87f48d)

このAPIServiceを使用する例
```ts
    this.idolApiService.findAll()
    .then(
      (response) => {
        this.param = response;
        // bodyの一覧を格納
        this.messageInfoList = this.param.body;
        console.log(this.messageInfoList);
        const obj = JSON.parse(this.messageInfoList);
        console.log(obj);
        obj.forEach( (idol:any) => {
          this.idols.push(new Idol(idol.name, "description"))
        });
    .catch(
      (error:any) => console.log(error)
    );
```

- [Angular]"No provider for xxx"エラーが発生するときの対処法

HttpClientを追加したばかりのころ、ヘッダー含めて何も表示されない状態になった。  
[公式](https://angular.jp/guide/http#httpclient)ドキュメント曰く、app.module.tsにHttpClientModuleをインポートする必要がある
```ts
import { HttpClientModule } from '@angular/common/http';

@NgModule({
  imports: [
    BrowserModule,
    AppRoutingModule,
    HttpClientModule
  ],
```

- "Access-Control-Allow-Origin": "*" つけてもエラーになる件

> has been blocked by CORS policy: No 'Access-Control-Allow-Origin' header is present on the requested resource.

~~ワイルドカードだとポリシー違反になる。~~

上手くいかなかったので、[Cross Domain](https://chrome.google.com/webstore/detail/cross-domain-cors/mjhpgnbimicffchbodmgfnemoghjakai/related?hl=ja)というのをインスコすればCORSのエラーはなくなった

- 200P手順7

ソースコードの大きな改変が必要
#### dynamoDBのcoverとdescriptionについて

goのscanでテーブルにアクセスする際に、キーが存在しない場合例外になる。
panicで例外捕捉するか、テーブルに必ずキーを追加して空文字にする必要がある。

下記のようなエラー
```json
{
  "errorMessage": "runtime error: invalid memory address or nil pointer dereference",
  "errorType": "errorString",
  "stackTrace": [
    {
      "path": "github.com/aws/aws-lambda-go@v1.26.0/lambda/errors.go",
      "line": 39,
      "label": "lambdaPanicResponse"
    },
    {
```

- postコマンド

> C:\Users\pisa0>curl -X POST -H "Content-Type:application/json" -d " {\"id\":\"14\",\"name\":\"Amana Osaki\",\"cover\":\"https://shinycolors.idolmaster.jp/pc/static/img/download/wallpaper/icon_alstroemeria_amana.jpg\",\"description\":\"大崎 姉妹の双子の妹。誰とでも分け隔てなく接する天真爛漫なギャル。今しかできないことを全力で楽しみたい今ドキの女の子。高校2\"}" https://i3w8zvfye4.execute-api.ap-northeast-1.amazonaws.com/staging/idols
{"Attributes":null,"ConsumedCapacity":null,"ItemCollectionMetrics":null}

コマンドプロンプトからのcurlだと日本語が文字化けする

- ng-bootstrapのエラー

[参考ページ](https://stackoverflow.com/questions/60824732/after-installing-material-design-i-am-unable-to-compile-the-angular-code-due-to)
>Error: Failed to compile entry-point @ng-bootstrap/ng-bootstrap (module as esm5) due to compilation errors:

1.一旦npm uninstall でbootstrapを削除
2.npm `add`で追加する
3.NgbModule.forRoot()を app.module.tsから削除する

- "item" is not a know element と表示されたとき

```ts

ng generateしたら必ず.tsファイルの slectorの名前を変えておこう

@Component({
  selector: 'new-idol',
  templateUrl: './new-idol.component.html',
  styleUrls: ['./new-idol.component.css']
})
```

- postのパラメータ

204ページのpostする movieはプロパティ3つでIdフィールドがない。この状態だとlambdaで引数エラーが返ってくるはず・・・

```json
{
    "id":"99",
  "name": "Amana Osaki",
  "cover": "https://shinycolors.idolmaster.jp/pc/static/img/download/wallpaper/icon_alstroemeria_amana.jpg",
  "description": "大崎姉妹の双子の妹。誰とでも分け隔てなく接する天真爛漫なギャル。今しかできないことを全力で楽しみたい今ドキの女の子。高校2年生。"
}
```

クラスのフィールド変数を追加して、json.Marshalする前にIdを追加する

```ts
    constructor(name: string, description: string, cover?: string, id?: string){
        this.name = name;
        this.description = description;
        this.cover = cover ? cover : "http://via.placeholder.com/185x287";
        this.id = id ? id : "999";
    }

  insert(idol: Idol){
    idol.addId("114");
    console.log("insert +", idol)
    return this.http
      .post(environment.api, JSON.stringify(idol))
      // .map(res => {
      //   return res
      // })
    }
    
```

- s3バケット作成の失敗例

大文字をbucket名に含めれない
> aws s3 mb s3://ServerlessApplicationMorichan.com
> make_bucket failed: s3://ServerlessApplicationMorichan.com An error occurred (InvalidBucketName) when calling the CreateBucket operation: The specified bucket is not valid.

```json
{
  "Id": "Policy1628968528756",
  "Version": "2012-10-17",
  "Statement": [
    {
      "Sid": "Stmt1628968526017",
      "Action": [
        "s3:GetObject"
      ],
      "Effect": "Allow",
      "Resource": "arn:aws:s3:::serverlessmorichan.com/*",
      "Principal": "*"
    }
  ]
}
```

- route53の設定まわり

ここを参考
https://blog.kozakana.net/2019/03/aws-dns-validate/