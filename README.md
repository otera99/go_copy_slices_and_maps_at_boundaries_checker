# ツールの概要

https://github.com/knsh14/uber-style-guide-ja/blob/master/guide.md#copy-slices-and-maps-at-boundaries にあるような、引数のスライス（もしくはマップ）で受け取ったスライス（もしくはマップ）がそのままフィールドに保存されている関数がある かつ その関数の引数に渡したスライス（もしくはマップ）があとで要素が変更されてたら警告をだすツールをskeletonを用いて作りました。