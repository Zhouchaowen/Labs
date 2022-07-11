// 模式名称：建造者模式
// 目的（What）：将一个复杂对象的构建分离成多个简单对象的构建组合
// 解决的问题（Why）：当一个复杂对象的创建过程基本不变（组合方式），
// 但是该对象的各个部件各自创建的算法经常在变的时候。例如一个复杂对象的创建工作，
// 通常由各个部分的子对象用一定的算法构成；由于需求的变化，这个复杂对象的各个部分经常面临着剧烈的变化，
// 但是将它们组合在一起的算法却相对稳定。
// 优点:
// 建造者独立，容易扩展，用户使用不同的具体建造者即可得到不同的产品对象，
// 可以更加精细地控制产品创建的过程
// 缺点：
// 当建造者很多时，会产生很多的类，难以维护,
// 建造者模式所创建的产品一般具有较多的共同点，其组成部分相似，若产品之间的差异性很大，则不适合使用该模式，因此其使用范围受到一定限制,
// 若产品的内部变化复杂，可能会导致需要定义很多具体建造者类来实现这种变化，导致系统变得很庞大。
package builder

import "fmt"

// 场景
// 现在需要生产一个汉堡
// 生产过程包括增加香肠,生菜,奶酪,番茄
// 且有固定的顺序，即必须先确定香肠，再确定生菜,奶酪，最后放番茄的顺序，顺序错了，
// 或者是少了其中任意一个步骤都不能做好一个汉堡。

// Burger 产品汉堡包
type Burger struct {
	// 香肠
	Pepperoni string
	// 🧀️
	Cheese string
	// 🥬
	Lettuce string
	// 🍅
	Tomato string
}

// GetDescription 获取描述
func (burger *Burger) GetDescription() {
	fmt.Printf("Burger 的成分是: %v,%v,%v,%v \n", burger.Pepperoni, burger.Cheese, burger.Lettuce, burger.Tomato)
}

// Builder 构建器接口
type Builder interface {
	AddPepperoni(Pepperoni string) Builder
	AddCheese(Cheese string) Builder
	AddLettuce(Lettuce string) Builder
	AddTomato(Tomato string) Builder
	Build() *Burger
}

// BurgerBuilder 具体的建造者
type BurgerBuilder struct {
	burger *Burger
}

// AddPepperoni 添加香肠
func (burgerBuilder *BurgerBuilder) AddPepperoni(Pepperoni string) Builder {
	burgerBuilder.burger.Pepperoni = Pepperoni
	return burgerBuilder
}

// AddCheese 添加 🧀️
func (burgerBuilder *BurgerBuilder) AddCheese(Cheese string) Builder {
	burgerBuilder.burger.Cheese = Cheese
	return burgerBuilder
}

// AddLettuce 添加 🥬
func (burgerBuilder *BurgerBuilder) AddLettuce(Lettuce string) Builder {
	burgerBuilder.burger.Lettuce = Lettuce
	return burgerBuilder
}

// AddTomato 添加 🍅
func (burgerBuilder *BurgerBuilder) AddTomato(Tomato string) Builder {
	burgerBuilder.burger.Tomato = Tomato
	return burgerBuilder
}

// Build 构建函数
func (burgerBuilder *BurgerBuilder) Build() *Burger {
	if burgerBuilder.burger.Pepperoni == "" || burgerBuilder.burger.Cheese == "" || burgerBuilder.burger.Lettuce == "" || burgerBuilder.burger.Tomato == "" {
		fmt.Println("🍔 成分不全,制作失败！")
		return nil
	}
	return burgerBuilder.burger
}

// NewBurgerBuilder 创建一个面包构建器
func NewBurgerBuilder() *BurgerBuilder {
	return &BurgerBuilder{&Burger{}}
}

// https://github.com/youyingxiang/design-patterns-examples/blob/main/examples/
