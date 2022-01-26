本文件用於本人自用學習筆記

資料來源:

- https://ithelp.ithome.com.tw/articles/10201706
- https://refactoringguru.cn/design-patterns/go

# Design Pattern

## 目錄

[TOC]



## 建立型模式

### **工廠方法模式(Factory Method Pattern)**

#### 定義

工廠方法模式( Factory Method )，定義一個用於建立物品的介面，讓子類決定實體化哪一個類別。工廠方法使一個類別的實例化延遲到其子類別。

#### UML

<img src="Design Pattern.assets/20112528fb3BbVITVH.png" style="zoom:70%;" />

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

------

### **抽象工廠模式 ( Abstract Factory )**

#### 定義

抽象工廠模式( Abstract Factory)，提供一個建立一系列相關或互相依賴物件的介面，而無需指定它們具體的類別。

#### UML

<img src="Design Pattern.assets/201125282oz5lmOvYH.png" style="zoom:70%;" />

- AbstractProductA：產品A系列的介面。
- ProductA1、ProductA2：A系列下的各個產品。
- AbsctractProductB：產品B系列的介面。
- ProductB1、ProductB2：B系列下的各個產品。
- AbstractFactory：抽象工廠。
- ConcreteFactory1、ConcreteFactory2：具體的工廠實現。

#### 優點

-  你可以確保同一工廠生成的產品相互匹配。
-  你可以避免客戶端和具體產品代碼的耦合。
-  單一職責原則。 你可以將產品生成代碼抽取到同一位置， 使得代碼易於維護。
-  開閉原則。 向應用程式中引入新產品變體時， 你無需修改客戶端代碼。

#### 缺點

-  由於採用該模式需要向應用中引入眾多介面和類， 代碼可能會比之前更加復雜。

#### Golang範例

##### 抽象工廠interface : iFactory.go

```go
package main

import "fmt"

type iFactory interface {
	makeShirt() iShirt
	makeShoe() iShoe
}

func getSportsFactory(brand string) (iFactory, error) {
	if brand == "UA" {
		return &UA{}, nil
	}

	if brand == "Nike" {
		return &Nike{}, nil
	}

	return nil, fmt.Errorf("Wrong brand type passed")
}
```

##### 具體工廠 : nike.go

```go
package main

type Nike struct {
}

func (n *Nike) makeShirt() iShirt {
	return &nikeShirt{
		shirt{
			logo: "nike",
			size: "XXL",
		},
	}
}

func (n *Nike) makeShoe() iShoe {
	return &nikeShoe{
		shoe{
			logo: "nike",
			size: "US_12",
		},
	}
}
```



##### 具體工廠 : under_armour.go

```go
package main

type UA struct {
}

func (ua *UA) makeShirt() iShirt {
	return &uaShirt{
		shirt{
			logo: "UA",
			size: "XL",
		},
	}
}

func (ua *UA) makeShoe() iShoe {
	return &uaShoe{
		shoe{
			logo: "UA",
			size: "US_11",
		},
	}
}
```

##### 抽象產品 : iShirt.go

```go
package main

type iShirt interface {
	setLogo(logo string)
	setSize(size string)
	getLogo() string
	getSize() string
}

type shirt struct {
	logo string
	size string
}

func (s *shirt) setLogo(logo string) {
	s.logo = logo
}

func (s *shirt) getLogo() string {
	return s.logo
}

func (s *shirt) setSize(size string) {
	s.size = size
}

func (s *shirt) getSize() string {
	return s.size
}
```

##### 具體產品 : nikeShirt.go

```go
package main

type nikeShirt struct {
	shirt
}
```

##### 具體產品 : uaShirt.go

```go
package main

type uaShirt struct {
	shirt
}
```

##### 抽象產品 : iShoe.go

```go
package main

type iShoe interface {
	setLogo(logo string)
	setSize(size string)
	getLogo() string
	getSize() string
}

type shoe struct {
	logo string
	size string
}

func (s *shoe) setLogo(logo string) {
	s.logo = logo
}

func (s *shoe) getLogo() string {
	return s.logo
}

func (s *shoe) setSize(size string) {
	s.size = size
}

func (s *shoe) getSize() string {
	return s.size
}
```

##### 具體產品 : nikeShoe.go

```go
package main

type nikeShoe struct {
	shoe
}
```

##### 具體產品 : uaShoe.go

```go
package main

type uaShoe struct {
	shoe
}
```

##### main.go

```go
package main

import "fmt"

func main() {
	nike_factory, err := getSportsFactory("Nike")
	if err != nil {
		fmt.Println(err)
	}
	nikeShoe := nike_factory.makeShoe()
	nikeShirt := nike_factory.makeShirt()
	printShirtDetails(nikeShirt)
	printShoeDetails(nikeShoe)

	ua_factory, err := getSportsFactory("UA")
	if err != nil {
		fmt.Println(err)
	}
	uaShirt := ua_factory.makeShirt()
	uaShoe := ua_factory.makeShoe()
	printShirtDetails(uaShirt)
	printShoeDetails(uaShoe)

}

func printShoeDetails(s iShoe) {
	fmt.Println("Shoe")
	fmt.Printf("Logo : %s\n", s.getLogo())
	fmt.Printf("Size : %s\n", s.getSize())
	fmt.Println("************************************************************************")
}

func printShirtDetails(s iShirt) {
	fmt.Println("Shirt")
	fmt.Printf("Logo : %s\n", s.getLogo())
	fmt.Printf("Size : %s\n", s.getSize())
	fmt.Println("************************************************************************")
}

/*
Output:
Shirt
Logo : nike
Size : XXL
************************************************************************
Shoe
Logo : nike
Size : US_12
************************************************************************
Shirt
Logo : UA
Size : XL
************************************************************************
Shoe
Logo : UA
Size : US_11
************************************************************************
*/
```

## 單例模式 (Singleton)

#### 定義

只有一個實例，而且自行實例化並向整個系統提供這個實例。

#### UML

<img src="Design Pattern.assets/20112528d85wEM7c50.png" alt="20112528d85wEM7c50" style="zoom:75%;" />

#### 優點

- 你可以保證一個類只有一個實例。
- 你獲得了一個指向該實例的全局訪問節點。
- 僅在首次請求單例對象時對其進行初始化。

#### 缺點

- 違反了_單一職責原則_。 該模式同時解決了兩個問題。
- 單例模式可能掩蓋不良設計， 比如程式各組件之間相互瞭解過多等。
- 該模式在多線程環境下需要進行特殊處理， 避免多個線程多次創建單例對象。
- 單例的客戶端代碼單元測試可能會比較困難，因為許多測試框架以基於繼承的方式創建模擬對象。由於單例類的構造函數是私有的，而且絕大部分語言無法重寫靜態方法，所以你需要想出仔細考慮模擬單例的方法。要麼乾脆不編寫測試代碼，或者不使用單例模式。

#### Golang範例

##### 單例 : sun.go

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

var lock = &sync.Mutex{}

type sun struct {
	description string
}

func (s *sun) Describe() string {
	return s.description
}

var sunInstance *sun

func getSun() *sun {
	if sunInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if sunInstance == nil {
			now := time.Now().Format("2006/01/02 15:04:05")
			fmt.Println("Create Sun.")
			sunInstance = &sun{fmt.Sprintf("created on : %s", now)}
			fmt.Println(sunInstance.Describe())
			return sunInstance
		} else {
			fmt.Println("Sun had created.")
			fmt.Println(sunInstance.Describe())
			return sunInstance
		}
	} else {
		fmt.Println("Sun had created.")
		fmt.Println(sunInstance.Describe())
		return sunInstance
	}
}

```

##### main.go

```go
package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		go getSun()
	}
	fmt.Scanln()
}

/*
Output:
Create Sun.
created on : 2022/01/26 09:02:41
Sun had created.
created on : 2022/01/26 09:02:41
Sun had created.
created on : 2022/01/26 09:02:41
Sun had created.
created on : 2022/01/26 09:02:41
Sun had created.
created on : 2022/01/26 09:02:41
Sun had created.
created on : 2022/01/26 09:02:41
Sun had created.
created on : 2022/01/26 09:02:41
Sun had created.
created on : 2022/01/26 09:02:41
Sun had created.
created on : 2022/01/26 09:02:41
Sun had created.
created on : 2022/01/26 09:02:41
*/
```

