# Meli Challenge Mutants API

![image](https://user-images.githubusercontent.com/25770844/169629362-56753248-df40-4301-a2b6-0397d1bfc361.png)

## Resumen

Este repositorio contiene el código de la API de Meli Challenge Mutants.

## Referencia de la API

### Verifica disponibilidad del servicio

**`GET /ping`**

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

**`POST /mutant/`**

#### Descripción

Detecta si un humano es mutante enviando la secuencia de ADN.

#### Parámetros

**Body parameters**:

| Parámetro | Tipo           | Descripción |
| :-------: | :-----------:  | :---------- |
| `dna`     | array (string) | Secuencia de ADN. <br> Solo pueden ser: (A,T,C,G), las cuales representa cada base nitrogenada del ADN. En caso de mandar la misma secuencia de ADN varias veces, se considera como una sola solicitud. |

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

#### Ejemplos de respuesta

**Status Code**: 400

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

**Status Code**: 500

**Headers**:

```bash
HTTP/1.1 500 Internal Server Error
Content-Type: application/json; charset=utf-8
X-Request-Id: 5d8f9c8a-f9b7-4c8b-b8f9-f8b8f8b8f8b8
Date: Fri, 01 Jan 2019 19:00:00 GMT
Server: Google Frontend
Content-Length: 40
```

**Body**:

```json
{
    "status": 500,
    "error": "database error"
}
```
