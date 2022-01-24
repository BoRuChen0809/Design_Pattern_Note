[TOC]

本文件用於本人自用學習筆記

資料來源:

- https://ithelp.ithome.com.tw/articles/10201706
- https://refactoringguru.cn/design-patterns/go

# Design Pattern

## 建立型模式

### 工廠方法模式(Factory Method Pattern)

#### 定義

工廠方法模式( Factory Method )，定義一個用於建立物品的介面，讓子類決定實體化哪一個類別。工廠方法使一個類別的實例化延遲到其子類別。

#### UML

<img src="Design Pattern.assets/20112528fb3BbVITVH.png" style="zoom:50%;" />

- Creator：創造者經由FactoryMethod創造產品
- Product：被創造的產品類別

#### 優點

- 你可以避免創建者和具體產品之間的緊密耦合。
-  單一職責原則。 你可以將產品創建代碼放在程式的單一位置， 從而使得代碼更容易維護。
-  開閉原則。 無需更改現有客戶端代碼， 你就可以在程式中引入新的產品類型。

#### 缺點

- 應用工廠方法模式需要引入許多新的子類，代碼可能會因此變得更復雜。 最好的情況是將該模式引入創建者類的現有層次結構中。

#### Golang範例

##### 產品interface : product.go

```go
type product interface {
	setName(name string)
	setDescribe(describe string)
	getName() string
	getDescribe() string
}
```

##### 具體產品 : food.go

```go
package main

type food struct {
	name     string
	describe string
}

func (f *food) setName(name string) {
	f.name = name
}

func (f *food) getName() string {
	return f.name
}

func (f *food) setDescribe(describe string) {
	f.describe = describe
}

func (f *food) getDescribe() string {
	return f.describe
}

```

##### 具體產品 : steak.go

```go
package main

type steak struct {
	food
}

func newSteak() product {
	return &steak{food: food{
		name:     "SteaKKKKKKKKKKK",
		describe: "Medium",
	}}
}

```

##### 具體產品 : pork_chop.go

```go
package main

type pork_chop struct {
	food
}

func newPork_Chop() product {
	return &pork_chop{food: food{
		name:     "pork chop",
		describe: "It is a pork chop.",
	}}
}

```

##### 工廠 : steakhouse.go

```go
package main

import "errors"

func getFood(foodType string) (product, error) {
	if foodType == "steak" {
		return newSteak(), nil
	}
	if foodType == "pork chop" {
		return newPork_Chop(), nil
	}
	notFound := errors.New("the type of product you want is not exist.")
	return nil, notFound
}

```

##### main.go

```go
package main

import "fmt"

func main() {
	product_1, err := getFood("steak")
	if err != nil {
		fmt.Println(err)
	} else {
		product_detail(product_1)
	}
	product_2, err := getFood("pork chop")
	if err != nil {
		fmt.Println(err)
	} else {
		product_detail(product_2)
	}
}

func product_detail(p product) {
	fmt.Printf("Product Type : %s\n", p.getName())
	fmt.Printf("Product Describe : %s\n", p.getDescribe())
	fmt.Println("********************************************************************************************")
}


/*
Output:
Product Type : SteaKKKKKKKKKKK
Product Describe : Medium
********************************************************************************************
Product Type : pork chop
Product Describe : It is a pork chop.
********************************************************************************************
*/
```

