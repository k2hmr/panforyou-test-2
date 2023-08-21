# panforyou-test-2

## アプリの概要

以下の要件を満たす graphQL サーバです。  
a. panforyou-test-1 で Cloud Firestore に保存したデータを一覧で取得できます。  
b. id を指定して、panforyou-test-1 で Cloud Firestore に保存したデータ中の任意のデータを取得できます。

## 環境構築

### モジュールのインストール

```
  go mod download
```

### Firebase でサービスアカウントの秘密鍵の作成

以下の流れで秘密鍵(JSON ファイル)を作成し、ルートディレクトリに `path/to/serviceAccount.json` として保存してください。

```
  ↓Firebase コンソール
  ↓ プロジェクトの設定
  ↓ サービスアカウント
  ↓ 新しい秘密鍵を生成ボタン押下
  →JSON ファイルのダウンロード
```

## アプリの使い方

### graphQL サーバの起動

以下のコマンドで graphQL サーバを起動します。

```
  go run ./server.go
```

### クエリの実行

`http://localhost:8080/`にアクセスし、GraphQL playground で以下の各クエリを記述・実行してください。

#### Cloud Firestore に保存したデータを一覧で取得する

```
  query {
    breads {
      id
      name
      createdAt
    }
  }
```

#### id を指定し、Cloud Firestore に保存したデータ中の任意のデータを取得する

```
  query {
    bread(id: "指定のID") {
      id
      name
      createdAt
    }
  }
```
