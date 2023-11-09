# Goroutines in Go

## Overview
Goroutines are a fundamental concept in Go for achieving concurrency. They are functions or methods that run concurrently with other functions or methods. Goroutines are lightweight and managed by the Go runtime.

## Creating a Goroutine
To start a goroutine, use the `go` keyword followed by the function call. This schedules the function to run concurrently.

## Synchronization
Goroutines execute independently, but sometimes you need to synchronize their execution. Go provides channels and synchronization primitives like `sync.WaitGroup` for that purpose.

## Channels
Channels are the preferred way for goroutines to communicate with each other. Data can be sent from one goroutine and received by another, allowing for safe and synchronized data exchange.

See examples of using goroutines in `goroutines.go`.


# Горутины в Go

## Обзор
Горутины — это основной концепт в Go для достижения конкурентности. Это функции или методы, выполняющиеся параллельно с другими функциями или методами. Горутины легковесны и управляются средой выполнения Go.

## Создание Горутины
Чтобы запустить горутину, используйте ключевое слово `go`, за которым следует вызов функции. Это планирует выполнение функции параллельно.

## Синхронизация
Горутины выполняются независимо, но иногда требуется синхронизировать их выполнение. Для этих целей Go предоставляет каналы и примитивы синхронизации, такие как `sync.WaitGroup`.

## Каналы
Каналы — предпочтительный способ коммуникации между горутинами. Данные могут быть отправлены одной горутиной и получены другой, что позволяет безопасно и синхронно обмениваться данными.

См. `goroutines.go`.

