# API для сокращения ссылок

Это простой API для сокращения ссылок. Он принимает URL в качестве аргумента (`url`) и возвращает сокращённую версию ссылки.

## Возможности

- Принимает ссылку в формате URL на вход.
- Возвращает уникальную сокращённую версию ссылки.
- Легко запускается с помощью Docker Compose.

## Требования

- Установленный Docker.
- Следующие порты должны быть свободны:
  - **8080**: Для API.
  - **5432**: Для базы данных.

## Эндпоинт API

API размещается на `localhost:8080` и принимает запросы на корневом пути `/`.

### Пример запроса

- **Метод**: `POST`
- **URL**: `http://localhost:8080/`
- **Заголовки**:
  ```json
  {
    "Content-Type": "application/json"
  }

Тело запроса:

  ```json
  {
    "url": "https://example.com"
  }
  ```
Пример ответа
Статус: 200 OK
Тело ответа:
  
  ```
    "http://localhost:8080/?url=https://example.com"
  ```
**Как запустить**:
Клонируйте репозиторий:

  ```bash
  git clone <ссылка на репозиторий>
  cd <имя репозитория>
  ```
Убедитесь, что docker-compose установлен и настроен.

Запустите приложение с помощью Docker Compose:

  ```bash
  docker-compose -f docker-compose.yaml up
  ```
API будет доступен по адресу http://localhost:8080.

Детали окружения
Для работы приложения требуется:

Веб-сервер, работающий на порту 8080.
База данных (PostgreSQL), доступная на порту 5432.
Перед запуском убедитесь, что указанные порты свободны.
