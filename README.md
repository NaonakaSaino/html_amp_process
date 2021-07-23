# html_amp_process

[MPcloudAMP化マニュアル](https://docs.google.com/presentation/d/1qm9DkH-dbQ3JrsmlEDJX0tlOcYdoTgMItDGKomsICAs/edit#slide=id.gd513a6735d_0_0)に基づいて実装

【実装済み。】<br>
・img.lazy要素の変換<br>
・amp-imgの必須属性を記述する<br>
・inline styleの削除<br>
・inline scriptの削除<br>

【未実装】<br>
・imgのパスをpapillonのメディアのIDに揃える<br>
 →現在はcdnの画像パスのみシェルに出力する実装。<br>
 →画像アップロードのAPIを叩けばここも自動化できる？<br>


<br>
※他メディアで使う場合の留意<br>

・[インラインスタイルの対応表](https://docs.google.com/spreadsheets/d/1UHE5LBzpFD2l5BVTXnPJAKq8qUGd1Mv9u2MnyR2zCVA/edit#gid=1702748687)に基づいて、現在はamp/lib/style_to_class.go内にMoneyTimes用の変換用mapを書いている。<br>
・amp/lib/html_divide.go内の125行目付近に、papillonの画像パスであることの判別に使う文字列として"admin.moneytimes.jp"を設定している。

<br>

※エラー<br>
・元のhtmlに、同一タグ内にclass属性が複数回設定されている箇所があると、加工後もそのまま。
