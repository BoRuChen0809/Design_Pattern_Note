本文件用於本人自用學習筆記

資料來源:

- https://ithelp.ithome.com.tw/articles/10201706
- https://refactoringguru.cn/design-patterns/go

# Design Pattern

## 目錄

[TOC]



## **建立型模式(Creational Pattern)**

### **工廠方法模式(Factory Method Pattern)**

#### 定義

工廠方法模式( Factory Method )，定義一個用於建立物品的介面，讓子類決定實體化哪一個類別。工廠方法使一個類別的實例化延遲到其子類別。

#### UML

<img src="Design Pattern.assets/20112528fb3BbVITVH.png" style="zoom:55%;"/>

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

### **單例模式 (Singleton)**

#### 定義

只有一個實例，而且自行實例化並向整個系統提供這個實例。

#### UML

<img src="Design Pattern.assets/20112528d85wEM7c50.png" alt="20112528d85wEM7c50" style="zoom:75%;" />

定義一個操作，允許客戶端使用這個操作取得他的唯一實例，GetInstance()負責建立他的唯一實例。

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

### **原形模式 ( Prototype Pattern )**

#### 定義

原形模式其實就是從一個物件再建立另外一個可訂製的物件，而且不需要知道任何建立的細節。

#### UML

<img src="Design Pattern.assets/201125280QEfTTOn9P.png" style="zoom:60%;" />

#### 優點

- 你可以克隆對象， 而無需與它們所屬的具體類相耦合。
- 你可以克隆預生成原型， 避免反復運行初始化代碼。
- 你可以更方便地生成復雜對象。
- 你可以用繼承以外的方式來處理復雜對象的不同配置。

#### 缺點

- 克隆包含循環引用的復雜對象可能會非常麻煩。

#### Golang範例

##### 原型interface : indoe.go

```go
package main

type inode interface {
	clone() inode
	print(string)
}
```

##### 具體原型 : file.go

```go
package main

import "fmt"

type file struct {
	name string
}

func (f *file) print(indentation string) {
	fmt.Println(indentation + f.name)
}

func (f *file) clone() inode {
	return &file{name: f.name + "_clone"}
}
```

##### 具體原型 :  folder.go

```go
package main

import "fmt"

type folder struct {
	children []inode
	name     string
}

func (f *folder) print(indentation string) {
	fmt.Println(indentation + f.name)
	for _, i := range f.children {
		i.print(indentation + indentation)
	}
}

func (f *folder) clone() inode {
	cloneFolder := &folder{name: f.name + "_clone"}
	var tempChildren []inode
	for _, i := range f.children {
		copy := i.clone()
		tempChildren = append(tempChildren, copy)
	}
	cloneFolder.children = tempChildren
	return cloneFolder
}
```

##### main.go

```go
package main

import "fmt"

func main() {
	file1 := &file{name: "File1"}
	file2 := &file{name: "File2"}
	file3 := &file{name: "File3"}

	folder1 := &folder{
		children: []inode{file1},
		name:     "Folder1",
	}

	folder2 := &folder{
		children: []inode{folder1, file2, file3},
		name:     "Folder2",
	}
	fmt.Println("\nPrinting hierarchy for Folder2")
	folder2.print("  ")

	cloneFolder := folder2.clone()
	fmt.Println("\nPrinting hierarchy for clone Folder")
	cloneFolder.print("  ")
}

/*
Output:
Printing hierarchy for Folder2
  Folder2
    Folder1
        File1
    File2
    File3

Printing hierarchy for clone Folder
  Folder2_clone
    Folder1_clone
        File1_clone
    File2_clone
    File3_clone
*/
```

### **建造者模式 ( Builder Pattern)**

#### 定義

建造者模式是設計來提供一個有彈性解決方案，目的是為了要分離一個複雜物品的建造和表示建造的方式，最終實現用同一標準的製造工序能產出不同產品。

#### UML

<img src="Design Pattern.assets/20112528PhfQoo1N4v.png" style="zoom:60%;" />

一個物品在建造完成的過程中，需要組成需要設定過多的元件，或裝配過程是有順序的，可以在一一設定完所需要的元件後再產生我們所需要的物件。主要的角色有`建造者`、`指揮者`和`產品`，指揮者控制建造的過程，並且可以用來區隔用戶和建造過程的關聯性。

#### 優點

- 你可以分步創建對象，暫緩創建步驟或遞歸運行創建步驟。
- 生成不同形式的產品時， 你可以復用相同的製造代碼。
- 單一職責原則。 你可以將復雜構造代碼從產品的業務邏輯中分離出來。

#### 缺點

- 由於該模式需要新增多個類， 因此代碼整體復雜程度會有所增加。

#### Golang範例

##### 建造者interface : iBuilder.go

```go
package main

type iBuilder interface {
    setWindowType()
    setDoorType()
    setNumFloor()
    getHouse() house
}

func getBuilder(builderType string) iBuilder {
    if builderType == "normal" {
        return &normalBuilder{}
    }

    if builderType == "igloo" {
        return &iglooBuilder{}
    }
    return nil
}
```

##### 具體建造者 : normalBuilder.go

```go
package main

type normalBuilder struct {
    windowType string
    doorType   string
    floor      int
}

func newNormalBuilder() *normalBuilder {
    return &normalBuilder{}
}

func (b *normalBuilder) setWindowType() {
    b.windowType = "Wooden Window"
}

func (b *normalBuilder) setDoorType() {
    b.doorType = "Wooden Door"
}

func (b *normalBuilder) setNumFloor() {
    b.floor = 2
}

func (b *normalBuilder) getHouse() house {
    return house{
        doorType:   b.doorType,
        windowType: b.windowType,
        floor:      b.floor,
    }
}
```

##### 具體建造者 : iglooBuilder.go

```go
package main

type iglooBuilder struct {
    windowType string
    doorType   string
    floor      int
}

func newIglooBuilder() *iglooBuilder {
    return &iglooBuilder{}
}

func (b *iglooBuilder) setWindowType() {
    b.windowType = "Snow Window"
}

func (b *iglooBuilder) setDoorType() {
    b.doorType = "Snow Door"
}

func (b *iglooBuilder) setNumFloor() {
    b.floor = 1
}

func (b *iglooBuilder) getHouse() house {
    return house{
        doorType:   b.doorType,
        windowType: b.windowType,
        floor:      b.floor,
    }
}
```

##### 產品 : house.go

```go
package main

type house struct {
    windowType string
    doorType   string
    floor      int
}
```

##### 指揮者 : dirctor.go

```go
package main

type director struct {
    builder iBuilder
}

func newDirector(b iBuilder) *director {
    return &director{
        builder: b,
    }
}

func (d *director) setBuilder(b iBuilder) {
    d.builder = b
}

func (d *director) buildHouse() house {
    d.builder.setDoorType()
    d.builder.setWindowType()
    d.builder.setNumFloor()
    return d.builder.getHouse()
}
```

##### main.go

```go
package main

import "fmt"

func main() {
    normalBuilder := getBuilder("normal")
    iglooBuilder := getBuilder("igloo")

    director := newDirector(normalBuilder)
    normalHouse := director.buildHouse()

    fmt.Printf("Normal House Door Type: %s\n", normalHouse.doorType)
    fmt.Printf("Normal House Window Type: %s\n", normalHouse.windowType)
    fmt.Printf("Normal House Num Floor: %d\n", normalHouse.floor)

    director.setBuilder(iglooBuilder)
    iglooHouse := director.buildHouse()

    fmt.Printf("\nIgloo House Door Type: %s\n", iglooHouse.doorType)
    fmt.Printf("Igloo House Window Type: %s\n", iglooHouse.windowType)
    fmt.Printf("Igloo House Num Floor: %d\n", iglooHouse.floor)
}

/*
Output:
Normal House Door Type: Wooden Door
Normal House Window Type: Wooden Window
Normal House Num Floor: 2

Igloo House Door Type: Snow Door
Igloo House Window Type: Snow Window
Igloo House Num Floor: 1
*/
```

## **結構型模式(Structural Pattern)**

### **外觀模式 ( Facade Pattern )**

#### 定義

外觀模式(Facade)，為子系統中的一組介面提供一個一致的介面，此模式定義了一個高階介面，這個介面使得這一子系統更加容易使用。

#### UML

#### <img src="Design Pattern.assets/20112528PhfQoo1N4v-16434229871491.png" style="zoom:50%;" />

客戶端(Client)只要統一從Facade這個介面，就可以不用管背後子類別的邏輯。

#### 優點

- 將自己的程式碼獨立於複雜的子系統

#### 缺點

- 外觀可能成為與所有子類別都耦合的God Object

#### Golang範例

##### 外觀 : restaurantFacade.go

```go
package main

type restaurant struct {
	waiter    *Waiter
	cleaner   *Cleaner
	chef      *Chef
	vegvendor *VegVendor
}

func newRestaurant() *restaurant {
	return &restaurant{
		waiter:    newWaiter(),
		cleaner:   newCleaner(),
		chef:      newChef(),
		vegvendor: newVegVendor(),
	}
}

func (r *restaurant) order_dishes() {
	r.vegvendor.purchase()
	r.waiter.order()
	r.chef.cook()
	r.waiter.service()
	r.cleaner.clean()
}

```

##### 子系統 : VegVendor.go

```go
package main

import "fmt"

type VegVendor struct {
}

func newVegVendor() *VegVendor {
	return &VegVendor{}
}

func (v *VegVendor) purchase() {
	fmt.Println("供應蔬菜......")
}

```

##### 子系統 : Waiter.go

```go
package main

import "fmt"

type Waiter struct {
}

func newWaiter() *Waiter {
	return &Waiter{}
}

func (w *Waiter) order() {
	fmt.Println("接待、入座、點菜......")
}

func (w *Waiter) service() {
	fmt.Println("上菜......")
}

```

##### 子系統 : Chef.go

```go
package main

import "fmt"

type Chef struct {
}

func newChef() *Chef {
	return &Chef{}
}

func (c *Chef) cook() {
	fmt.Println("下廚烹飪......")
}

```

##### 子系統 : Cleaner.go

```go
package main

import "fmt"

type Cleaner struct {
}

func newCleaner() *Cleaner {
	return &Cleaner{}
}

func (c *Cleaner) clean() {
	fmt.Println("清理桌面、洗碗......")
}

```

##### main.go

```go
package main

func main() {
	r := newRestaurant()

	r.order_dishes()
}

/*
Output:
供應蔬菜......
接待、入座、點菜......
下廚烹飪......
上菜......
清理桌面、洗碗......
*/

```

### **組合模式 ( Composite Pattern)**

#### 定義

組合模式 ( Composite )，將物件組合成樹形結構以表示「 部分 - 整體 」的層次結構。組合模式使得用戶對單個物件和組合物件的操作時會得到統一的回應。

#### UML

<img src="Design Pattern.assets/20112528y9tmrJlZ5C.png" style="zoom:50%;" />

#### 優點

- 你可以利用多態和遞歸機制更方便地使用復雜樹結構。
- 開閉原則。 無需更改現有代碼， 你就可以在應用中添加新元素， 使其成為對象樹的一部分。

#### 缺點

- 對於功能差異較大的類，提供公共介面或許會有困難。 在特定情況下，你需要過度一般化組件介面，使其變得令人難以理解。

#### Golang範例

##### Component interface : army.go

```go
package main

type army interface {
	showInfo()
	getLevel() int
	getPeople() int
	getNmae() string
}

```

##### Leaf : squard.go

```go
//班
package main

import "fmt"

type squard struct {
	Name   string
	People int
}

func (s *squard) showInfo() {
	fmt.Printf("the squard name is %s, there are %d people in here.\n", s.getNmae(), s.getPeople())
}

func (s *squard) getLevel() int {
	return 1
}

func (s *squard) getPeople() int {
	return s.People
}

func (s *squard) setPeople(num int) {
	s.People = num
}

func (s *squard) getNmae() string {
	return s.Name
}

func (s *squard) setName(name string) {
	s.Name = name
}

```

##### Composite : platoon.go

```go
//排
package main

import "fmt"

type platoon struct {
	children []army
	Name     string
	People   int
}

func (p *platoon) showInfo() {
	fmt.Printf("the platoon name is %s, there are %d people in here. ", p.getNmae(), p.getPeople())
	fmt.Printf("there are %d squards in this platoon, detail:\n", len(p.children))
	for _, a := range p.children {
		fmt.Print("。")
		a.showInfo()
	}
}

func (p *platoon) getLevel() int {
	return 2
}

func (p *platoon) setPeople() {
	num := 0
	for _, a := range p.children {
		num += a.getPeople()
	}
	p.People = num
}

func (p *platoon) getPeople() int {
	p.setPeople()
	return p.People
}

func (p *platoon) getNmae() string {
	return p.Name
}

func (p *platoon) setName(name string) {
	p.Name = name
}

func (p *platoon) add(a army) {
	if p.getLevel()-a.getLevel() != 1 {
		fmt.Println("this operation is invaild")
	} else {
		p.children = append(p.children, a)
	}
}

```

##### Composite : company.go

```go
//連
package main

import "fmt"

type company struct {
	children []army
	Name     string
	People   int
}

func (c *company) showInfo() {
	fmt.Printf("the company name is %s, there are %d people in here. ", c.getNmae(), c.getPeople())
	fmt.Printf("there are %d platoons in this platoon, detail:\n", len(c.children))
	for _, a := range c.children {
		fmt.Print("。")
		a.showInfo()
	}
}

func (c *company) getLevel() int {
	return 3
}

func (c *company) setPeople() {
	num := 0
	for _, a := range c.children {
		num += a.getPeople()
	}
	c.People = num
}

func (c *company) getPeople() int {
	c.setPeople()
	return c.People
}

func (c *company) getNmae() string {
	return c.Name
}

func (c *company) setName(name string) {
	c.Name = name
}

func (c *company) add(a army) {
	if c.getLevel()-a.getLevel() != 1 {
		fmt.Println("this operation is invaild")
	} else {
		c.children = append(c.children, a)
	}
}

```

##### main.go

```go
package main

func main() {
	var a army

	squard_1 := &squard{"1", 15}
	squard_2 := &squard{"2", 20}
	squard_3 := &squard{"3", 11}

	platoon_1 := &platoon{Name: "P1"}
	platoon_1.add(squard_1)
	platoon_1.add(squard_2)
	platoon_1.add(squard_3)

	platoon_2 := &platoon{Name: "P2"}
	platoon_2.add(&squard{"4", 18})
	platoon_2.add(&squard{"7", 19})
	platoon_2.add(&squard{"5", 10})

	platoon_3 := &platoon{Name: "P3"}
	platoon_3.add(&squard{"6", 14})
	platoon_3.add(&squard{"9", 13})
	platoon_3.add(&squard{"8", 17})

	C := &company{Name: "Company"}
	C.add(platoon_2)
	C.add(platoon_3)
	C.add(platoon_1)

	a = C
	a.showInfo()
}

/*
Output:
the company name is Company, there are 137 people in here. there are 3 platoons in this company, detail:
。the platoon name is P2, there are 47 people in here. there are 3 squards in this platoon, detail:
。the squard name is 4, there are 18 people in here.
。the squard name is 7, there are 19 people in here.
。the squard name is 5, there are 10 people in here.
。the platoon name is P3, there are 44 people in here. there are 3 squards in this platoon, detail:
。the squard name is 6, there are 14 people in here.
。the squard name is 9, there are 13 people in here.
。the squard name is 8, there are 17 people in here.
。the platoon name is P1, there are 46 people in here. there are 3 squards in this platoon, detail:
。the squard name is 1, there are 15 people in here.
。the squard name is 2, there are 20 people in here.
。the squard name is 3, there are 11 people in here.
*/
```

