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