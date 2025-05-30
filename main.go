package main

import (
	"fmt"
	"golang-design-pattern/chain"
	"golang-design-pattern/decorator"
	"golang-design-pattern/iterator"
	"golang-design-pattern/observer"
	"golang-design-pattern/singleton"
	"golang-design-pattern/state"
	"golang-design-pattern/strategy"
)

func main() {
	fmt.Println("=== Go 设计模式示例 ===")
	fmt.Println()

	// 1. 单例模式示例
	fmt.Println("1. 单例模式 (Singleton Pattern):")
	db1 := singleton.GetInstance()
	db2 := singleton.GetInstance()
	fmt.Printf("db1 == db2: %t\n", db1 == db2)
	fmt.Println(db1.Query("SELECT * FROM users"))
	fmt.Println()

	// 2. 观察者模式示例
	fmt.Println("2. 观察者模式 (Observer Pattern):")
	newsAgency := observer.NewNewsAgency()
	cnn := observer.NewNewsChannel("CNN")
	bbc := observer.NewNewsChannel("BBC")

	newsAgency.Subscribe(cnn)
	newsAgency.Subscribe(bbc)
	newsAgency.Notify("重大新闻：Go 1.22 正式发布！")
	fmt.Println()

	// 3. 策略模式示例
	fmt.Println("3. 策略模式 (Strategy Pattern):")
	amount := 150.0

	// 使用信用卡支付
	creditCard := strategy.NewCreditCardPayment("1234-5678-9012-3456")
	paymentContext := strategy.NewPaymentContext(creditCard)
	fmt.Println(paymentContext.ExecutePayment(amount))

	// 使用PayPal支付
	paypal := strategy.NewPayPalPayment("user@example.com")
	paymentContext.SetStrategy(paypal)
	fmt.Println(paymentContext.ExecutePayment(amount))
	fmt.Println()

	// 4. 装饰器模式示例
	fmt.Println("4. 装饰器模式 (Decorator Pattern):")
	coffee := &decorator.SimpleCoffee{}
	fmt.Printf("基础咖啡: %s - $%.2f\n", coffee.Description(), coffee.Cost())

	milkCoffee := decorator.NewMilkDecorator(coffee)
	fmt.Printf("加牛奶: %s - $%.2f\n", milkCoffee.Description(), milkCoffee.Cost())

	fullCoffee := decorator.NewWhipDecorator(decorator.NewSugarDecorator(milkCoffee))
	fmt.Printf("全配置: %s - $%.2f\n", fullCoffee.Description(), fullCoffee.Cost())
	fmt.Println()

	// 5. 状态模式示例
	fmt.Println("5. 状态模式 (State Pattern):")
	vendingMachine := state.NewVendingMachine()
	fmt.Println("自动售货机操作流程:")
	vendingMachine.InsertCoin()
	vendingMachine.SelectProduct()
	vendingMachine.DispenseProduct()
	fmt.Println()

	// 6. 责任链模式示例
	fmt.Println("6. 责任链模式 (Chain of Responsibility Pattern):")
	logger := chain.NewLogger()
	logger.Log(1, "系统启动成功")
	logger.Log(2, "内存使用率较高")
	logger.Log(3, "数据库连接失败")
	logger.Log(4, "系统崩溃")
	fmt.Println()

	// 7. 迭代器模式示例
	fmt.Println("7. 迭代器模式 (Iterator Pattern):")
	bookCollection := iterator.NewBookCollection()
	bookCollection.AddBook(iterator.Book{Title: "Go编程", Author: "张三"})
	bookCollection.AddBook(iterator.Book{Title: "设计模式", Author: "李四"})
	bookCollection.AddBook(iterator.Book{Title: "算法导论", Author: "王五"})

	fmt.Println("遍历书籍集合:")
	bookIterator := bookCollection.CreateIterator()
	for bookIterator.HasNext() {
		book := bookIterator.Next().(iterator.Book)
		fmt.Printf("- %s\n", book.String())
	}
}