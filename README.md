[![test](https://github.com/kenzo0107/shukujitsu/actions/workflows/test.yml/badge.svg)](https://github.com/kenzo0107/shukujitsu/actions/workflows/test.yml) [![lint](https://github.com/kenzo0107/shukujitsu/actions/workflows/lint.yml/badge.svg)](https://github.com/kenzo0107/shukujitsu/actions/workflows/lint.yml)

# shukujitsu

内閣府ホームページにて提供される日本の祝日データを元に、祝日判定・祝日名を取得する Go ライブラリです。


## データセット

[内閣府ホームページ](https://www8.cao.go.jp/chosei/shukujitsu/gaiyou.html) で提供される [shukujitsu.csv](https://www8.cao.go.jp/chosei/shukujitsu/syukujitsu.csv) をベースとしております。

GitHub Actions 内で定期的に取得します。

## サンプルコード

```go
package main

import (
	"fmt"
	"time"

	"github.com/kenzo0107/shukujitsu"
)

func main() {
	t, _ := time.Parse("2006-01-02", "2021-07-22")
	if shukujitsu.IsHoliday(t) {
		fmt.Println("there is a holiday")
	}
	fmt.Println(shukujitsu.HolidayName(t))
}
```
