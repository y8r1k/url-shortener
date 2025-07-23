# URL Shortener на Go

Простой и быстрый сервис сокращения ссылок, написанный на языке Go.

## 🚀 Возможности

- Генерация коротких URL из длинных ссылок
- Перенаправление на оригинальные URL
- Хранение в sqlite базе данных

## 📦 Установка

1. Клонируй репозиторий:

```bash
git clone https://github.com/твой-ник/url-shortener.git
cd url-shortener
```

2. Сборка и запуск:

```bash
go build -o bin/shortener cmd/url-shortener/main.go
CONFIG_PATH=./config/local.yaml ./bin/shortener
```