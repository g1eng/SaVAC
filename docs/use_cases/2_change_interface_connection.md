## 2. NICの接続先を変更する

ネットワークアプライアンスやルーターを[さくらのVPS](https://vps.sakura.ad.jp)に実装している場合、内部ネットワークのみに接続する収容サーバーを追加したり、除去したりしたい場合があります。
このような場合には、サーバーのネットワークインターフェース(NIC)が接続するスイッチを変更する必要があります。
さくらのVPSでは、API `v2.x`から、接続先スイッチの変更がフルサポートされ、インターフェース接続先スイッチの変更をAPI経由で実施できるようになりました。

ネットワークの組み替えをAPI経由で実施することにより、ローカルネットワークを用いる基盤の開発・運用の効率化が可能となります。
ただし、接続先の変更を行うためには、NICが紐付いているサーバーが停止している必要がありますので、事前に対象サーバーの電源をシャットダウンしておきましょう。

ここでは、新規に初期化したADサーバーをゲートウェイの配下に収容する操作について、SaVACのみを用いて実施します。

```shell
$ savac list -l 
  ID          NAME         STATUS     SPEC   ZONE       IPV4                       IPV6                 ZONE 
xxxxbcd  sys-gateway      power_on   2c/1GB  os3   133.2ab.xxx.yyy  2401:2500:102:1a11:133:2ab:xxx.yyy  os3   
xxxxdef  sys-adserver     power_on   3c/2GB  os3   133.2ab.xxx.yyy  2401:2500:102:aa04:133:2ab:xxx.yyy  os3   
$ savac stop sys-
$ sleep $wait_a_moment
$ savac info -E sys-
...(output omitted)

$ savac list -l 
  ID          NAME         STATUS     SPEC   ZONE       IPV4                       IPV6                 ZONE 
xxxxbcd  sys-gateway      power_off  2c/1GB  os3   133.2ab.xxx.yyy  2401:2500:102:1a11:133:2ab:xxx.yyy  os3   
xxxxdef  sys-adserver     power_off  3c/2GB  os3   133.2ab.xxx.yyy  2401:2500:102:aa04:133:2ab:xxx.yyy  os3   

$ savac interfaces -E \^sys-
server       NIC id    name       address       switchId   switch   
sys-gateway  4340abcd  eth0  9c:a3:aa:xx:xx:xx  shared     global   
sys-gateway  4340bcde  eth1  9c:a3:aa:xx:xx:xx  -          -        
sys-gateway  4340defa  eth2  9c:a3:aa:xx:xx:xx  -          -        
sys-adserver 3239abcd  eth0  9c:a3:aa:xx:xx:xx  shared     global   
sys-adserver 3239bcde  eth1  9c:a3:aa:xx:xx:xx  -          -
sys-adserver 3239defa  eth2  9c:a3:aa:xx:xx:xx  -          -        

$ savac sw list
(no switches)

$ savac sw create sys-svc-switch

$ savac sw list
   ID           NAME            CODE       ZONE  SERVER INTERFACES  EXTERNAL CONNECTION 
20000abcd  sys-svc-switch  VPSSW20000abcd  os3                                           

$ savac connect 4340bcde 20000abcd

$ savac connect 3239abcd 20000abcd

$ savac sw list
   ID           NAME            CODE       ZONE  SERVER INTERFACES  EXTERNAL CONNECTION 
20000abcd  sys-svc-switch  VPSSW20000abcd  os3   4340bcde,3239abcd 

$ savac interfaces \^sys-
server       NIC id    name       address       switchId       switch   
sys-gateway  4340abcd  eth0  9c:a3:aa:xx:xx:xx  shared     global   
sys-gateway  4340bcde  eth1  9c:a3:aa:xx:xx:xx  20000abcd  -        
sys-gateway  4340defa  eth2  9c:a3:aa:xx:xx:xx  -          -        
sys-adserver 3239abcd  eth0  9c:a3:aa:xx:xx:xx  20000abcd  sys-svc-switch  
sys-adserver 3239bcde  eth1  9c:a3:aa:xx:xx:xx  -          -
sys-adserver 3239defa  eth2  9c:a3:aa:xx:xx:xx  -          -        
```

> [!TIP]
> このケースではネットワークの組み替えが発生するため、 事前にゲートウェイとADサーバーの双方で組み替え後のネットワークに対応する設定を実施しておく必要があります。
>（事後的にNICの設定をしようとすると、結局コントロールパネルにログインしなければなりません)

NFSサーバーが所属するネットワークを変更したい場合には、この手順と同様の方法で実施できます。
