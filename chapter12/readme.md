## chapter12 Securing Your Sererless Application

#### memo
-  NotAuthorizedException: Unable to verify secret hash for client

一度チェックを付けた場合はクライアントを削除して再生成

 https://qiita.com/noobar/items/6615501b035e47792227

 既にユーザーが存在する場合、
 > UsernameExistsException: An account with the given email already exists.


- npm un angular-webstroage-service

anglular9だと互換性がないそうなのでビルドできない。`ngx-webstorage-service`に変更する。
>npm un angular-webstorage-service --save
>npm i ngx-webstorage-service --save 
(参考)[https://stackoverflow.com/questions/60507654/migration-issue-angular-7-to-9]

- ビルド成功後のエラー
> Uncaught ReferenceError: global is not defined
src/poclifill.ts に追記
https://dev.classmethod.jp/articles/angular6-referenceerror/

```html
// index.html 65L
 (window as any).global = window; // この行を追加
```

> Error: Bootstrap tooltips require Tether

bootstrapのバージョンがアルファだったので・・・
> Bootstrap 4 alphaのみ：
> ブートストラップ4alphaには Tether が必要なので、tether.min.jsを含める必要がありますbeforebootstrap.min.jsを含めます。
```html
<script src="https://npmcdn.com/tether@1.2.4/dist/js/tether.min.js"></script>
<script src="https://npmcdn.com/bootstrap@4.0.0-alpha.5/dist/js/bootstrap.min.js"></script>
```

- @viewChildが初期化されずにundefined

```ts
@ViewChild('login',{static: true}) public content:any;
```

- LOCAL_STORAGEの削除
f12-applicationから local_storageのkeyを削除する

![画像](.\images\cognito_2.png)