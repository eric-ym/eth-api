## 接口文档

1. 地址排名

   接口地址 ：/v1/api/home/address

   Method：POST

   Params：无

   Returns：

| 名称 | 类型 | 说明 |
| --- | --- | --- |
|  rank|  int| 排名 |
| account | string  | 地址 |
|bakance|float| 余额|
|percentage|string|占比|
|type|string|地址类型|
|lastUpdate|int|最后更新时间|
|recentTransfer|int|最后交易时间|

exp:

```
[
  {
    "rank": 1,
    "account": "0xa0e89b1bf287de207411f3523fe794b2ab67c845",
    "balance": 60325.062433434,
    "percentage": "",
    "type": 0,
    "lastUpdate": 1692090425,
    "recentTransfer": 1690959545
  },
  {
    "rank": 2,
    "account": "0x9967f1d491a9119c31465832b3b4aa1c80c53277",
    "balance": 59619.500021,
    "percentage": "",
    "type": 0,
    "lastUpdate": 1692090420,
    "recentTransfer": 0
  },
  {
    "rank": 3,
    "account": "0x38824715dc4aea51881586b9d1ed4e2798b2b037",
    "balance": 57540.312545565,
    "percentage": "",
    "type": 0,
    "lastUpdate": 1692090402,
    "recentTransfer": 1690959373
  },
  {
    "rank": 4,
    "account": "0xf65c2e5c73116c5c676d5436fadcb63b58edc5f0",
    "balance": 1e-9,
    "percentage": "",
    "type": 0,
    "lastUpdate": 1690959545,
    "recentTransfer": 1690959545
  }
]
```

2. 24小时出块奖励

   接口地址：/v1/api/home/24hours

   Method:POST

   Params：

   Returns：

   | 名称 | 类型 | 说明 |
   | --- | --- | --- |
   |  24hbonus|  int| 出块总奖励 |

   exp:

   ```
   {
   	"24hbonus": 12794
   }
   ```
3. Block

   接口地址：/v1/api/home/standard

   Method:POST

   Params：

   Returns：

   | 名称 | 类型 | 说明 |
   | --- | --- | --- |
   |  24hbonus|  int| 出块总奖励 |
   |lastHeight|int| 最后出块高度|
   |lastUpdate|int|最后出块时间|
   |hashPower| string| 全网有效算力|
   |power24|string| 24小时算力增涨|
   |productivity24|string| 24小时出块效率|


exp:

```
{
  "lastHeight": 88550,
  "lastUpdate": 1692090425,
  "hashPower": "2.74040M",
  "power24": "23.56000K",
  "productivity24": "0.36648562 xx/M"
}
```

4. 12小时算力趋势

   接口地址：/v1/api/home/trend_echarts

   Method:POST

   Params：

   Returns：

   | 名称 | 类型 | 说明 |
   | --- | --- | --- |
   |  hour|  int| 小时数 |
   |power| string| 算力|
   |incr| string | 较前一小时的算力增长|

   exp:

```
[
  {
    "hour": 6,
    "power": "2.78706M",
    "incr": "0.00000"
  },
  {
    "hour": 7,
    "power": "2.78706M",
    "incr": "0.00000"
  },
  {
    "hour": 8,
    "power": "2.71284M",
    "incr": "-74228.00000"
  },
  {
    "hour": 9,
    "power": "2.71284M",
    "incr": "0.00000"
  },
  {
    "hour": 10,
    "power": "2.71284M",
    "incr": "0.00000"
  },
  {
    "hour": 11,
    "power": "2.71284M",
    "incr": "0.00000"
  },
  {
    "hour": 12,
    "power": "2.71284M",
    "incr": "0.00000"
  },
  {
    "hour": 13,
    "power": "2.71284M",
    "incr": "0.00000"
  },
  {
    "hour": 14,
    "power": "2.71284M",
    "incr": "0.00000"
  },
  {
    "hour": 15,
    "power": "2.71284M",
    "incr": "0.00000"
  },
  {
    "hour": 16,
    "power": "2.71284M",
    "incr": "0.00000"
  },
  {
    "hour": 17,
    "power": "2.71284M",
    "incr": "0.00000"
  }
]
```

5. block

   接口地址：/v1/api/blocks

   Method:POST

   Params：

   | 名称 | 类型 | 说明 |
   | --- | --- | --- |
   |  start|  int| 查询的起始高度|
   | limit | int | 查询的长度|

   Returns：

   | 名称 | 类型 | 说明 |
   | --- | --- | --- |
   |  24hbonus|  int| 出块总奖励 |
   |blockID| string| block 的hash|
   |height| int| 高度|
   |time|int| 生成时间|
   |data| data | 块中消息列表|

   Data:

   | 名称 | 类型 | 说明 |
   | --- | --- | --- |
   |  cid|  string | 交易hash|
   |message | int| 消息数量|
   |reward|string|奖励|


exp:

```
[
  {
    "blockID": "0x192a32c860a87fcc4ba3f5c49fe8814825648a308577cd76fcafce0e819fd5b2",
    "height": 40323,
    "time": 1691481351,
    "data": []
  },
  {
    "blockID": "0x059df4d23a2366e4ce459c151c2b2e524165252fc488b3311e60513c2d5646b4",
    "height": 40324,
    "time": 1691481355,
    "data": []
  },
  {
    "blockID": "0x10fe7638a4ba2626b8841798e32e32ceca1222f9d680a88e72456097cf4b084c",
    "height": 40325,
    "time": 1691481370,
    "data": []
  },
  {
    "blockID": "0x324aaeda6c975ec7358a8ad3c831eda66cb87734ab4b8b6a5b7cf132083fa22b",
    "height": 40326,
    "time": 1691481374,
    "data": []
  },
  {
    "blockID": "0xe9d27b94f06b68a3a8982a40147109008e419343f3729e3580610c12c0394e79",
    "height": 40327,
    "time": 1691481380,
    "data": []
  },
  {
    "blockID": "0xdc5df3cc4623e2061a144c80f671bdbc328f9851df11a2e1e646527a4bd45f2d",
    "height": 40328,
    "time": 1691481411,
    "data": []
  },
  {
    "blockID": "0x748e7a1c840e3b786af64fbf2df64211bd63620c73d29fef4a597aac7c88415d",
    "height": 40329,
    "time": 1691481412,
    "data": []
  },
  {
    "blockID": "0x20f4242803cea2d3402bc3e94046fddcd6529b67a409dc599c3665bb2e21dcb1",
    "height": 40330,
    "time": 1691481423,
    "data": []
  },
  {
    "blockID": "0xf90b930681cf2c50df0142dd01a92d0ed31cfee2b0de6de1e6757803dee6238d",
    "height": 40331,
    "time": 1691481427,
    "data": []
  },
  {
    "blockID": "0xda4a2949d0d92e665d9a79b9fff554f3ddf469d4f50d7bf2a770ca5fe4e22591",
    "height": 40332,
    "time": 1691481441,
    "data": []
  }
]
```

6. 单个块的详细信息

   接口地址：/v1/api/blocks/detail/info

   Method:POST

   Params：

   | 名称 | 类型 | 说明 |
   | --- | --- | --- |
   |  hash|  string| block的hash |

   Returns：

   | 名称 | 类型 | 说明 |
   | --- | --- | --- |
   |  cid|  string| 块的hash|
   |height| int| 块的高度|
   |time| int| 块生成的时间|
   |message |int| 块中包含的消息数量|
   |parentCid| string| 父级块的hash值|
   | parentWeight| int| 父级块的高度|
   |reward| string| 获得的奖励|

exp:

```
{
  "Cid": "0x46247a5d98c2c5124864ca51deb5e2a4dbfa3efc21f5ca892617e8ce06d7be58",
  "height": 32,
  "time": 1690959373,
  "message": 1,
  "parentCid": "0xc0ae0c79833d1bcd53fdd2ccd23d62e486630995f1001f1c63eb0d03d15419fb",
  "parentWeight": 31,
  "stateRoot": "32",
  "Reward": "0.000021"
}
```

7. 单个块中的交易信息

   接口地址：/v1/api/blocks/detail/message

   Method:POST

   Params：

   | 名称 | 类型 | 说明 |
   | --- | --- | --- |
   |  hash|  string| block的hash |
   |page| int| 页数|
   |limit| int| 每页数据数|


   Returns：

| 名称 | 类型 | 说明 |
| --- | --- | --- |
|  total|  int| 数据总数|
|page| int| 返回的页数|
|limit| int| 返回的每页数据数|
|data |[]obejct| 具体数据|

data:

| 名称 | 类型 | 说明 |
| --- | --- | --- |
|id| string| 交易hash|
| height| int| 高度|
|time| int| 生成时间|
|from| string| 发起者|
|to| string| 接受者|
|value| string| 交易金额|
|status| string|交易状态|
|method| string| 方法|

	exp:

```
{
  "total": 1,
  "page": 1,
  "limit": 10,
  "data": [
    {
      "id": "0x2bcb55cdef0f5f8c3c105ed73c19e9e019713c0a61e1b34a3853b7c74d0e5053",
      "height": 32,
      "block": "",
      "time": 1690959373,
      "from": "0x38824715dc4aea51881586b9d1ed4e2798b2b037",
      "to": "0xa0e89b1bf287de207411f3523fe794b2ab67c845",
      "value": "50.0000000000",
      "status": "OK",
      "method": ""
    }
  ]
}
```

---

8. 交易列表

接口地址：/v1/api/transactions

Method:POST

Params：

| 名称 | 类型 | 说明 |
| --- | --- | --- |
|page| int| 页数|
|limit| int| 每页数据数|

Returns：


| 名称 | 类型 | 说明 |
| --- | --- | --- |
|  total|  int| 数据总数|
|page| int| 返回的页数|
|limit| int| 返回的每页数据数|
|data |[]obejct| 具体数据|

data:

| 名称 | 类型 | 说明 |
| --- | --- | --- |
|id| string| 交易hash|
| height| int| 高度|
|block| string| 交易所在块hash|
|time| int| 生成时间|
|from| string| 发起者|
|to| string| 接受者|
|value| string| 交易金额|
|status| string|交易状态|
|method| string| 方法|

```
exp:
```

```
{
  "total": 2,
  "page": 1,
  "limit": 10,
  "data": [
    {
      "id": "0x2bcb55cdef0f5f8c3c105ed73c19e9e019713c0a61e1b34a3853b7c74d0e5053",
      "height": 32,
      "block": "0x46247a5d98c2c5124864ca51deb5e2a4dbfa3efc21f5ca892617e8ce06d7be58",
      "time": 1690959373,
      "from": "0x38824715dc4aea51881586b9d1ed4e2798b2b037",
      "to": "0xa0e89b1bf287de207411f3523fe794b2ab67c845",
      "value": "50.0000000000",
      "status": "OK",
      "method": "Send"
    },
    {
      "id": "0x476861cad380a891a3afc719cff7d676ce0562d5edcf1d6d5e790f1dfe1bfd66",
      "height": 36,
      "block": "0x9476596819211317cc5263e78467fe4001586aa3dc6b0da15e332dd0fc57d065",
      "time": 1690959545,
      "from": "0xa0e89b1bf287de207411f3523fe794b2ab67c845",
      "to": "0xf65c2e5c73116c5c676d5436fadcb63b58edc5f0",
      "value": "0.0000000010",
      "status": "OK",
      "method": "Send"
    }
  ]
}
```

---

9. 地址列表

   接口地址：/v1/api/address

   Method:POST

   Params：

   | 名称 | 类型 | 说明 |
   | --- | --- | --- |
   |page| int| 页数|
   |limit| int| 每页数据数|

   Returns：

   | 名称 | 类型 | 说明 |
   | --- | --- | --- |
   |  total|  int| 数据总数|
   |page| int| 返回的页数|
   |limit| int| 返回的每页数据数|
   |data |[]obejct| 具体数据|

   data:

   | 名称 | 类型 | 说明 |
   | --- | --- | --- |
   |  rank|  int| 排名 |
   | account | string  | 地址 |
   |bakance|float| 余额|
   |percentage|string|占比|
   |type|string|地址类型|
   |lastUpdate|int|最后更新时间|
   |recentTransfer|int|最后交易时间|


exp:

```
{
  "total": 4,
  "page": 1,
  "limit": 10,
  "data": [
    {
      "rank": 1,
      "account": "0xa0e89b1bf287de207411f3523fe794b2ab67c845",
      "balance": 60325.062433434,
      "percentage": "",
      "type": 0,
      "lastUpdate": 1692090425,
      "recentTransfer": 1690959545
    },
    {
      "rank": 2,
      "account": "0x9967f1d491a9119c31465832b3b4aa1c80c53277",
      "balance": 59619.500021,
      "percentage": "",
      "type": 0,
      "lastUpdate": 1692090420,
      "recentTransfer": 0
    },
    {
      "rank": 3,
      "account": "0x38824715dc4aea51881586b9d1ed4e2798b2b037",
      "balance": 57540.312545565,
      "percentage": "",
      "type": 0,
      "lastUpdate": 1692090402,
      "recentTransfer": 1690959373
    },
    {
      "rank": 4,
      "account": "0xf65c2e5c73116c5c676d5436fadcb63b58edc5f0",
      "balance": 1e-9,
      "percentage": "",
      "type": 0,
      "lastUpdate": 1690959545,
      "recentTransfer": 1690959545
    }
  ]
}
```

---

10. 地址交易信息

   接口地址：/v1/api/address/detail/list/transaction

   Method:POST

Params：

| 名称 | 类型 | 说明 |
| --- | --- | --- |
|page| int| 页数|
|limit| int| 每页数据数|
|hash| string| 地址hash值|

Returns：

| 名称 | 类型 | 说明 |
| --- | --- | --- |
|  total|  int| 数据总数|
|page| int| 返回的页数|
|limit| int| 返回的每页数据数|
|data |[]obejct| 具体数据|

data:

| 名称 | 类型 | 说明 |
| --- | --- | --- |
|id| string| 交易hash|
| height| int| 高度|
|time| int| 生成时间|
|from| string| 发起者|
|to| string| 接受者|
|value| string| 交易金额|
|status| string|交易状态|
|method| string| 方法|

exp:

```
{
  "total": 1,
  "page": 1,
  "limit": 10,
  "data": [
    {
      "time": 1690959373,
      "id": "0x2bcb55cdef0f5f8c3c105ed73c19e9e019713c0a61e1b34a3853b7c74d0e5053",
      "from": "0x38824715dc4aea51881586b9d1ed4e2798b2b037",
      "to": "0xa0e89b1bf287de207411f3523fe794b2ab67c845",
      "value": "50.0000000000",
      "method": "Send"
    }
  ]
}
```

---

11. 地址消息列表

接口地址：/v1/api/address/detail/list/message

Method:POST

Params：

| 名称 | 类型 | 说明 |
| --- | --- | --- |
|page| int| 页数|
|limit| int| 每页数据数|
|hash| string| 地址hash值|

Returns：

| 名称 | 类型 | 说明 |
| --- | --- | --- |
|  total|  int| 数据总数|
|page| int| 返回的页数|
|limit| int| 返回的每页数据数|
|data |[]obejct| 具体数据|

data:

| 名称 | 类型 | 说明 |
| --- | --- | --- |
|id| string| 交易hash|
| height| int| 高度|
|time| int| 生成时间|
|from| string| 发起者|
|to| string| 接受者|
|value| string| 交易金额|
|status| string|交易状态|
|method| string| 方法|

exp:

```
{
  "total": 1,
  "page": 1,
  "limit": 10,
  "data": [
    {
      "time": 1690959373,
      "id": "0x2bcb55cdef0f5f8c3c105ed73c19e9e019713c0a61e1b34a3853b7c74d0e5053",
      "from": "0x38824715dc4aea51881586b9d1ed4e2798b2b037",
      "to": "0xa0e89b1bf287de207411f3523fe794b2ab67c845",
      "value": "50.0000000000",
      "method": "Send"
    }
  ]
}
```

---

12. 地址详情

接口地址：/v1/api/address/detail/info

Method:POST

Params：

| 名称 | 类型 | 说明 |
| --- | --- | --- |
|hash| string| 地址hash值|

Returns：

| 名称 | 类型 | 说明 |
| --- | --- | --- |
|address| string| 地址hash|
| accountID| int| 存储id|
|message| int| 消息数量|
|nounce| 未知| 未知|
|cid|未知|  未知|
|createTime| string| 交易金额|
|lastUpdate| int|更新时间|

exp:

```
{
  "address": "0x38824715dc4aea51881586b9d1ed4e2798b2b037",
  "accountID": 36,
  "balance": "57888.3750455650",
  "message": 1,
  "nounce": "",
  "cid": "",
  "createTime": 0,
  "lastUpdate": 1692090402
}
```

---

13. 地址连续7小时余额变化

接口地址：/v1/api/address/detail/trend

Method:POST

Params：

| 名称 | 类型 | 说明 |
| --- | --- | --- |
|hash| string| 地址hash值|

Returns：

exp:

```
{
  "xAxis": {
    "data": [
      "18:00",
      "17:00",
      "16:00",
      "15:00",
      "14:00",
      "13:00",
      "12:00",
      "11:00"
    ] // 时间
  },
  "yAxis": {
    "type": "value"
  },
  "series": {
    "data": [
      "0.000000",
      "57540.312546",
      "57352.312546",
      "57178.312546",
      "56964.312546",
      "56796.312546",
      "56590.312546",
      "56412.312546"
    ] // 余额
  }
}
```

---

14. 消息列表

接口地址：/v1/api/message

Method:POST

Params：



| 名称 | 类型 | 说明 |
| --- | --- | --- |
|page| int| 页数|
|limit| int| 每页数据数|


Returns：



| 名称 | 类型 | 说明 |
| --- | --- | --- |
|  total|  int| 数据总数|
|page| int| 返回的页数|
|limit| int| 返回的每页数据数|
|data |[]obejct| 具体数据|

data:

| 名称 | 类型 | 说明 |
| --- | --- | --- |
|id| string| 交易hash|
| height| int| 高度|
|time| int| 生成时间|
|block| string| 消息存储的块hash|
|from| string| 发起者|
|to| string| 接受者|
|value| string| 交易金额|
|status| string|交易状态|
|method| string| 方法|

exp:

```
{
  "total": 2,
  "page": 1,
  "limit": 10,
  "data": [
    {
      "id": "0x2bcb55cdef0f5f8c3c105ed73c19e9e019713c0a61e1b34a3853b7c74d0e5053",
      "height": 32,
      "block": "0x46247a5d98c2c5124864ca51deb5e2a4dbfa3efc21f5ca892617e8ce06d7be58",
      "time": 1690959373,
      "from": "0x38824715dc4aea51881586b9d1ed4e2798b2b037",
      "to": "0xa0e89b1bf287de207411f3523fe794b2ab67c845",
      "value": "50.0000000000",
      "status": "OK",
      "method": "Send"
    },
    {
      "id": "0x476861cad380a891a3afc719cff7d676ce0562d5edcf1d6d5e790f1dfe1bfd66",
      "height": 36,
      "block": "0x9476596819211317cc5263e78467fe4001586aa3dc6b0da15e332dd0fc57d065",
      "time": 1690959545,
      "from": "0xa0e89b1bf287de207411f3523fe794b2ab67c845",
      "to": "0xf65c2e5c73116c5c676d5436fadcb63b58edc5f0",
      "value": "0.0000000010",
      "status": "OK",
      "method": "Send"
    }
  ]
}
```

---

15. 消息详情

接口地址：/v1/api/message/detail/detail


Method:POST

Params：

| 名称 | 类型 | 说明 |
| --- | --- | --- |
|hash| string | 消息hash值|


Returns：

| 名称 | 类型 | 说明 |
| --- | --- | --- |
|version| string| 接口版本好|
| nonce| int| 高度|
|gasLimit| string| 剩余燃料|
|gasUsed| string| 消耗原料|
|gasFee| string| 消耗原料转换成币值|


exp:

```
{
  "version": "2.0",
  "nonce": "0x0",
  "gasLimit": "0.0000000042",
  "gasUsed": "0x5208",
  "gasFee": "0.0000210000",
}
```

---


16. 消息详情

接口地址：/v1/api/message/detail/info

Method:POST

Params：

| 名称 | 类型 | 说明 |
| --- | --- | --- |
|hash| string | 消息hash值|

Returns：



| 名称 | 类型 | 说明 |
| --- | --- | --- |
|id| string| 交易hash|
| height| int| 高度|
|block| string| 交易所在块hash|
|time| int| 生成时间|
|from| string| 发起者|
|to| string| 接受者|
|value| string| 交易金额|
|status| string|交易状态|
|method| string| 方法|

exp:

```
{
  "id": "0x2bcb55cdef0f5f8c3c105ed73c19e9e019713c0a61e1b34a3853b7c74d0e5053",
  "height": 32,
  "block": "0x46247a5d98c2c5124864ca51deb5e2a4dbfa3efc21f5ca892617e8ce06d7be58",
  "time": 1690959373,
  "from": "0x38824715dc4aea51881586b9d1ed4e2798b2b037",
  "to": "0xa0e89b1bf287de207411f3523fe794b2ab67c845",
  "value": "50.0000000000",
  "status": "OK",
  "method": "Send"
}
```

---

17. 消息详情

接口地址：/v1/api/message/detail/transaction

Method:POST

Params：

| 名称 | 类型 | 说明 |
| --- | --- | --- |
|hash| string | 消息hash值|

Returns：

[]object

| 名称 | 类型 | 说明 |
| --- | --- | --- |
|from| string| 发起者|
|to| string| 接受者|
|value| string| 交易金额|
|type| string|交易类型|

exp:

```
[
  {
    "from": "0xa0e89b1bf287de207411f3523fe794b2ab67c845",
    "to": "0xf65c2e5c73116c5c676d5436fadcb63b58edc5f0",
    "value": "0.0000000010",
    "type": "0x0"
  }
]
```

---
