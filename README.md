<h1 align="center">GeoIP 增强版：自由定制多种格式 GeoIP 文件</h1>

<p align="center">
  <img src="./assets/hero.png" alt="GeoIP project hero image">
</p>

<div align="center">

<a href="https://trendshift.io/repositories/5833" target="_blank"><img src="https://trendshift.io/api/badge/repositories/5833" alt="Loyalsoldier%2Fgeoip | Trendshift" style="width: 250px; height: 55px;" width="250" height="55"/></a>

</div>

<div align="center">
<a href="https://deepwiki.com/Loyalsoldier/geoip" target="_blank"><img src="https://deepwiki.com/badge.svg" alt="Ask DeepWiki badge"></a> <a href="https://www.jsdelivr.com/package/gh/Loyalsoldier/geoip" target="_blank"><img src="https://data.jsdelivr.com/v1/package/gh/Loyalsoldier/geoip/badge?style=rounded" alt="jsdelivr stats badge"></a>

<a href="https://shields.io" target="_blank"><img src="https://img.shields.io/github/downloads/Loyalsoldier/geoip/total?logo=github" alt="GitHub Downloads badge (all assets, all releases)"></a> <a href="https://shields.io" target="_blank"><img src="https://img.shields.io/github/downloads/Loyalsoldier/geoip/latest/total?logo=github" alt="GitHub Downloads badge (all assets, latest release)"></a>
</div>

## 名词解析

**GeoIP**，意为 ***IP geographic location***，即 IP 地址所对应的地理位置信息，例如所属的国家、地区等。GeoIP 数据文件则存储着 IP 地址所对应的地理位置信息。

**GeoIP**, which stands for ***IP geographic location***, refers to the geographic location information associated with an IP address, such as the country or region. And GeoIP data files store the geographic location information corresponding to IP addresses.

## 项目简介

本项目每天自动生成多种格式 GeoIP 数据文件，同时提供命令行界面（CLI）工具供用户自行定制 GeoIP 数据文件，包括但不限于 V2Ray `dat` 格式文件 `geoip.dat`、MaxMind-compatible `mmdb` 格式文件 `Country.mmdb`、sing-box `SRS` 格式文件、mihomo `MRS` 格式文件、Clash ruleset 规则文件、Surge ruleset 规则文件、Nginx allow list（允许访问列表）和 Nginx deny list（拒绝访问列表）。

This project releases various formats of GeoIP files automatically every day, and provides a command line interface (CLI) tool for users to customize their own GeoIP files, including but not limited to V2Ray `dat` format file `geoip.dat`, MaxMind-compatible `mmdb` format file `Country.mmdb`, sing-box `SRS` format file, mihomo `MRS` format file, Clash ruleset file, Surge ruleset file, Nginx allow list and Nginx deny list.

## GeoIP 类别

本项目默认使用 [IPInfo Lite 数据](https://ipinfo.io/developers/ipinfo-lite-database)生成各个国家和地区的 GeoIP 文件。发布的 `Country*.mmdb` 文件使用 IPInfo Lite 作为数据源，但输出为 MaxMind GeoLite2 Country 兼容字段。类别有：

- `geoip:cn`（`GEOIP,CN`）：**中国大陆**（Mainland China）
- `geoip:hk`（`GEOIP,HK`）：**香港**（Hong Kong）
- `geoip:mo`（`GEOIP,MO`）：**澳门**（Macau）
- `geoip:tw`（`GEOIP,TW`）：**台湾**（Taiwan）
- `geoip:us`（`GEOIP,US`）：**美国**（America）
- `geoip:jp`（`GEOIP,JP`）：**日本**（Japan）
- `geoip:kr`（`GEOIP,KR`）：**韩国**（Korea）
- `geoip:sg`（`GEOIP,SG`）：**新加坡**（Singapore）
- `geoip:mm`（`GEOIP,MM`）：**缅甸**（Myanmar）
- `geoip:ir`（`GEOIP,IR`）：**伊朗**（Iran）
- `geoip:ru`（`GEOIP,RU`）：**俄罗斯**（Russia）
- `geoip:by`（`GEOIP,BY`）：**白俄罗斯**（Belarus）
- `geoip:tm`（`GEOIP,TM`）：**土库曼斯坦**（Turkmenistan）
- `geoip:private`（`GEOIP,PRIVATE`）：**内网 IP 地址**、**保留 IP 地址**等特殊 IP 地址的集合
- 更多可用的类别（以两位英文字母表示国家或地区），请查看：[https://www.iban.com/country-codes](https://www.iban.com/country-codes)

### 与 IPInfo Lite 原始数据的区别

本项目对 IPInfo Lite 原始数据做了如下**修改**和**新增**：

- 中国大陆 IPv4 地址数据使用 [@gaoyifan/china-operator-ip](https://github.com/gaoyifan/china-operator-ip/blob/ip-lists/china.txt)
- 中国大陆 IPv6 地址数据使用 [@gaoyifan/china-operator-ip](https://github.com/gaoyifan/china-operator-ip/blob/ip-lists/china6.txt)
- 新增 ASN 类别优先使用 IPInfo Lite 的 `as_domain` / `as_name` 元数据归类，hard-coded ASN 清单只作为补充；两者冲突时以 IPInfo 元数据为准
- 新增类别（方便有特殊需求的用户使用）：
  - `geoip:cloudflare`（`GEOIP,CLOUDFLARE`）
  - `geoip:cloudfront`（`GEOIP,CLOUDFRONT`）
  - `geoip:facebook`（`GEOIP,FACEBOOK`）
  - `geoip:fastly`（`GEOIP,FASTLY`）
  - `geoip:google`（`GEOIP,GOOGLE`）
  - `geoip:netflix`（`GEOIP,NETFLIX`）
  - `geoip:spotify`（`GEOIP,SPOTIFY`）
  - `geoip:telegram`（`GEOIP,TELEGRAM`）
  - `geoip:twitter`（`GEOIP,TWITTER`）
  - `geoip:tor`（`GEOIP,TOR`）

其中 Spotify 类别包含 RIPE Database 登记给 Spotify AB / Spotify Technology S.A. 的 IP 段，并使用 IPInfo 元数据和 `AS8403` 作为补充。

## 下载地址与使用方法

> [!NOTE]
> 如果无法访问域名 `raw.githubusercontent.com`，可以使用第二个地址 `cdn.jsdelivr.net`。
> 如果无法访问域名 `cdn.jsdelivr.net`，可以将其替换为 `fastly.jsdelivr.net`。
>
> *.sha256sum 为校验文件。

<br/>

本项目发布的所有 GeoIP 文件，请查看 [release 分支](https://github.com/arsity/geoip/tree/release)。以下是部分格式 GeoIP 文件的下载地址和使用方法：

### V2Ray dat 使用方法

<details>
  <summary>点击查看在 <b>V2Ray</b> 和 <b>Xray-core</b> 中的使用方法</summary>
  <br/>
  <p>需要先下载 <code>.dat</code> 格式文件，并放置在程序目录内。</p>

```json
"routing": {
  "rules": [
    {
      "type": "field",
      "outboundTag": "Direct",
      "ip": [
        "geoip:cn",
        "geoip:private",
        "ext:cn.dat:cn",
        "ext:private.dat:private",
        "ext:geoip-only-cn-private.dat:cn",
        "ext:geoip-only-cn-private.dat:private"
      ]
    },
    {
      "type": "field",
      "outboundTag": "Proxy",
      "ip": [
        "geoip:us",
        "geoip:jp",
        "geoip:facebook",
        "geoip:telegram",
        "ext:geoip-asn.dat:facebook",
        "ext:geoip-asn.dat:telegram"
      ]
    }
  ]
}
```

</details>

<details>
  <summary>点击查看在 <b>mihomo</b> 中的使用方法</summary>

```yaml
geodata-mode: true
geox-url:
  geoip: "https://cdn.jsdelivr.net/gh/arsity/geoip@release/geoip.dat"
```

</details>

<details>
  <summary>点击查看在 <b>hysteria</b> 中的使用方法</summary>
  <br/>
  <p>需要先下载 <code>.dat</code> 格式文件，并放置在 hysteria 程序目录内。</p>

```
direct(geoip:cn)
proxy(geoip:telegram)
proxy(geoip:us)
```

</details>

<details>
  <summary>点击查看在 <b>Trojan-Go</b> 中的使用方法</summary>
  <br/>
  <p>需要先下载 <code>.dat</code> 格式文件，并放置在 Trojan-Go 程序目录内。</p>

```json
"router": {
  "enabled": true,
  "bypass": ["geoip:cn"],
  "proxy": ["geoip:telegram", "geoip:us"],
  "block": ["geoip:jp"],
  "default_policy": "proxy",
  "geoip": "./geoip.dat"
}
```

</details>

<details>
  <summary>点击查看在 <b>dae</b> 中的使用方法</summary>
  <br/>

点击前往查看：[《吃鹅直通手册》](https://github.com/daeuniverse/dae/blob/main/docs/zh/README.md)

</details>

---

### V2Ray dat 下载地址

> 适用于 [V2Ray](https://github.com/v2fly/v2ray-core)、[Xray-core](https://github.com/XTLS/Xray-core)、[mihomo](https://github.com/MetaCubeX/mihomo/tree/Meta)、[hysteria](https://github.com/apernet/hysteria)、[Trojan-Go](https://github.com/p4gefau1t/trojan-go)、[dae](https://github.com/daeuniverse/dae)。

> 此 dat 格式文件不能用于 Nginx。

所有**国家/地区**、**新增类别**的 dat 格式文件，请查看本项目 `release` 分支下的 [dat 目录](https://github.com/arsity/geoip/tree/release/dat)。

- **geoip.dat**：
  - [https://raw.githubusercontent.com/arsity/geoip/release/geoip.dat](https://raw.githubusercontent.com/arsity/geoip/release/geoip.dat)
  - [https://cdn.jsdelivr.net/gh/arsity/geoip@release/geoip.dat](https://cdn.jsdelivr.net/gh/arsity/geoip@release/geoip.dat)
- **geoip.dat.sha256sum**：
  - [https://raw.githubusercontent.com/arsity/geoip/release/geoip.dat.sha256sum](https://raw.githubusercontent.com/arsity/geoip/release/geoip.dat.sha256sum)
  - [https://cdn.jsdelivr.net/gh/arsity/geoip@release/geoip.dat.sha256sum](https://cdn.jsdelivr.net/gh/arsity/geoip@release/geoip.dat.sha256sum)
- **geoip-only-cn-private.dat**（精简版 GeoIP，只包含 `geoip:cn` 和 `geoip:private`）：
  - [https://raw.githubusercontent.com/arsity/geoip/release/geoip-only-cn-private.dat](https://raw.githubusercontent.com/arsity/geoip/release/geoip-only-cn-private.dat)
  - [https://cdn.jsdelivr.net/gh/arsity/geoip@release/geoip-only-cn-private.dat](https://cdn.jsdelivr.net/gh/arsity/geoip@release/geoip-only-cn-private.dat)
- **geoip-only-cn-private.dat.sha256sum**：
  - [https://raw.githubusercontent.com/arsity/geoip/release/geoip-only-cn-private.dat.sha256sum](https://raw.githubusercontent.com/arsity/geoip/release/geoip-only-cn-private.dat.sha256sum)
  - [https://cdn.jsdelivr.net/gh/arsity/geoip@release/geoip-only-cn-private.dat.sha256sum](https://cdn.jsdelivr.net/gh/arsity/geoip@release/geoip-only-cn-private.dat.sha256sum)
- **geoip-asn.dat**（精简版 GeoIP，只包含上述新增类别）：
  - [https://raw.githubusercontent.com/arsity/geoip/release/geoip-asn.dat](https://raw.githubusercontent.com/arsity/geoip/release/geoip-asn.dat)
  - [https://cdn.jsdelivr.net/gh/arsity/geoip@release/geoip-asn.dat](https://cdn.jsdelivr.net/gh/arsity/geoip@release/geoip-asn.dat)
- **geoip-asn.dat.sha256sum**：
  - [https://raw.githubusercontent.com/arsity/geoip/release/geoip-asn.dat.sha256sum](https://raw.githubusercontent.com/arsity/geoip/release/geoip-asn.dat.sha256sum)
  - [https://cdn.jsdelivr.net/gh/arsity/geoip@release/geoip-asn.dat.sha256sum](https://cdn.jsdelivr.net/gh/arsity/geoip@release/geoip-asn.dat.sha256sum)
- **cn.dat**（精简版 GeoIP，只包含 `geoip:cn`）：
  - [https://raw.githubusercontent.com/arsity/geoip/release/cn.dat](https://raw.githubusercontent.com/arsity/geoip/release/cn.dat)
  - [https://cdn.jsdelivr.net/gh/arsity/geoip@release/cn.dat](https://cdn.jsdelivr.net/gh/arsity/geoip@release/cn.dat)
- **cn.dat.sha256sum**：
  - [https://raw.githubusercontent.com/arsity/geoip/release/cn.dat.sha256sum](https://raw.githubusercontent.com/arsity/geoip/release/cn.dat.sha256sum)
  - [https://cdn.jsdelivr.net/gh/arsity/geoip@release/cn.dat.sha256sum](https://cdn.jsdelivr.net/gh/arsity/geoip@release/cn.dat.sha256sum)
- **private.dat**（精简版 GeoIP，只包含 `geoip:private`）：
  - [https://raw.githubusercontent.com/arsity/geoip/release/private.dat](https://raw.githubusercontent.com/arsity/geoip/release/private.dat)
  - [https://cdn.jsdelivr.net/gh/arsity/geoip@release/private.dat](https://cdn.jsdelivr.net/gh/arsity/geoip@release/private.dat)
- **private.dat.sha256sum**：
  - [https://raw.githubusercontent.com/arsity/geoip/release/private.dat.sha256sum](https://raw.githubusercontent.com/arsity/geoip/release/private.dat.sha256sum)
  - [https://cdn.jsdelivr.net/gh/arsity/geoip@release/private.dat.sha256sum](https://cdn.jsdelivr.net/gh/arsity/geoip@release/private.dat.sha256sum)

部分**国家/地区**类别：

- **中国大陆**（Mainland China）：
  - [https://raw.githubusercontent.com/arsity/geoip/release/dat/cn.dat](https://raw.githubusercontent.com/arsity/geoip/release/dat/cn.dat)
  - [https://cdn.jsdelivr.net/gh/arsity/geoip@release/dat/cn.dat](https://cdn.jsdelivr.net/gh/arsity/geoip@release/dat/cn.dat)
- **香港**（Hong Kong）：
  - [https://raw.githubusercontent.com/arsity/geoip/release/dat/hk.dat](https://raw.githubusercontent.com/arsity/geoip/release/dat/hk.dat)
  - [https://cdn.jsdelivr.net/gh/arsity/geoip@release/dat/hk.dat](https://cdn.jsdelivr.net/gh/arsity/geoip@release/dat/hk.dat)
- **澳门**（Macau）：
  - [https://raw.githubusercontent.com/arsity/geoip/release/dat/mo.dat](https://raw.githubusercontent.com/arsity/geoip/release/dat/mo.dat)
  - [https://cdn.jsdelivr.net/gh/arsity/geoip@release/dat/mo.dat](https://cdn.jsdelivr.net/gh/arsity/geoip@release/dat/mo.dat)
- **台湾**（Taiwan）：
  - [https://raw.githubusercontent.com/arsity/geoip/release/dat/tw.dat](https://raw.githubusercontent.com/arsity/geoip/release/dat/tw.dat)
  - [https://cdn.jsdelivr.net/gh/arsity/geoip@release/dat/tw.dat](https://cdn.jsdelivr.net/gh/arsity/geoip@release/dat/tw.dat)
- **美国**（America）：
  - [https://raw.githubusercontent.com/arsity/geoip/release/dat/us.dat](https://raw.githubusercontent.com/arsity/geoip/release/dat/us.dat)
  - [https://cdn.jsdelivr.net/gh/arsity/geoip@release/dat/us.dat](https://cdn.jsdelivr.net/gh/arsity/geoip@release/dat/us.dat)
- **日本**（Japan）：
  - [https://raw.githubusercontent.com/arsity/geoip/release/dat/jp.dat](https://raw.githubusercontent.com/arsity/geoip/release/dat/jp.dat)
  - [https://cdn.jsdelivr.net/gh/arsity/geoip@release/dat/jp.dat](https://cdn.jsdelivr.net/gh/arsity/geoip@release/dat/jp.dat)
- **韩国**（Korea）：
  - [https://raw.githubusercontent.com/arsity/geoip/release/dat/kr.dat](https://raw.githubusercontent.com/arsity/geoip/release/dat/kr.dat)
  - [https://cdn.jsdelivr.net/gh/arsity/geoip@release/dat/kr.dat](https://cdn.jsdelivr.net/gh/arsity/geoip@release/dat/kr.dat)
- **新加坡**（Singapore）：
  - [https://raw.githubusercontent.com/arsity/geoip/release/dat/sg.dat](https://raw.githubusercontent.com/arsity/geoip/release/dat/sg.dat)
  - [https://cdn.jsdelivr.net/gh/arsity/geoip@release/dat/sg.dat](https://cdn.jsdelivr.net/gh/arsity/geoip@release/dat/sg.dat)
- **缅甸**（Myanmar）：
  - [https://raw.githubusercontent.com/arsity/geoip/release/dat/mm.dat](https://raw.githubusercontent.com/arsity/geoip/release/dat/mm.dat)
  - [https://cdn.jsdelivr.net/gh/arsity/geoip@release/dat/mm.dat](https://cdn.jsdelivr.net/gh/arsity/geoip@release/dat/mm.dat)
- **伊朗**（Iran）：
  - [https://raw.githubusercontent.com/arsity/geoip/release/dat/ir.dat](https://raw.githubusercontent.com/arsity/geoip/release/dat/ir.dat)
  - [https://cdn.jsdelivr.net/gh/arsity/geoip@release/dat/ir.dat](https://cdn.jsdelivr.net/gh/arsity/geoip@release/dat/ir.dat)
- **俄罗斯**（Russia）：
  - [https://raw.githubusercontent.com/arsity/geoip/release/dat/ru.dat](https://raw.githubusercontent.com/arsity/geoip/release/dat/ru.dat)
  - [https://cdn.jsdelivr.net/gh/arsity/geoip@release/dat/ru.dat](https://cdn.jsdelivr.net/gh/arsity/geoip@release/dat/ru.dat)
- **白俄罗斯**（Belarus）：
  - [https://raw.githubusercontent.com/arsity/geoip/release/dat/by.dat](https://raw.githubusercontent.com/arsity/geoip/release/dat/by.dat)
  - [https://cdn.jsdelivr.net/gh/arsity/geoip@release/dat/by.dat](https://cdn.jsdelivr.net/gh/arsity/geoip@release/dat/by.dat)
- **土库曼斯坦**（Turkmenistan）：
  - [https://raw.githubusercontent.com/arsity/geoip/release/dat/tm.dat](https://raw.githubusercontent.com/arsity/geoip/release/dat/tm.dat)
  - [https://cdn.jsdelivr.net/gh/arsity/geoip@release/dat/tm.dat](https://cdn.jsdelivr.net/gh/arsity/geoip@release/dat/tm.dat)

**新增**类别：

- **cloudflare**：
  - [https://raw.githubusercontent.com/arsity/geoip/release/dat/cloudflare.dat](https://raw.githubusercontent.com/arsity/geoip/release/dat/cloudflare.dat)
  - [https://cdn.jsdelivr.net/gh/arsity/geoip@release/dat/cloudflare.dat](https://cdn.jsdelivr.net/gh/arsity/geoip@release/dat/cloudflare.dat)
- **cloudfront**：
  - [https://raw.githubusercontent.com/arsity/geoip/release/dat/cloudfront.dat](https://raw.githubusercontent.com/arsity/geoip/release/dat/cloudfront.dat)
  - [https://cdn.jsdelivr.net/gh/arsity/geoip@release/dat/cloudfront.dat](https://cdn.jsdelivr.net/gh/arsity/geoip@release/dat/cloudfront.dat)
- **facebook**：
  - [https://raw.githubusercontent.com/arsity/geoip/release/dat/facebook.dat](https://raw.githubusercontent.com/arsity/geoip/release/dat/facebook.dat)
  - [https://cdn.jsdelivr.net/gh/arsity/geoip@release/dat/facebook.dat](https://cdn.jsdelivr.net/gh/arsity/geoip@release/dat/facebook.dat)
- **fastly**：
  - [https://raw.githubusercontent.com/arsity/geoip/release/dat/fastly.dat](https://raw.githubusercontent.com/arsity/geoip/release/dat/fastly.dat)
  - [https://cdn.jsdelivr.net/gh/arsity/geoip@release/dat/fastly.dat](https://cdn.jsdelivr.net/gh/arsity/geoip@release/dat/fastly.dat)
- **google**：
  - [https://raw.githubusercontent.com/arsity/geoip/release/dat/google.dat](https://raw.githubusercontent.com/arsity/geoip/release/dat/google.dat)
  - [https://cdn.jsdelivr.net/gh/arsity/geoip@release/dat/google.dat](https://cdn.jsdelivr.net/gh/arsity/geoip@release/dat/google.dat)
- **netflix**：
  - [https://raw.githubusercontent.com/arsity/geoip/release/dat/netflix.dat](https://raw.githubusercontent.com/arsity/geoip/release/dat/netflix.dat)
  - [https://cdn.jsdelivr.net/gh/arsity/geoip@release/dat/netflix.dat](https://cdn.jsdelivr.net/gh/arsity/geoip@release/dat/netflix.dat)
- **spotify**：
  - [https://raw.githubusercontent.com/arsity/geoip/release/dat/spotify.dat](https://raw.githubusercontent.com/arsity/geoip/release/dat/spotify.dat)
  - [https://cdn.jsdelivr.net/gh/arsity/geoip@release/dat/spotify.dat](https://cdn.jsdelivr.net/gh/arsity/geoip@release/dat/spotify.dat)
- **telegram**：
  - [https://raw.githubusercontent.com/arsity/geoip/release/dat/telegram.dat](https://raw.githubusercontent.com/arsity/geoip/release/dat/telegram.dat)
  - [https://cdn.jsdelivr.net/gh/arsity/geoip@release/dat/telegram.dat](https://cdn.jsdelivr.net/gh/arsity/geoip@release/dat/telegram.dat)
- **twitter**：
  - [https://raw.githubusercontent.com/arsity/geoip/release/dat/twitter.dat](https://raw.githubusercontent.com/arsity/geoip/release/dat/twitter.dat)
  - [https://cdn.jsdelivr.net/gh/arsity/geoip@release/dat/twitter.dat](https://cdn.jsdelivr.net/gh/arsity/geoip@release/dat/twitter.dat)
- **tor**：
  - [https://raw.githubusercontent.com/arsity/geoip/release/dat/tor.dat](https://raw.githubusercontent.com/arsity/geoip/release/dat/tor.dat)
  - [https://cdn.jsdelivr.net/gh/arsity/geoip@release/dat/tor.dat](https://cdn.jsdelivr.net/gh/arsity/geoip@release/dat/tor.dat)

---

### MaxMind-compatible mmdb 使用方法

<details>
  <summary>点击查看在 <b>Clash</b> 中的使用方法</summary>
  <br/>
  <p>需要先下载 <code>.mmdb</code> 格式文件，命名为 <code>Country.mmdb</code>，并放置在 Clash 程序目录内。</p>

```yaml
rules:
  - GEOIP,PRIVATE,policy,no-resolve
  - GEOIP,FACEBOOK,policy
  - GEOIP,CN,policy,no-resolve
```

</details>

<details>
  <summary>点击查看在 <b>mihomo</b> 中的使用方法</summary>

```yaml
geodata-mode: false
geox-url:
  mmdb: "https://cdn.jsdelivr.net/gh/arsity/geoip@release/Country.mmdb"
```

</details>

<details>
  <summary>点击查看在 <b>Shadowrocket</b> 中的使用方法</summary>
  <br/>
  <p>需要将下载地址填入到 Shadowrocket 的设置中。</p>

```conf
[Rule]
GEOIP,PRIVATE,DIRECT
GEOIP,FACEBOOK,PROXY
GEOIP,CN,DIRECT
```

</details>

<details>
  <summary>点击查看在 <b>Quantumult X</b> 中的使用方法</summary>
  <br/>
  <p>需要将下载地址填入到 Quantumult X 的设置中。</p>

```conf
[filter_local]
GEOIP,PRIVATE,DIRECT
GEOIP,FACEBOOK,PROXY
GEOIP,CN,DIRECT
```

</details>

<details>
  <summary>点击查看在 <b>Surge</b> 中的使用方法</summary>
  <br/>
  <p>需要将下载地址填入到 Surge 的设置中。</p>

```conf
[Rule]
GEOIP,PRIVATE,policy,no-resolve
GEOIP,FACEBOOK,policy
GEOIP,CN,policy,no-resolve
```

</details>

---

### MaxMind-compatible mmdb 下载地址

<br/>

本项目生成的**国家/地区**类型 mmdb 文件：

> 适用于 [Clash](https://github.com/Dreamacro/clash)、[mihomo](https://github.com/MetaCubeX/mihomo/tree/Meta)、[Shadowrocket](https://apps.apple.com/us/app/id932747118)、[Quantumult X](https://apps.apple.com/us/app/id1443988620)、[Surge](https://nssurge.com)。

> 适用于 [Nginx](https://nginx.org)，需要配合 [ngx_http_geoip2_module](https://github.com/leev/ngx_http_geoip2_module) 模块使用。

> 本项目生成的 mmdb 格式文件使用 IPInfo Lite 作为默认数据源，并写出 MaxMind GeoLite2 Country 兼容字段。**国家/地区**类别包含 `continent` 和 `country.iso_code` 等字段；**新增类别**和 **`GEOIP,PRIVATE` 类别**只保留 `country.iso_code` 字段，用于兼容 `GEOIP,CN`、`GEOIP,CLOUDFLARE` 等规则。

- **Country-without-asn.mmdb**（传统版 GeoIP，只包含国家/地区类别和 `GEOIP,PRIVATE` 类别，不包含上述新增类别。建议优先使用）：
  - [https://raw.githubusercontent.com/arsity/geoip/release/Country-without-asn.mmdb](https://raw.githubusercontent.com/arsity/geoip/release/Country-without-asn.mmdb)
  - [https://cdn.jsdelivr.net/gh/arsity/geoip@release/Country-without-asn.mmdb](https://cdn.jsdelivr.net/gh/arsity/geoip@release/Country-without-asn.mmdb)
- **Country-without-asn.mmdb.sha256sum**：
  - [https://raw.githubusercontent.com/arsity/geoip/release/Country-without-asn.mmdb.sha256sum](https://raw.githubusercontent.com/arsity/geoip/release/Country-without-asn.mmdb.sha256sum)
  - [https://cdn.jsdelivr.net/gh/arsity/geoip@release/Country-without-asn.mmdb.sha256sum](https://cdn.jsdelivr.net/gh/arsity/geoip@release/Country-without-asn.mmdb.sha256sum)
- **Country.mmdb**（增强版 GeoIP，包含国家/地区类别、`GEOIP,PRIVATE` 类别，以及上述新增类别。但由于 MaxMind mmdb 格式限制，部分国家/地区类别的 IP 地址数据不如上述 **Country-without-asn.mmdb** 准确）：
  - [https://raw.githubusercontent.com/arsity/geoip/release/Country.mmdb](https://raw.githubusercontent.com/arsity/geoip/release/Country.mmdb)
  - [https://cdn.jsdelivr.net/gh/arsity/geoip@release/Country.mmdb](https://cdn.jsdelivr.net/gh/arsity/geoip@release/Country.mmdb)
- **Country.mmdb.sha256sum**：
  - [https://raw.githubusercontent.com/arsity/geoip/release/Country.mmdb.sha256sum](https://raw.githubusercontent.com/arsity/geoip/release/Country.mmdb.sha256sum)
  - [https://cdn.jsdelivr.net/gh/arsity/geoip@release/Country.mmdb.sha256sum](https://cdn.jsdelivr.net/gh/arsity/geoip@release/Country.mmdb.sha256sum)
- **Country-only-cn-private.mmdb**（精简版 GeoIP，只包含 `GEOIP,CN` 和 `GEOIP,PRIVATE`）：
  - [https://raw.githubusercontent.com/arsity/geoip/release/Country-only-cn-private.mmdb](https://raw.githubusercontent.com/arsity/geoip/release/Country-only-cn-private.mmdb)
  - [https://cdn.jsdelivr.net/gh/arsity/geoip@release/Country-only-cn-private.mmdb](https://cdn.jsdelivr.net/gh/arsity/geoip@release/Country-only-cn-private.mmdb)
- **Country-only-cn-private.mmdb.sha256sum**：
  - [https://raw.githubusercontent.com/arsity/geoip/release/Country-only-cn-private.mmdb.sha256sum](https://raw.githubusercontent.com/arsity/geoip/release/Country-only-cn-private.mmdb.sha256sum)
  - [https://cdn.jsdelivr.net/gh/arsity/geoip@release/Country-only-cn-private.mmdb.sha256sum](https://cdn.jsdelivr.net/gh/arsity/geoip@release/Country-only-cn-private.mmdb.sha256sum)
- **Country-asn.mmdb**（精简版 GeoIP，只包含上述新增类别）：
  - [https://raw.githubusercontent.com/arsity/geoip/release/Country-asn.mmdb](https://raw.githubusercontent.com/arsity/geoip/release/Country-asn.mmdb)
  - [https://cdn.jsdelivr.net/gh/arsity/geoip@release/Country-asn.mmdb](https://cdn.jsdelivr.net/gh/arsity/geoip@release/Country-asn.mmdb)
- **Country-asn.mmdb.sha256sum**：
  - [https://raw.githubusercontent.com/arsity/geoip/release/Country-asn.mmdb.sha256sum](https://raw.githubusercontent.com/arsity/geoip/release/Country-asn.mmdb.sha256sum)
  - [https://cdn.jsdelivr.net/gh/arsity/geoip@release/Country-asn.mmdb.sha256sum](https://cdn.jsdelivr.net/gh/arsity/geoip@release/Country-asn.mmdb.sha256sum)

---

### sing-box SRS 使用方法

<details>
  <summary>点击查看在 <b>sing-box</b> 中的使用方法</summary>

```json
"route": {
  "rules": [
    {
      "rule_set": "geoip-cn",
      "outbound": "direct"
    },
    {
      "rule_set": "geoip-us",
      "outbound": "block"
    }
  ],
  "rule_set": [
    {
      "tag": "geoip-cn",
      "type": "remote",
      "format": "binary",
      "url": "https://cdn.jsdelivr.net/gh/arsity/geoip@release/srs/cn.srs"
    },
    {
      "tag": "geoip-us",
      "type": "remote",
      "format": "binary",
      "url": "https://cdn.jsdelivr.net/gh/arsity/geoip@release/srs/us.srs"
    }
  ]
}
```

</details>

---

### sing-box SRS 下载地址

> 适用于 [sing-box](https://github.com/SagerNet/sing-box)。

所有**国家/地区**、**新增类别**的 SRS 格式文件，请查看本项目 `release` 分支下的 [srs 目录](https://github.com/arsity/geoip/tree/release/srs)。

部分**国家/地区**类别：

- **中国大陆**（Mainland China）：
  - [https://raw.githubusercontent.com/arsity/geoip/release/srs/cn.srs](https://raw.githubusercontent.com/arsity/geoip/release/srs/cn.srs)
  - [https://cdn.jsdelivr.net/gh/arsity/geoip@release/srs/cn.srs](https://cdn.jsdelivr.net/gh/arsity/geoip@release/srs/cn.srs)
- **香港**（Hong Kong）：
  - [https://raw.githubusercontent.com/arsity/geoip/release/srs/hk.srs](https://raw.githubusercontent.com/arsity/geoip/release/srs/hk.srs)
  - [https://cdn.jsdelivr.net/gh/arsity/geoip@release/srs/hk.srs](https://cdn.jsdelivr.net/gh/arsity/geoip@release/srs/hk.srs)
- **澳门**（Macau）：
  - [https://raw.githubusercontent.com/arsity/geoip/release/srs/mo.srs](https://raw.githubusercontent.com/arsity/geoip/release/srs/mo.srs)
  - [https://cdn.jsdelivr.net/gh/arsity/geoip@release/srs/mo.srs](https://cdn.jsdelivr.net/gh/arsity/geoip@release/srs/mo.srs)
- **台湾**（Taiwan）：
  - [https://raw.githubusercontent.com/arsity/geoip/release/srs/tw.srs](https://raw.githubusercontent.com/arsity/geoip/release/srs/tw.srs)
  - [https://cdn.jsdelivr.net/gh/arsity/geoip@release/srs/tw.srs](https://cdn.jsdelivr.net/gh/arsity/geoip@release/srs/tw.srs)
- **美国**（America）：
  - [https://raw.githubusercontent.com/arsity/geoip/release/srs/us.srs](https://raw.githubusercontent.com/arsity/geoip/release/srs/us.srs)
  - [https://cdn.jsdelivr.net/gh/arsity/geoip@release/srs/us.srs](https://cdn.jsdelivr.net/gh/arsity/geoip@release/srs/us.srs)
- **日本**（Japan）：
  - [https://raw.githubusercontent.com/arsity/geoip/release/srs/jp.srs](https://raw.githubusercontent.com/arsity/geoip/release/srs/jp.srs)
  - [https://cdn.jsdelivr.net/gh/arsity/geoip@release/srs/jp.srs](https://cdn.jsdelivr.net/gh/arsity/geoip@release/srs/jp.srs)
- **韩国**（Korea）：
  - [https://raw.githubusercontent.com/arsity/geoip/release/srs/kr.srs](https://raw.githubusercontent.com/arsity/geoip/release/srs/kr.srs)
  - [https://cdn.jsdelivr.net/gh/arsity/geoip@release/srs/kr.srs](https://cdn.jsdelivr.net/gh/arsity/geoip@release/srs/kr.srs)
- **新加坡**（Singapore）：
  - [https://raw.githubusercontent.com/arsity/geoip/release/srs/sg.srs](https://raw.githubusercontent.com/arsity/geoip/release/srs/sg.srs)
  - [https://cdn.jsdelivr.net/gh/arsity/geoip@release/srs/sg.srs](https://cdn.jsdelivr.net/gh/arsity/geoip@release/srs/sg.srs)
- **缅甸**（Myanmar）：
  - [https://raw.githubusercontent.com/arsity/geoip/release/srs/mm.srs](https://raw.githubusercontent.com/arsity/geoip/release/srs/mm.srs)
  - [https://cdn.jsdelivr.net/gh/arsity/geoip@release/srs/mm.srs](https://cdn.jsdelivr.net/gh/arsity/geoip@release/srs/mm.srs)
- **伊朗**（Iran）：
  - [https://raw.githubusercontent.com/arsity/geoip/release/srs/ir.srs](https://raw.githubusercontent.com/arsity/geoip/release/srs/ir.srs)
  - [https://cdn.jsdelivr.net/gh/arsity/geoip@release/srs/ir.srs](https://cdn.jsdelivr.net/gh/arsity/geoip@release/srs/ir.srs)
- **俄罗斯**（Russia）：
  - [https://raw.githubusercontent.com/arsity/geoip/release/srs/ru.srs](https://raw.githubusercontent.com/arsity/geoip/release/srs/ru.srs)
  - [https://cdn.jsdelivr.net/gh/arsity/geoip@release/srs/ru.srs](https://cdn.jsdelivr.net/gh/arsity/geoip@release/srs/ru.srs)
- **白俄罗斯**（Belarus）：
  - [https://raw.githubusercontent.com/arsity/geoip/release/srs/by.srs](https://raw.githubusercontent.com/arsity/geoip/release/srs/by.srs)
  - [https://cdn.jsdelivr.net/gh/arsity/geoip@release/srs/by.srs](https://cdn.jsdelivr.net/gh/arsity/geoip@release/srs/by.srs)
- **土库曼斯坦**（Turkmenistan）：
  - [https://raw.githubusercontent.com/arsity/geoip/release/srs/tm.srs](https://raw.githubusercontent.com/arsity/geoip/release/srs/tm.srs)
  - [https://cdn.jsdelivr.net/gh/arsity/geoip@release/srs/tm.srs](https://cdn.jsdelivr.net/gh/arsity/geoip@release/srs/tm.srs)

**新增**类别：

- **cloudflare**：
  - [https://raw.githubusercontent.com/arsity/geoip/release/srs/cloudflare.srs](https://raw.githubusercontent.com/arsity/geoip/release/srs/cloudflare.srs)
  - [https://cdn.jsdelivr.net/gh/arsity/geoip@release/srs/cloudflare.srs](https://cdn.jsdelivr.net/gh/arsity/geoip@release/srs/cloudflare.srs)
- **cloudfront**：
  - [https://raw.githubusercontent.com/arsity/geoip/release/srs/cloudfront.srs](https://raw.githubusercontent.com/arsity/geoip/release/srs/cloudfront.srs)
  - [https://cdn.jsdelivr.net/gh/arsity/geoip@release/srs/cloudfront.srs](https://cdn.jsdelivr.net/gh/arsity/geoip@release/srs/cloudfront.srs)
- **facebook**：
  - [https://raw.githubusercontent.com/arsity/geoip/release/srs/facebook.srs](https://raw.githubusercontent.com/arsity/geoip/release/srs/facebook.srs)
  - [https://cdn.jsdelivr.net/gh/arsity/geoip@release/srs/facebook.srs](https://cdn.jsdelivr.net/gh/arsity/geoip@release/srs/facebook.srs)
- **fastly**：
  - [https://raw.githubusercontent.com/arsity/geoip/release/srs/fastly.srs](https://raw.githubusercontent.com/arsity/geoip/release/srs/fastly.srs)
  - [https://cdn.jsdelivr.net/gh/arsity/geoip@release/srs/fastly.srs](https://cdn.jsdelivr.net/gh/arsity/geoip@release/srs/fastly.srs)
- **google**：
  - [https://raw.githubusercontent.com/arsity/geoip/release/srs/google.srs](https://raw.githubusercontent.com/arsity/geoip/release/srs/google.srs)
  - [https://cdn.jsdelivr.net/gh/arsity/geoip@release/srs/google.srs](https://cdn.jsdelivr.net/gh/arsity/geoip@release/srs/google.srs)
- **netflix**：
  - [https://raw.githubusercontent.com/arsity/geoip/release/srs/netflix.srs](https://raw.githubusercontent.com/arsity/geoip/release/srs/netflix.srs)
  - [https://cdn.jsdelivr.net/gh/arsity/geoip@release/srs/netflix.srs](https://cdn.jsdelivr.net/gh/arsity/geoip@release/srs/netflix.srs)
- **spotify**：
  - [https://raw.githubusercontent.com/arsity/geoip/release/srs/spotify.srs](https://raw.githubusercontent.com/arsity/geoip/release/srs/spotify.srs)
  - [https://cdn.jsdelivr.net/gh/arsity/geoip@release/srs/spotify.srs](https://cdn.jsdelivr.net/gh/arsity/geoip@release/srs/spotify.srs)
- **telegram**：
  - [https://raw.githubusercontent.com/arsity/geoip/release/srs/telegram.srs](https://raw.githubusercontent.com/arsity/geoip/release/srs/telegram.srs)
  - [https://cdn.jsdelivr.net/gh/arsity/geoip@release/srs/telegram.srs](https://cdn.jsdelivr.net/gh/arsity/geoip@release/srs/telegram.srs)
- **twitter**：
  - [https://raw.githubusercontent.com/arsity/geoip/release/srs/twitter.srs](https://raw.githubusercontent.com/arsity/geoip/release/srs/twitter.srs)
  - [https://cdn.jsdelivr.net/gh/arsity/geoip@release/srs/twitter.srs](https://cdn.jsdelivr.net/gh/arsity/geoip@release/srs/twitter.srs)
- **tor**：
  - [https://raw.githubusercontent.com/arsity/geoip/release/srs/tor.srs](https://raw.githubusercontent.com/arsity/geoip/release/srs/tor.srs)
  - [https://cdn.jsdelivr.net/gh/arsity/geoip@release/srs/tor.srs](https://cdn.jsdelivr.net/gh/arsity/geoip@release/srs/tor.srs)

---

### mihomo MRS 使用方法

<details>
  <summary>点击查看在 <b>mihomo</b> 中的使用方法</summary>

```yaml
rule-providers:
  cn-cidr:
    type: http
    behavior: ipcidr
    format: mrs
    url: "https://cdn.jsdelivr.net/gh/arsity/geoip@release/mrs/cn.mrs"
    path: ./mrs/geoip/cn.mrs
    interval: 86400

  google-cidr:
    type: http
    behavior: ipcidr
    format: mrs
    url: "https://cdn.jsdelivr.net/gh/arsity/geoip@release/mrs/google.mrs"
    path: ./mrs/geoip/google.mrs
    interval: 86400

rules:
  - RULE-SET,cn-cidr,DIRECT
  - RULE-SET,google-cidr,PROXY,no-resolve
```

</details>

---

### mihomo MRS 下载地址

> 适用于 [mihomo](https://github.com/MetaCubeX/mihomo/tree/Meta)。

所有**国家/地区**、**新增类别**的 MRS 格式文件，请查看本项目 `release` 分支下的 [mrs 目录](https://github.com/arsity/geoip/tree/release/mrs)。

部分**国家/地区**类别：

- **中国大陆**（Mainland China）：
  - [https://raw.githubusercontent.com/arsity/geoip/release/mrs/cn.mrs](https://raw.githubusercontent.com/arsity/geoip/release/mrs/cn.mrs)
  - [https://cdn.jsdelivr.net/gh/arsity/geoip@release/mrs/cn.mrs](https://cdn.jsdelivr.net/gh/arsity/geoip@release/mrs/cn.mrs)
- **香港**（Hong Kong）：
  - [https://raw.githubusercontent.com/arsity/geoip/release/mrs/hk.mrs](https://raw.githubusercontent.com/arsity/geoip/release/mrs/hk.mrs)
  - [https://cdn.jsdelivr.net/gh/arsity/geoip@release/mrs/hk.mrs](https://cdn.jsdelivr.net/gh/arsity/geoip@release/mrs/hk.mrs)
- **澳门**（Macau）：
  - [https://raw.githubusercontent.com/arsity/geoip/release/mrs/mo.mrs](https://raw.githubusercontent.com/arsity/geoip/release/mrs/mo.mrs)
  - [https://cdn.jsdelivr.net/gh/arsity/geoip@release/mrs/mo.mrs](https://cdn.jsdelivr.net/gh/arsity/geoip@release/mrs/mo.mrs)
- **台湾**（Taiwan）：
  - [https://raw.githubusercontent.com/arsity/geoip/release/mrs/tw.mrs](https://raw.githubusercontent.com/arsity/geoip/release/mrs/tw.mrs)
  - [https://cdn.jsdelivr.net/gh/arsity/geoip@release/mrs/tw.mrs](https://cdn.jsdelivr.net/gh/arsity/geoip@release/mrs/tw.mrs)
- **美国**（America）：
  - [https://raw.githubusercontent.com/arsity/geoip/release/mrs/us.mrs](https://raw.githubusercontent.com/arsity/geoip/release/mrs/us.mrs)
  - [https://cdn.jsdelivr.net/gh/arsity/geoip@release/mrs/us.mrs](https://cdn.jsdelivr.net/gh/arsity/geoip@release/mrs/us.mrs)
- **日本**（Japan）：
  - [https://raw.githubusercontent.com/arsity/geoip/release/mrs/jp.mrs](https://raw.githubusercontent.com/arsity/geoip/release/mrs/jp.mrs)
  - [https://cdn.jsdelivr.net/gh/arsity/geoip@release/mrs/jp.mrs](https://cdn.jsdelivr.net/gh/arsity/geoip@release/mrs/jp.mrs)
- **韩国**（Korea）：
  - [https://raw.githubusercontent.com/arsity/geoip/release/mrs/kr.mrs](https://raw.githubusercontent.com/arsity/geoip/release/mrs/kr.mrs)
  - [https://cdn.jsdelivr.net/gh/arsity/geoip@release/mrs/kr.mrs](https://cdn.jsdelivr.net/gh/arsity/geoip@release/mrs/kr.mrs)
- **新加坡**（Singapore）：
  - [https://raw.githubusercontent.com/arsity/geoip/release/mrs/sg.mrs](https://raw.githubusercontent.com/arsity/geoip/release/mrs/sg.mrs)
  - [https://cdn.jsdelivr.net/gh/arsity/geoip@release/mrs/sg.mrs](https://cdn.jsdelivr.net/gh/arsity/geoip@release/mrs/sg.mrs)
- **缅甸**（Myanmar）：
  - [https://raw.githubusercontent.com/arsity/geoip/release/mrs/mm.mrs](https://raw.githubusercontent.com/arsity/geoip/release/mrs/mm.mrs)
  - [https://cdn.jsdelivr.net/gh/arsity/geoip@release/mrs/mm.mrs](https://cdn.jsdelivr.net/gh/arsity/geoip@release/mrs/mm.mrs)
- **伊朗**（Iran）：
  - [https://raw.githubusercontent.com/arsity/geoip/release/mrs/ir.mrs](https://raw.githubusercontent.com/arsity/geoip/release/mrs/ir.mrs)
  - [https://cdn.jsdelivr.net/gh/arsity/geoip@release/mrs/ir.mrs](https://cdn.jsdelivr.net/gh/arsity/geoip@release/mrs/ir.mrs)
- **俄罗斯**（Russia）：
  - [https://raw.githubusercontent.com/arsity/geoip/release/mrs/ru.mrs](https://raw.githubusercontent.com/arsity/geoip/release/mrs/ru.mrs)
  - [https://cdn.jsdelivr.net/gh/arsity/geoip@release/mrs/ru.mrs](https://cdn.jsdelivr.net/gh/arsity/geoip@release/mrs/ru.mrs)
- **白俄罗斯**（Belarus）：
  - [https://raw.githubusercontent.com/arsity/geoip/release/mrs/by.mrs](https://raw.githubusercontent.com/arsity/geoip/release/mrs/by.mrs)
  - [https://cdn.jsdelivr.net/gh/arsity/geoip@release/mrs/by.mrs](https://cdn.jsdelivr.net/gh/arsity/geoip@release/mrs/by.mrs)
- **土库曼斯坦**（Turkmenistan）：
  - [https://raw.githubusercontent.com/arsity/geoip/release/mrs/tm.mrs](https://raw.githubusercontent.com/arsity/geoip/release/mrs/tm.mrs)
  - [https://cdn.jsdelivr.net/gh/arsity/geoip@release/mrs/tm.mrs](https://cdn.jsdelivr.net/gh/arsity/geoip@release/mrs/tm.mrs)

**新增**类别：

- **cloudflare**：
  - [https://raw.githubusercontent.com/arsity/geoip/release/mrs/cloudflare.mrs](https://raw.githubusercontent.com/arsity/geoip/release/mrs/cloudflare.mrs)
  - [https://cdn.jsdelivr.net/gh/arsity/geoip@release/mrs/cloudflare.mrs](https://cdn.jsdelivr.net/gh/arsity/geoip@release/mrs/cloudflare.mrs)
- **cloudfront**：
  - [https://raw.githubusercontent.com/arsity/geoip/release/mrs/cloudfront.mrs](https://raw.githubusercontent.com/arsity/geoip/release/mrs/cloudfront.mrs)
  - [https://cdn.jsdelivr.net/gh/arsity/geoip@release/mrs/cloudfront.mrs](https://cdn.jsdelivr.net/gh/arsity/geoip@release/mrs/cloudfront.mrs)
- **facebook**：
  - [https://raw.githubusercontent.com/arsity/geoip/release/mrs/facebook.mrs](https://raw.githubusercontent.com/arsity/geoip/release/mrs/facebook.mrs)
  - [https://cdn.jsdelivr.net/gh/arsity/geoip@release/mrs/facebook.mrs](https://cdn.jsdelivr.net/gh/arsity/geoip@release/mrs/facebook.mrs)
- **fastly**：
  - [https://raw.githubusercontent.com/arsity/geoip/release/mrs/fastly.mrs](https://raw.githubusercontent.com/arsity/geoip/release/mrs/fastly.mrs)
  - [https://cdn.jsdelivr.net/gh/arsity/geoip@release/mrs/fastly.mrs](https://cdn.jsdelivr.net/gh/arsity/geoip@release/mrs/fastly.mrs)
- **google**：
  - [https://raw.githubusercontent.com/arsity/geoip/release/mrs/google.mrs](https://raw.githubusercontent.com/arsity/geoip/release/mrs/google.mrs)
  - [https://cdn.jsdelivr.net/gh/arsity/geoip@release/mrs/google.mrs](https://cdn.jsdelivr.net/gh/arsity/geoip@release/mrs/google.mrs)
- **netflix**：
  - [https://raw.githubusercontent.com/arsity/geoip/release/mrs/netflix.mrs](https://raw.githubusercontent.com/arsity/geoip/release/mrs/netflix.mrs)
  - [https://cdn.jsdelivr.net/gh/arsity/geoip@release/mrs/netflix.mrs](https://cdn.jsdelivr.net/gh/arsity/geoip@release/mrs/netflix.mrs)
- **spotify**：
  - [https://raw.githubusercontent.com/arsity/geoip/release/mrs/spotify.mrs](https://raw.githubusercontent.com/arsity/geoip/release/mrs/spotify.mrs)
  - [https://cdn.jsdelivr.net/gh/arsity/geoip@release/mrs/spotify.mrs](https://cdn.jsdelivr.net/gh/arsity/geoip@release/mrs/spotify.mrs)
- **telegram**：
  - [https://raw.githubusercontent.com/arsity/geoip/release/mrs/telegram.mrs](https://raw.githubusercontent.com/arsity/geoip/release/mrs/telegram.mrs)
  - [https://cdn.jsdelivr.net/gh/arsity/geoip@release/mrs/telegram.mrs](https://cdn.jsdelivr.net/gh/arsity/geoip@release/mrs/telegram.mrs)
- **twitter**：
  - [https://raw.githubusercontent.com/arsity/geoip/release/mrs/twitter.mrs](https://raw.githubusercontent.com/arsity/geoip/release/mrs/twitter.mrs)
  - [https://cdn.jsdelivr.net/gh/arsity/geoip@release/mrs/twitter.mrs](https://cdn.jsdelivr.net/gh/arsity/geoip@release/mrs/twitter.mrs)
- **tor**：
  - [https://raw.githubusercontent.com/arsity/geoip/release/mrs/tor.mrs](https://raw.githubusercontent.com/arsity/geoip/release/mrs/tor.mrs)
  - [https://cdn.jsdelivr.net/gh/arsity/geoip@release/mrs/tor.mrs](https://cdn.jsdelivr.net/gh/arsity/geoip@release/mrs/tor.mrs)

---

### Clash ruleset 使用方法

<details>
  <summary>点击查看在 <b>Clash Premium</b> 和 <b>mihomo</b> 中的使用方法</summary>

```yaml
rule-providers:
  cn-cidr:
    type: http
    behavior: ipcidr
    format: yaml
    url: "https://cdn.jsdelivr.net/gh/arsity/geoip@release/clash/ipcidr/cn.txt"
    path: ./ruleset/ipcidr/cn.yaml
    interval: 86400

  telegram-cidr:
    type: http
    behavior: ipcidr
    format: yaml
    url: "https://cdn.jsdelivr.net/gh/arsity/geoip@release/clash/ipcidr/telegram.txt"
    path: ./ruleset/ipcidr/telegram.yaml
    interval: 86400

rules:
  - RULE-SET,cn-cidr,DIRECT
  - RULE-SET,telegram-cidr,PROXY,no-resolve
```

</details>

---

### Clash ruleset 下载地址

> 适用于 [Clash Premium](https://github.com/Dreamacro/clash)、[mihomo](https://github.com/MetaCubeX/mihomo/tree/Meta)。

所有**国家/地区**、**新增类别**的 Clash ruleset 格式文件，请查看本项目 `release` 分支下的 [clash 目录](https://github.com/arsity/geoip/tree/release/clash)。

部分**国家/地区**类别：

- **中国大陆**（Mainland China）：
  - [https://raw.githubusercontent.com/arsity/geoip/release/clash/ipcidr/cn.txt](https://raw.githubusercontent.com/arsity/geoip/release/clash/ipcidr/cn.txt)
  - [https://cdn.jsdelivr.net/gh/arsity/geoip@release/clash/ipcidr/cn.txt](https://cdn.jsdelivr.net/gh/arsity/geoip@release/clash/ipcidr/cn.txt)
- **香港**（Hong Kong）：
  - [https://raw.githubusercontent.com/arsity/geoip/release/clash/ipcidr/hk.txt](https://raw.githubusercontent.com/arsity/geoip/release/clash/ipcidr/hk.txt)
  - [https://cdn.jsdelivr.net/gh/arsity/geoip@release/clash/ipcidr/hk.txt](https://cdn.jsdelivr.net/gh/arsity/geoip@release/clash/ipcidr/hk.txt)
- **澳门**（Macau）：
  - [https://raw.githubusercontent.com/arsity/geoip/release/clash/ipcidr/mo.txt](https://raw.githubusercontent.com/arsity/geoip/release/clash/ipcidr/mo.txt)
  - [https://cdn.jsdelivr.net/gh/arsity/geoip@release/clash/ipcidr/mo.txt](https://cdn.jsdelivr.net/gh/arsity/geoip@release/clash/ipcidr/mo.txt)
- **台湾**（Taiwan）：
  - [https://raw.githubusercontent.com/arsity/geoip/release/clash/ipcidr/tw.txt](https://raw.githubusercontent.com/arsity/geoip/release/clash/ipcidr/tw.txt)
  - [https://cdn.jsdelivr.net/gh/arsity/geoip@release/clash/ipcidr/tw.txt](https://cdn.jsdelivr.net/gh/arsity/geoip@release/clash/ipcidr/tw.txt)
- **美国**（America）：
  - [https://raw.githubusercontent.com/arsity/geoip/release/clash/ipcidr/us.txt](https://raw.githubusercontent.com/arsity/geoip/release/clash/ipcidr/us.txt)
  - [https://cdn.jsdelivr.net/gh/arsity/geoip@release/clash/ipcidr/us.txt](https://cdn.jsdelivr.net/gh/arsity/geoip@release/clash/ipcidr/us.txt)
- **日本**（Japan）：
  - [https://raw.githubusercontent.com/arsity/geoip/release/clash/ipcidr/jp.txt](https://raw.githubusercontent.com/arsity/geoip/release/clash/ipcidr/jp.txt)
  - [https://cdn.jsdelivr.net/gh/arsity/geoip@release/clash/ipcidr/jp.txt](https://cdn.jsdelivr.net/gh/arsity/geoip@release/clash/ipcidr/jp.txt)
- **韩国**（Korea）：
  - [https://raw.githubusercontent.com/arsity/geoip/release/clash/ipcidr/kr.txt](https://raw.githubusercontent.com/arsity/geoip/release/clash/ipcidr/kr.txt)
  - [https://cdn.jsdelivr.net/gh/arsity/geoip@release/clash/ipcidr/kr.txt](https://cdn.jsdelivr.net/gh/arsity/geoip@release/clash/ipcidr/kr.txt)
- **新加坡**（Singapore）：
  - [https://raw.githubusercontent.com/arsity/geoip/release/clash/ipcidr/sg.txt](https://raw.githubusercontent.com/arsity/geoip/release/clash/ipcidr/sg.txt)
  - [https://cdn.jsdelivr.net/gh/arsity/geoip@release/clash/ipcidr/sg.txt](https://cdn.jsdelivr.net/gh/arsity/geoip@release/clash/ipcidr/sg.txt)
- **缅甸**（Myanmar）：
  - [https://raw.githubusercontent.com/arsity/geoip/release/clash/ipcidr/mm.txt](https://raw.githubusercontent.com/arsity/geoip/release/clash/ipcidr/mm.txt)
  - [https://cdn.jsdelivr.net/gh/arsity/geoip@release/clash/ipcidr/mm.txt](https://cdn.jsdelivr.net/gh/arsity/geoip@release/clash/ipcidr/mm.txt)
- **伊朗**（Iran）：
  - [https://raw.githubusercontent.com/arsity/geoip/release/clash/ipcidr/ir.txt](https://raw.githubusercontent.com/arsity/geoip/release/clash/ipcidr/ir.txt)
  - [https://cdn.jsdelivr.net/gh/arsity/geoip@release/clash/ipcidr/ir.txt](https://cdn.jsdelivr.net/gh/arsity/geoip@release/clash/ipcidr/ir.txt)
- **俄罗斯**（Russia）：
  - [https://raw.githubusercontent.com/arsity/geoip/release/clash/ipcidr/ru.txt](https://raw.githubusercontent.com/arsity/geoip/release/clash/ipcidr/ru.txt)
  - [https://cdn.jsdelivr.net/gh/arsity/geoip@release/clash/ipcidr/ru.txt](https://cdn.jsdelivr.net/gh/arsity/geoip@release/clash/ipcidr/ru.txt)
- **白俄罗斯**（Belarus）：
  - [https://raw.githubusercontent.com/arsity/geoip/release/clash/ipcidr/by.txt](https://raw.githubusercontent.com/arsity/geoip/release/clash/ipcidr/by.txt)
  - [https://cdn.jsdelivr.net/gh/arsity/geoip@release/clash/ipcidr/by.txt](https://cdn.jsdelivr.net/gh/arsity/geoip@release/clash/ipcidr/by.txt)
- **土库曼斯坦**（Turkmenistan）：
  - [https://raw.githubusercontent.com/arsity/geoip/release/clash/ipcidr/tm.txt](https://raw.githubusercontent.com/arsity/geoip/release/clash/ipcidr/tm.txt)
  - [https://cdn.jsdelivr.net/gh/arsity/geoip@release/clash/ipcidr/tm.txt](https://cdn.jsdelivr.net/gh/arsity/geoip@release/clash/ipcidr/tm.txt)

**新增**类别：

- **cloudflare**：
  - [https://raw.githubusercontent.com/arsity/geoip/release/clash/ipcidr/cloudflare.txt](https://raw.githubusercontent.com/arsity/geoip/release/clash/ipcidr/cloudflare.txt)
  - [https://cdn.jsdelivr.net/gh/arsity/geoip@release/clash/ipcidr/cloudflare.txt](https://cdn.jsdelivr.net/gh/arsity/geoip@release/clash/ipcidr/cloudflare.txt)
- **cloudfront**：
  - [https://raw.githubusercontent.com/arsity/geoip/release/clash/ipcidr/cloudfront.txt](https://raw.githubusercontent.com/arsity/geoip/release/clash/ipcidr/cloudfront.txt)
  - [https://cdn.jsdelivr.net/gh/arsity/geoip@release/clash/ipcidr/cloudfront.txt](https://cdn.jsdelivr.net/gh/arsity/geoip@release/clash/ipcidr/cloudfront.txt)
- **facebook**：
  - [https://raw.githubusercontent.com/arsity/geoip/release/clash/ipcidr/facebook.txt](https://raw.githubusercontent.com/arsity/geoip/release/clash/ipcidr/facebook.txt)
  - [https://cdn.jsdelivr.net/gh/arsity/geoip@release/clash/ipcidr/facebook.txt](https://cdn.jsdelivr.net/gh/arsity/geoip@release/clash/ipcidr/facebook.txt)
- **fastly**：
  - [https://raw.githubusercontent.com/arsity/geoip/release/clash/ipcidr/fastly.txt](https://raw.githubusercontent.com/arsity/geoip/release/clash/ipcidr/fastly.txt)
  - [https://cdn.jsdelivr.net/gh/arsity/geoip@release/clash/ipcidr/fastly.txt](https://cdn.jsdelivr.net/gh/arsity/geoip@release/clash/ipcidr/fastly.txt)
- **google**：
  - [https://raw.githubusercontent.com/arsity/geoip/release/clash/ipcidr/google.txt](https://raw.githubusercontent.com/arsity/geoip/release/clash/ipcidr/google.txt)
  - [https://cdn.jsdelivr.net/gh/arsity/geoip@release/clash/ipcidr/google.txt](https://cdn.jsdelivr.net/gh/arsity/geoip@release/clash/ipcidr/google.txt)
- **netflix**：
  - [https://raw.githubusercontent.com/arsity/geoip/release/clash/ipcidr/netflix.txt](https://raw.githubusercontent.com/arsity/geoip/release/clash/ipcidr/netflix.txt)
  - [https://cdn.jsdelivr.net/gh/arsity/geoip@release/clash/ipcidr/netflix.txt](https://cdn.jsdelivr.net/gh/arsity/geoip@release/clash/ipcidr/netflix.txt)
- **spotify**：
  - [https://raw.githubusercontent.com/arsity/geoip/release/clash/ipcidr/spotify.txt](https://raw.githubusercontent.com/arsity/geoip/release/clash/ipcidr/spotify.txt)
  - [https://cdn.jsdelivr.net/gh/arsity/geoip@release/clash/ipcidr/spotify.txt](https://cdn.jsdelivr.net/gh/arsity/geoip@release/clash/ipcidr/spotify.txt)
- **telegram**：
  - [https://raw.githubusercontent.com/arsity/geoip/release/clash/ipcidr/telegram.txt](https://raw.githubusercontent.com/arsity/geoip/release/clash/ipcidr/telegram.txt)
  - [https://cdn.jsdelivr.net/gh/arsity/geoip@release/clash/ipcidr/telegram.txt](https://cdn.jsdelivr.net/gh/arsity/geoip@release/clash/ipcidr/telegram.txt)
- **twitter**：
  - [https://raw.githubusercontent.com/arsity/geoip/release/clash/ipcidr/twitter.txt](https://raw.githubusercontent.com/arsity/geoip/release/clash/ipcidr/twitter.txt)
  - [https://cdn.jsdelivr.net/gh/arsity/geoip@release/clash/ipcidr/twitter.txt](https://cdn.jsdelivr.net/gh/arsity/geoip@release/clash/ipcidr/twitter.txt)
- **tor**：
  - [https://raw.githubusercontent.com/arsity/geoip/release/clash/ipcidr/tor.txt](https://raw.githubusercontent.com/arsity/geoip/release/clash/ipcidr/tor.txt)
  - [https://cdn.jsdelivr.net/gh/arsity/geoip@release/clash/ipcidr/tor.txt](https://cdn.jsdelivr.net/gh/arsity/geoip@release/clash/ipcidr/tor.txt)

---

### Surge ruleset 使用方法

<details>
  <summary>点击查看在 <b>Surge</b> 中的使用方法</summary>

```conf
[Rule]
RULE-SET,https://cdn.jsdelivr.net/gh/arsity/geoip@release/surge/us.txt,REJECT
RULE-SET,https://cdn.jsdelivr.net/gh/arsity/geoip@release/surge/cn.txt,DIRECT
RULE-SET,https://cdn.jsdelivr.net/gh/arsity/geoip@release/surge/telegram.txt,PROXY,no-resolve
```

</details>

---

### Surge ruleset 下载地址

> 适用于 [Surge](https://nssurge.com)。

所有**国家/地区**、**新增类别**的 Surge ruleset 格式文件，请查看本项目 `release` 分支下的 [surge 目录](https://github.com/arsity/geoip/tree/release/surge)。

部分**国家/地区**类别：

- **中国大陆**（Mainland China）：
  - [https://raw.githubusercontent.com/arsity/geoip/release/surge/cn.txt](https://raw.githubusercontent.com/arsity/geoip/release/surge/cn.txt)
  - [https://cdn.jsdelivr.net/gh/arsity/geoip@release/surge/cn.txt](https://cdn.jsdelivr.net/gh/arsity/geoip@release/surge/cn.txt)
- **香港**（Hong Kong）：
  - [https://raw.githubusercontent.com/arsity/geoip/release/surge/hk.txt](https://raw.githubusercontent.com/arsity/geoip/release/surge/hk.txt)
  - [https://cdn.jsdelivr.net/gh/arsity/geoip@release/surge/hk.txt](https://cdn.jsdelivr.net/gh/arsity/geoip@release/surge/hk.txt)
- **澳门**（Macau）：
  - [https://raw.githubusercontent.com/arsity/geoip/release/surge/mo.txt](https://raw.githubusercontent.com/arsity/geoip/release/surge/mo.txt)
  - [https://cdn.jsdelivr.net/gh/arsity/geoip@release/surge/mo.txt](https://cdn.jsdelivr.net/gh/arsity/geoip@release/surge/mo.txt)
- **台湾**（Taiwan）：
  - [https://raw.githubusercontent.com/arsity/geoip/release/surge/tw.txt](https://raw.githubusercontent.com/arsity/geoip/release/surge/tw.txt)
  - [https://cdn.jsdelivr.net/gh/arsity/geoip@release/surge/tw.txt](https://cdn.jsdelivr.net/gh/arsity/geoip@release/surge/tw.txt)
- **美国**（America）：
  - [https://raw.githubusercontent.com/arsity/geoip/release/surge/us.txt](https://raw.githubusercontent.com/arsity/geoip/release/surge/us.txt)
  - [https://cdn.jsdelivr.net/gh/arsity/geoip@release/surge/us.txt](https://cdn.jsdelivr.net/gh/arsity/geoip@release/surge/us.txt)
- **日本**（Japan）：
  - [https://raw.githubusercontent.com/arsity/geoip/release/surge/jp.txt](https://raw.githubusercontent.com/arsity/geoip/release/surge/jp.txt)
  - [https://cdn.jsdelivr.net/gh/arsity/geoip@release/surge/jp.txt](https://cdn.jsdelivr.net/gh/arsity/geoip@release/surge/jp.txt)
- **韩国**（Korea）：
  - [https://raw.githubusercontent.com/arsity/geoip/release/surge/kr.txt](https://raw.githubusercontent.com/arsity/geoip/release/surge/kr.txt)
  - [https://cdn.jsdelivr.net/gh/arsity/geoip@release/surge/kr.txt](https://cdn.jsdelivr.net/gh/arsity/geoip@release/surge/kr.txt)
- **新加坡**（Singapore）：
  - [https://raw.githubusercontent.com/arsity/geoip/release/surge/sg.txt](https://raw.githubusercontent.com/arsity/geoip/release/surge/sg.txt)
  - [https://cdn.jsdelivr.net/gh/arsity/geoip@release/surge/sg.txt](https://cdn.jsdelivr.net/gh/arsity/geoip@release/surge/sg.txt)
- **缅甸**（Myanmar）：
  - [https://raw.githubusercontent.com/arsity/geoip/release/surge/mm.txt](https://raw.githubusercontent.com/arsity/geoip/release/surge/mm.txt)
  - [https://cdn.jsdelivr.net/gh/arsity/geoip@release/surge/mm.txt](https://cdn.jsdelivr.net/gh/arsity/geoip@release/surge/mm.txt)
- **伊朗**（Iran）：
  - [https://raw.githubusercontent.com/arsity/geoip/release/surge/ir.txt](https://raw.githubusercontent.com/arsity/geoip/release/surge/ir.txt)
  - [https://cdn.jsdelivr.net/gh/arsity/geoip@release/surge/ir.txt](https://cdn.jsdelivr.net/gh/arsity/geoip@release/surge/ir.txt)
- **俄罗斯**（Russia）：
  - [https://raw.githubusercontent.com/arsity/geoip/release/surge/ru.txt](https://raw.githubusercontent.com/arsity/geoip/release/surge/ru.txt)
  - [https://cdn.jsdelivr.net/gh/arsity/geoip@release/surge/ru.txt](https://cdn.jsdelivr.net/gh/arsity/geoip@release/surge/ru.txt)
- **白俄罗斯**（Belarus）：
  - [https://raw.githubusercontent.com/arsity/geoip/release/surge/by.txt](https://raw.githubusercontent.com/arsity/geoip/release/surge/by.txt)
  - [https://cdn.jsdelivr.net/gh/arsity/geoip@release/surge/by.txt](https://cdn.jsdelivr.net/gh/arsity/geoip@release/surge/by.txt)
- **土库曼斯坦**（Turkmenistan）：
  - [https://raw.githubusercontent.com/arsity/geoip/release/surge/tm.txt](https://raw.githubusercontent.com/arsity/geoip/release/surge/tm.txt)
  - [https://cdn.jsdelivr.net/gh/arsity/geoip@release/surge/tm.txt](https://cdn.jsdelivr.net/gh/arsity/geoip@release/surge/tm.txt)

**新增**类别：

- **cloudflare**：
  - [https://raw.githubusercontent.com/arsity/geoip/release/surge/cloudflare.txt](https://raw.githubusercontent.com/arsity/geoip/release/surge/cloudflare.txt)
  - [https://cdn.jsdelivr.net/gh/arsity/geoip@release/surge/cloudflare.txt](https://cdn.jsdelivr.net/gh/arsity/geoip@release/surge/cloudflare.txt)
- **cloudfront**：
  - [https://raw.githubusercontent.com/arsity/geoip/release/surge/cloudfront.txt](https://raw.githubusercontent.com/arsity/geoip/release/surge/cloudfront.txt)
  - [https://cdn.jsdelivr.net/gh/arsity/geoip@release/surge/cloudfront.txt](https://cdn.jsdelivr.net/gh/arsity/geoip@release/surge/cloudfront.txt)
- **facebook**：
  - [https://raw.githubusercontent.com/arsity/geoip/release/surge/facebook.txt](https://raw.githubusercontent.com/arsity/geoip/release/surge/facebook.txt)
  - [https://cdn.jsdelivr.net/gh/arsity/geoip@release/surge/facebook.txt](https://cdn.jsdelivr.net/gh/arsity/geoip@release/surge/facebook.txt)
- **fastly**：
  - [https://raw.githubusercontent.com/arsity/geoip/release/surge/fastly.txt](https://raw.githubusercontent.com/arsity/geoip/release/surge/fastly.txt)
  - [https://cdn.jsdelivr.net/gh/arsity/geoip@release/surge/fastly.txt](https://cdn.jsdelivr.net/gh/arsity/geoip@release/surge/fastly.txt)
- **google**：
  - [https://raw.githubusercontent.com/arsity/geoip/release/surge/google.txt](https://raw.githubusercontent.com/arsity/geoip/release/surge/google.txt)
  - [https://cdn.jsdelivr.net/gh/arsity/geoip@release/surge/google.txt](https://cdn.jsdelivr.net/gh/arsity/geoip@release/surge/google.txt)
- **netflix**：
  - [https://raw.githubusercontent.com/arsity/geoip/release/surge/netflix.txt](https://raw.githubusercontent.com/arsity/geoip/release/surge/netflix.txt)
  - [https://cdn.jsdelivr.net/gh/arsity/geoip@release/surge/netflix.txt](https://cdn.jsdelivr.net/gh/arsity/geoip@release/surge/netflix.txt)
- **spotify**：
  - [https://raw.githubusercontent.com/arsity/geoip/release/surge/spotify.txt](https://raw.githubusercontent.com/arsity/geoip/release/surge/spotify.txt)
  - [https://cdn.jsdelivr.net/gh/arsity/geoip@release/surge/spotify.txt](https://cdn.jsdelivr.net/gh/arsity/geoip@release/surge/spotify.txt)
- **telegram**：
  - [https://raw.githubusercontent.com/arsity/geoip/release/surge/telegram.txt](https://raw.githubusercontent.com/arsity/geoip/release/surge/telegram.txt)
  - [https://cdn.jsdelivr.net/gh/arsity/geoip@release/surge/telegram.txt](https://cdn.jsdelivr.net/gh/arsity/geoip@release/surge/telegram.txt)
- **twitter**：
  - [https://raw.githubusercontent.com/arsity/geoip/release/surge/twitter.txt](https://raw.githubusercontent.com/arsity/geoip/release/surge/twitter.txt)
  - [https://cdn.jsdelivr.net/gh/arsity/geoip@release/surge/twitter.txt](https://cdn.jsdelivr.net/gh/arsity/geoip@release/surge/twitter.txt)
- **tor**：
  - [https://raw.githubusercontent.com/arsity/geoip/release/surge/tor.txt](https://raw.githubusercontent.com/arsity/geoip/release/surge/tor.txt)
  - [https://cdn.jsdelivr.net/gh/arsity/geoip@release/surge/tor.txt](https://cdn.jsdelivr.net/gh/arsity/geoip@release/surge/tor.txt)

---

### Nginx `allow` 和 `deny` 格式文件

所有**国家/地区**、**新增类别**的 Nginx `allow` 和 `deny` 格式文件，请查看本项目 `release` 分支下的 [nginx 目录](https://github.com/arsity/geoip/tree/release/nginx)。

---

### 纯文本 txt 格式文件

所有**国家/地区**、**新增类别**的纯文本 txt 格式文件，请查看本项目 `release` 分支下的 [text 目录](https://github.com/arsity/geoip/tree/release/text)。

部分**国家/地区**类别：

- **中国大陆**（Mainland China）：
  - [https://raw.githubusercontent.com/arsity/geoip/release/text/cn.txt](https://raw.githubusercontent.com/arsity/geoip/release/text/cn.txt)
  - [https://cdn.jsdelivr.net/gh/arsity/geoip@release/text/cn.txt](https://cdn.jsdelivr.net/gh/arsity/geoip@release/text/cn.txt)
- **香港**（Hong Kong）：
  - [https://raw.githubusercontent.com/arsity/geoip/release/text/hk.txt](https://raw.githubusercontent.com/arsity/geoip/release/text/hk.txt)
  - [https://cdn.jsdelivr.net/gh/arsity/geoip@release/text/hk.txt](https://cdn.jsdelivr.net/gh/arsity/geoip@release/text/hk.txt)
- **澳门**（Macau）：
  - [https://raw.githubusercontent.com/arsity/geoip/release/text/mo.txt](https://raw.githubusercontent.com/arsity/geoip/release/text/mo.txt)
  - [https://cdn.jsdelivr.net/gh/arsity/geoip@release/text/mo.txt](https://cdn.jsdelivr.net/gh/arsity/geoip@release/text/mo.txt)
- **台湾**（Taiwan）：
  - [https://raw.githubusercontent.com/arsity/geoip/release/text/tw.txt](https://raw.githubusercontent.com/arsity/geoip/release/text/tw.txt)
  - [https://cdn.jsdelivr.net/gh/arsity/geoip@release/text/tw.txt](https://cdn.jsdelivr.net/gh/arsity/geoip@release/text/tw.txt)
- **美国**（America）：
  - [https://raw.githubusercontent.com/arsity/geoip/release/text/us.txt](https://raw.githubusercontent.com/arsity/geoip/release/text/us.txt)
  - [https://cdn.jsdelivr.net/gh/arsity/geoip@release/text/us.txt](https://cdn.jsdelivr.net/gh/arsity/geoip@release/text/us.txt)
- **日本**（Japan）：
  - [https://raw.githubusercontent.com/arsity/geoip/release/text/jp.txt](https://raw.githubusercontent.com/arsity/geoip/release/text/jp.txt)
  - [https://cdn.jsdelivr.net/gh/arsity/geoip@release/text/jp.txt](https://cdn.jsdelivr.net/gh/arsity/geoip@release/text/jp.txt)
- **韩国**（Korea）：
  - [https://raw.githubusercontent.com/arsity/geoip/release/text/kr.txt](https://raw.githubusercontent.com/arsity/geoip/release/text/kr.txt)
  - [https://cdn.jsdelivr.net/gh/arsity/geoip@release/text/kr.txt](https://cdn.jsdelivr.net/gh/arsity/geoip@release/text/kr.txt)
- **新加坡**（Singapore）：
  - [https://raw.githubusercontent.com/arsity/geoip/release/text/sg.txt](https://raw.githubusercontent.com/arsity/geoip/release/text/sg.txt)
  - [https://cdn.jsdelivr.net/gh/arsity/geoip@release/text/sg.txt](https://cdn.jsdelivr.net/gh/arsity/geoip@release/text/sg.txt)
- **缅甸**（Myanmar）：
  - [https://raw.githubusercontent.com/arsity/geoip/release/text/mm.txt](https://raw.githubusercontent.com/arsity/geoip/release/text/mm.txt)
  - [https://cdn.jsdelivr.net/gh/arsity/geoip@release/text/mm.txt](https://cdn.jsdelivr.net/gh/arsity/geoip@release/text/mm.txt)
- **伊朗**（Iran）：
  - [https://raw.githubusercontent.com/arsity/geoip/release/text/ir.txt](https://raw.githubusercontent.com/arsity/geoip/release/text/ir.txt)
  - [https://cdn.jsdelivr.net/gh/arsity/geoip@release/text/ir.txt](https://cdn.jsdelivr.net/gh/arsity/geoip@release/text/ir.txt)
- **俄罗斯**（Russia）：
  - [https://raw.githubusercontent.com/arsity/geoip/release/text/ru.txt](https://raw.githubusercontent.com/arsity/geoip/release/text/ru.txt)
  - [https://cdn.jsdelivr.net/gh/arsity/geoip@release/text/ru.txt](https://cdn.jsdelivr.net/gh/arsity/geoip@release/text/ru.txt)
- **白俄罗斯**（Belarus）：
  - [https://raw.githubusercontent.com/arsity/geoip/release/text/by.txt](https://raw.githubusercontent.com/arsity/geoip/release/text/by.txt)
  - [https://cdn.jsdelivr.net/gh/arsity/geoip@release/text/by.txt](https://cdn.jsdelivr.net/gh/arsity/geoip@release/text/by.txt)
- **土库曼斯坦**（Turkmenistan）：
  - [https://raw.githubusercontent.com/arsity/geoip/release/text/tm.txt](https://raw.githubusercontent.com/arsity/geoip/release/text/tm.txt)
  - [https://cdn.jsdelivr.net/gh/arsity/geoip@release/text/tm.txt](https://cdn.jsdelivr.net/gh/arsity/geoip@release/text/tm.txt)

**新增**类别：

- **cloudflare**：
  - [https://raw.githubusercontent.com/arsity/geoip/release/text/cloudflare.txt](https://raw.githubusercontent.com/arsity/geoip/release/text/cloudflare.txt)
  - [https://cdn.jsdelivr.net/gh/arsity/geoip@release/text/cloudflare.txt](https://cdn.jsdelivr.net/gh/arsity/geoip@release/text/cloudflare.txt)
- **cloudfront**：
  - [https://raw.githubusercontent.com/arsity/geoip/release/text/cloudfront.txt](https://raw.githubusercontent.com/arsity/geoip/release/text/cloudfront.txt)
  - [https://cdn.jsdelivr.net/gh/arsity/geoip@release/text/cloudfront.txt](https://cdn.jsdelivr.net/gh/arsity/geoip@release/text/cloudfront.txt)
- **facebook**：
  - [https://raw.githubusercontent.com/arsity/geoip/release/text/facebook.txt](https://raw.githubusercontent.com/arsity/geoip/release/text/facebook.txt)
  - [https://cdn.jsdelivr.net/gh/arsity/geoip@release/text/facebook.txt](https://cdn.jsdelivr.net/gh/arsity/geoip@release/text/facebook.txt)
- **fastly**：
  - [https://raw.githubusercontent.com/arsity/geoip/release/text/fastly.txt](https://raw.githubusercontent.com/arsity/geoip/release/text/fastly.txt)
  - [https://cdn.jsdelivr.net/gh/arsity/geoip@release/text/fastly.txt](https://cdn.jsdelivr.net/gh/arsity/geoip@release/text/fastly.txt)
- **google**：
  - [https://raw.githubusercontent.com/arsity/geoip/release/text/google.txt](https://raw.githubusercontent.com/arsity/geoip/release/text/google.txt)
  - [https://cdn.jsdelivr.net/gh/arsity/geoip@release/text/google.txt](https://cdn.jsdelivr.net/gh/arsity/geoip@release/text/google.txt)
- **netflix**：
  - [https://raw.githubusercontent.com/arsity/geoip/release/text/netflix.txt](https://raw.githubusercontent.com/arsity/geoip/release/text/netflix.txt)
  - [https://cdn.jsdelivr.net/gh/arsity/geoip@release/text/netflix.txt](https://cdn.jsdelivr.net/gh/arsity/geoip@release/text/netflix.txt)
- **spotify**：
  - [https://raw.githubusercontent.com/arsity/geoip/release/text/spotify.txt](https://raw.githubusercontent.com/arsity/geoip/release/text/spotify.txt)
  - [https://cdn.jsdelivr.net/gh/arsity/geoip@release/text/spotify.txt](https://cdn.jsdelivr.net/gh/arsity/geoip@release/text/spotify.txt)
- **telegram**：
  - [https://raw.githubusercontent.com/arsity/geoip/release/text/telegram.txt](https://raw.githubusercontent.com/arsity/geoip/release/text/telegram.txt)
  - [https://cdn.jsdelivr.net/gh/arsity/geoip@release/text/telegram.txt](https://cdn.jsdelivr.net/gh/arsity/geoip@release/text/telegram.txt)
- **twitter**：
  - [https://raw.githubusercontent.com/arsity/geoip/release/text/twitter.txt](https://raw.githubusercontent.com/arsity/geoip/release/text/twitter.txt)
  - [https://cdn.jsdelivr.net/gh/arsity/geoip@release/text/twitter.txt](https://cdn.jsdelivr.net/gh/arsity/geoip@release/text/twitter.txt)
- **tor**：
  - [https://raw.githubusercontent.com/arsity/geoip/release/text/tor.txt](https://raw.githubusercontent.com/arsity/geoip/release/text/tor.txt)
  - [https://cdn.jsdelivr.net/gh/arsity/geoip@release/text/tor.txt](https://cdn.jsdelivr.net/gh/arsity/geoip@release/text/tor.txt)

## 自行定制 GeoIP 文件

> [!NOTE]
> 默认配置使用 IPInfo Lite 数据。在线生成需要在 GitHub Actions secrets 中添加 `IPINFO_TOKEN`；本地生成需要提前准备 `./ipinfo/ipinfo_lite.csv` 和 `./ipinfo/ipinfo_lite.mmdb`。

### 定制方式

- **在线生成**：[Fork](https://github.com/Loyalsoldier/geoip/fork) 本仓库后，根据 [`configuration.md`](https://github.com/Loyalsoldier/geoip/blob/HEAD/configuration.md) 配置说明文档，修改自己仓库内的配置文件 `config.json` 和 GitHub Workflow `.github/workflows/build.yml`
- **本地生成**：
  - 安装 [Golang](https://go.dev/dl/) 和 [Git](https://git-scm.com)
  - 拉取项目代码: `git clone https://github.com/Loyalsoldier/geoip.git`
  - 进入项目根目录：`cd geoip`
  - 下载 IPInfo Lite CSV 和 MMDB 文件，放置为 `./ipinfo/ipinfo_lite.csv` 和 `./ipinfo/ipinfo_lite.mmdb`
  - 根据 [`configuration.md`](https://github.com/Loyalsoldier/geoip/blob/HEAD/configuration.md) 配置说明文档，修改配置文件 `config.json`
  - 运行代码：`go run ./ convert -c ./config.json`

### 配置文件概念解析

本项目配置文件 `config.json` 有两个概念：`input` 和 `output`。`input` 指数据源（data source）及其输入格式，`output` 指数据的去向（data destination）及其输出格式。CLI 的作用就是通过读取配置文件中的选项，聚合用户提供的所有数据源，去重，将其转换为目标格式，并输出到文件。

These two concepts in configuration file `config.json` are notable: `input` and `output`. The `input` is the data source and its input format, whereas the `output` is the destination of the converted data and its output format. What the CLI does is to aggregate all input format data, then convert them to output format and write them to GeoIP files by using the options in the config file.

### 支持的格式

关于每种格式所支持的配置选项，查看本项目 [`configuration.md`](https://github.com/Loyalsoldier/geoip/blob/HEAD/configuration.md) 配置说明文档。

支持的 `input` 输入格式：

- **text**：纯文本 IP 和 CIDR（例如：`1.1.1.1` 或 `1.0.0.0/24`）
- **stdin**：从 standard input 获取纯文本 IP 和 CIDR（例如：`1.1.1.1` 或 `1.0.0.0/24`）
- **private**：局域网和私有网络 CIDR（例如：`192.168.0.0/16` 和 `127.0.0.0/8`）
- **cutter**：用于裁剪前置步骤中的数据
- **json**：JSON 数据格式
- **v2rayGeoIPDat**：V2Ray GeoIP dat 数据格式（`geoip.dat`）
- **maxmindMMDB**：MaxMind GeoLite2 country mmdb 数据格式（`GeoLite2-Country.mmdb`）
- **maxmindGeoLite2ASNCSV**：MaxMind GeoLite2 ASN CSV 数据格式（`GeoLite2-ASN-CSV.zip`）
- **maxmindGeoLite2CountryCSV**：MaxMind GeoLite2 country CSV 数据格式（`GeoLite2-Country-CSV.zip`）
- **dbipCountryMMDB**：DB-IP country mmdb 数据格式（`dbip-country-lite.mmdb`）
- **ipinfoCountryMMDB**：IPInfo Lite mmdb 数据格式（`ipinfo_lite.mmdb`）
- **ipinfoLiteCountryCSV**：IPInfo Lite CSV 数据格式（国家/地区数据，`ipinfo_lite.csv`）
- **ipinfoLiteASNCSV**：IPInfo Lite CSV 数据格式（ASN 数据，`ipinfo_lite.csv`）
- **mihomoMRS**：mihomo MRS 数据格式（`geoip-cn.mrs`）
- **singboxSRS**：sing-box SRS 数据格式（`geoip-cn.srs`）
- **clashRuleSetClassical**：[classical 类型的 Clash RuleSet](https://wiki.metacubex.one/config/rule-providers/content/#classical)
- **clashRuleSet**：[ipcidr 类型的 Clash RuleSet](https://wiki.metacubex.one/config/rule-providers/content/#ipcidr)
- **surgeRuleSet**：[Surge RuleSet](https://manual.nssurge.com/rule/ruleset.html)

支持的 `output` 输出格式：

- **text**：纯文本 CIDR（例如：`1.0.0.0/24`）
- **stdout**：将纯文本 CIDR 输出到 standard output（例如：`1.0.0.0/24`）
- **lookup**：从指定的列表中查找指定的 IP 或 CIDR
- **v2rayGeoIPDat**：V2Ray GeoIP dat 数据格式（`geoip.dat`）
- **maxmindMMDB**：MaxMind GeoLite2 country mmdb 数据格式（`GeoLite2-Country.mmdb`）
- **dbipCountryMMDB**：DB-IP country mmdb 数据格式（`dbip-country-lite.mmdb`）
- **ipinfoCountryMMDB**：使用 IPInfo 元数据补全的 MaxMind-compatible country mmdb 数据格式（`Country.mmdb`）
- **mihomoMRS**：mihomo MRS 数据格式（`geoip-cn.mrs`）
- **singboxSRS**：sing-box SRS 数据格式（`geoip-cn.srs`）
- **clashRuleSetClassical**：[classical 类型的 Clash RuleSet](https://wiki.metacubex.one/config/rule-providers/content/#classical)
- **clashRuleSet**：[ipcidr 类型的 Clash RuleSet](https://wiki.metacubex.one/config/rule-providers/content/#ipcidr)
- **surgeRuleSet**：[Surge RuleSet](https://manual.nssurge.com/rule/ruleset.html)

### 注意事项

由于 MaxMind、DB-IP、IPInfo 的 mmdb 文件格式的限制，当不同列表的 IP 或 CIDR 数据有交集或重复项时，后写入的列表的 IP 或 CIDR 数据会覆盖（overwrite）之前已写入的列表的数据。譬如，IP `1.1.1.1` 同属于列表 `AU` 和列表 `Cloudflare`。如果 `Cloudflare` 在 `AU` 之后写入，则 IP `1.1.1.1` 归属于列表 `Cloudflare`。

为了确保某些指定的列表、被修改的列表一定囊括属于它的所有 IP 或 CIDR 数据，可在 `output` 相应输出格式的配置中增加选项 `overwriteList`，该选项中指定的列表会在最后逐一写入，列表中最后一项优先级最高。若已设置选项 `wantedList`，则无需设置 `overwriteList`。`wantedList` 中指定的列表会在最后逐一写入，列表中最后一项优先级最高。

## CLI 功能展示

可通过 `go install -v github.com/Loyalsoldier/geoip@latest` 直接安装 CLI 工具。

CLI 提供的功能如下：

- 列出支持的 `input` 和 `output` 格式（`list`）
- GeoIP 数据格式转换（`convert`）
- 查找 IP 或 CIDR 所在类别（`lookup`）
- 去重和合并 IP 与 CIDR（`merge`）

### 总览

```bash
$ ./geoip
geoip is a convenient tool to merge, convert and lookup IP & CIDR from various formats of geoip data.

Usage:
  geoip [command]

Available Commands:
  convert     Convert geoip data from one format to another by using config file
  help        Help about any command
  list        List all available input and output formats
  lookup      Lookup specified IP or CIDR in specified lists
  merge       Merge plaintext IP & CIDR from standard input, then print to standard output

Flags:
  -h, --help   help for geoip

Use "geoip [command] --help" for more information about a command.
```

### 列出支持的 `input` 和 `output` 格式（`list`）

```bash
$ ./geoip list
All available input formats:
  - clashRuleSet (Convert ipcidr type of Clash RuleSet to other formats)
  - clashRuleSetClassical (Convert classical type of Clash RuleSet to other formats (just processing IP & CIDR lines))
  - cutter (Remove data from previous steps)
  - dbipCountryMMDB (Convert DB-IP country mmdb database to other formats)
  - ipinfoCountryMMDB (Convert IPInfo Lite mmdb database to other formats)
  - ipinfoLiteASNCSV (Convert IPInfo Lite ASN CSV data to other formats)
  - ipinfoLiteCountryCSV (Convert IPInfo Lite country CSV data to other formats)
  - json (Convert JSON data to other formats)
  - maxmindGeoLite2ASNCSV (Convert MaxMind GeoLite2 ASN CSV data to other formats)
  - maxmindGeoLite2CountryCSV (Convert MaxMind GeoLite2 country CSV data to other formats)
  - maxmindMMDB (Convert MaxMind mmdb database to other formats)
  - mihomoMRS (Convert mihomo MRS data to other formats)
  - private (Convert LAN and private network CIDR to other formats)
  - singboxSRS (Convert sing-box SRS data to other formats)
  - stdin (Accept plaintext IP & CIDR from standard input, separated by newline)
  - surgeRuleSet (Convert Surge RuleSet to other formats (just processing IP & CIDR lines))
  - test (Convert specific CIDR to other formats (for test only))
  - text (Convert plaintext IP & CIDR to other formats)
  - v2rayGeoIPDat (Convert V2Ray GeoIP dat to other formats)

All available output formats:
  - clashRuleSet (Convert data to ipcidr type of Clash RuleSet)
  - clashRuleSetClassical (Convert data to classical type of Clash RuleSet)
  - dbipCountryMMDB (Convert data to DB-IP country mmdb database format)
  - ipinfoCountryMMDB (Convert data to MaxMind-compatible mmdb format with IPInfo metadata)
  - lookup (Lookup specified IP or CIDR from various formats of data)
  - maxmindMMDB (Convert data to MaxMind mmdb database format)
  - mihomoMRS (Convert data to mihomo MRS format)
  - singboxSRS (Convert data to sing-box SRS format)
  - stdout (Convert data to plaintext CIDR format and output to standard output)
  - surgeRuleSet (Convert data to Surge RuleSet)
  - text (Convert data to plaintext CIDR format)
  - v2rayGeoIPDat (Convert data to V2Ray GeoIP dat format)
```

### 去重和合并 IP 与 CIDR（`merge`）

```bash
$ curl -s https://core.telegram.org/resources/cidr.txt | ./geoip merge -t ipv4
91.105.192.0/23
91.108.4.0/22
91.108.8.0/21
91.108.16.0/21
91.108.56.0/22
149.154.160.0/20
185.76.151.0/24
```

### GeoIP 数据格式转换（`convert`）

```bash
$ ./geoip convert -c config.json
2021/08/29 12:11:35 ✅ [v2rayGeoIPDat] geoip.dat --> output/dat
2021/08/29 12:11:35 ✅ [v2rayGeoIPDat] geoip-only-cn-private.dat --> output/dat
2021/08/29 12:11:35 ✅ [v2rayGeoIPDat] geoip-asn.dat --> output/dat
2021/08/29 12:11:35 ✅ [v2rayGeoIPDat] cn.dat --> output/dat
2021/08/29 12:11:35 ✅ [v2rayGeoIPDat] private.dat --> output/dat
2021/08/29 12:11:39 ✅ [maxmindMMDB] Country.mmdb --> output/maxmind
2021/08/29 12:11:39 ✅ [maxmindMMDB] Country-only-cn-private.mmdb --> output/maxmind
2021/08/29 12:11:39 ✅ [text] netflix.txt --> output/text
2021/08/29 12:11:39 ✅ [text] telegram.txt --> output/text
2021/08/29 12:11:39 ✅ [text] cn.txt --> output/text
2021/08/29 12:11:39 ✅ [text] cloudflare.txt --> output/text
2021/08/29 12:11:39 ✅ [text] cloudfront.txt --> output/text
2021/08/29 12:11:39 ✅ [text] facebook.txt --> output/text
2021/08/29 12:11:39 ✅ [text] fastly.txt --> output/text
2021/08/29 12:11:45 ✅ [singboxSRS] netflix.txt --> output/srs
2021/08/29 12:11:45 ✅ [singboxSRS] telegram.txt --> output/srs
2021/08/29 12:11:45 ✅ [singboxSRS] cn.txt --> output/srs
2021/08/29 12:11:45 ✅ [singboxSRS] cloudflare.txt --> output/srs
2021/08/29 12:11:45 ✅ [singboxSRS] cloudfront.txt --> output/srs
2021/08/29 12:11:45 ✅ [singboxSRS] facebook.txt --> output/srs
2021/08/29 12:11:45 ✅ [singboxSRS] fastly.txt --> output/srs
2021/08/29 12:11:50 ✅ [mihomoMRS] netflix.txt --> output/mrs
2021/08/29 12:11:50 ✅ [mihomoMRS] telegram.txt --> output/mrs
2021/08/29 12:11:50 ✅ [mihomoMRS] cn.txt --> output/mrs
2021/08/29 12:11:50 ✅ [mihomoMRS] cloudflare.txt --> output/mrs
2021/08/29 12:11:50 ✅ [mihomoMRS] cloudfront.txt --> output/mrs
2021/08/29 12:11:50 ✅ [mihomoMRS] facebook.txt --> output/mrs
2021/08/29 12:11:50 ✅ [mihomoMRS] fastly.txt --> output/mrs
```

### 查找 IP 或 CIDR 所在类别（`lookup`）

可能的返回结果：

- 查询的字符串不是有效的 IP 或 CIDR，返回 `false`
- 查询的 IP 或 CIDR 不存在于任何一个类别中，返回 `false`
- 查询的 IP 或 CIDR 存在于某种格式文件的单个类别中：
  - 若该格式文件只包含一个类别，返回 `true`
  - 若该格式文件包含多个类别，返回匹配的类别名称
- 查询的 IP 或 CIDR 存在于多个类别中，返回以英文逗号分隔的类别名称，如 `au,cloudflare`

```bash
# ================= One-time Mode ================= #

# 从 text 格式的本地文件（只包含一个类别）中查找某个 IP 地址
# lookup IP from local file (with only one list) in text format
$ ./geoip lookup -f text -u ./cn.txt 1.0.1.1
true


# 从 text 格式的本地文件（只包含一个类别）中查找某个 IP 地址
# lookup IP from local file (with only one list) in text format
$ ./geoip lookup -f text -u ./cn.txt 2.2.2.2
false


# 从 text 格式的本地文件（只包含一个类别）中查找某个 CIDR
# lookup CIDR from local file (with only one list) in text format
$ ./geoip lookup -f text -u ./cn.txt 1.0.1.1/24
true


# 从 text 格式的本地文件（只包含一个类别）中查找某个 CIDR
# lookup CIDR from local file (with only one list) in text format
$ ./geoip lookup -f text -u ./cn.txt 1.0.1.1/23
false


# 从 text 格式的远程 URL（只包含一个类别）中查找某个 IP 地址
# lookup IP from remote URL (with only one list) in text format
$ ./geoip lookup -f text -u https://example.com/cn.txt 1.0.1.1
true


# 从 v2rayGeoIPDat 格式的本地文件（只包含一个类别）中查找某个 IP 地址
# lookup IP from local file (with only one list) in v2rayGeoIPDat format
$ ./geoip lookup -f v2rayGeoIPDat -u ./cn.dat 1.0.1.1
true


# 从 v2rayGeoIPDat 格式的本地文件（包含多个类别）中查找某个 IP 地址
# lookup IP from local file (with multiple lists) in v2rayGeoIPDat format
$ ./geoip lookup -f v2rayGeoIPDat -u ./geoip.dat 1.0.1.1
cn


# 从 v2rayGeoIPDat 格式的本地文件（包含多个类别）中查找某个 IP 地址
# lookup IP from local file (with multiple lists) in v2rayGeoIPDat format
$ ./geoip lookup -f v2rayGeoIPDat -u ./geoip.dat 1.0.0.1
au,cloudflare


# 从 v2rayGeoIPDat 格式的远程 URL（包含多个类别）中查找某个 CIDR
# lookup CIDR from remote URL (with multiple lists) in v2rayGeoIPDat format
$ ./geoip lookup -f v2rayGeoIPDat -u https://example.com/geoip.dat 1.0.0.1/24
au,cloudflare




# ================= REPL Mode ================= #

# 从 text 格式的本地文件（只包含一个类别）中查找某个 IP 地址或 CIDR
# lookup IP or CIDR from local file (with only one list) in text format
$ ./geoip lookup -f text -u ./cn.txt
Enter IP or CIDR (type "exit" to quit):
>> 1.0.1.1
true

>> 1.0.1.1/24
true

>> 1.0.1.1/23
false

>> 2.2.2.2
false

>> 2.2.2.2/24
false

>> 300.300.300.300
false

>> 300.300.300.300/24
false

>> exit


# 从 text 格式的远程 URL（只包含一个类别）中查找某个 IP 地址或 CIDR
# lookup IP or CIDR from remote URL (with only one list) in text format
$ ./geoip lookup -f text -u https://example.com/cn.txt
Enter IP or CIDR (type "exit" to quit):
>> 1.0.1.1
true

>> 1.0.1.1/24
true

>> 1.0.1.1/23
false

>> 2.2.2.2
false

>> 2.2.2.2/24
false

>> 300.300.300.300
false

>> 300.300.300.300/24
false

>> exit


# 从 v2rayGeoIPDat 格式的本地文件（只包含一个类别）中查找某个 IP 地址或 CIDR
# lookup IP or CIDR from local file (with only one list) in v2rayGeoIPDat format
$ ./geoip lookup -f v2rayGeoIPDat -u ./cn.dat
Enter IP or CIDR (type "exit" to quit):
>> 1.0.1.1
true

>> 1.0.1.1/24
true

>> 1.0.1.1/23
false

>> 2.2.2.2
false

>> 2.2.2.2/24
false

>> 300.300.300.300
false

>> 300.300.300.300/24
false

>> exit


# 从 v2rayGeoIPDat 格式的远程 URL（包含多个类别）中查找某个 IP 地址或 CIDR
# lookup IP or CIDR from remote URL (with multiple list) in v2rayGeoIPDat format
$ ./geoip lookup -f v2rayGeoIPDat -u https://example.com/geoip.dat
Enter IP or CIDR (type "exit" to quit):
>> 1.0.1.1
cn

>> 1.0.1.1/24
cn

>> 1.0.1.1/23
false

>> 1.0.0.1
au,cloudflare

>> 1.0.0.1/24
au,cloudflare

>> 300.300.300.300
false

>> 300.300.300.300/24
false

>> exit
```

## 使用本项目的项目

- [@Loyalsoldier/v2ray-rules-dat](https://github.com/Loyalsoldier/v2ray-rules-dat)
- [@Loyalsoldier/clash-rules](https://github.com/Loyalsoldier/clash-rules)
- [@Loyalsoldier/surge-rules](https://github.com/Loyalsoldier/surge-rules)

## License

[CC-BY-SA-4.0](https://creativecommons.org/licenses/by-sa/4.0/) and [GPL-3.0](https://github.com/Loyalsoldier/geoip/blob/master/LICENSE-GPL)

Default generated artifacts include IPInfo Lite data created by [IPInfo](https://ipinfo.io). The CLI also supports MaxMind GeoLite2 data when users provide their own MaxMind source files.

## 项目 Star 数增长趋势

<a href="https://www.star-history.com/?repos=Loyalsoldier%2Fgeoip&type=date&legend=top-left">
 <picture>
   <source media="(prefers-color-scheme: dark)" srcset="https://api.star-history.com/chart?repos=Loyalsoldier/geoip&type=date&theme=dark&legend=top-left&sealed_token=esf8pGs-3KVdrX__Wt8TDQgH-o8msKwSDVxZqJfNGHLlsNEjsUx8QPPOe7eW5m8eL0wQb81hwELIri8RWP2XCXrRnrN4sq4F_gb4Vc244ZdIvo2iA1Fk0w" />
   <source media="(prefers-color-scheme: light)" srcset="https://api.star-history.com/chart?repos=Loyalsoldier/geoip&type=date&legend=top-left&sealed_token=esf8pGs-3KVdrX__Wt8TDQgH-o8msKwSDVxZqJfNGHLlsNEjsUx8QPPOe7eW5m8eL0wQb81hwELIri8RWP2XCXrRnrN4sq4F_gb4Vc244ZdIvo2iA1Fk0w" />
   <img alt="Star History Chart" src="https://api.star-history.com/chart?repos=Loyalsoldier/geoip&type=date&legend=top-left&sealed_token=esf8pGs-3KVdrX__Wt8TDQgH-o8msKwSDVxZqJfNGHLlsNEjsUx8QPPOe7eW5m8eL0wQb81hwELIri8RWP2XCXrRnrN4sq4F_gb4Vc244ZdIvo2iA1Fk0w" />
 </picture>
</a>
