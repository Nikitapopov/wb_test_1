package task_6

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"time"
)

// Реализовать все возможные способы остановки выполнения горутины.

// PS Не уверен, что понял правильно задание, поэтому реализовал все возможные
// варианты, которые могут повлиять на корректную работу горутины

func Execute() {
	// Закрытие горутины с помощью передачи по каналу
	quit := make(chan struct{})
	go func(quit <-chan struct{}) {
		select {
		case <-quit:
			return
		}
	}(quit)
	quit <- struct{}{}

	// Частный вариант предыдущего примера. Закрытие горутины через определенное время
	timeout := time.After(1 * time.Second)
	go func(timeout <-chan time.Time) {
		select {
		case <-timeout:
			return
		}
	}(timeout)

	// Еще один частный вариант первого примера. Закрытие горутины через определенное время с помощью контекста
	ctx, ctxCancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer ctxCancel()
	go func(ctx context.Context) {
		select {
		case <-ctx.Done():
			return
		}
	}(ctx)

	// Остановка работы путем перемещения горутины в очередь ожидания планировщика из-за ожидания чтения данных в канале
	// Аналогично с ожиданием записи или буферизированным каналом
	ch := make(chan struct{})
	go func(ch <-chan struct{}) {
		<-ch
	}(ch)

	// Остановка работы из-за ожидания разблокировки мьютекса
	mtx := sync.Mutex{}
	mtx.Lock()
	go func(mtx *sync.Mutex) {
		mtx.Lock()
	}(&mtx)

	// Остановка работы из-за ожидания группы
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		wg.Wait()
	}(&wg)

	// Бесконечный цикл
	go func() {
		select {}
		fmt.Println("Won't be printed")
	}()
	go func() {
		for true {
		}
		fmt.Println("Won't be printed")
	}()

	// Остановка из-за работы сборщика мусора
	// Аналогично конкурентные системные вызовы
	go func() {
		// Этот код может быть заблокирован на время, пока выполняется сборка мусора
	}()
	go func() {
		runtime.GC()
	}()

	// Преждевременное завершение горутины из-за завершения главной
	go func() {
		ch := time.After(1 * time.Second)
		go func(ch <-chan time.Time) {
			select {
			case <-ch:
				fmt.Println("Won't be printed")
			}
		}(ch)
	}()
}
