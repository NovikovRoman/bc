# BC (BESTCHANGE)

> Библиотека для работы с данными версии 2.03 bestchange.ru

- [Использование](#использование)
- [Данные с bestchange.ru](#данные-с-bestchangeru)
  - [bm_bcodes.dat](#bm_bcodesdat)
  - [bm_brates.dat](#bm_bratesdat)
  - [bm_cities.dat](#bm_citiesdat)
  - [bm_cy.dat](#bm_cydat)
  - [bm_cycodes.dat](#bm_cycodesdat)
  - [bm_exch.dat](#bm_exchdat)
  - [bm_info.dat](#bm_infodat)
  - [bm_news.dat](#bm_newsdat)
  - [bm_rates.dat](#bm_ratesdat)
  - [bm_top.dat](#bm_topdat)

## Использование

```shell
go get github.com/NovikovRoman/bc
```

```go
package main

import (
    "context"
    "fmt"
    "log"

    "github.com/NovikovRoman/bc"
)

func main() {
    dirData := "./bcData"
    ctx := context.Background()

    // Загрузка данных с bestchange.ru
    err := bc.Download(ctx, dirData, nil)
    if err != nil {
        log.Fatalln("Download", err)
    }

    // Получить список обменников
    e, err := bc.NewExchanges(dirData)
    if err != nil {
        log.Fatalln("NewExchanges", err)
    }

    fmt.Printf("%+v\n", e)
}

```

## Данные с bestchange.ru

[!] Все файлы имеют кодировку `CP1251`.

### bm_bcodes.dat

> Коды валют

Каждая строка содержит данные:

- ID кода
- Символьный код
- Наименование
- Какого банка

### bm_brates.dat

> Курсы валют

Каждая строка содержит данные:

- ID валюты 1
- ID валюты 2
- Количество валюты 1 к валюте 2

### bm_cities.dat

> Список городов

Каждая строка содержит данные:

- ID города
- Название

### bm_cy.dat

> Список типов обменов

Каждая строка содержит данные:

- ID
- Символьный код
- ID валюты
- Тип обмена
- Название
- Альтернативное название
- Возможность обмена между типами (1|0). Битовая маска. Типы обмена отсортированны по ID.

### bm_cycodes.dat

> Коды типов обменов

Каждая строка содержит данные:

- ID типа
- Символьный код

### bm_exch.dat

> Список обменников

Каждая строка содержит данные:

- ID
- Название
- ???
- WebMoney Business Level
- Резерв

### bm_info.dat

> Информация о версии API.

Пример файла:

```plain
last_update=17:14:58, 4 июля
current_version=2.03
compatible_version=2.02
```

### bm_news.dat

> Список новостей

Каждая строка содержит данные:

- Заголовок
- Контент
- Дата

### bm_rates.dat

> Список курсов обменников

Каждая строка содержит данные:

- ID валюты отдаете
- ID валюты получаете
- ID обменника
- Отдаете сумму
- Получаете сумму
- Резерв обменника (получаемой валюты)
- Количество незакрытых претензий к обменнику
- Количество отзывов
- Активность
- Минимальная сумма для обмена
- Максимальная сумма для обмена
- ID города (как правило для приема/получения наличных)

### bm_top.dat

> ТОП типов обмена (за сутки)

Каждая строка содержит данные:

- ID типа обмена 1
- ID типа обмена 2
- процент использования из типа 1 в тип 2
