# go-rest-chi
base project for go rest api with chi

## 使用ライブラリ
- wire
- viper
- swag
- gomock
- go-playground validator
- chi
- direnv

## 実装内容
- Chiでのルーディング
- Middleware
  - context value
    - request id
    - request info
  - CORS
  - logging
  - エラーハンドリング（recover）
  - タイムアウト
  - 認証
  - tracing
  - metrics
- カスタムエラー 
- カスタムロガー
- バリデーター
- wire ジェネレート
- swagger定義のジェネレート
- mock ジェネレート
- telemetry
  - datadog
- CI
  - lint
  - test
  - Dockerfile/docker-compose

### README に載せたい情報
- swagger の使い方
- mock generator の使い方
- wire の使い方
- datadog container の使い方


## このプロジェクトは以下のディレクトリ構造に基づいています：
```
.
├── cmd/
│   └── api/                 # アプリケーションのエントリーポイント
├── docs/
│   └── swagger/             # API仕様書（OpenAPI/Swagger）
├── internal/
│   ├── adapter/            # 外部システムとのインターフェース層
│   │   ├── primary/           # 外部からのリクエストを受け付ける側のアダプター
│   │   └── secondary/         # 外部システムへのアクセスを行う側のアダプター
│   ├── domain/                # ビジネスロジックの中心
│   │   ├── common/            # 共通ユーティリティ
│   │   ├── service/          # サービス
│   │   └── usecase/          # アプリケーションのユースケース
│   ├── infrastructure/      # 横断的・技術的な実装詳細
│   │   ├── apperror/         # カスタムエラー
│   │   ├── config/            # 設定管理
│   │   ├── logger/            # ログ機能
│   │   ├── telemetry/         # 監視・計測（メトリクス、トレーシング）
│   │   └── validator/         # バリデーション
│   └── mock/               # モックオブジェクト
├── script/                 # ビルド、デプロイ、その他のスクリプト
└── tools/                   # 開発ツール関連
```

## Setup

### node
```shell
npm install
```

### direnv
1. direnv をインストール
```bash
brew install direnv
```
- シェルの設定を追加（~/.zshrc や ~/.bashrc）
```
eval "$(direnv hook zsh)" 
```

2. 環境変数テンプレートをコピー
```bash
cp .envrc.example .envrc
```

3. `.envrc` を編集

4. direnv の許可
```bash
direnv allow
```
