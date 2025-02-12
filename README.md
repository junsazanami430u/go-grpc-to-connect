# go-grpc-to-connect

このプロジェクトは、既存のgRPCベースのサービスをConnectフレームワークに移行することを目的としています。
詳しくは、[zenn]()を参照してください。

## インストール

1. ルートディレクトリから依存関係をインストールします。

   ```bash
   make deps
   ```

## 使用方法
`grpc`または`connect`に移動して、下記を実行します。

1. go mod tidyを実行して、依存関係を整理します。

    ```bash
    make tidy
    ```

2. サーバを起動し、クライアントからリクエストを送信して、動作を確認します。

    サーバーの起動
    ```bash
    make serve
    ```

    クライアントからリクエストを送信
    ```bash
    make client
    ```

## test

```bash
make test
```

## lint

```bash
make lint
```