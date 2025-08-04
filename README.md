# Effective Mobile Subscription API

## Описание
Этот проект представляет собой API для управления подписками. Он разработан на языке Go и использует PostgreSQL в качестве базы данных. API поддерживает создание, обновление, удаление и получение информации о подписках, а также вычисление общей стоимости подписок.

## Структура проекта
- **cmd/api**: Точка входа для запуска API.
- **config**: Конфигурационные файлы.
- **docs**: Swagger-документация API.
- **internal/modules**: Основная бизнес-логика, разделённая на модули.
  - **application**: DTO, мапперы и use-case'ы.
  - **domain**: Модели, репозитории и value objects.
  - **infrastructure**: Реализация репозиториев.
  - **interfaces**: REST-интерфейсы и обработчики.
- **internal/shared**: Общие компоненты, такие как middleware, роутеры и утилиты.
- **migrations**: SQL-скрипты для миграции базы данных.

## Установка и запуск

### Предварительные требования
- Go 1.24.3 или выше
- Docker и Docker Compose

### Шаги для запуска
1. Склонируйте репозиторий:
   ```bash
   git clone <URL репозитория>
   cd effective_mobile_go
   ```

2. Скопируйте файл `.env.example` в `.env` и заполните необходимые переменные окружения:
   ```bash
   cp .env.example .env
   ```

3. Соберите и запустите контейнеры Docker:
   ```bash
   make compose-up
   ```

4. Сгенерируйте Swagger-документацию (опционально):
   ```bash
   make swagger
   ```

5. Запустите приложение:
   ```bash
   make run
   ```

Приложение будет доступно по адресу `http://localhost:<APP_PORT>`.

### Внимание: убедитесь, что в файле Dockerfile указан правильный порт приложения, соответствующий переменной `APP_PORT` в файле `.env`.

## Команды Makefile
- `make build`: Сборка приложения.
- `make run`: Сборка и запуск приложения.
- `make clean`: Очистка собранных файлов.
- `make tidy`: Упорядочивание зависимостей.
- `make migration name=<имя>`: Создание новой миграции.
- `make migrate direction=<up|down>`: Применение или откат миграций.
- `make swagger`: Генерация Swagger-документации.
- `make swagger-fmt`: Форматирование Swagger-документации.
- `make swagger-clean`: Удаление сгенерированных Swagger-файлов.
- `make compose-up`: Запуск Docker Compose.
- `make compose-down`: Остановка Docker Compose.
- `make compose-logs`: Просмотр логов Docker Compose.

## Используемые технологии
- Go
- PostgreSQL
- Docker и Docker Compose
- Swagger для документации API

## Контакты

Поддержка API - x3.na.tri@gmail.com

Ссылка на
проект: [https://github.com/x3-developer/effective-mobile--subscription-api](https://github.com/x3-developer/effective-mobile--subscription-api)