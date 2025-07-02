# DNSアプライアンスを操作する

Sakura Internet IaaSでは、マネージドDNSサーバーを提供しています。
このサービスは「DNSアプライアンス」として知られており、日本国内でGSLB分散されたDNSサーバーの REST API を通じた管理を可能とするものです。
ホスト可能なゾーンは、国際化gTLDを除く全てのccTLDおよびgTLDゾーンで、IDNドメイン名のホスティングもサポートされています。

`SaVAC`は、DNSアプライアンスの主要な操作をフルサポートするほか、DNSゾーンファイルのインポート・エクスポートをサポートします。

> [!WARNING] 
> この機能は[Tier 2 サポート](/POLICY.md)の対象です。
> 一部の引数の組み合わせでは不具合が発生したり、不十分なエラーメッセージしか示されなかったりする場合があります。
> 問題点を見つけた場合には、[イシュー](https://github.com/g1eng/savac/issues)でご連絡ください。

## 1. DNSアプライアンスを作成する

新規のDNSアプライアンスを作成するには、次のコマンドを実行します。

```shell
savac dns create test.example.com
```

> [!NOTE] 
> 実際に実行する場合、あなたが管理権限を有するゾーン名を入力してください。
> 一定期間ゾーンの委任が行われなかった場合、DNSアプライアンスが削除される可能性があります。

## 2. DNSアプライアンスの一覧を表示する

```shell
savac dns list
```


## 3. DNSレコードを追加する

`SaVAC`では、DNSアプライアンスのデータをレコード単位で編集することができます。
まず、新規ホストの`A`, `AAAA`, `CNAME`レコードを設定してみます。

```shell
savac dns rr 12736716 --ttl 180 A nekoneko-fast 192.0.2.5
savac dns rr 12736716 --ttl 180 A nekoneko-fast 192.0.2.6
savac dns rr 12736716 --ttl 180 AAAA nekoneko-fast db02::5
savac dns rr 12736716 --ttl 180 AAAA nekoneko-fast db02::6
savac dns rr 12736716 --ttl 180 CNAME nekoneko-slowlife common-blog-svc.example.net.
```

レコードが設定できたら、`dig`コマンド等を用いて、サーバーがRRを認識できていることを確認しておきましょう。

```shell
dig @ns1.gslbX.sakura.ad.jp nekoneko-fast.test.example.com. A +rd
dig @ns1.gslbX.sakura.ad.jp nekoneko-fast.test.example.com. AAAA +rd
dig @ns1.gslbX.sakura.ad.jp nekoneko-slowlife.test.example.com. CNAME +rd
```

> 問い合わせ先の権威DNSサーバーのホスト名は、DNSアプライアンスによって異なります。
> あなたのDNSアプライアンスのFQDNを確認した上で、正しい値を指定してください。

次に、`MX`レコードと`TXT`レコードを設定してみます。

```shell
savac dns rr 12736716 --ttl 180 MX @ 10 nekoneko-fast.test.example.com.
savac dns rr 12736716 --ttl 180 TXT @ "v=spf1 aaaa:nekoneko-fast.test.example.com a:nekoneko-fast.test.example.com ~all"
savac dns rr 12736716 --ttl 180 TXT @ "v=dkim p=none;=r;adsp="
savac dns rr 12736716 --ttl 180 TXT @ "v=dmarc report: m@test.example.com"
```

## 4. DNSレコードをインポート・エクスポートする

多数のDNSレコードを管理する作業は骨が折れるものです。
また、上記のようなRR単位での逐次的なゾーン設定は手間がかかるため、レコード数が多い場合には現実的な管理方法とは言えません。
こうした場合にはSaVACのゾーンファイルインポート・エクスポート機能を利用することで、ゾーン内の多数のレコードを効率的に管理することができます。

#### 4.1. ゾーンファイルをインポートする

作成済みのゾーンファイルをDNSアプライアンスにインポートするには、次のように実行します。
SaVACはゾーンファイル内のSOAレコードを無視します。 また、DNSアプライアンスがサポートしないRRTYPEが指定された場合にエラーを返しますので、ご注意ください。
インポート可能なファイル形式は、JSONまたはBIND形式のゾーンファイルです。

```shell
savac dns import zonefile.txt
```

> [!WARNING] 
> SaVACは、ゾーンファイル内の`ORIGIN`属性を参照して、操作対象のDNSアプライアンスを特定します。
> 同一のゾーンを保持するDNSアプライアンスが複数存在する場合には、インポート対象のアプライアンスが一意に定まりませんので、正しく動作しない可能性があります。


#### 4.2. ゾーンファイルをエクスポートする

あなたやあなたの同僚がゾーンファイルの取り扱いに習熟しており、DNSアプライアンスの設定情報をファイルとして管理したい場合、
リソースレコードの設定内容をゾーンファイルとして保管し、運用担当者が常に参照できるようにしておくことも考慮しておくとよいでしょう。
万が一、担当者のミスや不慮の事故でDNSアプライアンスが消去されてしまっても、ゾーンファイルを用いてリソースレコードを簡単に復元することができます。

以下のように実行すると、現時点でDNSアプライアンスに設定されているリソースレコードの一覧をBIND互換形式のゾーンファイルとしてエクスポートできます。

```shell
# デフォルトでは標準出力に書き出し
savac dns export test.exmaple.com 
# ファイルへの書き込み
savac dns export test.exmaple.com > zonefile.txt
```
