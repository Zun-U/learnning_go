package synthetictype

import (
	"fmt"
)

func  SyntheticTypeMain() {

	fmt.Println("======== 合成型 ========")
	// 『合成型』--「複合データ型」、あるいは「コンテナ型」とも呼ばれる

	// 配列の宣言方法
	fmt.Printf("%v:%T\n", xa, xa)

	// 値を指定した配列の宣言
	fmt.Printf("%v:%T\n", xar, xar)

	// 途中に空きがある配列の宣言
	fmt.Printf("%v:%T\n", xarr, xarr)

	// 配列の初期化
	fmt.Printf("%v:%T\n", xarra, xarra)

	// 配列の比較(「==」と「!=」が使える)
	fmt.Println(ax == ay)

	// 疑似多次元配列
	fmt.Println(xaaa)

	// 配列のインデックスを指定して値を代入
	AssignmentArray()

	// 組み込み関数(len)を使用して、配列の長さを調べることができる
	fmt.Println(len(xa))

	// スライスの宣言
	fmt.Printf("%v:%T\n", xtr, xtr)

	// インデックスを指定したスライスリテラルを用いた宣言
	fmt.Println(xtrtr)

	// インデックスを指定して要素の値を変更、取得
	SliceIndex()

	// 多次元スライス
	fmt.Println(xslice)

	// スライスの初期値はnil
	fmt.Println(xsliceZero)

	// スライスの比較はnilかどうかしか行えない
	fmt.Println(xsliceZero == nil) // true

	// スライスに要素の追加、「append」
	SliceAppend()

	// すでに値が入っているスライスに「append」
	SliceAppendVal()

	// 複数の要素の値の追加
	MultiSliceAppend()

	// スライスに別のスライスの全要素を追加する
	SliceAppendSlice()

	// Len & Cap
	LenAndCap()

	// Make
	MakeExp()

	// Fale make pattern
	FailMakePattern()

	// Make specify cap
	MakeCap()
}

// ❐❐❐　配列　❐❐❐
//
// Go言語にも配列(array)があるが、配列が直接使われることは多くない
//
// ◇◆◇　配列の宣言方法　◇◆◇
// 『配列の各要素は "指定された型" でなければいけない』(ただし、全要素が常に同じ型だということは意味しない)
// いくつか宣言方法があるが、最初の形式は、次のように『配列の大きさ』と『要素の型』を指定するもの
var xa [3]int // [3] => 配列の大きさ、　int => 要素の型
// ・この宣言で３つの整数を持つ配列が生成される
// ・値が指定されていないので、すべての要素はintのゼロ値(すなわち「0」)を持つ
//
// 値を指定する場合は、次のように「配列リテラル」を用いる
var xar = [3]int{10, 20, 30} // 「10」「20」「30」が配列リテラルに当たる
//
// 途中に「空き」がある配列の値を指定したい場合は、配列のインデックス(添字)とその値を次のように指定する
var xarr = [12]int{1, 5: 4, 6, 10: 100, 15}
//
// 配列の初期化に配列リテラルを使う場合は、大きさを表す整数の代わりに「...」が使える
var xarra = [...]int{10, 20, 30}
//
// ☆「==」と「!=」を使って配列の比較ができる
var ax = [...]int{1, 2, 3}
var ay = [3]int{1, 2, 3}
//
// ☆◆☆ Go言語には一次元配列しかないが、次のようにして多次元配列のように使うことができる ☆◆☆
var xaaa [2][3]int //[[0, 0, 0], [0, 0, 0]]　★[0, 0, 0]が型になる
// これにより、長さ『2』の配列になるが、その型は長さ3の整数配列ということになる
//
// 多くの言語と同様、次のようにインデックスを使用する
func AssignmentArray(){

	var x [3]int

	x[0] = 10 // 先頭要素として10を代入

	fmt.Println(x[0])

}
//
//
// 最後の要素を "超える" インデックスを指定したり、インデックスとして "負の値" を指定したりは出来ない
// 定数あるいはリテラルでこのような指定をした場合は、コンパイル時のエラーになる
//
// ★ インデックスを変数で指定した場合にはコンパイルされるが、値が範囲外になると実行に失敗し、「パニック」となる
//
// 組み込みの関数lenを使って、配列の長さ(大きさ、length)を調べることができる
//
//
// Goでは「配列が直接使われる事は多くない」と書いたが、その理由は「配列にはかなりの制限がある」から
//
// ☆★☆ Goでは配列の大きさを配列の「型」の一部としてみなす
//       つまり、[3]intと宣言した配列と、[4]intと宣言した配列では "型が異なる"
//       ※これは、配列の大きさを指定するのに変数を使えないことを意味する
//			 ◆◇◆ 配列の型は、コンパイル時に決定できなければいけない
//
// 長さが異なる配列への型変換もできない
// 異なる長さの配列を相互に変換できない為、任意の長さの配列を扱う関数が書けない
// さらには、一つの変数に異なる長さの配列を代入することはできない
//
// ☆★ このような制限がある為、配列を使うのは事前にサイズの分かる場合のみになります
// ※配列は、Goでよく使われる「スライス」の「後方支援」の為存在している



// ❐❐❐　スライス　❐❐❐
//
// 一連の値を保持する為のデータ構造が必要なら、ほとんどの場合は「可変長の配列」ともいえる『スライス(slice)』を使うのが正解
// スライスの型にはサイズ(長さ)が含まれないため、配列の様な制限はない
//
// 任意サイズのスライスを処理する関数をかけるし、スライスのサイズは必要に応じて変化する
//
// スライスが配列と大きく異なるところは、『宣言時に大きさを指定しない』
var xtr = []int{10, 20, 30} // スライスリテラルを使用して、3個の整数からなるスライスを生成
//
// [n]あるいは[...]と書くと配列になる。(nは正の整数)
// []だけ書くとスライスになる
//
// 配列リテラル同様、スライスリテラルでもインデックスを指定して要素の一部の値だけを指定できる
var xtrtr = []int{1, 2, 5: 4, 6, 7, 10: 100, 5}
//
// 配列同様、次のようにインデックスを指定して要素の値の取得や変更ができますが、「インデックスの範囲を超える要素の読み書きは "できない"」
func SliceIndex(){

	var x = []int{1,2,3}

	x[0] = 10

	fmt.Println(x[0])

}
//
// また、配列同様、多次元のスライスをシミュレートすることができる
var xslice [][]int
//
// リテラル以外を使ってスライスを宣言すると、配列との違いがよくわかる
var xsliceZero []int // スライスのゼロ値、すなわち「nil」が初期値になる(nilに初期化される)
//
// 値が一つも代入されていないスライスは、スライスのゼロ値、すなわち「nil」が挿入される
// Goにおいて「nil」はいくつかの型において『値がない』事を示す識別子となる
//
// ★☆★ 数値定数と同様、『nilには型がない』
// 　　　 従って、異なる方に代入したり、異なる方と比較したりできる
//
// ◇◆◇ スライスは「比較可能」では "ない"
//       2つのスライスが同じか違うかを判定するのに、「==」や「!=」を使うとコンパイル時のエラーになる
// 　　　 ※ スライスと比較できるのは「nil」だけ
//

// ❐　len　❐
//
// 関数「len」で配列の長さ(サイズ)が分かるが、この関数はスライスにも使える
// そして、nilスライスを「len」に渡すと0を返す

// ❐　append　❐
//
// スライスの要素を増やすには、関数「append」を使用する
func SliceAppend(){

	var x []int

	x = append(x, 10) // 「nil」の後ろに要素の追加

	fmt.Println(x)

}
//
// 「append」には少なくとも2個の引数(任意の型のスライスと『同じ型』の値)をしていする
//
// ☆ 戻り値は『同じ型のスライス』になる
//
// 第1引数に指定されたスライスにappendの結果が代入されている
// 上記の例では「nil」スライスの後ろに追加されているが、次の例のようにすでに要素があるスライスの『最後尾』への追加もできる
func SliceAppendVal() {

	var x = []int{1, 2, 3}

	x = append(x, 4) // 最後尾へ要素の追加

	fmt.Println(x)

}
//
// 複数の値も追加できる
func MultiSliceAppend() {

	var x = []int{1, 2, 3}

	x = append(x, 4, 5, 6)

	fmt.Println(x)

}
//
// ☆ 演算子「...」を使うことでスライスの個々の値を展開することができ、これを使うことでスライスの後に『別のスライスの全要素』を追加できる
func SliceAppendSlice() {

	x := []int{2, 3 ,4}

	y := []int{20, 30, 40}

	x = append(x, y...) // 「y...」でスライス「y」をスライス「x」の要素の最後尾に追加している

	fmt.Println(x)

}
//
// ☆★☆ 「append」の戻り値を変数に代入したりせずに無視すると、コンパイル時のエラーになる
// 				「=」の左辺と右辺の両方にxを書かなければならないのは面倒だと思うかもしれないが、Go言語は「値呼び出し」の言語
// 				❐❐❐ 関数に引数を渡す際には、 "必ず" 『値のコピーが作られてから』渡される
//
// ※※※ 「append」にスライスを渡す際にも、実際に渡されるのは『コピー』
//        「append」はスライスのコピーに値を追加したものを返す
// 				従って、変数の後ろに値を追加したい場合は、戻された結果を。改めてその変数に代入することになる

// ❐ スライスのキャパシティ ❐
//
// ☆ スライスは値が連続したもの
// ☆☆		スライスの各要素はメモリ内の連続した領域に記憶されることで、素早い読み書きができるようになってる
//
// 　　★ 各スライスはキャパシティ(容量)を持っている
// ◆--◆ 『あらかじめ』、一定個の連続する領域が確保されているのである ◆--◆
//　　　　この値はlen関数が返す値よりも大きい場合がある
//
// スライスに「append」を行うと、一個以上の値がスライスの最後に追加される
// 関数「len」の値は、値が一つ追加されるごとに1ずつ増えていく
//
// ★★★　「len」の値がキャパシティ(容量)に達するとそれ以上記憶できなくなる
// 	☆☆☆ このような状態で『さらに』値を追加しようとするとGoのランタイムがより大きなキャパシティを持つスライスの領域を確保する。
// 　　　　最初のスライスのすべての値が新しいスライスにコピーされたうえで、「append」に指定された値が最後に追加され、
// 　　　　その結果できた新しいスライスが「append」から『戻される』ことになる
//
// 📝 Goのランタイム
// 		高レベルの言語はすべて、プログラムの実行の為に一群のライブラリにいぞんしており、Go言語もその例外ではない
// 		Goのランタイムはメモリの確保(アロケーション)やガベージコレクション、並行性のサポート、ネットワーキング、組み込みの型や関数の実装などのサービスを提供している
// 		Goのランタイムは、すべてのGoのバイナリファイルに『コンパイル時』に組み込まれる
// 		バイナリにランタイムが含まれている為Goのプログラムは簡単に配布できる
//
//	「append」によってスライスが大きくなってくると、コピーの為に時間がかかるようになる
// (以前確保されていたメモリのガベージコレクションも必要になる)
// ☆ この為Goのランタイムはスライスのサイズを増やす際には、ある程度の余裕をもってメモリを確保する
//
// ※ Go1.17の段階でのルールはキャパシティが1,024未満の場合は2倍にし、それ以降は25％以上増やすというものになっていた
//
// 「len」はスライスの "現在" のサイズを返すが、組み込み関数「cap」はスライスの現在のキャパシティを返す
// 「len」よりも使われる頻度は少ないだろうが、新しいデータを保持するのに十分なサイズがあるか確認したい場合は「cap」が使える
//
// ☆★ 関数「cap」の引数として配列を渡すこともできるが、この場合は「cap」の値と「len」の値は同じになる
//
func LenAndCap() {
	var x []int
	fmt.Println(x, len(x), cap(x)) // [] 0 0
	x = append(x, 10)
	fmt.Println(x, len(x), cap(x)) // [10] 1 1
	x = append(x, 20)
	fmt.Println(x, len(x), cap(x)) // [10 20] 2 2
	x = append(x, 30)
	fmt.Println(x, len(x), cap(x)) // [10 20 30] 3 4 >>> 「2」のcap(容量)を超えたため、その2倍の「4」のcapがメモリ確保のため割り当てられた
	x = append(x, 40)
	fmt.Println(x, len(x), cap(x)) // [10 20 30 40] 4 4
	x = append(x, 50)
	fmt.Println(x, len(x), cap(x)) // [10 20 30 40 50] 5 8　>>> 「4」のcap(容量)を超えたため、その2倍の「8」のcapがメモリ確保のため割り当てられた

}
//
// スライスが自動的にキャパシティを増やしてくれるのは悪いことではないが、
// あらかじめ『最大のサイズ』が分かっているのなら、それを指定しておいた方が効率的である
// このためには『make』を使う

// ❐ make ❐
//
// これまで見たスライスの宣言方法ではキャパシティをあらかじめ指定することはできないが、
// 「make」を使えば型や長さ、それに(オプションで)キャパシティを指定できる
func MakeExp() {

	x := make([]int, 5)

	fmt.Println(x, len(x), cap(x)) // [0 0 0 0 0] 長さ：5 キャパシティ：5　のintのスライス

	// ☆ 長さは「5」なので、x[0]やx[4]といった表現は有効で、いずれも値として「0」を持っている

}
//
// 「make」を使って生成した配列の要素に値を入れるのに「append」を使用する初心者がいるが、それは誤り
func FailMakePattern() {

	x := make([]int, 5)
	fmt.Println(x, len(x), cap(x)) // [0 0 0 0 0] 5 5

	x = append(x, 10)
	fmt.Println(x, len(x), cap(x)) // [0 0 0 0 0 10] 6 10

	// このコードを実行すると、x[4]の後に新たな要素が追加されることになるから、xの値は
	//
	// 0|1|2|3|4|5
	// 0|0|0|0|0|10
	//
	// となる
	// このとき、len(x)は6、cap(x)は10になる
}
//
// こんどは「make」にキャパシティを指定する
func MakeCap() {

	x := make([]int, 5, 10) // 「長さ:5、キャパシティ:10」のスライス
	fmt.Println(x, len(x), cap(x))

}