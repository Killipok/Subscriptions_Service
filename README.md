# Subscriptions Service

REST API для управления подписками пользователей.  
Реализован на Go с использованием Gin, PostgreSQL и Docker.  

---

## Технологии

- Go 1.25  
- Gin (HTTP framework)  
- PostgreSQL 15  
- Docker / Docker Compose  
- Swagger для документации API  

---

## Возможности

- Создание, получение, обновление и удаление подписок (CRUD)  
- Получение общей стоимости всех подписок  
- Swagger UI для визуализации и тестирования API  

---

## Запуск проекта

1 - Клонируем репозиторий
2 - Запускаем докер: docker compose up --build (Остановить докер ctrl + c и удалить docker compose down)

## 1️ - Создание подписки
curl -X POST localhost:8080/subscriptions \
-H "Content-Type: application/json" \
-d '{
  "service_name":"Netflix",
  "price":700,
  "user_id":"60601fee-2bf1-4721-ae6f-7636e79a0cba",
  "start_date":"2025-07-01"
}'

## 2️ - Получение всех подписок
curl localhost:8080/subscriptions

## 3️ - Получение подписки по ID
curl localhost:8080/subscriptions/1

## 4️ - Обновление подписки
curl -X PUT localhost:8080/subscriptions/1 \
-H "Content-Type: application/json" \
-d '{
  "service_name":"Netflix Premium",
  "price":1200,
  "user_id":"60601fee-2bf1-4721-ae6f-7636e79a0cba",
  "start_date":"2025-07-01",
  "end_date":"2026-07-01"
}'

## 5️ - Удаление подписки
curl -X DELETE localhost:8080/subscriptions/1

## 6️ - Получение общей стоимости всех подписок
curl localhost:8080/subscriptions/total

## Swagger

Так же команды можно ввести через swagger 
Swagger UI доступен по адресу:

http://localhost:8080/swagger/index.html
