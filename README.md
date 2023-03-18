# TranslatorWeb

## Установка через бат файл
---
1. Запустить `extractor.bat`, после открывшегося окна выбрать extractor.py и нажать convert .py to .exe. Затем скачать и установить [tesseract](https://github.com/UB-Mannheim/tesseract/wiki) по пути Python\Extractor\lib\Tesseract (надо самому создать папки, если нету)
2. Запустить `translator.bat`, после открывшегося окна выбрать translator.py и нажать convert .py to .exe
3. Запустить `go.bat`
4. Перейти по пути TranslatorWeb\Go
5. Запустить powershell
6. `go run .`
*Если в extractor или translator не открывается окно с конвертацией в exe, то установите отдельно и конвертируйте .py*

## Установка без бат файла
---
### Extractor
1. Перейти по пути TranslatorWeb\Python\Extractor
2. Запустить powershell
3. `python -m venv venv`
4. Затем внутреннюю внутренную среду в venv
5. `pip install -r "requirements.txt"`
6. `auto-py-to-exe`
7. Выбрать main.py и сгенерировать exe
8. Скачать [tesseract](https://github.com/UB-Mannheim/tesseract/wiki)
9. Установить в путь Python\Extractor\lib\Tesseract и обязательно выбрать Русский язык в опциях при установке

### Translator
1. Перейти по пути TranslatorWeb\Python\Translator
2. Запустить powershell
3. `python -m venv venv`
4. Затем активировать внутреннюю среду в venv
5. `pip install -r "requirements.txt"`
6. `auto-py-to-exe`
7. Выбрать main.py и сгенерировать exe

### Go (основной проект)
1. Перейти по пути TranslatorWeb\Go
2. Запустить powershell
3. `go get .`
4. `go run .`
