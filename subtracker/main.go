package main

func main() {
	if sub1 := newSubscription("service1", true, "weekly", 132.34); sub1 != nil {
		printSubscriptionInfo(*sub1)
	}
	if sub2 := newSubscription("service2", false, "a", 1); sub2 != nil {
		printSubscriptionInfo(*sub2)
	}
	if sub3 := newSubscription("service3", true, "yearly", 1321); sub3 != nil {
		printSubscriptionInfo(*sub3)
	}
	if sub4 := newSubscription("service4", true, "monthly", 11); sub4 != nil {
		printSubscriptionInfo(*sub4)
	}
	if sub5 := newSubscription("service5", false, "monthly", 11); sub5 != nil {
		printSubscriptionInfo(*sub5)
	}
	if sub6 := newSubscription("", false, "monthly", 11); sub6 != nil {
		printSubscriptionInfo(*sub6)
	}
}
