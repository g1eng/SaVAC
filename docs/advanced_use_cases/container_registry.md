# コンテナレジストリを操作する

「コンテナレジストリ」は、さくらのクラウドでOCI互換なコンテナイメージを保管できるマネージドサービスで、
dockerやpodmanでビルドしたコンテナイメージやWASMワークロードを保管したり、配布したりすることができます。
レジストリに保管する成果物は公開することもできますし、非公開にすることで認証済みのユーザーのみにアクセスさせることもできます。
2025年6月現在、このサービスβ版として提供されており、無料で利用できます。

この記事ではsavacでコンテナレジストリを操作する方法について開設します。

## 1. コンテナレジストリを作成する

次のようにすると、さくらのクラウドのサブドメインを利用するコンテナレジストリを作成できます。
このパターンでは、独自のドメインを用意せずに利用できるため、手軽です。

```shell
$ savac container-registry create mussle-honke
{
  "AccessLevel": "none",
  "Availability": "available",
  "CreatedAt": "2025-06-18T12:11:42+09:00",
  "Description": "",
  "FQDN": "mussle-honke.sakuracr.jp",
  "ID": 113701XXXXXX,
  "IconID": 0,
  "ModifiedAt": "2025-06-18T12:11:42+09:00",
  "Name": "mussle-honke",
  "SettingsHash": "563b2646670924XXXXXXXXXXXXXXXXXX",
  "SubDomainLabel": "mussle-honke",
  "Tags": [],
  "VirtualDomain": ""
}
```

独自ドメインを用いてコンテナレジストリにアクセスしたい場合には、次のように`--domain`オプションを指定します。
（ただし、この形態でコンテナレジストリをデプロイするとHTTPSがサポートされません。）

```shell
$ savac container-registry create --domain someiyoshino-2.example.jp next-example
{
  "AccessLevel": "none",
  "Availability": "available",
  "CreatedAt": "2025-06-18T12:27:53+09:00",
  "Description": "",
  "FQDN": "next-example.sakuracr.jp",
  "ID": 113701XXXXXX,
  "IconID": 0,
  "ModifiedAt": "2025-06-18T12:27:53+09:00",
  "Name": "next-example",
  "SettingsHash": "a7e621dd4bb6b7XXXXXXXXXXXXXXXXXX",
  "SubDomainLabel": "next-example",
  "Tags": [],
  "VirtualDomain": "someiyoshino-2.example.jp"
}
```

コンテナレジストリ作成時には、レジストリに対するアクセス権限を指定することができます。
デフォルトでは`none`が指定され、認証済みユーザーのみがアクセスできますが、
`readonly`を指定することで、アップロードしたイメージを誰でもpullできるようになります。

```shell
$ savac co create --permission readonly ukeyasuo3
{
  "AccessLevel": "readonly",
  "Availability": "available",
  "CreatedAt": "2025-06-18T12:32:56+09:00",
  "Description": "",
  "FQDN": "ukeyasuo3.sakuracr.jp",
  "ID": 1137016XXXXX,
  "IconID": 0,
  "ModifiedAt": "2025-06-18T12:32:56+09:00",
  "Name": "ukeyasuo3",
  "SettingsHash": "a0e1ac9ec678fa3e4XXXXXXXXXXXXXXX",
  "SubDomainLabel": "ukeyasuo2",
  "Tags": [],
  "VirtualDomain": ""
}
```

## 2. コンテナレジストリの認証情報を設定する

アクセス権限として`none`を指定したコンテナレジストリのイメージを利用したり、作成したイメージをアップロードするためには、
コンテナレジストリのユーザーを作成する必要があります。
ユーザーIDとパスワードをセットで指定し、コンテナレジストリへのアクセス時にはこれらの認証情報を指定します。

以下のようにすると、コンテナレジストリの読書権限のある認証情報を設定できます。

```shell
$ savac co user add \
    --user ari-33-taigun \
    --password L2m3Bx0Ja-aLoyRn \
    --permission readwrite ukeyasuo3 
```

読込専用の権限を設定するには、次のようにします。

```shell
$ savac co user add \
    --user kama-kiri33 \
    --password NZXc9vN214-AZSdh \
    --permission readonly ukeyasuo3
```

設定したユーザーの一覧は次のコマンドで確認できます。

```shell
savac co user list 113701684393
{
  "Users": [
    {
      "Permission": "readwrite",
      "UserName": "ari-33-taigun"
    },
    {
      "Permission": "readonly",
      "UserName": "kama-kiri33"
    }
  ]
}
```

dockerでコンテナレジストリにイメージをpushする場合には、認証情報を次のように設定します。
非公開レジストリの場合には、pullする場合にも同様の認証情報設定が必要になります。

```shell
$ docker login ukeyasuo3.sakuracr.jp
Username: ari-33-taigun
Password: [intput password]
Login Succeeded
$ docker build -t  ukeyasuo3.sakuracr.jp/sample-img:tag .
 [+] Building 14.3s (14/14) FINISHED                                                             docker:desktop-linux
 => [internal] load build definition from Dockerfile                                                            0.0s
 ....(output omitted) ....
 => => exporting attestation manifest sha256:4c47384ca2fa738aa91af5286d0da6ab74d6fc976316ebc4c27e85fafd1b9d1c   0.0s 
 => => exporting manifest list sha256:6bca6896b1feb0c1adcb300b2e777f45572b00a4a92ef3d6da22841f8164728d          0.0s
 => => naming to ukeyasuo3.sakuracr.jp/sample-img:tag                                                           0.0s
 => => unpacking to ukeyasuo3.sakuracr.jp/sample-img:tag                                                        0.1s
$ docker push  ukeyasuo3.sakuracr.jp/sample-img:tag  
The push refers to repository [ukeyasuo3.sakuracr.jp/sample-img]
6ce5a687cb63: Pushed 
dc34890d8ac1: Pushed 
tag: digest: sha256:6bca6896b1feb0c1adcb300b2e777f45572b00a4a92ef3d6da22841f8164728d size: 855
```

最後に、ユーザーを削除するには次のようにします。ユーザー名を指定するのを忘れないで下さい。

```shell
$ savac co user del --user kama-kiri33 113701684393
```

## 3. コンテナレジストリを削除する

利用しなくなったコンテナレジストリは削除しておきましょう。
savacでは、次のようにしてコンテナレジストリを削除します。

```shell
$ savac co delete ukeyasuo3
# もしくはIDで削除する
$ savac co delete 113701XXXXX
```

