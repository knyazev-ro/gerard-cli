# 🧱 Gerard — Go CLI для быстрого создания чистых и модульных микросервисов

**Сохраняйте чистоту архитектуры. Ускоряйте командную разработку.**

**Gerard** — лёгкий и понятный CLI-инструмент для быстрого scaffolding API-модулей на Go по паттерну MVC.
Использует [Gorilla MUX](https://github.com/gorilla/mux) под капотом для маршрутизации.

С помощью Gerard вы можете создавать новые модули, добавлять контроллеры, сервисы, middleware и многое другое всего одной командой.

---

## 🚀 Основные возможности

* ✅ Быстро scaffold модульной MVC-архитектуры
* ✅ Автоматическая генерация контроллеров, сервисов, моделей, middleware, репозиториев, интерфейсов, конфигов
* ✅ Поддержка удобных шаблонов и кастомизации структуры проекта
* ✅ Приведение названий компонентов к единому виду (например, `BlogController`, `blog_controller`, `blog-controller` → `blog_controller.go` с PascalCase структурами и camelCase переменными)
* ✅ Поддержка ключа `--force` для принудительной перезаписи файлов
* ✅ Подсветка вывода успешных действий и ошибок через [fatih/color](https://github.com/fatih/color)
* ✅ Управление доступными командами, шаблонами и структурой проекта через `settings.yaml`

---

## 🛠️ Быстрый старт

### 1. Скачайте или склонируйте Gerard в корень вашего Go-проекта

```bash
git clone https://github.com/knyazev-ro/gerard-cli.git
```

Или просто скопируйте папку `gerard/` в ваш проект.

---

### 2. Соберите CLI

```bash
cd gerard
go build -o ../gerard.exe .
cd ..
```

Теперь в корне проекта появится `gerard.exe`.

---

### 3. Используйте команды Gerard!

```bash
gerard.exe create:module <module_name>
```

Создаёт новый модуль с базовой структурой и шаблонами.

---

## 💡 Доступные команды

```plaintext
gerard.exe create:module <module_name>                   - Создать новый модуль
gerard.exe create:controller <name> <module>           - Добавить контроллер в модуль
gerard.exe create:middleware <name> <module>           - Добавить middleware в модуль
gerard.exe create:model <name> <module>                 - Создать модель в модуле
gerard.exe create:repository <name> <module>            - Создать репозиторий в модуле
gerard.exe create:service <name> <module>               - Создать сервис в модуле
gerard.exe create:interface <name> <module>             - Создать интерфейс в модуле
gerard.exe create:config <name> <module>                - Создать конфиг в модуле
gerard.exe create:test <name> <module>                - Создать конфиг в модуле
gerard.exe remove:module <module_name>                   - Удалить модуль

gerard.exe help                                         - Показать это сообщение

Пример использования:

gerard.exe create:controller user blog --force
```

---

## ⚙️ Управление командами и структурой проекта

Вы можете гибко управлять доступными командами, шаблонами и структурой проекта в файле `gerard-cli/settings.yaml`.

```yaml
commands:
  create-init: true
  create-model: true
  create-config: true
  create-service: true
  create-interface: true
  create-controller: true
  create-middleware: true
  create-repository: true
  remove-module: true

templates:
  service: "gerard-cli/templates/service.tmpl"
  controller: "gerard-cli/templates/controller.tmpl"
  model: "gerard-cli/templates/model.tmpl"
  interface: "gerard-cli/templates/interface.tmpl"
  config: "gerard-cli/templates/config.tmpl"
  config-utils: "gerard-cli/templates/config-utils.tmpl"
  config-base: "gerard-cli/templates/config-base.tmpl"
  config-database: "gerard-cli/templates/config-database.tmpl"
  config-server: "gerard-cli/templates/config-server.tmpl"
  module: "gerard-cli/templates/module.tmpl"
  middleware: "gerard-cli/templates/middleware.tmpl"
  repository: "gerard-cli/templates/repository.tmpl"
  dockerfile: "gerard-cli/templates/dockerfile.tmpl"
  gitignore: "gerard-cli/templates/gitignore.tmpl"
  readme: "gerard-cli/templates/readme.tmpl"
  route: "gerard-cli/templates/route.tmpl"
  env-example: "gerard-cli/templates/env-example.tmpl"
  github-workflows: "gerard-cli/templates/github-workflows.tmpl"

generated-file-structure:
  docs: "docs"
  scripts: ".scripts"
  configs: "configs"
  config_utils: "configs/utils"
  tests: "tests"
  docker: "docker"
  src: "src"
  utils: "src/utils"
  enums: "src/enums"
  models: "src/models"
  routes: "src/routes"
  services: "src/services"
  interfaces: "src/interfaces"
  middlewares: "src/middlewares"
  controllers: "src/controllers"
  repositories: "src/repositories"
  github-workflows: ".github/workflows"
```

---

## 📂 Структура проекта после инициализации

```
your_project/
├── gerard.exe
├── <module_name>/
│   ├── src/
│   │   ├── controllers/
│   │   │   └── user_controller.go
│   │   ├── middlewares/
│   │   │   └── auth_middleware.go
│   │   ├── routes/
│   │   │   └── routes.go
│   │   ├── models/
│   │   ├── repositories/
│   │   ├── services/
│   │   ├── interfaces/
│   │   ├── utils/
│   │   └── enums/
│   ├── configs/
│   ├── tests/
│   ├── scripts/
│   ├── docker/
│   └── docs/
```

---

## 📌 Требования

* Go 1.18 или выше
* Поддержка Windows (используйте `gerard.exe`) или сборка для других платформ без параметра `-o`

---

## ⚡ Дополнительно

* Все имена автоматически нормализуются и приводятся к snake\_case для файлов и папок, PascalCase — для структур, camelCase — для переменных, независимо от того, как вы вводите (`BlogController`, `blog-controller`, `blog_controller` → `blog_controller.go` и `BlogController` struct)
* Ключ `--force` позволяет перезаписывать файлы при генерации
* Подсветка ошибок и успешных сообщений делается с помощью библиотеки `github.com/fatih/color`
* В будущем планируется поддержка генерации моделей, Swagger документации и автокомплита CLI

---

## 🧪 Генерация тестов вместе с командами `create:`

Для любой команды `create:` (например, `create:controller`, `create:service`, `create:model` и др.) можно указать флаг `--with-tests`, чтобы автоматически создать тестовый файл в папке `tests`.

**Правила генерации:**

* Имя тестового файла формируется на основе имени создаваемого компонента, с добавлением суффикса `_test`.
  Например:

  ```
  gerard.exe create:controller user_controller blog --with-tests
  ```

  создаст:

  ```
  tests/user_controller_test.go
  ```
* Внутри файла теста будет базовый шаблон теста, в котором название теста формируется по правилу `Test` + PascalCase название компонента:

  ```
  func TestUserController(t *testing.T) {
      // TODO: implement test
  }
  ```

**Пример:**

```bash
gerard.exe create:service payment_service shop --with-tests
```

Результат:

```
src/services/payment_service.go
tests/payment_service_test.go
```