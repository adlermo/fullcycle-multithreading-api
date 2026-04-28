# Fullcycle - Multithreading API

Busca CEP em duas APIs simultaneamente e retorna a resposta mais rápida.

---

## 🚀 Como rodar

```bash
go run main.go 01001000
```

##### __01001000 = número de CEP com 8 digitos sem traço__

---

## ⚙️ Como funciona

* Duas requisições são feitas em paralelo:

  * BrasilAPI
  * ViaCEP
* A primeira resposta vence
* Timeout global de 1 segundo

---

## 📦 Tecnologias

* Goroutines
* Channels
* Select
* net/http
* context

---

## ⚠️ Observações

* Apenas a resposta mais rápida é considerada
* A outra requisição é descartada
* Em caso de timeout, uma mensagem de erro é exibida
