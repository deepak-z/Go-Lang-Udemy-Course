package main

import (
	"fmt"
	"sort"
	//"sort"
)
	
type person struct{
	first string
	age int
}

type ByAge []person

func (a ByAge) Len() int  { return len(a) }
func (a ByAge) Swap(i, j int) {a[i], a[j] = a[j], a[i]}
func (a ByAge) Less(i,j int) bool{ return a[i].age < a[j].age}

type ByName []person

func (a ByName) Len() int  { return len(a) }
func (a ByName) Swap(i, j int) {a[i], a[j] = a[j], a[i]}
func (a ByName) Less(i,j int) bool{ return a[i].first < a[j].first}

	TakerFeeVolume      float64        `json:"taker_fee_volume"`
	ExecutionPrice      float64        `json:"execution_price"`
	MakerFeePercentage  float32        `json:"maker_fee_percentage"`
	TakerFeePercentage  float32        `json:"taker_fee_percentage"`
	QodenExecutionInfo  datatypes.JSON `json:"qoden_execution_info"`
	//TODO Run a migration to remove them
	// CryptoTransactionID *int64         `json:"crypto_transaction_id"`
	// FiatTransactionID   *int64         `json:"fiat_transaction_id"`
	RequestId           uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4()" json:"request_id"`
			 
func main() {
		
	// xi := []int{4,7,3,42,99,18,16,52};
	// xs := []string{"What","are","you","doing"}

	// fmt.Println(xi)
	// fmt.Println(xs)

	// sort.Ints(xi);
	// fmt.Println(xi)
	
	// sort.Strings(xs)
	// fmt.Println(xs)


	// * Sort Custom

	p1 := person{"Deepak",22}
	p2 := person{"asdf",100}
	p3 := person{"gggg",1}
	p4 := person{"zzzz",77}

	people := []person{p1,p2,p3,p4}

	fmt.Println(people)

	sort.Sort(ByAge(people))

	fmt.Println(people)

	sort.Sort(ByName(people))

	fmt.Println(people)




	
			 
}
