# 01 Filename Ext

## Task

Требуется написать программу, которая по введенному полному пути выводит имя файла и расширение

## Пример работы программы

```
./filename-ext /home/robotomize/1.txt
filename: 1
extension: txt
```

## Сборка

`go build -o filename-ext`

## Сценарии

```
./filename-ext /home/robotomize/1.txt

filename: 1
extension: txt
```

```
./filename-ext /home/robotomize/1.txt.txt

filename: 1.txt
extension: txt
```

```
./filename-ext /home/robotomize/1

filename: 1
extension:
```