# 1. 基本機能

`savac`は「さくらのVPS」をAPI経由で操作するためのフルサポートの提供を目指しています。

本ツールがサポートする基本的な操作は以下の通りです。

以下に、現バージョンの`savac`がサポートするリソース操作の一覧を示します。

## 1.1. サーバー操作

| サブコマンド                                         | 説明                     |
|------------------------------------------------|------------------------|
| server list                                    | サーバーの一覧を表示します(※1)      |
| server info `target`                           | サーバーの詳細情報を表示します        |
| server hostname `target` `[name]`              | サーバーの名前を表示/編集します       |
| server description `target` `[name]`           | サーバーの説明を表示/編集します       |
| server tag `target` `[key]` `[val]`            | サーバーのタグを設定・表示します(※2)   |
| server ptr `target` `[name]`                   | サーバーのPTRレコードを表示/編集します  |
| server interfaces `target`                     | サーバーのインターフェース一覧を表示します  |
| server connect `[-d]` `interfaceId` `switchId` | インターフェースの接続先を変更します(※3) |
| server start `target`                          | サーバーを起動します             |
| server stop `target`                           | サーバーを停止します             |
| server reboot `target`                         | サーバーを強制再起動します          |

## 1.2. NFS操作

| サブコマンド                          | 説明                       |
|---------------------------------|--------------------------|
| nfs list                        | NFSサーバーの一覧を表示します(※1)     |
| nfs info `target`               | NFSサーバーの一覧を表示します         |
| nfs interfaces `target`         | NFSサーバーのインターフェース一覧を表示します |
| nfs connect `target` `switchId` | NFSサーバーの接続先スイッチを変更します    |

## 1.3. スイッチ操作

| サブコマンド                                 | 説明             |
|----------------------------------------|----------------|
| switch create `name` --zone `zoneCode` | スイッチを作成します     |
| switch delete `switchId`               | スイッチを削除します     |
| switch name `switchId` `[name]`        | スイッチ名を表示/変更します |
| switch list                            | スイッチ一覧を表示します   |

## 1.4. VPS監視設定

| サブコマンド                                                      | 説明                     |
|-------------------------------------------------------------|------------------------|
| monitoring list                                             | サーバー監視設定の一覧を表示します      |
| monitoring info `target` `[monitoringId]`                   | 個別サーバーの監視設定の詳細を表示します   |
| monitoring ping `target` `name` `[email\|webhook]` `[url]`  | サーバーのping監視設定を新規作成します  |
| monitoring tcp `target` `name` `[email\|webhook]` `[url]`   | サーバーのtcp監視設定を新規作成します   |
| monitoring http `target` `name` `[email\|webhook]` `[url]`  | サーバーのhttp監視設定を新規作成します  |
| monitoring https `target` `name` `[email\|webhook]` `[url]` | サーバーのhttps監視設定を新規作成します |
| monitoring smtp `target` `name` `[email\|webhook]` `[url]`  | サーバーのsmtp監視設定を新規作成します  |
| monitoring pop3 `target` `name` `[email\|webhook]` `[url]`  | サーバーのpop3監視設定を新規作成します  |
| monitoring delete `serverId` `[monitoringId]`               | 指定サーバーの監視設定を削除します      |

* ※1 APIエンドポイントからのサーバー一覧の戻り値はキャッシュされた値です。電源状態を更新するには`savac info server-name`
  などとして、キャッシュの内容を更新してください。(`info` サブコマンドと正規表現オプション`-E`等を併用して複数リソースを参照する場合、対象サーバーのすべてに対する再帰的アクセスが発生します。
  多数のサーバー契約がある場合にはレート制限に抵触する場合がありますので、ご注意ください。
  API接続のレートリミットについては、[上流のドキュメント](http://manual.sakura.ad.jp/vps/api/api-doc/index.html) を参照してください。)
* ※2 タグは改行区切りで説明フィールドに格納されることに注意してください
* ※3 インターフェースをインターネットに接続する際は`switchId`に`1`を指定してください。`-d`
  オプションを指定するとインターフェースをネットワークから切断できます。

## 1.5. APIキー関連機能

| サブコマンド                                                               | 説明                       |
|----------------------------------------------------------------------|--------------------------|
| apikey create --role `roleId` `name`                                 | APIキーを作成して、指定したロールに紐付けます |
| apikey list                                                          | APIキーの一覧を表示します           |
| apikey rotate `keyId`                                                | APIキーを更新します              |
| apikey delete `keyId`                                                | APIキーを削除します              |
| role list                                                            | ロールの一覧を表示します             |
| role create `[--perm permExpr]` `[--resource resourceExpr]` roleName | ロールを作成します                |
| role update `[--perm permExpr]` `[--resource resourceExpr]` roleName | ロールを更新します                |
| role delete `roleId`                                                 | ロールを削除します                |
| perm list                                                            | 権限の一覧を表示します              |

* ※1 ロール作成・更新時に指定できる`--perm`および`--resource`オプションの引数は正規表現として解釈されます。ロールはパターンに一致する権限・リソースに紐付けられます。
* ※2 1つ以上のAPIキーに紐付いているロールは削除できません。このようなロールを削除する場合、事前にAPIキーを削除してください。

# 2. 高度な機能

`savac`は「さくらのクラウド」のIaaSリソース操作を部分的にサポートします。

IaaS技術はコモディティ化した技術です。したがって、ここでいう高度な機能とは「<ruby>
雲は高い場所にある<rt>CLOUD ON ANYWHERE HIGH</rt></ruby>」という意味です。

以下のIaaS操作機能の一部は`usacloud`の簡易的なコピーですが、`usacloud`
や[さくらのクラウド Terraform Provider](https://github.com/sacloud/terraform-provider-sakuracloud)
がサポートしていないリソース操作を中心に実装しているため、これらのツールと補完的に運用することができるでしょう。

(VPS基盤も含めた Sakura Internet IaaS
のすべてのリソースをIaCで管理できるようになれば、こうした機能群は不要となるかもしれません。)

## 2.1. DNSアプライアンス操作機能

DNS操作機能では、JSONおよびBIND形式でゾーンファイルを管理することが可能です。

| サブコマンド                   | 説明                                                                                      |
|--------------------------|-----------------------------------------------------------------------------------------|
| dns create `zoneName`    | ゾーンを指定してDNSアプライアンスを作成する                                                                 |
| dns read `pattern`       | `pattern`で指定したDNSアプライアンスの設定値を表示する<br/>`--output-type=text`を指定した場合はゾーン情報をzonefile形式で出力する |
| dns list                 | DNSアプライアンスの一覧を表示する                                                                      |
| dns export `pattern`     | `pattern`で指定したDNSアプライアンスのゾーン情報をzonefile形式で出力する                                          |
| dns import -f `filePath` | DNSゾーンを含むDNSアプライアンス設定情報をインポートする<br/>JSONまたはzonefile形式でのインポートが可能                         |

## 2.2. コンテナレジストリ操作機能

コンテナレジストリの作成・削除に加えて、ユーザーの管理機能をサポートします。

| サブコマンド                                       | 説明                    |
|----------------------------------------------|-----------------------|
| co create `name`                             | コンテナレジストリを作成する        |
| co list                                      | コンテナレジストリ一覧を表示する      |
| co delete `registryId`                       | コンテナレジストリを削除する        |
| co user add `registryId` -u `user` -p `pass` | コンテナレジストリのユーザーを追加する   |
| co user list                                 | コンテナレジストリのユーザー一覧を表示する |
| co user delete `registryId` -u `user`        | コンテナレジストリのユーザー一覧を表示する |

## 2.3. オブジェクトストレージ操作機能

オブジェクトストレージの作成・削除を含む全機能のAPI操作をサポートします。
ただし、オブジェクトストレージの一覧を表示するには、特権アカウントキーを認証情報として設定する必要があります。
また、パーミッションの作成・更新機能を利用するためには、特権アカウントキーもしくはオブジェクトストレージサービスの管理権限がある[クラウドアクセストークン](CREDENTIALS.md)を認証情報として設定する必要があります。

| サブコマンド                                       | 説明                                                                                                                                                     |
|----------------------------------------------|--------------------------------------------------------------------------------------------------------------------------------------------------------|
| o12 sites                                    | ストレージサイト一覧を表示する                                                                                                                                        |
| o12 access-key create                        | 特権アクセスキーを発行する                                                                                                                                          |
| o12 access-key list                          | 特権アクセスキー一覧を表示する                                                                                                                                        |
| o12 access-key delete `keyId`                | 特権アクセスキーを削除する                                                                                                                                          |
| o12 permission create `name`                 | 権限を作成する <br/>--rw `name` read+write権限を付与するバケットをカンマ区切りで指定<br/>--ro `name` readonly権限を付与するバケットをカンマ区切りで指定<br/> --wo `name` writeonly権限を付与するバケットをカンマ区切りで指定 |
| o12 permission update `name`                 | 権限を更新する (--rw, --ro, --wo ルールを指定可能)                                                                                                                    |
| o12 permission delete `name`                 | 権限を削除する                                                                                                                                                |
| o12 permission list                          | 権限一覧を表示する                                                                                                                                              |
| o12 permission key create                    | 権限キーを作成する                                                                                                                                              |
| o12 permission key list                      | 権限キー一覧を表示する                                                                                                                                            |
| o12 permission key delete `permName` `keyId` | 権限キーを削除する                                                                                                                                              |
| o12 mb `bucket-name`                         | 新規のバケットを作成する                                                                                                                                           |
| o12 ls                                       | バケット一覧を表示する                                                                                                                                            |
| o12 rb `bucket-name`                         | バケットを削除する                                                                                                                                              |
| o12 ls `s3://bucket-name/...`                | プレフィクスの一致するオブジェクトの一覧を表示する                                                                                                                              |
| o12 put `src` `s3://bucket-name/dst`         | コンテンツをバケット内にコピーする。                                                                                                                                     |
| o12 get `s3://bucket-name/src` `dst`         | バケット内のコンテンツをダウンロードする。                                                                                                                                  |
| o12 check `s3://bucket-name/target`          | オブジェクトの存在を検証する。(正常終了時のステータスは`0`)                                                                                                                       |

## 2.4. ウェブアクセラレーター操作機能

ウェブアクセラレーターの管理機能全般をサポートします。

| サブコマンド                                                             | 説明                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                       |
|--------------------------------------------------------------------|--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| wa list                                                            | サイト一覧を表示する                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                               |
| wa create `name`                                                   | サイトを新規作成する<br/>   `--domain-type string` ドメインの種類。`own_domain`か`subdomain`を指定。<br/>`--domain string` 独自ドメイン名<br/>`--request-protocol string`, リクエストプロトコル (default: "http+https")<br/>`--default-cache-ttl int`     デフォルトのキャッシュ保持期間（秒） (default: 0)<br/>`--cors string` CORS許可オリジン (can be multiply used)<br/>`--vary` Varyサポートを有効化する (default: false)<br/>`--accept-encoding string`  サポートする圧縮形式の種類 (default gzip)<br/>`--origin-type string` オリジンの種類。`web`または`bucket`. (default: "web")<br/>`--origin-protocol string` オリジンプロトコル。`http`または`https`。 (default: "https")<br/>`--origin web`  オリジンのホスト名またはIPアドレス (option for a web origin)<br/>`--host-header string` オリジンへの接続で使用するHostヘッダ (option for a web origin)<br/>`--bucket bucket` オリジンバケット名 (option for a bucket origin)<br/>`--endpoint bucket` S3エンドポイント (option for a bucket origin) (default: "s3.isk01.sakurastorage.jp") <br/>`--region bucket` S3リージョン (option for a bucket origin) <br/>`--access-key string` アクセスキー (option for a bucket origin) <br/>`--access-secret string` シークレットアクセスキー (option for a bucket origin)<br/>`--docindex` ドキュメントインデックスを有効化する (option for a bucket origin) (default: false) |
| wa read   `siteId`                                                 | サイト詳細を表示する                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                               |
| wa update `siteId`                                                 | サイトを更新する<br/> <br/>`--request-protocol string`, リクエストプロトコル (default: "http+https")<br/>`--default-cache-ttl int`     デフォルトのキャッシュ保持期間（秒） (default: 0)<br/>`--cors string` CORS許可オリジン (can be multiply used)<br/>`--vary` Varyサポートを有効化する (default: false)<br/>`--accept-encoding string`  サポートする圧縮形式の種類 (default gzip)<br/>`--origin-type string` オリジンの種類。`web`または`bucket`. (default: "web")<br/>`--origin-protocol string` オリジンプロトコル。`http`または`https`。 (default: "https")<br/>`--origin web`  オリジンのホスト名またはIPアドレス (option for a web origin)<br/>`--host-header string` オリジンへの接続で使用するHostヘッダ (option for a web origin)<br/>`--bucket bucket` オリジンバケット名 (option for a bucket origin)<br/>`--endpoint bucket` S3エンドポイント (option for a bucket origin) (default: "s3.isk01.sakurastorage.jp") <br/>`--region bucket` S3リージョン (option for a bucket origin) <br/>`--access-key string` アクセスキー (option for a bucket origin) <br/>`--access-secret string` シークレットアクセスキー (option for a bucket origin)<br/>`--docindex` ドキュメントインデックスを有効化する (option for a bucket origin) (default: false)                                                                                              |
| wa enable `siteId`                                                 | サイトを有効化する                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                |
| wa disable `siteId`                                                | サイトを無効化する                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                |
| wa origin-guard create `siteId`                                    | サイトのオリジンガードトークンを発行する<br/>`--next` 次期オリジンガードトークンを発行し、移行準備モードに移行する。<br/>      準備モードで再度`create`コマンドを実行すると、次期オリジンガードトークンが設定される。                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              |
| wa origin-guard read `siteId`                                      | サイトのオリジンガードトークンを取得する                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     |
| wa origin-guard delete `siteId`                                    | サイトのオリジンガードトークンを削除する<br/>`--next` オリジンガードトークンの移行モードを解除し、次期トークンのみを削除する。                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                   |
| wa one-time-url secret `siteId` `secret`                           | サイト全体ワンタイムURLのためのシークレットを指定する。<br/>`--purge`  サイト全体ワンタイムURLを解除する。                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         |
| wa one-time-url generate `<url>`                                   | サイトのワンタイムURL文字列を生成する<br/>`--expired timeSpec`<br/>  `timeSpec`: 1min, 3hr, 1day, 1745489388 (timestamp from EPOCH)                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                       |
| wa certificate import --file `certFile` [--key `keyFile`] `siteId` | サイトのTLS証明書を更新する                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                          |
| wa certificate auto-renew [--disable] `siteId`                     | サイトのLet's Encrypt自動更新を設定する<br/>--disable 自動更新を停止する。                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                      |
| wa purge-cache `url...`                                            | キャッシュを削除する。<br/>--all 指定したドメインの全てのキャッシュを削除する。                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                            |
| wa acl read `siteId`                                               | サイトACL一覧を取得する。                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           |
| wa acl apply `siteId`                                              | サイトACLを追加する。<br/>--allow `prefixes` 許可対象のIPv4プレフィクスをカンマ区切りで指定<br/>--deny `prefixes` 拒否対象のIPv4プレフィクスをカンマ区切りで指定                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                            |
| wa acl clear `siteId`                                              | サイトACLを除去する。                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                             |
