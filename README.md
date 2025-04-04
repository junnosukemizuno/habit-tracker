# habit-tracker

habit-tracker/ 
├── go.mod # モジュール定義ファイル（module名、依存パッケージ） 
├── go.sum # go.modに対応する依存パッケージのバージョン情報 
├── main.go # エントリーポイント。Ginサーバーを起動 
├── handlers/ # エンドポイントごとのハンドラ（処理）を定義するパッケージ 
│ └── habit.go # /habits POSTリクエストの処理を記述 
└── README.md # このプロジェクトの概要と使い方