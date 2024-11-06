#!/bin/bash

echo "топ 10 - ip адресов с которых идут запросы:" > lab.txt
grep -Eo '^\d{1,3}(\.\d{1,3}){3}' access.log | sort | uniq -c | sort -nr | head -10 >> lab.txt
echo "" >> lab.txt

echo "подсчет количества методов: POST, GET, PUT, DELETE, PATCH:" >> lab.txt
grep -Eo '"(GET|POST|PUT|DELETE|PATCH)' access.log | sort | uniq -c >> lab.txt
echo "" >> lab.txt

echo "кол-во операционных систем с которых выполняются запросы.  перечислить в порядке убывания (Windows, MacOS и т.д.):" >> lab.txt
grep -Eo '\((Windows|Linux|Macintosh|Android|iPhone)[^)]*\)' access.log | sort | uniq -c | sort -nr >> lab.txt
echo "" >> lab.txt

echo "cамое популярное устройство по количеству сделанных запросов (наименование устройства):" >> lab.txt
grep -Eo '\((Linux|Android|iPhone|Windows|Macintosh|Other);' access.log | sort | uniq -c | sort -nr | head -1 >> lab.txt
echo "" >> lab.txt

echo "топ 5 популярных браузеров включая их версию:" >> lab.txt
grep -Eo '"Mozilla[^"]+' access.log | sort | uniq -c | sort -nr | head -5 >> lab.txt
echo "" >> lab.txt

echo "Готово"
