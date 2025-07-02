## 4. WEBサーバーに対する監視設定を追加する

さくらのVPSが標準提供するサーバー監視機能をサーバーの状態把握に役立てましょう。
PrometheusやSaaSのモニタリングツールを用いて統合監視を実施している場合であっても、サーバー障害の際の保険的監視機構を用意することができます。
以下では、一般的なWebサーバーで`ping` `http` `https`監視を設定する方法を見ていきましょう。

監視サブコマンド`savac monitoring`の必須パラメタは以下の通りです。

**savac monitoring** `monitoringType` `serverId` `monitoringName`

**savac mon** `monitoringType` `serverId` `monitoringName`


なお、`http` `https`監視を設定する場合は、追加の`--host`オプション指定が必要です。このオプションは、サーバーへのhttp(s)リクエストの`Host`ヘッダーの値を指定するものです。
その他の指定可能なオプション一覧は、各サブコマンドのドキュメント(`--help`)をご確認ください。

ここでは、ホスト名`next-neko.example.com`でhttps通信を受け付けるWebサーバーと、
`6012/tcp`で`/health`エンドポイントを露出するPrometheus exporterが稼働しているケースを想定します。

適用する通知ポリシとしては、
<u>(1)異常を検知した場合にメール通知</u>し、
<u>(2)ICMP echo-replyを返却できなくなった場合にはDiscordにも通知</u>することとしましょう。

この条件を満たす監視設定を実施するには、以下のコマンドを実行します。

```shell
# ping monitoring
$ savac mon test-server ping  #標準ではメール通知が適用されます
$ savac mon test-server ping 4460131 ping-mon-1 webhook https://discord/channel_key.../channel_secret.../slack
# https monitoring the web server
$ savac mon test-server https --host next-neko.example.com --path /protected --status 401 --sni 4460131 Neko-https
# http monitoring the prometheus exporter
$ savac mon test-server http --host promx01.example.com --port 6012 --path /health --status 200 4460131 Prom-http
```

> [!TIP]
> この監視が正常に動作するためには、VPSの`443/tcp` `6012/tcp` が[さくらインターネットに対して解放](https://manual.sakura.ad.jp/vps/network/servermonitor.html?highlight=%E7%9B%A3%E8%A6%96#id4)されていることが条件となります。
> OSのパケットフィルターだけでなく、VPSサービスのマネージドパケットフィルターが正しく機能しているかどうか、十分に確認しておきましょう。
