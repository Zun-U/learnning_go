module LearnGo

go 1.20

// モジュール化
// 「go mod init モジュール名」でモジュールの初期化を行う。

// 「go mod tidy」
// ソースファイルを解析して必要なサードパーティライブラリのダウンロードや不要になったファイルの削除などを行ってくれるコマンド。

// 「go build」
// 再利用できるように、実行形式のバイナリファイルを作るコマンド。
// windowsだと「モジュール名.exe」ファイルが作成される。
// 名前を変更したい場合はフラグ『-o』を使用して、「go build -o hello_world setup.go」のようにする。

// 「go install」
// インストールしたいプロジェクトのソースコードリポジトリの場所を指定し、「@」に続けてツールのバージョンを指定する。
// ※最新バージョンを指定したい場合は、「~@latest」と指定する。
// 例)
// go install github.com/rakyll/hey@latest
//
// コマンドを実行すると、そのツールのダウンロードとコンパイルを行い、$GOPATH/binの元(あるいはgo/bin)にインストールする。