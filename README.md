# Containers-Monitoring

 Приложение осуществляет мониторинг состояние контейнеров (alive/not alive).
 
 В составе приложения реализованы 4 сервиса согласно схеме:
 
 ![Image alt](https://github.com/lvrdn/Containers-Monitoring/blob/main/app_structure.png)

 1. **Frontend-service** - отображение таблицы с перечнем адресов, временем последнего пинга и временем последнего успешного пинга по адресу **http://localhost:80**. Таблица реализована с помощью библиотеки Ant Design.
    Настроен nginx сервер, отдающий статические файлы и осуществляющий роутинг запросов к api.
 2. **API-service** обладает двумя rest маршрутами:
    * **"GET /api/containers"** получение данных о всех отслеживаемых маршрутах. Данные передаются в формате JSON в отсортированном виде, маршрут используется frontend-service.
    * **"POST /api/containers"** обновление данных маршрутов. Данные передаются в формате JSON, маршрут используется pinger-service.
    Запускается после запуска database.
 3. **Pinger-service** осуществляет пинг отслеживаемых адресов. Пинг, сбор информации и отправление в API осуществляется асинхронно по адресам. В сервисе реализованы механизмы TIMEOUT (максимальное время ожидания отслеживаемого адреса) и FREQUENCY (частота опроса адреса).
    Значение FREQUENCY должно быть больше значения TIMEOUT, параметры задаются в app.env.
    Запускается после запуска API-service.
 5. **Database** представлен базой данных PostgreSQL, в таблице которой осуществляется хранение информации об отслеживаемых адресах.

 ## Конфигурирование осуществляется в файле **app.env** , основные параметры:
 * **PING_ADDRESSES** - перечень отслеживаемых адресов, значения вводятся через ",". Для проверки работоспособности в файле уже записаны 4 адреса: 2 случайных адреса и 2 адреса самого приложения (база данных и api).
 * **PING_TIMEOUT** - максимальное время ожидания во время выполнения ping, задается в формате "1s", где число - значение в секундах, s - секунды.
 * **PING_FREQUENCY** - частота ping, задается в формате "1s", где число - значение в секундах, s - секунды.
 * остальные параметры в файле app.env не требуется изменять, с помощью них осуществляется связь сервисов между друг другом.

 ## Запуск

 Для того чтобы собрать и запустить приложение необходимо убедиться, что на рабочей машине запущен Docker Engine.
 Далее в консоли в корне проекта необходимо выполнить команду **"make compose_up"**. Осуществится сборка приложения и его запуск в docker-контейнерах.
 Для запуска приложения (в случае когда контейнеры уже созданы) можно воспользоваться командой **make on** в консоли в корне проекта или интерфейсом Docker Desktop.
 Для остановки приложения можно воспользоваться командой **make off** в консоли в корне проекта или интерфейсом Docker Desktop.

 ## Пример вывода таблицы отслеживаемых адресов **http://localhost:80** :
