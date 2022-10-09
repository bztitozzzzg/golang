/*
Exported names
Goでは、最初の文字が大文字で始まる名前は、外部のパッケージから参照できるエクスポート（公開）された名前( exported name )です。 例えば、 Pi は math パッケージでエクスポートされています。

小文字ではじまる pi や hoge などはエクスポートされていない名前です。

パッケージをインポートすると、そのパッケージがエクスポートしている名前を参照することができます。 エクスポートされていない名前(小文字ではじまる名前)は、外部のパッケージからアクセスすることはできません。

コードを実行し、エラーを確認してみましょう。

エラーを修正するために、 math.pi を math.Pi に書き換え、もう一度実行してみてください。
*/
package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(math.Pi)
}
