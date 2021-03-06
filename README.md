# Meli Challenge Mutants API

![image](https://user-images.githubusercontent.com/25770844/169629362-56753248-df40-4301-a2b6-0397d1bfc361.png)

## Resumen

Este repositorio contiene el código de la API de Meli Challenge Mutants, instrucciones para ejecutar, y una pequeña referencia de la API.

___

## Tabla de contenido

1. [Requerimientos](#requerimientos)
2. [Arquitectura](#arquitectura)
    * [Cloud](#cloud)
    * [Proyecto](#proyecto)
    * [Logs](#logs)
3. [Instalación](#instalación)
    * [Clona el proyecto](#clona-el-proyecto)
    * [Instala las dependencias](#instala-las-dependencias)
    * [Ejecuta las pruebas](#ejecuta-las-pruebas)
    * [Configura variables de entorno](#configura-variables-de-entorno)
    * [Ejecuta el proyecto](#ejecuta-el-proyecto)
4. [Endpoints](#endpoints)
5. [Referencia de la API](#referencia-de-la-api)
    * [Verifica disponibilidad del servicio](#verifica-disponibilidad-del-servicio)
    * [Verifica si un humano es mutante](#verifica-si-un-humano-es-mutante)
    * [Genera estadísticas de mutantes y humanos](#genera-estadísticas-de-mutantes-y-humanos)
6. [Autor](#autor)

___

## Requerimientos

* [Golang 1.17.x](https://go.dev/dl/)
* [Git](https://git-scm.com/downloads)
* [PostgreSQL 13](https://www.postgresql.org/download/)
* [Postman](https://www.postman.com/downloads/) (Para probar la API)
* [Docker](https://docs.docker.com/get-docker/) (Solo para Cloud Run)
* [PgAdmin](https://www.pgadmin.org/download/) (Opcional)

___

## Arquitectura

### Cloud

Meli Challenge Mutants API está desplegada en [Cloud Run](https://cloud.google.com/run/) en Google Cloud Platform. 
Utiliza Git para versionar el código y Git Actions para automatizar las pruebas y el despliegue.

![Cloud](https://user-images.githubusercontent.com/25770844/169662642-c3631ef5-8851-40c7-a387-26528ba44fdd.png)

### Proyecto

El proyecto utiliza el enfoque de [Domain-Driven Design](https://es.wikipedia.org/wiki/Dise%C3%B1o_guiado_por_el_dominio) y [Arquitectura hexagonal](https://en.wikipedia.org/wiki/Hexagonal_architecture_(software)) para la arquitectura de la API.

![Stats](https://user-images.githubusercontent.com/25770844/169662591-1d50a189-0ea1-445d-9273-e581d02cb1a7.png)

![Mutant](https://user-images.githubusercontent.com/25770844/169662640-519e984f-a5ef-4de0-be2f-059cd211a7b2.png)

### Logs

Para registrar los eventos de la API en cada `request` se utiliza [zerolog](https://github.com/rs/zerolog). Zerolog es un paquete de código abierto que permite escribir logs de forma sencilla en formato JSON. Sin embargo, provee una configuración que permite escribir logs en un formato mas *humano*.

*Ejemplo de logs en formato JSON:*

```json
{
  "caller": "/app/services/humans_service.go:38",
  "msg": "Error validating input",
  "request_id": "72efb5eb-18f1-4cdd-a08d-59a66b62b5ab", # x-request-id
  "dna": [
    "ATGCGA",
    "CAGTGC",
    "TTATGT",
    "AGAAGG",
    "CCCCTA",
    "TCACTGq"
  ],
  "time": 1653172298, # Unix timestamp
  "error": "dna must be a square matrix",
  "level": "error"
}
```

*Ejemplo de logs en formato humano:*

![Logs](https://user-images.githubusercontent.com/25770844/169671986-431a3109-6dcf-422a-8e6e-c64ca65009d9.png)

## Instalación

### Clona el proyecto

Dentro de tu workspace, ejecuta el siguiente comando:

```bash
git clone https://github.com/getmiranda/meli-challenge-api
cd meli-challenge-api
```

### Instala las dependencias

```bash
go mod tidy
```

### Ejecuta las pruebas

Habitualmente, las pruebas se ejecutan con el comando `go test`.

```bash
go test ./... -v -cover
```

Para ejecutar las pruebas y visualizar el reporte de cobertura de código, se debe ejecutar el siguiente comando:

```bash
go test ./... -coverprofile coverage.out -covermode count && go tool cover -func coverage.out
```

Segun la configuración de [Git Actions](https://github.com/getmiranda/meli-challenge-api/blob/main/.github/workflows/coverage.yml), el actual coverage es de **100%**.

```bash
Quality Gate: checking test coverage is above threshold ...
Threshold             : 80 %
Current test coverage : 100.0 %
OK
```

### Configura variables de entorno

Es necesario configurar las variables de entorno para que la API funcione correctamente.

```bash
LOG_LEVEL='info'                 # 'info', 'debug' or 'error'. Default: 'info'
LOG_ENCODE_OUTPUT='console'      # 'console' or 'json'. Default: 'console'

DB_MUTANTS_HOST='127.0.0.1'      # Database host
DB_MUTANTS_USER='myuser'         # Database user
DB_MUTANTS_PASSWORD='mypassword' # Database password
DB_MUTANTS_DBNAME='dbname'       # Database name
DB_MUTANTS_PORT='1234'           # Database port
```

### Ejecuta el proyecto

```bash
go run main.go
```

> **Nota**
> Si no tienes una base de datos en la nube, puedes ejecutar el proyecto en local con la siguiente configuración usando Docker: [PostgreSQL 13 con Docker](https://gist.github.com/getmiranda/57957134f7144429bc195a50c91d003f)

___

## Endpoints

Local
: `http://localhost:8080`

Producción
: `https://meli-challenge-api-mjpjf6v63q-uc.a.run.app`

> **Nota**
> Puedes descargar la colección de postman para probar la API aquí: [Meli Challenge Mutants API.postman_collection.json](https://gist.github.com/getmiranda/c6a84f1eb247ee06afc016d13d88fa05)

___

## Referencia de la API

### Verifica disponibilidad del servicio

#### **`GET /ping`**

#### Descripción

Verifica si el servicio está disponible.

#### Parámetros

No hay parametros.

#### Códigos de estado de respuesta HTTP

| Codigo | Descripción       |
| ------ | ----------------- |
| **200**    | Respuesta exitosa <br> **Headers** <br> `x-request-id` (string): Id unico por petición, sirve para identificar los logs asociados a la petición. |

#### Ejemplo de respuesta

**Headers**:

```bash
HTTP/1.1 200 OK
Content-Type: text/plain; charset=utf-8
X-Request-Id: 5d8f9c8a-f9b7-4c8b-b8f9-f8b8f8b8f8b8
Date: Fri, 01 Jan 2019 19:00:00 GMT
Server: Google Frontend
Content-Length: 15
```

**Body**:

```bash
pong
```

### Verifica si un humano es mutante

#### **`POST /mutant/`**

#### Descripción

Detecta si un humano es mutante enviando la secuencia de ADN. El tamaño de la secuencia de ADN debe cumplir con la definición de [*Matriz Cuadrada*](https://es.wikipedia.org/wiki/Matriz_cuadrada). En caso de no cumplir con la definición de Matriz Cuadrada, la API devolverá un código de estado de respuesta HTTP 400.

#### Parámetros

**Headers**:

| Parámetro | Tipo           | Descripción |
| :-------: | :-----------:  | :---------- |
| `Content-Type` | `string` | `application/json` |

**Body parameters**:

| Parámetro | Tipo           | Descripción |
| :-------: | :-----------:  | :---------- |
| `dna`     | `array (string)` | Secuencia de ADN. `Required ` <br> Solo pueden ser: (A,T,C,G), las cuales representa cada base nitrogenada del ADN. En caso de mandar la misma secuencia de ADN varias veces, se considera como una sola solicitud. |

#### Ejemplo de petición

**Body**:

```json
{
  "dna": [
    "ATGCGA",
    "CAGTGC",
    "TTATGT",
    "AGAAGG",
    "CCCCTA",
    "TCACTG"
  ]
}
```

#### Códigos de estado de respuesta HTTP

| Codigo | Descripción |
| :----: | :---------- |
| **200** | El ADN es mutante. <br> **Headers** <br> `x-request-id` (string): Id unico por petición, sirve para identificar los logs asociados a la petición. <br> **Body** <br> Body vacío |
| **400** | Petición Incorrecta.  <br> **Headers** <br> `x-request-id` (string): Id unico por petición, sirve para identificar los logs asociados a la petición. |
| **403** | El ADN no es mutante. <br> **Headers** <br> `x-request-id` (string): Id unico por petición, sirve para identificar los logs asociados a la petición. <br> **Body** <br> Body vacío |
| **500** | Error interno del servidor. <br> **Headers** <br> `x-request-id` (string): Id unico por petición, sirve para identificar los logs asociados a la petición. |

#### Ejemplo de respuesta

*Status Code*: 400

**Headers**:

```bash
HTTP/1.1 400 Bad Request
Content-Type: application/json; charset=utf-8
X-Request-Id: 5d8f9c8a-f9b7-4c8b-b8f9-f8b8f8b8f8b8
Date: Fri, 01 Jan 2019 19:00:00 GMT
Server: Google Frontend
Content-Length: 40
```

**Body**:

```json
{
    "status": 400,
    "error": "dna must be composed only of 'A', 'T', 'C' and 'G'"
}
```

### Genera estadísticas de mutantes y humanos

#### **`GET /stats/`**  

#### Descripción

Genera estadísticas básicas de numero de mutantes y humanos, y el ratio de mutantes sobre humanos.

#### Parámetros

No hay parámetros.

#### Códigos de estado de respuesta HTTP

| Codigo | Descripción |
| :----: | :---------- |
| **200** | Respuesta exitosa <br> **Headers** <br> `x-request-id` (string): Id unico por petición, sirve para identificar los logs asociados a la petición. |
| **500** | Error interno del servidor. <br> **Headers** <br> `x-request-id` (string): Id unico por petición, sirve para identificar los logs asociados a la petición. |

#### Ejemplo de respuesta

*Status Code*: 200

**Headers**:

```bash
HTTP/1.1 200 OK
Content-Type: application/json; charset=utf-8
X-Request-Id: 5d8f9c8a-f9b7-4c8b-b8f9-f8b8f8b8f8b8
Date: Fri, 01 Jan 2019 19:00:00 GMT
Server: Google Frontend
Content-Length: 40
```

**Body**:

```json
{
    "count_mutant_dna": 12,
    "count_human_dna": 15,
    "ratio": 0.80
}
```

## Autor

* [Jose Manuel Miranda V.](https://www.linkedin.com/in/getmiranda/) - Desarrollo, testing, documentación y despliegue.
