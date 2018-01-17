---
weight: 12
title: API Reference
---

# NEO钱包

## 创建新钱包

> 创建新的NEO钱包:

```java
package com.inwecrypto.test

public class App {
    public static void main(String args[]) {
        neomobile.Wallet neowallet = neomobile.new_();
    }
}
```
```objc
    
```

#### 请求参数


Parameter | Default | Description
--------- | ------- | -----------



## 通过WIF字符串创建钱包

> 通过WIF字符串创建钱包:

```java
package com.inwecrypto.test

public class App {
    public static void main(String args[]) {
        neomobile.Wallet neowallet = neomobile.fromWIF("xxxxxx");
    }
}
```

### 请求参数


Parameter | Type | Description
--------- | ---- | -----------
wif | string | WIF字符串


## 通过读取web3 keystore字符串创建钱包

> 读取keystore:

```java
package com.inwecrypto.test

public class App {
    public static void main(String args[]) {
        neomobile.Wallet neowallet = neomobile.fromKeyStore("xxxxxx","xxxxx");
    }
}
```
> keystore的格式类似下面的json代码：

```json
{
    "version": 3,
    "id": "1b3cb7fc-306f-4ec3-b753-831cb9e18984",
    "address": "00e773ad3fa1481bc4222277f324d57f35f06b60",
    "crypto": {
        "ciphertext": "423abc4fea2f1f58543b456f9a67f60eb7be076a79471d284d2777c1ce5ee2cd",
        "cipherparams": {
            "iv": "a737c70e4541a9eb053a49e9103d7ccc"
        },
        "cipher": "aes-128-ctr",
        "kdf": "pbkdf2",
        "kdfparams": {
            "dklen": 32,
            "salt": "0de73ba9540afa6424f05d159575a665da3aefa751cd7c56fc5dd87aeac4ea6b",
            "c": 65536,
            "prf": "hmac-sha256"
        },
        "mac": "53b2cf7744e91d2718f9aa586dac8c5fb647b3b912b5c4ba3c1eafd7a99346e3"
    }
}
```

### 请求参数


Parameter | Type | Description
--------- | ---- | -----------
keystore | string | keystore json 字符串
password | string | keystore 秘钥

## 通过助记词创建钱包

> 读取助记词:

```java
package com.inwecrypto.test

public class App {
    public static void main(String args[]) {
        neomobile.Wallet neowallet = neomobile.fromMnemonic("xxxxxx","zh_CN");
    }
}
```

### 请求参数


Parameter | Type | Description
--------- | ---- | -----------
mnemonic | string | 空格分割的助记词字符串
lang | string | 助记词语言，当前支持 zh_CN ， en_US


## 转账

> 创建钱包并转账:

```java
package com.inwecrypto.test

public class App {
    public static void main(String args[]) {
        neomobile.Wallet neowallet = neomobile.fromMnemonic("xxxxxx");

        neowallet.createAssertTx("xxxx","from","to","",1,"xxxxxx")
    }
}
```

> unspent 参数示例：

```json
[{
	"txid": "0x07537a82ef57610931cffe3e31ca5f38dbadc9daa2c2c881a49d9a984539ed06",
	"vout": {
		"Address": "AMpupnF6QweQXLfCtF4dR45FDdKbTXkLsr",
		"Asset": "0xc56f33fc6ecfcd0c225c4ab356fee59390af8560be0e930faebe74a6daff7c9b",
		"N": 0,
		"Value": "1"
	},
	"createTime": "2017-11-23T12:58:31Z",
	"spentTime": "",
	"block": 809098,
	"spentBlock": -1
},{
        "txid": "0x07537a82ef57610931cffe3e31ca5f38dbadc9daa2c2c881a49d9a984539ed06",
	"vout": {
		"Address": "AMpupnF6QweQXLfCtF4dR45FDdKbTXkLsr",
		"Asset": "0xc56f33fc6ecfcd0c225c4ab356fee59390af8560be0e930faebe74a6daff7c9b",
		"N": 0,
		"Value": "1"
	},
	"createTime": "2017-11-23T12:58:31Z",
	"spentTime": "",
	"block": 809098,
	"spentBlock": -1
}]
```

### 请求参数


Parameter | Type | Description
--------- | ---- | -----------
assert | string | 资产类型字符串
from | string | 转账源地址
to | string | 转账目标地址
amount | string | 转账金额
unspent | string | 转账源地址所持有的未花费的utxo列表，json格式


### 返回值


Parameter | Type | Description
--------- | ---- | -----------
tx | object | 包含 txid 以及 raw tx string


## 获取Gas

> Claim Gas:

```java
package com.inwecrypto.test

public class App {
    public static void main(String args[]) {
        neomobile.Wallet neowallet = neomobile.fromMnemonic("xxxxxx");

        neowallet.createClaimTx(1,"address","xxxxxx")
    }
}
```

> unspent 参数示例：

```json
[{
    "txid": "0x07537a82ef57610931cffe3e31ca5f38dbadc9daa2c2c881a49d9a984539ed06",
    "vout": {
        "Address": "AMpupnF6QweQXLfCtF4dR45FDdKbTXkLsr",
        "Asset": "0xc56f33fc6ecfcd0c225c4ab356fee59390af8560be0e930faebe74a6daff7c9b",
        "N": 0,
        "Value": "1"
    },
    "createTime": "2017-11-23T12:58:31Z",
    "spentTime": "2017-11-24T06:38:04Z",
    "block": 809098,
    "spentBlock": 811513
}, {
    "txid": "0x07537a82ef57610931cffe3e31ca5f38dbadc9daa2c2c881a49d9a984539ed06",
    "vout": {
        "Address": "AMpupnF6QweQXLfCtF4dR45FDdKbTXkLsr",
        "Asset": "0xc56f33fc6ecfcd0c225c4ab356fee59390af8560be0e930faebe74a6daff7c9b",
        "N": 0,
        "Value": "1"
    },
    "createTime": "2017-11-23T12:58:31Z",
    "spentTime": "2017-11-24T06:38:04Z",
    "block": 809098,
    "spentBlock": 811513
}]
```

### 请求参数


Parameter | Type | Description
--------- | ---- | -----------
amount | string | 转账金额
address | string | 转账目标地址
unspent | string | unclaimed utxo 列表，json格式





### 返回值


Parameter | Type | Description
--------- | ---- | -----------
tx | object | 包含 txid 以及 raw tx string




## 获取地址

> get neo address:

```java
package com.inwecrypto.test

public class App {
    public static void main(String args[]) {
        neomobile.Wallet neowallet = neomobile.fromMnemonic("xxxxxx");

        neowallet.address()
    }
}
```

### 请求参数

### 返回值


Parameter | Type | Description
--------- | ---- | -----------
address | string | neo 地址


## 生成助记词

> generate mnemonic:

```java
package com.inwecrypto.test

public class App {
    public static void main(String args[]) {
        neomobile.Wallet neowallet = neomobile.fromMnemonic("xxxxxx");

        neowallet.Mnemonic()
    }
}
```

### 请求参数

Parameter | Type | Description
--------- | ---- | -----------
lang | string | 助记词语言，当前支持 zh_CN， en_US

### 返回值


Parameter | Type | Description
--------- | ---- | -----------
Mnemonic | string | neo keystore 助记词



## NEP5转账

> nep5 transfer:

```java
package com.inwecrypto.test

public class App {
    public static void main(String args[]) {
        neomobile.Wallet neowallet = neomobile.fromMnemonic("xxxxxx");

        neowallet.CreateNep5Tx("script hash","from address","to address",10.1,1000,unspent)
    }
}
```

> unspent 参数示例：

```json
[{
    "txid": "0x07537a82ef57610931cffe3e31ca5f38dbadc9daa2c2c881a49d9a984539ed06",
    "vout": {
        "Address": "AMpupnF6QweQXLfCtF4dR45FDdKbTXkLsr",
        "Asset": "0xc56f33fc6ecfcd0c225c4ab356fee59390af8560be0e930faebe74a6daff7c9b",
        "N": 0,
        "Value": "1"
    },
    "createTime": "2017-11-23T12:58:31Z",
    "spentTime": "2017-11-24T06:38:04Z",
    "block": 809098,
    "spentBlock": 811513
}, {
    "txid": "0x07537a82ef57610931cffe3e31ca5f38dbadc9daa2c2c881a49d9a984539ed06",
    "vout": {
        "Address": "AMpupnF6QweQXLfCtF4dR45FDdKbTXkLsr",
        "Asset": "0xc56f33fc6ecfcd0c225c4ab356fee59390af8560be0e930faebe74a6daff7c9b",
        "N": 0,
        "Value": "1"
    },
    "createTime": "2017-11-23T12:58:31Z",
    "spentTime": "2017-11-24T06:38:04Z",
    "block": 809098,
    "spentBlock": 811513
}]
```

### 请求参数

Parameter | Type | Description
--------- | ---- | -----------
scritpHash | string | nep5 合约地址
from | string | 转账源地址
to | string | 转账目标地址
gas| float64| 转账消耗的gas费用
amount| int64| 转账代币数目
unspent|json|当前账户可用全局资产列表


### 返回值


Parameter | Type | Description
--------- | ---- | -----------
tx | object | 包含 txid 以及 raw tx string

## NEP5 ICO

> nep5 mintToken:

```java
package com.inwecrypto.test

public class App {
    public static void main(String args[]) {
        neomobile.Wallet neowallet = neomobile.fromMnemonic("xxxxxx");

        neowallet.MintToken("script hash",10.1,1000,unspent)
    }
}
```

> unspent 参数示例：

```json
[{
    "txid": "0x07537a82ef57610931cffe3e31ca5f38dbadc9daa2c2c881a49d9a984539ed06",
    "vout": {
        "Address": "AMpupnF6QweQXLfCtF4dR45FDdKbTXkLsr",
        "Asset": "0xc56f33fc6ecfcd0c225c4ab356fee59390af8560be0e930faebe74a6daff7c9b",
        "N": 0,
        "Value": "1"
    },
    "createTime": "2017-11-23T12:58:31Z",
    "spentTime": "2017-11-24T06:38:04Z",
    "block": 809098,
    "spentBlock": 811513
}, {
    "txid": "0x07537a82ef57610931cffe3e31ca5f38dbadc9daa2c2c881a49d9a984539ed06",
    "vout": {
        "Address": "AMpupnF6QweQXLfCtF4dR45FDdKbTXkLsr",
        "Asset": "0xc56f33fc6ecfcd0c225c4ab356fee59390af8560be0e930faebe74a6daff7c9b",
        "N": 0,
        "Value": "1"
    },
    "createTime": "2017-11-23T12:58:31Z",
    "spentTime": "2017-11-24T06:38:04Z",
    "block": 809098,
    "spentBlock": 811513
}]
```

### 请求参数

Parameter | Type | Description
--------- | ---- | -----------
scritpHash | string | nep5 合约地址
gas| float64| 转账消耗的gas费用
amount| int64| 转账代币数目
unspent|json|当前账户可用全局资产列表


### 返回值


Parameter | Type | Description
--------- | ---- | -----------
tx | object | 包含 txid 以及 raw tx string


## 获取钱包公钥

```java
package com.inwecrypto.test

public class App {
    public static void main(String args[]) {
        neomobile.Wallet neowallet = neomobile.fromMnemonic("xxxxxx");

        neowallet.PubKey()
    }
}
```

### 返回值


Parameter | Type | Description
--------- | ---- | -----------
pubkey | string | 钱包公钥

## NEO地址转换为Hash160格式

> NEO address to Hash160:

```java
package com.inwecrypto.test

public class App {
    public static void main(String args[]) {
        neomobile.DecodeAddress("ATLoURz25z4PpsrzZmnowRT3dya44LGEpS")
    }
}
```


### 请求参数

Parameter | Type | Description
--------- | ---- | -----------
address | string | NEO 地址


### 返回值


Parameter | Type | Description
--------- | ---- | -----------
address | hash160 string | hash160 address string


## Hash160格式地址转换为NEO地址

> NEO address to Hash160:

```java
package com.inwecrypto.test

public class App {
    public static void main(String args[]) {
        neomobile.EncodeAddress("bfc469dd56932409677278f6b7422f3e1f34481d")
    }
}
```


### 请求参数

Parameter | Type | Description
--------- | ---- | -----------
address | hash160 string | hash160 address string




### 返回值


Parameter | Type | Description
--------- | ---- | -----------
address | string | NEO 地址