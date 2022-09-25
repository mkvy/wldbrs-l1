package main

import "fmt"

const (
	ethToBtc = 0.069
	btcToEth = 14.49
)

// Клиент
type Client struct{}

// Клиент работает только с BTC
func (c *Client) PayViaBTCToClient(cps ClientPaymentSystem, amount float32) {
	res := cps.PayViaBTC(amount, "0x1A8ac1B70A96266126766cDc4861bA0f8d35e4AA")
	if !res {
		fmt.Println("Unsuccessful payment")
		return
	}
	fmt.Println("Payment is successful")
}

// Клиентский интерфейс принимает оплату только в биткоинах
type ClientPaymentSystem interface {
	PayViaBTC(amount float32, address string) bool
}

// Биткоин кошелек и его имплементированный интерфейс оплаты через биткоин
type BitcoinWallet struct {
	balance float32
}

func (b *BitcoinWallet) PayViaBTC(amount float32, address string) bool {
	if amount > b.balance {
		return false
	}
	fmt.Printf("transfered %g BTC to address %s\n", amount, address)
	b.balance -= amount
	return true
}

// Структура, которую необходимо адаптировать к клиентскому интерфейсу
type EthereumWallet struct {
	balance float32
}

func (e *EthereumWallet) PayViaETH(amount float32, address string) bool {
	if amount > e.balance {
		return false
	}
	fmt.Printf("transfered %g ETH to address %s\n", amount, address)
	e.balance -= amount
	return true
}

// Структура адаптера, содержит в себе ссылку на адаптируемую структуру
type ETHToBTCAdapter struct {
	ethWallet *EthereumWallet
}

// Адаптированный интерфейс взаимодействия с клиентом
func (a *ETHToBTCAdapter) PayViaBTC(amount float32, address string) bool {
	b := BitcoinWallet{
		balance: a.ethWallet.balance * ethToBtc,
	}
	//конвертация эфириума в биткоин
	res := b.PayViaBTC(amount*ethToBtc, address)
	if res {
		a.ethWallet.balance -= amount
	}
	return res
}

func main() {
	//создаем клиента
	client := &Client{}
	btcWallet := &BitcoinWallet{balance: 15.1}

	fmt.Println("---Пытаемся отправить больше, чем в кошельке---")
	client.PayViaBTCToClient(btcWallet, 16)
	fmt.Println("---Отправляем меньше, чем в кошельке---")
	client.PayViaBTCToClient(btcWallet, 1.5)
	fmt.Println()
	fmt.Println()
	//иницализируем eth кошелек и адаптер
	ethWallet := &EthereumWallet{balance: 5555.34}
	fmt.Println("Баланс эфириума до: ", ethWallet.balance)
	adapter := &ETHToBTCAdapter{ethWallet: ethWallet}
	fmt.Println("---Отправляем через адаптер ETH---")
	//передаем клиенту адаптер, в котором имплементирован необходимый интерфейс
	client.PayViaBTCToClient(adapter, 1453.2)
	fmt.Println("Баланс после использования адаптера: ", ethWallet.balance)
}
