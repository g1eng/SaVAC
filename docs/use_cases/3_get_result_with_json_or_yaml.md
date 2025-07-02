## 3. JSON, YAML, Text形式の出力を使い分ける

VPS APIの生のレスポンスはすべてJSONです。オリジナルのJSONが欲しい時には、`-j`オプションを使いましょう。
`-y`オプションを用いるとYAMLとしての出力も可能です。

```shell
$ savac list -j
{
  "count": 12,
  "next": null,
  "previous": null,
  "results": [
    {
      "contract": {
        "plan_code": 3440,
        "plan_name": "さくらのVPS(v5)  2G IK01 ストレージ変更オプション付き",
        "service_code": "11000abcdefg"
      },
      "cpu_cores": 3,
      "description": "",
      "id": 442abcd,
      "ipv4": {
        "address": "133.2ab.xxx.yyy",
        "gateway": "133.2ab.xxx.p",
        "hostname": "ik1-xxx-yyyyy.vs.sakura.ne.jp",
        "nameservers": [
          "133.2ab.0.a",
          "133.2ab.0.b"
        ],
        "netmask": "255.255.xxx.0",
        "ptr": "ik1-xxx-yyyyy.vs.sakura.ne.jp"
      },
      "ipv6": {
        "address": "2401:2500:1ac:1111:133:2ab:xxx:yyy",
        "gateway": "fe80::s:t",
        "hostname": "133.2ab.xxx.yyy.v6.sakura.ne.jp",
        "nameservers": [
          "2401:2500::p"
        ],
        "prefixlen": c,
        "ptr": "133.2ab.xxx.yyy.v6.sakura.ne.jp"
      },
      "memory_mebibytes": 2048,
      "name": "samorau-extreme-high-capacity-1",
      "options": [],
      "power_status": "power_off",
      "service_type": "linux",
      "storage": [
        {
          "port": 0,
          "size_gibibytes": 12000,
          "type": "ssd"
        }
      ],
      "version": "v5",
      "zone": {
        "code": "is1",
        "name": "石狩第1"
      }
    },
    ...
```
