# Folios

[![Product Name](image.png)](https://www.youtube.com/watch?v=rwhL1l9rNwk)

## 製品概要
### IT Engineer x Tech
### 背景（製品開発のきっかけ、課題等）
- プログラマがポートフォリオサイトをデザインするのは難しい
- 他の方のQiita,GitHub,ブログなどが頭の中で関連付かない
- 手軽に名刺がわりにできたら良い

### 製品説明（具体的な製品の説明）

### 特長

#### 1. Githubのリポジトリをマイニングしてポートフォリオを作成する

#### 2. Qiitaやはてなブログの更新情報を表示する

#### 3. Githubでフォローしている人を一覧表示する


### 解決出来ること
- デザインが苦手でも見栄えの良いポートフォリオサイトを作成することができる

### 今後の展望
今回は実現できなかったが、今後改善すること、どのように展開していくことが可能かについて記載をしてください。
- Webだけでなくpdfに書き出し、本当の名刺のように使える
- テンプレートの切り替え

## 開発内容・開発技術
### 活用した技術
#### API・データ
今回スポンサーから提供されたAPI、製品などの外部技術があれば記述をして下さい。
* Github REST API
* Hatena RSS
* Qiita API v2

#### フレームワーク・ライブラリ・モジュール
* Vue.js
* React.js
* Gin
* goQuery
* gofeed
* Bootstrap 4
* jQuery
* ASP.NET Core
* Azure Web Apps
* C#
* XUnit, Moq
* Azure Devops

### 独自開発技術（Hack Dayで開発したもの）
#### 2日間に開発した独自の機能・技術
* GitHub,Qiitaの記事を取得してそれぞれ人気のものが上位に来るような表示とした
* 本来別々に存在するサイトを1つのページでみられるようにした
* 友達の情報を簡単に見られるようにした

#### Github などの OAuth を利用する者に対する認証基盤を作成した。

実運用で使われることを考慮した設計を盛り込んだ具体的には、
- 重要なセキュリティ情報については環境変数より取得するようにした
- Repository パターンや DI(依存性の注入）などを利用して拡張性の高い構成にした
- 実環境でいつでも動かせるようにするために CI/CD などを導入した。 実環境にすでにデプロイされているというのは評価して欲しいです。

### 注意事項
- 認証基盤, 設定保存DB とのやりとりは server/github のブランチにSubRepository の形で参照が貼ってあります。最新のを見てくださいませ。
