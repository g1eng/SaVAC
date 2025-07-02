## 1. サーバーを一括操作する

`SaVAC`のリソース操作は、引数がIDもしくは名前に完全一致するリソースを対象として実施されます。
2つ以上の引数が与えられない限り、デフォルトでは複数リソースの一括操作はできません。
これにより、関係のないリソースを操作してしまうリスクを軽減できる作りとなっています。

しかし、実際の運用では、しばしば特定のグループに属するリソースの一括操作を行いたい場合があるでしょう。


例えば、Blue-Green deploymentで公開サーバーを変更した後、非公開となったサーバーをいったん停止して追加の対処をしたい場合などは、
こうしたケースに該当します。

`SaVAC`では、操作対象リソースの文字列フィルタおよび正規表現フィルタを実装しています。
`-s`ないしは`--search`オプションを指定すると、引数の文字列がリソース名に部分一致するかどうかを判定します。
また、`-E`オプションを指定する場合には、引数の文字列を正規表現としてパースし、リソースの名前がパターンに一致するかどうかを判定します。

これらのフィルタは様々な操作に適用することができ、名前がマッチしたリソースに対する一括操作を行うことができます。

以下の例は、`green`サーバーのデプロイ直後の状態を再現しています。`blue`になっているサーバーを一括停止したい場合には、次のようにすればよいでしょう。

```shell
$ savac list -l
  ID        NAME           STATUS     SPEC   ZONE       IPV4                       IPV6                 ZONE 
xxxxbcd  k8s-blue-1       power_on   4c/3GB  os3   133.2ab.xxx.yyy  2401:2500:102:1a11:133:2ab:xxx.yyy  os3   
xxxxdef  db-blue-1        power_on   3c/2GB  os3   133.2ab.xxx.yyy  2401:2500:102:aa04:133:2ab:xxx.yyy  os3   
xxxxbcd  k8s-blue-2       power_on   4c/3GB  os3   133.2ab.xxx.yyy  2401:2500:102:1a11:133:2ab:xxx.yyy  os3   
xxxxdef  db-blue-2        power_on   3c/2GB  os3   133.2ab.xxx.yyy  2401:2500:102:aa04:133:2ab:xxx.yyy  os3   
xxxxbcd  k8s-green-1      power_on   4c/3GB  os3   133.2ab.xxx.yyy  2401:2500:102:1a11:133:2ab:xxx.yyy  os3   
xxxxdea  db-green-1       power_on   3c/2GB  os3   133.2ab.xxx.yyy  2401:2500:102:aa04:133:2ab:xxx.yyy  os3   
xxxxbad  k8s-green-2      power_on   4c/3GB  os3   133.2ab.xxx.yyy  2401:2500:102:1a11:133:2ab:xxx.yyy  os3   
xxxxaaf  db-green-2       power_on   3c/2GB  os3   133.2ab.xxx.yyy  2401:2500:102:aa04:133:2ab:xxx.yyy  os3   

$ savac stop -E "^(k8s|db)-blue-"

# wait a moment

$ savac info -E "^(k8s|db)-blue-" > /dev/null

$ savac list -l -s blue
  ID        NAME           STATUS     SPEC   ZONE       IPV4                       IPV6                 ZONE 
xxxxbcd  k8s-blue-1       power_off  4c/3GB  os3   133.2ab.xxx.yyy  2401:2500:102:1a11:133:2ab:xxx.yyy  os3   
xxxxdef  db-blue-1        power_off  3c/2GB  os3   133.2ab.xxx.yyy  2401:2500:102:aa04:133:2ab:xxx.yyy  os3   
xxxxbcd  k8s-blue-2       power_off  4c/3GB  os3   133.2ab.xxx.yyy  2401:2500:102:1a11:133:2ab:xxx.yyy  os3   
xxxxdef  db-blue-2        power_off  3c/2GB  os3   133.2ab.xxx.yyy  2401:2500:102:aa04:133:2ab:xxx.yyy  os3   
```

> [!NOTE]
> 個別のサーバーに対する処理を実施したいときには、オプションを指定せずにサーバーIDや名前を指定した方がよいでしょう。

