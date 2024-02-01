package golangcontext

import (
	"context"
	"fmt"
	"runtime"
	"testing"
	"time"
)

/**
Membuat Context

● Karena Context adalah sebuah interface, untuk membuat context kita butuh sebuah struct yang
sesuai dengan kontrak interface Context
● Namun kita tidak perlu membuatnya secara manual
● Di Golang package context terdapat function yang bisa kita gunakan untuk membuat Context
*/

/**
Function Membuat Context

context.Background() : Membuat context kosong. Tidak pernah dibatalkan, tidak pernah timeout, dan tidak memiliki value apapun. Biasanya
digunakan di main function atau dalam test, atau dalam awal proses request terjadi.

context.TODO() : Membuat context kosong seperti Background(), namun biasanya menggunakan ini ketika belum jelas context apa yang ingin digunakan
*/

func TestContex(t *testing.T) {
	background := context.Background()
	fmt.Println(background)

	todo := context.TODO()
	fmt.Println(todo)
}

/**
Context With Value

● Pada saat awal membuat context, context tidak memiliki value
● Kita bisa menambah sebuah value dengan data Pair (key - value) ke dalam context
● Saat kita menambah value ke context, secara otomatis akan tercipta child context baru, artinya
original context nya tidak akan berubah sama sekali
● Untuk membuat menambahkan value ke context, kita bisa menggunakan function
context.WithValue(parent, key, value)
*/

func TestWithValue(t *testing.T) {
	contextA := context.Background()

	contextB := context.WithValue(contextA, "b", "B")
	contextC := context.WithValue(contextA, "c", "C")

	contextD := context.WithValue(contextB, "d", "D")
	contextE := context.WithValue(contextB, "d", "D")

	contextF := context.WithValue(contextC, "f", "F")

	fmt.Println(contextA)
	fmt.Println(contextB)
	fmt.Println(contextC)
	fmt.Println(contextD)
	fmt.Println(contextE)
	fmt.Println(contextF)

	fmt.Println(contextF.Value("f"))
	fmt.Println(contextF.Value("c"))
	fmt.Println(contextF.Value("b"))

	fmt.Println(contextA.Value("b"))
}

/**
Context With Cancel

● Selain menambahkan value ke context, kita juga bisa menambahkan sinyal cancel ke context
● Kapan sinyal cancel diperlukan dalam context?
● Biasanya ketika kita butuh menjalankan proses lain, dan kita ingin bisa memberi sinyal cancel ke
proses tersebut
● Biasanya proses ini berupa goroutine yang berbeda, sehingga dengan mudah jika kita ingin
membatalkan eksekusi goroutine, kita bisa mengirim sinyal cancel ke context nya
● Namun ingat, goroutine yang menggunakan context, tetap harus melakukan pengecekan terhadap
context nya, jika tidak, tidak ada gunanya
● Untuk membuat context dengan cancel signal, kita bisa menggunakan function
context.WithCancel(parent)

contoh untuk goroutine leak
*/

func CreateCounterLeak() chan int {
	destination := make(chan int)

	go func() {
		defer close(destination)

		counter := 1
		for {
			destination <- counter
			counter++
		}
	}()
	return destination
}
func CreateCounter(ctx context.Context) chan int {
	destination := make(chan int)

	go func() {
		defer close(destination)

		counter := 1
		for {
			select {
			case <-ctx.Done():
				return
			default:
				destination <- counter
				counter++
				time.Sleep(1 * time.Second) //simulate slow process
			}
		}
	}()
	return destination
}

func TestContextWithCancleLeak(t *testing.T) {
	fmt.Println("total goroutine :", runtime.NumGoroutine())

	destination := CreateCounterLeak()
	for n := range destination {
		fmt.Println("counter", n)
		if n == 10 {
			break
		}
	}
	fmt.Println("total goroutine :", runtime.NumGoroutine())
}

func TestContextWithCancle(t *testing.T) {
	fmt.Println("total goroutine :", runtime.NumGoroutine())

	//solve go leak
	parent := context.Background()
	ctx, cancle := context.WithCancel(parent)

	destination := CreateCounter(ctx)
	fmt.Println("total goroutine :", runtime.NumGoroutine())

	for n := range destination {
		fmt.Println("counter", n)
		if n == 10 {
			break
		}
	}
	cancle() //mengirim sinyal cancle ke context

	time.Sleep(2 * time.Second)
	fmt.Println("total goroutine :", runtime.NumGoroutine())
}

/**
Context With Timeout

● Selain menambahkan value ke context, dan juga sinyal cancel, kita juga bisa menambahkan sinyal
cancel ke context secara otomatis dengan menggunakan pengaturan timeout
● Dengan menggunakan pengaturan timeout, kita tidak perlu melakukan eksekusi cancel secara
manual, cancel akan otomatis di eksekusi jika waktu timeout sudah terlewati
● Penggunaan context dengan timeout sangat cocok ketika misal kita melakukan query ke database
atau http api, namun ingin menentukan batas maksimal timeout nya
● Untuk membuat context dengan cancel signal secara otomatis menggunakan timeout, kita bisa
menggunakan function context.WithTimeout(parent, duration)
*/

func TestContextWithTimeout(t *testing.T) {
	fmt.Println("total goroutine :", runtime.NumGoroutine())

	//solve go leak
	parent := context.Background()
	ctx, cancle := context.WithTimeout(parent, 5*time.Second)
	defer cancle() //mengirim sinyal cancle ke context

	destination := CreateCounter(ctx)
	fmt.Println("total goroutine :", runtime.NumGoroutine())

	for n := range destination {
		fmt.Println("counter", n)
	}

	time.Sleep(2 * time.Second)
	fmt.Println("total goroutine :", runtime.NumGoroutine())
}

/**
Context With Deadline

● Selain menggunakan timeout untuk melakukan cancel secara otomatis, kita juga bisa menggunakan
deadline
● Pengaturan deadline sedikit berbeda dengan timeout, jika timeout kita beri waktu dari sekarang,
kalo deadline ditentukan kapan waktu timeout nya, misal jam 12 siang hari ini
● Untuk membuat context dengan cancel signal secara otomatis menggunakan deadline, kita bisa
menggunakan function context.WithDeadline(parent, time)
*/

func TestContextWithDeadline(t *testing.T) {
	fmt.Println("total goroutine :", runtime.NumGoroutine())

	//solve go leak
	parent := context.Background()
	ctx, cancle := context.WithDeadline(parent, time.Now().Add(5*time.Second))
	defer cancle() //mengirim sinyal cancle ke context

	destination := CreateCounter(ctx)
	fmt.Println("total goroutine :", runtime.NumGoroutine())

	for n := range destination {
		fmt.Println("counter", n)
	}

	time.Sleep(2 * time.Second)
	fmt.Println("total goroutine :", runtime.NumGoroutine())
}
