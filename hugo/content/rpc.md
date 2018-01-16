---
weight: 13
title: API Reference
---

# neo-insight

加强版的 NEO jsonrpc 2.0 远程服务

## 支持特性

1. 获取用户可用资产列表 utxo
2. 获取 claim utxo 列表用于 claim transaction 调用

所有扩展调用都位于**/extend**路径下面


## 其它接口

<aside class="success">
注意 — NEO支持的标准RPC接口参考<a href="http://docs.neo.org/zh-cn/node/api.html">这里</a>
</aside>


## 获取用户可用资产列表

> 请求用户的资产utxo列表 :

```json 
{
  "jsonrpc": "2.0",
  "method": "balance",
  "params": ["AMpupnF6QweQXLfCtF4dR45FDdKbTXkLsr","0xc56f33fc6ecfcd0c225c4ab356fee59390af8560be0e930faebe74a6daff7c9b"],
  "id": 0
}
```

> 返回值 :

```json
[
	{
		"txid": "0xc503927ac7541c45578e975d4971f2e5cad951f33a591bb59e9303548bf65fe3",
		"vout": {
			"Address": "AMpupnF6QweQXLfCtF4dR45FDdKbTXkLsr",
			"Asset": "0xc56f33fc6ecfcd0c225c4ab356fee59390af8560be0e930faebe74a6daff7c9b",
			"N": 0,
			"Value": "1"
		},
		"createTime": "2017-11-24T06:38:04Z",
		"spentTime": "",
		"block": 811513,
		"spentBlock": -1
	}
]
```

### 请求参数


Parameter | Type | Description
--------- | ------- | -----------
address| string | neo 地址
asset| string | 资产id


### 返回值


Parameter | Type | Description
--------- | ------- | -----------
utxo | object[] | 未使用的utxo列表


## 获取Claim utxo列表

> 请求 :

```json 
{
  "jsonrpc": "2.0",
  "method": "balance",
  "params": ["AMpupnF6QweQXLfCtF4dR45FDdKbTXkLsr"],
  "id": 0
}
```

> 返回值 :

```json
{
	"Unavailable": "0.00000032",
	"Available": "0.00000736",
	"Claims": [
		{
			"txid": "0xc503927ac7541c45578e975d4971f2e5cad951f33a591bb59e9303548bf65fe3",
			"vout": {
				"Address": "AMpupnF6QweQXLfCtF4dR45FDdKbTXkLsr",
				"Asset": "0xc56f33fc6ecfcd0c225c4ab356fee59390af8560be0e930faebe74a6daff7c9b",
				"N": 0,
				"Value": "1"
			},
			"createTime": "2017-11-24T06:38:04Z",
			"spentTime": "2017-11-24T07:18:12Z",
			"block": 811513,
			"spentBlock": 811605
		}
	]
}
```

### 请求参数


Parameter | Type | Description
--------- | ------- | -----------
address| string | neo 地址


### 返回值


Parameter | Type | Description
--------- | ------- | -----------
Claims | object[] | 可提现Gas的UTXO列表
Available | string | 可提现Gas
Unavailable | string | 锁定Gas