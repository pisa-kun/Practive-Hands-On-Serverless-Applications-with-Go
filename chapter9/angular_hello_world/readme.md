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