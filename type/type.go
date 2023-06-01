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

// ❐　論理型bool　❐
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
//
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