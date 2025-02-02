package observable

import (
	"fmt"
	"observer-pat/observer"
)

// IphoneObservableImpl represents the concrete observable that manages the stock count
type IphoneObservableImpl struct {
	observeList []observer.NotificationAlertObserver
	stockCount  int
}

// NewIphoneObservableImpl creates and initializes a new IphoneObservableImpl
func NewIphoneObservableImpl() *IphoneObservableImpl {
	return &IphoneObservableImpl{
		observeList: []observer.NotificationAlertObserver{},
	}
}

// Add adds an observer to the list
func (i *IphoneObservableImpl) Add(o observer.NotificationAlertObserver) {
	i.observeList = append(i.observeList, o)
}

// Remove removes an observer from the list
func (i *IphoneObservableImpl) Remove(o observer.NotificationAlertObserver) {
	for index, observer := range i.observeList {
		if observer == o {
			i.observeList = append(i.observeList[:index], i.observeList[index+1:]...)
			break
		}
	}
}

// NotifySubscribers notifies all the observers of the stock count update
func (i *IphoneObservableImpl) NotifySubscribers() {
	for _, observer := range i.observeList {
		observer.Update()
	}
}

// SetStockCount updates the stock count and notifies the observers
func (i *IphoneObservableImpl) SetStockCount(newStockAdded int) {
	i.stockCount += newStockAdded
	fmt.Printf("Stock updated: %d items\n", i.stockCount)
	i.NotifySubscribers()
}

// GetStockCount returns the current stock count as a string
func (i *IphoneObservableImpl) GetStockCount(str string) string {
	if str == "stockCount" {
		return fmt.Sprintf("Current stock count: %d", i.stockCount)
	}
	return "Invalid request"
}
