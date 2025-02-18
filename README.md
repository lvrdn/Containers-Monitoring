# Containers-Monitoring

 Приложение осуществляет мониторинг состояния контейнеров (alive/not alive).
 
 В составе приложения реализованы 4 сервиса согласно схеме:
 
 ![Image alt](https://github.com/lvrdn/Containers-Monitoring/blob/main/app_structure.png)

 1. **Frontend-service** написан на JS с использоваем React. Осуществляет отображение таблицы с перечнем адресов, временем последнего пинга и временем последнего успешного пинга. Таблица реализована с помощью библиотеки Ant Design.
    В случае долгого ответа от бекенда вместо таблицы будет отображаться loading... , в случае ошибки получения данных - error. Настроен nginx сервер, отдающий статические файлы и осуществляющий роутинг запросов к api.
 2. **API-service** написан на Golang, обладает двумя rest эндпоинтами:
    * **"GET /api/containers"** получение данных о всех отслеживаемых адресах. Данные передаются в формате JSON в отсортированном виде, эндпоинт используется frontend-service.
    * **"POST /api/containers"** обновление данных маршрутов. Данные передаются в формате JSON, эндпоинт используется pinger-service.
    API-service запускается после запуска database с помощью скрипта wait-for-it.sh.
 3. **Pinger-service** написан на Golang, осуществляет пинг отслеживаемых адресов. Пинг, сбор информации и отправление в API осуществляется асинхронно по адресам. В сервисе реализованы механизмы TIMEOUT (максимальное время ожидания отслеживаемого адреса) и FREQUENCY (частота опроса адреса).
    Значение FREQUENCY должно быть больше значения TIMEOUT, параметры задаются в app.env.
    Pinger-service запускается после запуска API-service с помощью скрипта wait-for-it.sh.
 4. **Database** - в качестве базы данных используется PostgreSQL. Информации об отслеживаемых адресах хранится в таблице. Первичное заполнение таблицы происходит во время запуска API-service.

    Сборка и запуск сервисов в составе приложения осуществляется с использованием технологии Docker.

 ## Конфигурирование осуществляется в файле **app.env** , основные параметры:
 * **PING_ADDRESSES** - перечень отслеживаемых адресов, значения вводятся через ",". Для проверки работоспособности в файле уже записаны 4 адреса: 2 случайных адреса и 2 адреса самого приложения (база данных и api).
 * **PING_TIMEOUT** - максимальное время ожидания во время выполнения ping, задается в формате "1s", где число - значение в секундах, s - секунды.
 * **PING_FREQUENCY** - частота ping, задается в формате "1s", где число - значение в секундах, s - секунды.
 * остальные параметры в файле app.env не требуют изменений, с их помощью осуществляется связь сервисов между друг другом.

 ## Запуск

 Для того чтобы собрать и запустить приложение необходимо убедиться, что на рабочей машине запущен Docker Engine. Убедитесь, что 80 порт свободен.
 
 Далее в консоли в корне проекта необходимо выполнить команду **"make compose_up"**. Осуществится сборка приложения и его запуск в docker-контейнерах. Откройте страницу в браузере по адресу **http://localhost:80**.
 
 Для запуска приложения (в случае когда контейнеры уже созданы) можно воспользоваться командой **"make on"** в консоли в корне проекта или интерфейсом Docker Desktop.
 
 Для остановки приложения можно воспользоваться командой **"make off"** в консоли в корне проекта или интерфейсом Docker Desktop.

 Остальные команды Makefile использовались для одиночного тестирования сервисов.

 ## Пример вывода таблицы :

![Image alt](https://github.com/lvrdn/Containers-Monitoring/blob/main/table_example.png)

В данном примере два первых адреса с самого начала мониторинга не доступны, поэтому в последнем столбце нет информации.

Следующие два адреса доступны, поэтому время во втором и третьем столбце совпадают.

В случае если адрес был доступен какое-то время, а потом нет, то обновлятся будут только данные во втором столбце.
