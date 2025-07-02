---
layout: default
title: クラウドリソースの操作
permalink: /advanced_use_cases/
---

SaVAC は「さくらのクラウド」のIaaSリソース操作を部分的にサポートします。

IaaS技術はコモディティ化した技術です。したがって、ここでいう高度な機能とは「<ruby>
雲は高い場所にある<rt>CLOUD ON ANYWHERE HIGH</rt></ruby>」という意味です。

SaVAC のIaaS操作機能の一部は`usacloud`の簡易的なコピーですが、`usacloud`
や[さくらのクラウド Terraform Provider](https://github.com/sacloud/terraform-provider-sakuracloud)
がサポートしていないリソース操作を中心に実装しているため、これらのツールを補完する上で有用です。


[DNSアプライアンスを操作する](dns.md)

[オブジェクトストレージを操作する](object_storage.md)

[ウェブアクセラレーターを操作する](webaccel.md)
