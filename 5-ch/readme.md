# Инструкция по запуску двух реплик ClickHouse: локальная и Docker с ZooKeeper

## Описание

Мы настроим **две реплики ClickHouse**:
1. **Локальная реплика ClickHouse** на хостовой машине.
2. **Реплика ClickHouse в Docker**, работающая с **ZooKeeper** в контейнере.

### Шаги:

## 1. Запуск ZooKeeper и Clickhouse реплики в Docker

### Создание `docker-compose.yml`

### Создание конфигов в папке `configs`:

```
/configs/config.xml
/configs/users.xml
```
### Запуск:
```
docker-compose up -d
```

## 2. Настройка локальной реплики ClickHouse

### Локальная установка clickhouse
```
brew install clickhouse
```

### Создание конфигов в папке `config-locals`:

```
/config-locals/config.xml
/config-locals/users.xml
```

### Подключение к локальной реплики clickhouse
```
./clickhouse server -C local-config/config.xml     
clickhouse-client --host=localhost --port=9000    
```


## 3. Проверка репликации

### Создание БД
```
CREATE TABLE replicated_212 (
    id UInt64,
    value String
) ENGINE = ReplicatedMergeTree('/clickhouse/tables/replicated_212', '{replica}')
ORDER BY id;
```

### Добавление значений
```
INSERT INTO replicated_212 VALUES (1, 'value1'), (2, 'value2');
```

### Проверка значений
```
SELECT * FROM replicated_212;
```

### Данные из ЗК
```
[zk: 127.0.0.1:9181(CONNECTED) 11] ls /clickhouse/tables/test_table/replicated_212
[replica_docker, replica_local]
```