---
weight: 11
title: API Reference
---

# ETH钱包




## 创建新的ETH钱包

> 创建新的ETH钱包:


```java
package com.inwecrypto.test

public class App {
    public static void main(String args[]) {
        ethmobile.Wallet wallet = ethmobile.new_();
    }
}
```
```objc
    
```

#### 请求参数

Parameter | Default | Description
--------- | ------- | -----------


## 通过读取web3 keystore字符串创建钱包

> 读取keystore:

```java
package com.inwecrypto.test

public class App {
    public static void main(String args[]) {
        ethmobile.Wallet ethwallet = ethmobile.fromKeyStore("xxxxxx","xxxxx");
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
        ethmobile.Wallet ethwallet = ethmobile.fromMnemonic("xxxxxx","zh_CN");
        ethwallet.Transfer("","","","","")
    }
}
```

### 请求参数


Parameter | Type | Description
--------- | ---- | -----------
mnemonic | string | 空格分割的助记词字符串
lang | string | 助记词语言，当前支持 zh_CN ， en_US

## 通过私钥创建钱包

> 读取助记词:

```java
package com.inwecrypto.test

public class App {
    public static void main(String args[]) {
        ethmobile.Wallet ethwallet = ethmobile.fromPrivateKey("xxxxxx");
        ethwallet.Transfer("","","","","")
    }
}
```

### 请求参数


Parameter | Type | Description
--------- | ---- | -----------
privateKey | string | hex形式私钥


## 全局资产转账（ETH)

> 转账:

```java
package com.inwecrypto.test

public class App {
    public static void main(String args[]) {
        ethmobile.Wallet ethwallet = ethmobile.fromMnemonic("xxxxxx","zh_CN");
        ethwallet.Transfer("","","","","")
    }
}
```


### 请求参数


Parameter | Type | Description
--------- | ---- | -----------
nonce | string | 服务器获取的nonce
to | string | 转入地址
amount | string | 转入数量
gasPrice | string | 燃料费价格
gasLimits | string | 燃料最高限额



## ERC20资产转账

> 转账:

```java
package com.inwecrypto.test

public class App {
    public static void main(String args[]) {
        ethmobile.Wallet ethwallet = ethmobile.fromMnemonic("xxxxxx","zh_CN");
        ethwallet.TransferERC20("","","","","","")
    }
}
```


### 请求参数


Parameter | Type | Description
--------- | ---- | -----------
contract | string | 合约地址
nonce | string | 服务器获取的nonce
to | string | 转入地址
amount | string | 转入数量
gasPrice | string | 燃料费价格
gasLimits | string | 燃料最高限额

## ERC20代币授权

> 示例:

```java
package com.inwecrypto.test

public class App {
    public static void main(String args[]) {
        ethmobile.Wallet ethwallet = ethmobile.fromMnemonic("xxxxxx","zh_CN");
        ethwallet.Approve("","","","","","")
    }
}
```


### 请求参数


Parameter | Type | Description
--------- | ---- | -----------
contract | string | 合约地址
nonce | string | 服务器获取的nonce
to | string | 授权地址
value | string | 授权额度
gasPrice | string | 燃料费价格
gasLimits | string | 燃料最高限额

## ERC20代币第三方转账，需要转出地址先授权

> 示例:

```java
package com.inwecrypto.test

public class App {
    public static void main(String args[]) {
        ethmobile.Wallet ethwallet = ethmobile.fromMnemonic("xxxxxx","zh_CN");
        ethwallet.TransferFrom("","","","","","","")
    }
}
```


### 请求参数


Parameter | Type | Description
--------- | ---- | -----------
contract | string | 合约地址
nonce | string | 服务器获取的nonce
from | string | 转出地址
to | string | 转入地址
value | string | 数量
gasPrice | string | 燃料费价格
gasLimits | string | 燃料最高限额

## NFT代币DecentraLand Land转账接口

> 示例:

```java
package com.inwecrypto.test

public class App {
    public static void main(String args[]) {
        ethmobile.Wallet ethwallet = ethmobile.fromMnemonic("xxxxxx","zh_CN");
        ethwallet.TransferLand("","","","","","","")
    }
}
```


### 请求参数


Parameter | Type | Description
--------- | ---- | -----------
contract | string | 合约地址
nonce | string | 服务器获取的nonce
to | string | 转入地址
x | string | 土地X坐标
y | string | 土地Y坐标
gasPrice | string | 燃料费价格
gasLimits | string | 燃料最高限额

## NFT代币红包 新红包接口，需要转出地址先授权

> 示例:

```java
package com.inwecrypto.test

public class App {
    public static void main(String args[]) {
        ethmobile.Wallet ethwallet = ethmobile.fromMnemonic("xxxxxx","zh_CN");
        ethwallet.NewRedPacket("","","","","","","","","","","")
    }
}
```

### 请求参数

Parameter | Type | Description
--------- | ---- | -----------
redcontract | string | NFT红包合约地址
nonce | string | 服务器获取的nonce
erc20contract  | string | 要发红包的ERC20代币合约地址
tokenId  | string | 红包的ID,需要事先生成的不重复的ID
from | string | 转出地址
amount | string | 发红包收取的手续费
value | string | 红包中包含的ERC20代币总数
count | string | 红包个数
command | string | 领取红包的口令（目前都设置为0）
gasPrice | string | 燃料费价格
gasLimits | string | 燃料最高限额

## 获取ERC20代币的Decimals

> 示例:

```java
package com.inwecrypto.test

public class App {
    public static void main(String args[]) {
        ethmobile.EthCall call = ethmobile.NewEthCall();
        call.Decimals("");
    }
}
```

### 请求参数

Parameter | Type | Description
--------- | ---- | -----------
contract  | string | ERC20代币合约地址


## 获取ERC20代币的TotalSupply

> 示例:

```java
package com.inwecrypto.test

public class App {
    public static void main(String args[]) {
        ethmobile.EthCall call = ethmobile.NewEthCall();
        call.TotalSupply("");
    }
}
```

### 请求参数

Parameter | Type | Description
--------- | ---- | -----------
contract  | string | ERC20代币合约地址

## 获取ERC20代币的Name

> 示例:

```java
package com.inwecrypto.test

public class App {
    public static void main(String args[]) {
        ethmobile.EthCall call = ethmobile.NewEthCall();
        call.Name("");
    }
}
```

### 请求参数

Parameter | Type | Description
--------- | ---- | -----------
contract  | string | ERC20代币合约地址

## 获取ERC20代币的余额

> 示例:

```java
package com.inwecrypto.test

public class App {
    public static void main(String args[]) {
        ethmobile.EthCall call = ethmobile.NewEthCall();
        call.BalanceOf("","");
    }
}
```

### 请求参数

Parameter | Type | Description
--------- | ---- | -----------
contract  | string | ERC20代币合约地址
address  | string | 地址

## NFT代币DecentraLand 解码tokenID

> 示例:

```java
package com.inwecrypto.test

public class App {
    public static void main(String args[]) {
        ethmobile.EthCall call = ethmobile.NewEthCall();
        call.LandDecodeTokenId("","");
    }
}
```

### 请求参数

Parameter | Type | Description
--------- | ---- | -----------
contract  | string | 合约地址
value  | string | tokenID

## NFT代币DecentraLand 编码tokenID

> 示例:

```java
package com.inwecrypto.test

public class App {
    public static void main(String args[]) {
        ethmobile.EthCall call = ethmobile.NewEthCall();
        call.LandEncodeTokenId("","","");
    }
}
```

### 请求参数

Parameter | Type | Description
--------- | ---- | -----------
contract  | string | 合约地址
x  | string | x坐标值
y  | string | y坐标值

## NFT代币DecentraLand LandData

> 示例:

```java
package com.inwecrypto.test

public class App {
    public static void main(String args[]) {
        ethmobile.EthCall call = ethmobile.NewEthCall();
        call.LandData("","","");
    }
}
```

### 请求参数

Parameter | Type | Description
--------- | ---- | -----------
contract  | string | 合约地址
x  | string | x坐标值
y  | string | y坐标值

## NFT代币DecentraLand 查询地址所拥有的土地

> 示例:

```java
package com.inwecrypto.test

public class App {
    public static void main(String args[]) {
        ethmobile.EthCall call = ethmobile.NewEthCall();
        call.LandOf("","");
    }
}
```

### 请求参数

Parameter | Type | Description
--------- | ---- | -----------
contract  | string | 合约地址
address  | string | 地址

## NFT代币DecentraLand 查询土地的拥有者地址

> 示例:

```java
package com.inwecrypto.test

public class App {
    public static void main(String args[]) {
        ethmobile.EthCall call = ethmobile.NewEthCall();
        call.OwnerOfLand("","","");
    }
}
```

### 请求参数

Parameter | Type | Description
--------- | ---- | -----------
contract  | string | 合约地址
x  | string | x坐标值
y  | string | y坐标值

## NFT代币DecentraLand的描述信息

> 示例:

```java
package com.inwecrypto.test

public class App {
    public static void main(String args[]) {
        ethmobile.EthCall call = ethmobile.NewEthCall();
        call.Description("");
    }
}
```

### 请求参数

Parameter | Type | Description
--------- | ---- | -----------
contract  | string | 合约地址

## NFT代币 地址所持有的所有资产

> 示例:

```java
package com.inwecrypto.test

public class App {
    public static void main(String args[]) {
        ethmobile.EthCall call = ethmobile.NewEthCall();
        call.TokensOf("","");
    }
}
```

### 请求参数

Parameter | Type | Description
--------- | ---- | -----------
contract  | string | 合约地址
address  | string | 地址

## NFT代币 代币是否存在

> 示例:

```java
package com.inwecrypto.test

public class App {
    public static void main(String args[]) {
        ethmobile.EthCall call = ethmobile.NewEthCall();
        call.Exists("","");
    }
}
```

### 请求参数

Parameter | Type | Description
--------- | ---- | -----------
contract  | string | 合约地址
value  | string | 代币

## NFT代币 代币元数据

> 示例:

```java
package com.inwecrypto.test

public class App {
    public static void main(String args[]) {
        ethmobile.EthCall call = ethmobile.NewEthCall();
        call.TokenMetadata("","");
    }
}
```

### 请求参数

Parameter | Type | Description
--------- | ---- | -----------
contract  | string | 合约地址
value  | string | 代币

## NFT代币 根据持有者的索引查询代币

> 示例:

```java
package com.inwecrypto.test

public class App {
    public static void main(String args[]) {
        ethmobile.EthCall call = ethmobile.NewEthCall();
        call.TokenOfOwnerByIndex("","","");
    }
}
```

### 请求参数

Parameter | Type | Description
--------- | ---- | -----------
contract  | string | 合约地址
address  | string | 地址
value  | string | 索引编码

## NFT代币 查询代币持有者地址

> 示例:

```java
package com.inwecrypto.test

public class App {
    public static void main(String args[]) {
        ethmobile.EthCall call = ethmobile.NewEthCall();
        call.OwnerOf("","");
    }
}
```

### 请求参数

Parameter | Type | Description
--------- | ---- | -----------
contract  | string | 合约地址
value  | string | 代币

## NFT代币 发放每个红包的手续费

> 示例:

```java
package com.inwecrypto.test

public class App {
    public static void main(String args[]) {
        ethmobile.EthCall call = ethmobile.NewEthCall();
        call.RedPacketTaxCost("");
    }
}
```

### 请求参数

Parameter | Type | Description
--------- | ---- | -----------
contract  | string | 合约地址

## NFT代币 最大可以发放多少个红包

> 示例:

```java
package com.inwecrypto.test

public class App {
    public static void main(String args[]) {
        ethmobile.EthCall call = ethmobile.NewEthCall();
        call.RedPacketMaxCount("");
    }
}
```

### 请求参数

Parameter | Type | Description
--------- | ---- | -----------
contract  | string | 合约地址

## NFT代币 红包的详情

> 示例:

```java
package com.inwecrypto.test

public class App {
    public static void main(String args[]) {
        ethmobile.EthCall call = ethmobile.NewEthCall();
        call.RedPacketOpenDetail("","");
    }
}
```

### 请求参数

Parameter | Type | Description
--------- | ---- | -----------
contract  | string | 合约地址
value  | string | 红包ID