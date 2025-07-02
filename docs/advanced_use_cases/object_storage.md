# オブジェクトストレージを操作する

「さくらのオブジェクトストレージ」は、さくらインターネットが提供するS3互換なオブジェクトストレージです。

> [!WARNING]
> この機能は[Tier 2 サポート](/POLICY.md)の対象です。
> 一部の引数の組み合わせでは不具合が発生したり、不十分なエラーメッセージしか示されなかったりする場合があります。
> 問題点を見つけた場合には、[イシュー](https://github.com/g1eng/savac/issues)でご連絡ください。


## 0. サイト開設・サイトアカウント作成

あまり知られていませんが、オブジェクトストレージサービスの利用開始のためだけに必要な操作があります。
この操作は、プロジェクトで初めてオブジェクトストレージを利用するときにだけ実行する必要があり、その他の場合には不要です。

```shell
savac o12 site create
savac o12 site account create
```

## 1. サイトアクセスキーを管理する

さくらのオブジェクトストレージには、site-wideなリソースの管理が可能な特権キーが用意されており、「サイトアクセスキー」と呼ばれています。
このサイトアクセスキーはAWSの`S3Administrator`と`IAMAdministrator`に相当する、取り消し不能な特権ロールが割り当てられるものであり、慎重な取り扱いが必要です。
このサイトアクセスキーは、1プロジェクトにつき1つしか作成できません。

サイトアクセスキーで実現できる操作は以下の通りです。

* バケットの作成・削除
* バケットの一覧表示 (*)
* パーミッションの作成・更新・削除
* パーミッション鍵の作成・削除
* バケット内オブジェクトに対する全操作

> バケット一覧表示は、S3クレデンシャルとしてサイトアクセスキーを設定した場合にのみ実現可能なS3互換命令`ListBuckets`です。

プロジェクトでサイトアクセスキーを作成するには、次のように実行します。

```shell
savac o12 account key create
```

> コマンド出力に含まれるアクセスシークレットは、再度取得することができません。
> シークレットの値が分からなくなった場合は、一旦アクセスキーを削除してから再作成する必要があるでしょう。

作成済みのサイトアクセスキーのIDを確認する場合は、次のように実行します。

```shell
savac o12 account key read
```

サイトアクセスキーは次のようにして削除できます。

```shell
savac o12 account key delete
```

> [!IMPORTANT]
> オブジェクトストレージへのアクセス権限があるクラウドAPIトークンやサイトアクセスキーは、
> オブジェクトストレージの特権操作が可能であり、課金可能なリソースを無制限に作成できます。
> 権限キーにはこの特徴がなく、オブジェクトの操作に限った権限のみが付与されます。
> 
> したがって、こうした特権的な操作が運用上不要な場合には、サイトアクセスキーを削除しておいてもよいでしょう。


## 2. バケットを作成する

新規のバケットを作成するには、次のように実行します。

```shell
savac o12 mb sample-bucket-name
```

> なお、さくらのオブジェクトストレージは pay-as-you-go 方式の課金体系を採用しており、バケット単位で最低利用料金が発生します。
コマンド実行直後から月額課金が発生しますので、ご注意ください。

## 3. バケット一覧を表示する

サイトアクセスキーをを認証情報として設定すると、プロジェクト内で利用しているバケットの一覧を取得することができます。

```shell
savac o12 ls
```

## 4. バケットオブジェクトを管理する

SaVACはS3互換APIを通じたオブジェクトの操作をサポートします。
AWS CLI等のS3クライアントが存在しない環境でも、SaVACをインストール済みであれば、
バケットオブジェクトのアップロードやダウンロードを行うことができます。

SaVACは `aws-sdk-go-v2` を利用してPure Goで実装されています。
オブジェクトのバージョニングやMultipart Upload にも対応しており、5GiBを超えるファイルの転送や標準入力からのリダイレクトもサポートしています。

> [!NOTE] 
> 2025年5月現在、AWS CLIの最新バージョンをで「さくらのクラウド オブジェクトストレージ」のバケットオブジェクトを操作する場合、
> 一部の操作でオブジェクトが破損する事象が報告されています。
> SaVACはこの不具合の影響を受けないSDK `ask-sdk-go-v2@1.32.8`を利用しています。

#### 4.1. オブジェクトをアップロードする

```shell
savac o12 put /path/to/local-file.txt s3://sample-bucket-name/the/local/file/path-of-text.txt
savac o12 put -r /path/to/local-dir s3://sample-bucket-name/the/local/file/local-dir
# 最終引数としてディレクトリパスを与えた場合、同一名称のファイルがディレクトリ内に作成されます。
savac o12 put /path/to/other-file.csv s3://sample-bucket-name/the/local/file/
# 大容量ファイルの場合には、標準入力からのリダイレクトを利用できます。
savac o12 put - s3://sample-bucket-name/dev/my-model-000x.safetensor <  my-model-000x.safetensor
```

#### 4.2. オブジェクトをダウンロードする

```shell
savac o12 get s3://sample-bucket-name/the/local/file/other-file.csv  .
savac o12 get s3://sample-bucket-name/the/local/file/path-of-text.txt  my-backup.txt
```

※ savac は、現時点でオブジェクトのバージョニングをサポートしていません。特定のオブジェクトに複数バージョンが存在する場合、最新バージョンがダウンロードされます。

#### 4.3. オブジェクトが存在するかどうかを確認する

```shell
savac o12 check s3://sample-bucket-name/the/local/file/other-file.csv 
```

> [!NOTE] 
> このコマンドは、オブジェクトが存在する場合には終了ステータス`0`のみを返却し、標準出力には何も表示しません

#### 4.4. オブジェクトのメタデータを取得する

> [!WARNING] 
> この機能は動作しません。

```shell
savac o12 info s3://sample-bucket-name/the/local/file/other-file.csv 
```

#### 4.5. オブジェクトの一覧を表示する

```shell
savac o12 ls s3://sample-bucket-name/the/local/file
```


#### 4.6. オブジェクトを削除する

```shell
savac o12 delete s3://sample-bucket-name/the/local/file/other-file.csv 
```

※ savac は、現時点でオブジェクトのバージョニングをサポートしていません。特定のオブジェクトに複数バージョンが存在する場合、完全に削除するために複数回の削除実行が必要になる可能性があります。

