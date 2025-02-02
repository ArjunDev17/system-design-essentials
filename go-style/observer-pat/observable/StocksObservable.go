package observable

import (
	"observer-pat/observer"
)

type StockObservable interface {
	Add(observer observer.NotificationAlertObserver)
	Remove(observer observer.NotificationAlertObserver)
	NotifySubscribers()
	SetStockCount(newStockAdded int)
	GetStockCount(str string) string
}
