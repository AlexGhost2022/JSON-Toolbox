# JSON-Toolbox


```markdown
<div align="center">

# 🧰 DevTools Toolbox

### Четыре мощных веб-инструмента в одном месте

**JSON Formatter · Key Formatter · Base64 · URL Encoder**

![GitHub License](https://img.shields.io/badge/license-MIT-blue.svg)
![GitHub Stars](https://img.shields.io/github/stars/YOUR_USERNAME/devtools-toolbox?style=social)
![Made with](https://img.shields.io/badge/made%20with-%E2%9D%A4%EF%B8%8F-red)
![Privacy](https://img.shields.io/badge/privacy-100%25%20local-green)

[🚀 Live Demo](#) • [✨ Features](#-features) • [🔒 Security](#-security) • [📖 Docs](#-quick-start)

</div>

---

## 🎯 Overview

> **DevTools Toolbox** — это единый веб-тулбокс с четырьмя профессиональными инструментами для разработчиков. Всё работает **локально в браузере**: ваши данные никуда не отправляются.

```
┌─────────────────────────────────────────────────────────┐
│  🔀  [ JSON Formatter ]  [ Key Formatter ]  [ Base64 ]  │
│                                                          │
│   ┌─────────────────────────────────────────────────┐   │
│   │                                                 │   │
│   │          ← Единый стиль, мгновенное             │   │
│   │          ← переключение, независимые данные    │   │
│   │                                                 │   │
│   └─────────────────────────────────────────────────┘   │
└─────────────────────────────────────────────────────────┘
```

---

## ✨ Features

### 🔹 Tab 1: JSON Formatter
*Полный набор функций из Pro-версии*

| Функция | Описание |
|---------|----------|
| 🎨 **Format / Minify / Clean** | Мгновенное форматирование, минификация и очистка JSON |
| 🌳 **Tree View** | Раскрытие/сворачивание узлов, кнопки *«Развернуть всё»* / *«Свернуть всё»* |
| 📝 **Raw JSON** | Просмотр исходного формата |
| 💾 **Download / Copy** | Скачивание с пользовательским именем файла и копирование в буфер |

<details>
<summary><b>Пример работы:</b></summary>

```json
// ❌ До
{"name":"Санёк","age":25,"tags":["dev","tools"]}

// ✅ После Format
{
  "name": "Санёк",
  "age": 25,
  "tags": ["dev", "tools"]
}
```
</details>

---

### 🔹 Tab 2: Key Formatter 🚀
*Улучшенная версия с рекурсивным обходом*

> 🔑 **Главная фишка:** находит `data.key` на **любом** уровне вложенности — больше не привязан к `response.fields[]`

**Где находит ключи:**
- ✅ `response.fields[].data.key`
- ✅ `response.nested.deep.fields[].data.key`
- ✅ `any.path.you.want[].data.key`

**Пример трансформации:**

```diff
- data__company_id
+ data.company.id

- data__changes__is_flagged__old
+ data.changes.is.flagged.old

- data__changed_by__member_id
+ data.changed.by.member.id
```

**Дополнительно:**
- 📋 Лог всех изменений с ID полей
- 🧪 Кнопка **«Пример»** с готовым тестовым JSON

---

### 🔹 Tab 3: Base64 Encode / Decode
*Полная поддержка UTF-8*

| Режим | Действие |
|-------|----------|
| 🔒 **Encode** | Кодирует любой текст в Base64 |
| 🔓 **Decode** | Декодирует Base64 обратно |
| 🔄 **Swap** | Меняет местами input/output **+ автоматически переключает режим** |

**Работает с:**
- ✅ Кириллицей
- ✅ Эмодзи 🚀🔥💎
- ✅ Любыми спецсимволами

> 💡 Использует `TextEncoder` / `TextDecoder` — самый надёжный способ обработки UTF-8 в браузере

```js
// Пример
"Привет, Санёк! 👋" → "0J/RgNC40LLQtdGCINCh0JvQldC10YDQutCwISDwn5GL"
```

---

### 🔹 Tab 4: URL Encoder / Decode
*Безопасный `encodeURIComponent`*

```js
// Encode
"https://example.com/Санёк"
→ "https%3A%2F%2Fexample.com%2F%D0%A1%D0%B0%D0%BD%D1%91%D0%BA"

// Decode
"https%3A%2F%2Fexample.com%2F%D0%A1%D0%B0%D0%BD%D1%91%D0%BA"
→ "https://example.com/Санёк"
```

**Возможности:**
- 🔄 **Swap** с автопереключением режима
- 💾 Скачивание результата в `.txt`
- 🛡️ Самый безопасный вариант для URL-параметров

---

## 🎨 Unified Interface

```
┌──────────────────────────────────────────┐
│         🎨  Единый дизайн                │
├──────────────────────────────────────────┤
│  ⚡ Мгновенное переключение табов       │
│  💾 Данные каждого инструмента           │
│     сохраняются независимо               │
│  🎯 Единая цветовая схема и стиль        │
└──────────────────────────────────────────┘
```

---

## 🚀 Quick Start

```bash
# Клонируй репозиторий
git clone https://github.com/AlexGhost2022/devtools-toolbox.git
cd devtools-toolbox

# Установи зависимости
npm install

# Запусти dev-сервер
npm run dev
```

Или просто открой `index.html` в браузере — **никакой сборки не требуется** ✨

---

## 🔒 Security & Privacy

<div align="center">

```
╔═══════════════════════════════════════════╗
║  🛡️ 100% CLIENT-SIDE                      ║
║  Никаких запросов на сервер               ║
║  Никакой телеметрии                       ║
║  Никаких cookies                          ║
║                                           ║
║  Все данные остаются ТОЛЬКО у вас        ║
╚═══════════════════════════════════════════╝
```

</div>

| ✅ Работает | ❌ Не делает |
|------------|-------------|
| Локально в браузере | Не отправляет данные на сервер |
| Без регистрации | Не требует API-ключей |
| Работает offline | Не трекает пользователя |

---

## 📚 Tech Stack

![HTML5](https://img.shields.io/badge/HTML5-E34F26?logo=html5&logoColor=white)
![CSS3](https://img.shields.io/badge/CSS3-1572B6?logo=css3&logoColor=white)
![JavaScript](https://img.shields.io/badge/JavaScript-F7DF1E?logo=javascript&logoColor=black)
![Vite](https://img.shields.io/badge/Vite-646CFF?logo=vite&logoColor=white)

---

## 🤝 Contributing

Приветствуются любые вклады!

1. 🍴 Fork репозитория
2. 🔨 Создай ветку: `git checkout -b feature/amazing-feature`
3. 💾 Commit: `git commit -m 'Add amazing feature'`
4. 📤 Push: `git push origin feature/amazing-feature`
5. 🔁 Открой Pull Request

---

## 📄 License

Проект лицензирован под [MIT License](LICENSE) — используй свободно 🎉

---

<div align="center">



**Made with ❤️ by Санёк**

[⬆ Back to top](#-devtools-toolbox)

</div>
```

---



```markdown
## 🗺️ Roadmap

- [x] JSON Formatter
- [x] Key Formatter
- [x] Base64 Encode/Decode
- [x] URL Encoder/Decode
- [ ] 🌙 Dark Mode
- [ ] ⌨️ Keyboard Shortcuts
- [ ] 🔧 JWT Decoder
- [ ] 📦 History of transformations
```

Готово! Теперь у тебя профессиональный README, который не стыдно показать 🚀
