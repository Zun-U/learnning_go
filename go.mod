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
// ※httpサーバのロードテストを行えるツール
//
// コマンドを実行すると、そのツールのダウンロードとコンパイルを行い、$GOPATH/binの元(あるいはgo/bin)にインストールする。

// コードのフォーマット
// 「goimports」(標準に搭載されている「go fmt」の強化版)
// go install golang.org/x/tools/cmd/goimports@latest
//
// 「goimports -l -w .」←プロジェクト全体に実行
// フラグ「-l」を指定することで、フォーマットが正しくないファイルをコンソールに表示
// フラグ「-w」を指定することで、ファイルを直接書き換える
// 「.」の部分はチェックするディレクトリの指定（ファイル名を書くことができる）。この場合は、現行のディレクトリとすべてのサブディレクトリにあるファイルが対象になる。

// ❐　実は、GOにはセミコロン(;)が必要だが、コンパイラが規則に従って自動で挿入してくれている。　❐
// 規則
// ・識別子(intやfloat64等も含む)
// ・基本型のリテラル
// ・いずれかのトークン
// breake
// continue
// fallthrough
// return
// ++
// --
// )
// }
//
// ****************************************************************************************

// Goの開発者は「Effective Go」(https://oreil.ly/GBRut)
// と「Go Code Review Comments」(https://oreil.ly/FHi_h)(goのwiki)を読むべき。

// 「staticcheck」
// https://staticcheck.io/
//
// go install honnef.co/go/tools/cmd/staticcheck@latest
//
// go.modのあるディレクトリなら次のようにするだけで。staticcheckがプロジェクト全体に関して実行される。
// 「staticcheck」で実行

// 「go vet」
// 構文以外のエラーを検知する

// 「golangci-lint」
// 「staticcheck」や「go vet」などの複数のコードチェックツールを一緒に実行する。
//
// go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
// 「golangci-lint run」で実行
//
// 「golangci.yml」でどのリンターを実行し、どのファイルを分析するかを指定可能。
// ※https://oreil.ly/vufj1　を参照

// goは複数のバージョンが共存できる
// ※詳しくは公式参照
// https://go.dev/doc/manage-install

// Makefile
// windowsでは「make」をインストールする必要がある。
// 「scoop install make」
