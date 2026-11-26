package main

import "fmt"

type Subscription struct {
	serviceName string
	weeklyCost  float64
	monthlyCost float64
	yearlyCost  float64
	isActive    bool
}

func newSubscription(serviceName string, isActive bool, periodType string, periodCost float64) *Subscription {
	if ok := validateSubscription(serviceName, periodType, periodCost); !ok {
		return nil
	}
	sub := Subscription{
		serviceName: serviceName,
		isActive:    isActive,
	}
	if ok := calcSubscriptionCost(&sub, periodType, periodCost); !ok {
		return nil
	}
	return &sub
}

func validateSubscription(serviceName string, periodType string, periodCost float64) bool {
	if serviceName == "" {
		fmt.Printf(
			"\nError in validateSubscription: service name is empty\n\n",
		)
		return false
	}
	if periodCost < 0 {
		fmt.Printf(
			"\nError in validateSubscription: cost must be >= 0\n\n",
		)
		return false
	}
	switch periodType {
	case "weekly", "monthly", "yearly":
		return true
	default:
		fmt.Printf(
			"\nError in validateSubscription: undefined period type (not in [weekly, monthly, yearly])\n\n",
		)
		return false
	}
}

func calcSubscriptionCost(subPtr *Subscription, periodType string, periodCost float64) bool {
	if subPtr == nil {
		fmt.Printf(
			"\nError in calcSubscriptionCost: passed pointer is nil\n\n",
		)
		return false
	}

	if !subPtr.isActive {
		resetSubscription(subPtr)
		return true
	}

	switch periodType {
	case "weekly":
		subPtr.weeklyCost = periodCost
		subPtr.yearlyCost = periodCost * 52
		subPtr.monthlyCost = subPtr.yearlyCost / 12
		return true
	case "monthly":
		subPtr.monthlyCost = periodCost
		subPtr.yearlyCost = periodCost * 12
		subPtr.weeklyCost = subPtr.yearlyCost / 52
		return true
	case "yearly":
		subPtr.yearlyCost = periodCost
		subPtr.monthlyCost = periodCost / 12
		subPtr.weeklyCost = periodCost / 52
		return true
	default:
		fmt.Printf(
			"\nError in calcSubscriptionCost: undefined period type (not in [weekly, monthly, yearly])\n\n",
		)
		return false
	}
}

func resetSubscription(subPtr *Subscription) {
	if subPtr == nil {
		fmt.Printf(
			"\nError in resetSubscription: passed pointer is nil\n\n",
		)
		return
	}
	subPtr.isActive = false
	subPtr.weeklyCost = 0
	subPtr.monthlyCost = 0
	subPtr.yearlyCost = 0
}

func printSubscriptionInfo(sub Subscription) {
	const rowHeaderFmt = "%-12s  %-22s  %-13s  %-14s  %-13s\n"
	const rowValuesFmt = "%-12s  %-22t  %-13.2f  %-14.2f  %-13.2f\n\n"
	fmt.Printf(rowHeaderFmt, "Service name", "Subscription is active", "Cost per week", "Cost per month", "Cost per year")
	fmt.Printf(rowHeaderFmt, "------------", "----------------------", "-------------", "--------------", "-------------")
	fmt.Printf(rowValuesFmt, sub.serviceName, sub.isActive, sub.weeklyCost, sub.monthlyCost, sub.yearlyCost)
}
