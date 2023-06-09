package learn_type

import (
	"fmt"
)

func TypeMain() {
	fmt.Println("Chapter2 primitive type and variable")

	// 論理型
	fmt.Println(flag, ":論理型のゼロ値(宣言されたが値が割り当てられていない変数)は、『false』")
	fmt.Println(isAwsome)

	// 整数型
	fmt.Println(integerType, ":整数型のゼロ値は『0』")

	// 整数リテラルのデフォルトの型
	fmt.Printf("%v:%T  整数リテラルのデフォルトは「int」型\n", intLiteralDefault, intLiteralDefault)

	// 結果を同じ変数に代入
	reAssignmentInt()

	// 浮動点小数点のゼロ値
	fmt.Printf("%v:%T  :浮動小数点型のゼロ値は『0』\n", xf, xf)

	// 浮動小数点数リテラルのデフォルトの型
	fmt.Printf("%v:%T  浮動小数点数リテラルのデフォルトは「float64」型\n", xft, xft)

	// 文字列のゼロ値は「""」
	fmt.Printf("%v:%T  :文字列のゼロ値は『\"\"』(空文字列)\n", StringZeroValue, StringZeroValue)

	// runeリテラルのデフォルトの型
	fmt.Printf("%v:%T  :runeリテラルのデフォルトの型は「rune」型\n", RuneDefaultType, RuneDefaultType)

	// 文字列リテラルのデフォルトの型
	fmt.Printf("%v:%T  :文字列リテラルのデフォルトの型は「string」型\n", StringDefaultType, StringDefaultType)

	// 型変換
	TypeChange()

	// 変数の宣言「var」
	fmt.Printf("「var」による変数の宣言(変数xxを例にする)： 型を左辺(%T)、初期値(%v)を右辺に指定する\n", xx, xx)

	// 整数リテラルのデフォルトの型による型指定の省略
	fmt.Printf("変数yyの初期値である(%v)は、整数リテラルのデフォルトの型である(%T)であると分かる為、右辺の(%T)の型指定を省略できる\n", yy, yy, yy)

	// 変数の宣言時、型のみ指定すると、その変数にゼロ値を代入できる
	fmt.Printf("ゼロ値の代入は型指定のみ： 型 => (%T)、値 => (%v)(ゼロ値)\n", ZeroValueValiable, ZeroValueValiable)

	// 多重代入
	fmt.Printf("変数「tt」の値は(%v)、型は(%T): 変数「uu」の値は(%v)、型は(%T)\n", tt, tt, uu, uu)

	// 多重代入(ゼロ値)
	fmt.Printf("変数ssの初期値(%v)、型(%T)： 変数ddの初期値(%v)、型(%T)\n", ss, ss, dd, dd)

	// 異なる型の同時に変数宣言
	fmt.Printf("異なる型の変数の宣言 : value = %v, type = %T : value = %v, type = %T\n", gg, gg, jj, jj)

	// 宣言リスト
	fmt.Printf("%v:%T, %v:%T, %v:%T, %v:%T, %v:%T, %v:%T, %v:%T\n", xa, xa, ya, ya, za, za, da, da, ea, ea, fa, fa, ga, ga)

	// 関数ないの変数の宣言
	VariableDeclaration()

	// 複数の変数の宣言
	MultiVariableDeclaration()

	// 関数内の変数の宣言(暗黙的)
	VariableDeclarationShort()

	// 定数の宣言
	ConstTest()

	// 型なし定数
	fmt.Printf("%v:%T\n", yys, yys)
	fmt.Printf("%v:%T\n", zzs, zzs)
	fmt.Printf("%v:%T\n", dds, dds)

	// 型付き定数
	fmt.Printf("定数typedXの型は「%T」で値は「%v」\n変数zzzの型は「%T」で代入された値は「%v」\n", typedX, typedX, zzz, zzz)

	// 変数に代入された値が使われなくてもエラーが起きない
	UnUsedVariable()

	// 利用可能な変数名
	AbairableVariable()

	// 似たような変数名
	SeemsLikeSame()
}

// ❐ 基本型(primitive type) ❐
//
// 4つの型がある。
// ・論理型
// ・整数型
// ・浮動小数点型
// ・文字列型

// ❐ ゼロ値　❐
//
// 「宣言されたが値が割り当てられていない変数」に対してデフォルトのゼロ値(zero value)が割り当てられる。

// ❐　リテラル　❐
//
// 数値や文字、文字列などをじかに示したもの
// Goでは4種類ある

//　1.整数リテラル
// 通常は10進数として表す
// 他には「0b」を頭に置くことで2進数(binary)、「0o」で8進数(octal)、「0x」で16進数(hexadeciaml)を表せる
//
// Goでは整数リテラルの任意の箇所にアンダースコア「_」を書ける
// 例)「1_234」
// ※「1_2_3_4」は読みにくくなるから推奨しない
//
// ☆☆ アンダースコア「_」は10進数を3桁ごとに区切るのに用いたり、2,8,16進数で2文字、4文字ごとに区切るのに用いることが推奨される ☆☆

// 2.浮動小数点リテラル
// 「3.14」や「6.03e23」の様な指数を使った表記も可能
//
// 「0x」を頭につけることで16進数で表現できる
//
// 整数リテラル同様、浮動小数点リテラルでも桁を区切る目的で「_」が使える

// 3.runeリテラル
// 文字を表す
//
// ☆　先頭と最後にシングルクォテーション「'」を置く ☆
//
// runeリテラルは次の形式で表現される
//
// ・1文字のUnicode文字 (例：'a')
// ・8ビット8進数 (例：'\141')
// ・8ビット16進数 (例：'\x61')
// ・16ビット16進数 (例：'\u0061')
// ・32ビット Unicode (例：'\U00000061')
//
// この他に、バックスラッシュ「\」でエスケープされたruneリテラルもある
// 主なものは
//
// ・改行 ('\n')
// ・タブ ('\t')
// ・一重引用符 ('\'') ※シングルクオート
// ・二重引用符 ('\"') ※ダブルクオート
// ・バックスラッシュ ('\')

// 4.文字列リテラル
//
// 文字列リテラルを表すのに2つの方法がある
//
// 〇 解釈済みの文字列リテラル (interpreted string literal)
// ◇ 生(ロー)文字列リテラル (raw string literal)
//
//　ほとんどの場合は「"」の組を使って「解釈済み文字列リテラル」を作る
// この場合は0個以上のruneリテラルを含むことができる
// ここで使えない文字は、エスケープされていない「\」、改行、それに「"」がある
//
// ☆ 解釈済み文字列リテラルを使って2行にわたる文字列を表示したい場合は、「\n」を使って改行する必要がある
//
//
// 文字列中に「\」、改行、「"」を含めたい場合は、「`」(バッククォート)で囲まれた生(ロー)の文字列リテラルを使用する
// ロー文字リテラルには「`」(バッククォート)以外の任意の文字を含めることができる
// 次のようにすることで改行や二重引用符を文字列中に書くことができる
//   'バッククオートを使った"ロー文字列リテラル"を使うことで改行や二重引用符(ダブルクォート)を文字列中にかける'
//

// *************************************************
// GOの言語使用により、ソースコードは常にUTF-8で書かれる
// 16進数のエスケープを使わない限り、文字列リテラルUTF-8で書かれる
// *************************************************

// ❐ リテラルと型 ❐
//
// ・サイズの異なる整数については、加算さえもできない
// ・一方、整数リテラルを「浮動小数点数」の式で使ったり、整数リテラルを「浮動小数点数」の変数に代入できる
//
// ☆ Go言語において、リテラルは『型がない(untype)』ためで、"リテラルと『互換性を持つ』任意の変数に代入できる"
//
// ・ただし、「文字列リテラル」を数値型の変数に代入できないし、「数値リテラル」を文字列の変数に代入することもできない。
//
// ☆ リテラルが型を持たないのは、GO言語がプログラマーが型を指定するよう設計されている為（プログラミングがしやすい）
//
// ・指定された変数をあふれさせてしまうようなリテラルは、コンパイル時にエラーになる
// 例)「1000」をbyte型の変数に代入しようとするとエラーになる
//
// ★ 型が明示的に宣言されない場合がある。この場合は、リテラルに対して『デフォルトの型』が使用される。

//
// -*-*-*-*-*-*-*- 型(type) -*-*-*-*-*-*-*-
//

// ❐❐　論理型bool　❐❐
// 真偽値(論理値、bool値)を持つ変数の型は、論理型(boolean)になる
//
// ・ 論理型の変数は「true」「false」のいずれかの値を持つ
// ・ boolのゼロ値は『false』
//
// 例)
var flag bool // 値が代入されないとfalseになる。
var isAwsome = true

//
// ☆ 変数の型の多くの場合は、変数の宣言で決まる

// ❐❐　数値型　❐❐
// GOには「12種類」の数値を表す型がある
// これらは「3つのグループに分けられる」
//
// :::::: 整数型 ::::::
// Go言語で整数を表す型には、1バイトから4バイトまでの大きさの、『符号付き』の整数と『符号なしの整数』がある
//
// 「符号付き整数」
// int8    -128 ~ 127
// int16   -32768 ~ 32767
// int32   -2147483648 ~ 2147483647
// int64   -9223372036854775808 ~ 9223372036854775807
//
// 「符号なし整数」
// unit8	  0 ~ 255
// uint16   0 ~ 65535
// uint32   0 ~ 4294967295
// uint64   0 ~ 18446744073709551615
//
// ☆ すべての整数型のゼロ値は『0』
var integerType int // ゼロ値は『0』
// ◇ 特別な整数型 ◇
//
// 整数型の中に特別な名前を持つ物がある
//
// ・「byte」  --unit8の別名(※別名のことを『エイリアス』という)
// ・「int」   --CPUによって64ビットあるいは32ビットになる(int64あるいはint32と『同じサイズ』)
// ・「uint」  --intの符号がないもの(0以上)
// ・「rune」  --ひとつのコードポイントを表現する
// ・「uintptr」  --メモリを操作して、C言語の様なポインタ操作などを可能にするためのもの
//
// 「byte」はuint8のエイリアスで、byteとuint8の間の代入、比較、数値演算が可能。
// しかし、uint8は"めったに"使われない
//
// ☆☆ byteを使用するのが一般的
//
// 「int」は32ビットのCPUにおいては、32ビットの符号付き整数となるため、この場合のintは「int32」と同じサイズになる
// ほとんどの64ビットのCPUにおいては、intは64ビットの符号付き整数となるので、「int64」と同じサイズになる
// intは「プラットフォームによって大きさが異なることになる」ので、intと、int32やint64の間の代入、比較、数値演算は型変換しない限りコンパイル時にエラーとなる
//
// ☆☆　整数リテラルはデフォルトでint型となる　☆☆
var intLiteralDefault = 1

// ◇ 整数型の選択 ◇
//
// 整数型選択するための単純なルール
//
// ・ 「特定の大きさ」と「特定の符号の整数バイナリファイル」あるいはネットワークプロトコルを処理しているのならば、対応する整数型を使う
// ・ すべての整数型について機能するライブラリ関数を書く場合は2つの関数を書く。一つは引数と戻り値が「int64」のもの、もう一つが「uint64」のもの
// ・ その他の場合は「int」を使用する
//
// ☆ パフォーマンスや他のシステムの統合といった目的で、符号の有無やサイズについて明示する必要があるという場合を除けば、
// ☆ 「int型」を使うのが良い
// ☆ 理由なく他の型を使うのは「早すぎる最適化」

// ◇ 整数関連の演算子 ◇
//
// Go言語では、「+」「-」「*」「/」の四則演算子と、剰余(あまり)を計算する「%」の各演算子が使える
//
// ・ 「整数」の割り算の結果は、「整数」になる
// ・ 「浮動小数点数」の結果が欲しい場合は、型変換をして「浮動小数点数」に変換してから計算する
//
// ★ 「0」による割り算をしてしまうと『パニック』になるので避ける
// ※ 「整数」の割り算は「0」の方向に丸められる
//
// ある変数に対して、演算を行って「結果を同じ変数に代入」する場合は。「+=」「-=」「*=」「/=」「%=」の各演算子が使える
func reAssignmentInt() {

	var x int = 10

	x *= 2 // 関数内でのみ書ける書き方

	fmt.Printf("%v:%T 演算を行って結果を同じ変数に代入\n", x, x)

}

//
// ・ 「整数」の比較には「==」「!=」「>」「>=」「<」「<=」を使う
//
// ・ 「整数」に"対して"はビット演算を行うことができる
//
// ・   左シフトは「<<」、右シフトは「>>」で行う
//
// ・   ビット単位の「&」(AND)、「|」(OR)、「^」(XOR)、「&^」(AND NOT)の各演算も可能
//
// ☆★☆ 算術演算子と同様、「=」の前に演算子を置いて『左辺の変数に対する演算と代入』ができる
//       「&=」「|=」「^=」「&^=」「<<=」「>>=」

// ❐ 浮動点小数点 ❐
//
// # Go言語の浮動点小数点は「float32」と「float64」の2種類がある
//
// ☆ 整数型と同様、浮動小数点型のゼロ値は「0」
var xf float64

// ☆ 浮動小数点数リテラルのデフォルトの型は「float64」
var xft = 3.14

//
// ※ 浮動小数点数の扱いは『IEEE 754』に準拠している
//
//
// ・ 既存のフォーマットとの互換性を気にする必要が無ければ、「float64」を使用する
//
//
// ・ 浮動小数点数リテラルのデフォルトの型が「float64」なので、 『いつもfloat64』というのが最も簡単な選択
// ☑ float32は10進数で6~7桁の精度しかないので、『精度を考えても』float64を推奨する
//
// ※※　そもそも、浮動小数点数を使う必要があるか要検討する　※※
// -- ほとんどの場合、その必要がない
//
//
// ・ Goにおいて、浮動点小数点があらわす範囲が広く、あくまで『近似値』でしかない
// (グラフィクスや物理計算などの誤差があっても許容される場合)
//
// ☑ 浮動小数点数は、『10進数を正確に表現できない』
//    金額など誤差なしで10進数を表現しなければならない場合には「使わない」
//
// ・ 浮動小数点数については、標準的な四則演算子や比較演算子が使える(「%」を除く)
//
// 割り算についてはいくつか留意する点がある
// -- 非ゼロの浮動小数点数を「0」で割ると、+Inf、または-Inf(正あるいは負の無限大)になる
// -- 0が記憶された浮動小数点数型の変数を「0」で割ると「NaN」(Not a Number)になる
//
//
// ☆★☆ 浮動小数点数は常に「近似値」
//       その為、比較ではなく2つの浮動小数点数の差が『一定の範囲内に収まっているか』を判定する
// ※ 『一定の範囲』どのくらいにするべきかは求められる精度に依存する(この範囲のことを"イプシロン"と呼ぶ場合がある)

// 複素数型
//
// 大規模なプログラムの一部でマンデルブロ集合を計算する必要や、二次方程式を解く必要があるといった場合の為、Go言語では複素数の計算はできるようになっている

// ❐❐　文字列型とrune型　❐❐
//
// 文字列のゼロ値は「""」(空文字列:くうもじれつ)
var StringZeroValue string

// ・ GoはUnicodeをサポートしているため、文字列にUnicode文字を含めることができる
//
// ☆ 整数や浮動小数点数と同様に、
//
//	--文字列の比較にも「==」あるいは「!=」を用いる
//	--順序の比較には「>」「>=」「<」「=<」を用いる
//
// 　 --文字列連結には演算子『+』を用いる
//
// ★★ Goの文字列はイミュータブル(変更不可能)である
//
//	したがって、文字列型の変数に現在記憶されているものとは『別の値を代入する』ことはできるが、
//	そうした変数に代入された『文字列の値』を変えることはできない
var ImutableWords string = "文字列はイミュータブルである。"

// ☆★ Go言語にはひとつの「コードポイント」を表現する型であるrune型もある
//
//	:::: 「rune」型は「int32」の別名(エイリアス)
//	runeリテラルのデフォルトの型はrune(int32)型
//	そして、文字列リテラルのデフォルトはstring型(文字列)
var RuneDefaultType = 'a'
var StringDefaultType = "a"

//
//
// ※ コードポイントは、 (Unicode などの) テキストを表現するシステムにおいて、抽象文字を表現するために割り当てられた『番号』のこと
// 例)「あ」 ⇒ 「U+3042」
//    「a」 ⇒ 「U+0061」
//
//
// ☆☆☆ 文字列を参照している場合は、「int32」ではなく「rune」を使用する
//        ※ コンパイラにとっては同じかもしれないが、『コードの意図が明確になる』

// ❐❐　明示的型変換　❐❐
//
// Goは明快さと可読性に重きを置く言語のため、変数間の『暗黙的』型変換は "行わない"
//
// ★★ 型が違う場合は、必ず型変換を行う
//
//	整数や浮動小数点数で『サイズが異なる "だけ" 』の場合も変換が必要
func TypeChange() {

	var x int = 10 // 整数型
	fmt.Printf("%v:%T 整数型\n", x, x)

	var y float64 = 30.2 // 浮動小数点数型
	fmt.Printf("%v:%T 浮動小数点型\n", y, y)

	var z float64 = float64(x) + y // 変数x(整数型)を浮動小数点数型に型変換

	var d int = x + int(y) // 変数y(浮動小数点数型)を整数型に型変換

	fmt.Printf("%v:%T, %v:%T\n", z, z, d, d)
}

// 型に対するこのような厳密さは、bool(論理型)についても適用される
//
// ☆☆☆☆☆ Goの型変換はすべて『明示的』に行われる為、論理型以外の型の変数などを論理型として扱う事は出来ない
// 　　　　　 ※ 多くの言語では非ゼロの数値や、空でない文字列を『true』として扱うことができるがGoはそうではない
//
// ★★★★★ 論理型に変換できるのは論理型しかない ★★★★★
//
// ※ 他のデータ型を『論理型』に変換したい場合は、比較演算子を用いるしか方法がない
//    xが0と等しいかどうか判定する方法は「X == 0」しかないし、文字列sが空文字列かを判定する方法は「s == ""」しかない
//
// Goでのこのような型変換は、単純さと明快さを得るために、冗長さを招いている点だと言える
// Goではこのような「トレードオフ」をいくつか有している
//
// ★ Goでは簡潔さよりも理解しやすさに重きが置かれている

// :::::::::::  変数の宣言 :::::::::::
//
// ・もっとも詳細な指定は『var』を使用する
var xx int = 10 // 型と初期値を指定する
// 「=」の右辺の型(上記の例でいうと "10" )が「変数の想定する型」になっている場合は、型指定(上の例ではint)を省略できる
// 例えば、整数リテラルのデフォルトの型は「int」なので、左辺のintは省略できる
var yy = 10 // 整数リテラルのデフォルトの型はintなので、intの指定を省略できる
// ☆☆☆ 変数を宣言してその変数にゼロ値を代入したいのなら、型だけを指定して「=」から後を省略することもできる
var ZeroValueValiable int

// 「var」を使って、同じ型の複数の変数を宣言できるし、同時に初期値を代入(☆多重代入)することもできる
var tt, uu int = 10, 20
var ss, dd int // ゼロ値の複数宣言
// 型の異なる変数を宣言することも可能
var gg, jj = 10, "hello" // 整数のリテラルと、文字列のリテラルのデフォルトの型により、型指定の省略している
// 複数の宣言を、下記の様な宣言リストにまとめることができる
var (
	xa     int
	ya         = 20
	za     int = 30
	da, ea     = 40, "hello"
	fa, ga string
)

// ☆★☆ 関数『内』ではvarによる宣言の代わりに演算子「:=」を使うことができる
// 　　　 次の2文はまったく同じことになり、いずれも型が推定される
func VariableDeclaration() {

	var x = 10

	// 関数内でのみ使用できる変数宣言　『:=』演算子
	xx := 10

	// どちらも同じことになる(どちらも同じように型推定される)
	fmt.Printf("%v:%T, %v:%T\n", x, x, xx, xx)

}

// 「var」の時と同様に、「:=」を使って複数の変数宣言が可能
func MultiVariableDeclaration() {

	var x, y = 10, "hello"

	xx, yy := 10, "hello"

	fmt.Printf("%v:%T, %v:%T, %v:%T, %v:%T\n", x, x, y, y, xx, xx, yy, yy)

}

// 「:=」には「var」にはできない事ができる
// ★ 既存の変数に値を代入できる
// ☆☆「:=」の左辺に新しい変数が一つあれば、他の変数は既存のものも使える
func VariableDeclarationShort() {

	x := 10

	// 既存の変数に値を代入できる
	x, y := 30, "hello"

	fmt.Printf("%v:%T, %v:%T\n", x, x, y, y)

}

//
// 「:=」には一つ制限があり、パッケージ(package)レベルで変数を宣言する場合は、varを使わなければいけない
// ☆★☆ 「:=」は関数の外では使えない
//
//
// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
//
// -- 関数内で最もよく使われるのは「:=」
// -- 関数の外では、パッケージレベルの複数の変数を宣言するのなら宣言リストを使用する(このようなケースはあまりない)
//
// ☆ ただし、関数内でも「:=」を避けるべき時がある
//
// ・変数を『ゼロ値に初期化する』場合は『var x int』の形式を使用すること
//	**** ゼロ値が意図されていることが明確になる ****
//
// ・型が指定されていない "リテラル" を変数に代入するとき、そのリテラルのデフォルトの型が「希望する変数の型と異なる時」には
//   型を指定して長い形式の「var」を使用する
// 　※型が指定されていない "定数" の場合も同様
// ☆☆☆ 値の型を指定して「x := byte(20)」と型変換を使って書くのはエラーにはならないが、「var x byte = 20」と書く方がイディオム的
//
// ・「:=」は新しい変数に代入する為だけではなく、既存の変数に代入するためにも使える
// 　(★ ただし、既存の変数に代入できるのは、一番内側のブロックで宣言されたものだけ)
//   このため、既存の変数を再利用しようとして新しい変数を作ってしまう場合がある
// 　こうした状況で「どの変数が新しいものであるかを明示するため」、「var」を使って新しい変数を明示的に宣言しておいてから
// 　「=」を使って変数に値を代入するようにする(変数のシャドーイング)
//
//
// ☆☆☆ 「var」や「:=」を使うと複数の変数の宣言を１行で行えるが、このスタイルを使うのは複数の値を返す関数(あるいは「カンマ ok イディオム」)からの戻り値を代入するときだけにする
//
//
// ※ 関数の外で変数の宣言をする必要は(ほとんど)ないはず
// ※ 関数の外側は「パッケージブロック」と呼ばれる
// ※ ★ 基本的にパッケージレベルでの変数の値を変更するのは避ける
// ※ ☆ 関数の外にある変数は値の追跡が "困難" でその結果データフローの理解が妨げられることになり、バグにつながる
//
// ☆☆☆　一般的なルールとして、パッケージブロックにおいては基本的に『イミュータブル(可変不可)』な変数のみ宣言するべき

// ❐❐ 定数 ❐❐
//
// Goでは値がイミュータブルなものを宣言する方法が「const」になる
const xax int64 = 10

const (
	idKey   = "id"
	nameKey = "name"
)

const zaz = 20 * 10

func ConstTest() {

	const yay = "hello"

	fmt.Println(xax)
	fmt.Println(yay)

	// 再代入できない、何故な定数は「リテラルに名前を付ける」ものだから
	// xax = xax + 1
	// yay = "bye"

	fmt.Printf("%v\n", xax)
	fmt.Printf("%v\n", yay)
	fmt.Printf("%v\n", zaz)
	fmt.Printf("%v\t%v\n", idKey, nameKey)

}

//
// 定数を宣言するのは、パッケージレベルあるいは関数内になる
// 「var」と同様、リストの様に()で囲んで関連する複数の定数を宣言できる(☆ このほうが一か所にまとまる)
//
// ☆☆☆ Goのconstは機能が限られていて、リテラルに「名前」を付けるためのもの ☆☆☆
// ★★★ コンパイル時に決定できる値を保持できるだけ ★★★
// ◇つまり代入できるのは次のものに限られる
// ・数値リテラル
// ・true および false
// ・文字列
// ・rune
// ・次にあげる組み込み関数の「結果」 --- complex, real, imag, len, cap
//   --「len」　配列の "現在" の長さ(length)を調べる関数
//   --「cap」　スライスの容量(capacity)(☆ "あらかじめ確保されているサイズ")を調べる関数
// ・上であげた「値と演算子から構成される式」
//
// ※ 「iota」と呼ばれる値も「const」で使える
//
// **** Go言語には、実行時に計算された値がイミュータブルであることを指定する方法がない
// 　　 イミュータブルな「配列」「スライス」「マップ」「構造体」はない
// 　　 また、「構造体」のフィールドがイミュータブルであることを宣言する方法もない
// 　　 ☆ これは。関数の中で変数の値が変更されているかどうかは明白である"べき"だから、イミュータブルかどうかはそれほど重要ではない

// ◇ 型付きの定数と型のない定数
//
// 定数には型があるもの(type)と型がないもの(untyped)がある
//
// ★★ 型のない定数はリテラルとまったく同じように動作する
//
//	　型のない定数もデフォルトの型をもっており、型を推定する必要がある場合はデフォルトの型が使用される
//
// ☆☆ 型付きの定数は、「同じ型の変数」に対してのみ、直接代入できる
//
// ☆★☆ 定数を型付きにするかどうかは、定数が宣言される『理由』による
// 　　　 複数の数値型と比較等をする可能性がある定数を定義している場合は、定数を形無しにする
//
//	☆ 一般に、定数は形無しのままにしておく方が柔軟になるが、定数に型を強制したくなる場合もある
//
// ＊ 型なし定数の宣言 ＊
const xxs = 10

// 次のような代入も許される
var yys int = xxs
var zzs float64 = xxs
var dds byte = xxs

// ＊ 型付き定数の宣言 ＊
const typedX int = 10

// この定数を直接代入できるのは『intの変数のみ』
// 例えば以下の様に他の型に変数を代入しようとするとコンパイルエラーになる
// var zzz float64 = typedX // compile erroe
var zzz int = typedX // ok

// ◇未使用定数
//
// Goの一つのゴールとして「大規模なチームによる共同開発を容易にする」というものがあり、このためのユニークなルールがいくつかある
// 例)Goのプログラムは「go fmt」によって決まった形式でフォーマットされなければならない(コードを操作するツールの作成を容易にし、またコーディング標準を提供するため)
//
// ☆☆☆☆☆☆ Goには「宣言された局所変数(ローカル変数)はすべて使われなければならない」というルールがある
//
//	宣言されただけで、一度でもその値が使われない局所変数があるとコンパイル時にエラーになる ☆☆☆☆☆☆
//
// 未使用変数のコンパイラによるチェックはそれほど厳密ではなく、『変数が一度でも使われればよい』
//
// たとえば次の例の様に「変数に値が代入されたがその値が使われなかった」
func UnUsedVariable() {
	x := 10
	x = 20
	fmt.Println(x)
	x = 30 // 値が代入されたが、使われなくてもエラーにならない
}

// 「コンパイラ」および「go vet」は変数xへの10や30の代入について文句を言いません
// ただし、「golangci - lint」は警告を発します
//
// ☆ パッケージレベルの変数が使われない場合はコンパイル時のエラーにならない

// ***********************************************************************
// 　　　　　　　　　　　　　　　未使用の定数
//
// 驚くべきことといえると思うが、constを使って宣言された定数が使われていなくても
// エラーにならない
// ☆☆ 理由は、Goの定数はコンパイル時に計算され副作用がない
//	　　ある定数が使われていなければコンパイル後にバイナリファイルに含める必要が
// 　　 なくなるだけの話
//
// ***********************************************************************

// ◇◆◇◆◇◆◇　変数および定数の名前　◇◆◇◆◇◆◇
//
//	変数と定数の名前の付け方には違いがある
//
// ☆☆ ほとんどの言語同様、変数の名前は「文字」あるいは「_」(アンダースコア)ではじまり、「数字」あるいは「_」あるいは「文字」が続く
//
// 「文字」と「数字」の定義は他の言語より少し広くなっている
// ☆☆☆☆ Unicode文字で「文字」あるいは「数字」とみなされているものが許されている　★★★★
// 例えば、以下の変数はすべて使用可能である
func AbairableVariable() {

	_0 := 0_0 // 「0_0」の「_」は桁区切り用
	_𝟙 := 20
	π := 3
	ａ := "hello" //「全角」の「a」
	数値1 := 22    // 漢字 + 全角の「1」

	fmt.Printf("変数_0の値「%v」:型「%T」\n", _0, _0)

	fmt.Printf("変数_𝟙 の値「%v」:型「%T」\n", _𝟙, _𝟙)

	fmt.Printf("変数πの値「%v」:型「%T」\n", π, π)

	fmt.Printf("変数ａの値「%v」:型「%T」\n", ａ, ａ)

	fmt.Printf("変数数値1の値「%v」:型「%T」\n", 数値1, 数値1)

}

// ☑ このコードは動作するが、このような変数の命名法は推奨できず、「イディオム的」ではないとみなされる
//
// こうした変数は区別がつきにくいし、環境によっては入力が難しい場合もある
// 次のリストも同様で、似たような文字でも、まったく別の変数を表すことになる
func SeemsLikeSame() {

	ａ := "こんにちは"     // 全角の「a」(Unicode:U+FF41)
	a := "konichiwa" // 半角の「a」(Unicode:U+0061)

	fmt.Printf("変数ａの値は「%v」で型は「%T」\n", ａ, ａ)

	fmt.Printf("変数aの値は「%v」で型は「%T」\n", a, a)
}

//
//
// \\\\\\\\\\\\\\\\\\\\\\\\\\\\\
// 変数名に「_」を "含める" ことは出来るが、基本的にはこのような用途には使われない
//
// イディオム的Goではスネークケースは使わず、キャメルケースを使用する
//
// ・「index_center」「number_tries」 ⇒ ☓
// ・「indexCenter」「numberTries」   ⇒ 〇
//
// ※「_」"だけ" からなる変数は特別な識別子として使われる
// 　	関数が複数の値を返すが、そのうちの一つ以上の値を使う必要がない(無視したい)とき、『使わない値を「_」に代入する』
//
// \\\\\\\\\\\\\\\\\\\\\\\\\\\\\
//
//
// ]]]]]]]]]]]]]]]]]]]]]]]]]]]]]
// 多くの言語では、定数は「INDEX_CENTER」あるいは「NUMBER_TRIES」のように大文字だけを使って、単語の区切りに「_」を用いて表す
// Goではこのようなパターンは用いない
//
// ★★★★ パッケージレベルの宣言において、名前の先頭文字が大文字か小文字かでそれがパッケージの外からアクセスできるかどうか御決めることになっている ★★★★
//
// ]]]]]]]]]]]]]]]]]]]]]]]]]]]]]
//
//
// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>
// 基本的に関数内では短い名前を使う
// 変数のスコープ(有効範囲)が狭ければ狭いほど短い名前が使われる
// 1文字の変数が使われることも珍しくない
// ★ たとえば、for-rangeループではキー(key)を表すのに「k」を、値(value)を表すのに「v」を使うのが一般的
// ☆ 標準的なforループの場合なら、「i」や「j」がループ変数(ループカウンタ)として使われる
//
// ☆★☆ Goは強い方システムを持つ言語のため、もともとの型(基底型)が分かるようにしておく必要はない ☆★☆
//
// ★☆★ ただし、1文字の変数名については次のような文字を用いるという慣例があるので、以下のパターンを踏襲する
// ・整数(integer)は「i」
// ・浮動小数点数(float)は「f」
// ・論理値(boolean)は「b」
//
//
// |||||||||||  短い名前の変数の用途が分からなくなるようなら、そのブロックの処理が複雑すぎることになる ||||||||||||
//
// ☆☆☆ パッケージブロックの変数や定数の名前は、用途が分かるよう説明的なもの(したがって長めのもの)にする
// 				※ ここでも、型に関する情報を名前に含める必要はない
//
//
// ★★☆ すべての型(type)に基底型(underlying type)がある
//
// ・型Tが基本型(論理型、数値型、文字列型)あるいは型リテラル(type literal)の場合、Tの基底型はT(自分自身)になる
// ・それ以外の場合、Tの宣言で参照してる型が「T」の基底型になる
//
// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>
