package observer

import "testing"

func TestObserver(t *testing.T) {
	newsAgency := NewNewsAgency()
	
	channel1 := NewNewsChannel("CNN")
	channel2 := NewNewsChannel("BBC")
	
	newsAgency.Subscribe(channel1)
	newsAgency.Subscribe(channel2)
	
	newsAgency.Notify("Breaking News: Go 1.21 Released!")
	
	if channel1.GetNews() != "Breaking News: Go 1.21 Released!" {
		t.Error("Channel1 did not receive the correct news")
	}
	
	if channel2.GetNews() != "Breaking News: Go 1.21 Released!" {
		t.Error("Channel2 did not receive the correct news")
	}
	
	newsAgency.Unsubscribe(channel1)
	newsAgency.Notify("Update: Go 1.22 in development")
	
	if channel1.GetNews() == "Update: Go 1.22 in development" {
		t.Error("Channel1 should not have received the news after unsubscribing")
	}
	
	if channel2.GetNews() != "Update: Go 1.22 in development" {
		t.Error("Channel2 did not receive the updated news")
	}
}