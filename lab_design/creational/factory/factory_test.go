package main

import "testing"

// command: go test -v factory_method_test.go factory_method.go
func TestFactoryMethod(t *testing.T) {
	// 开发者面试
	developer := &Developer{}
	// 获取一个给开发者面试的面试官
	developmentManager := NewHiringManager(developer)
	// 询问开发者问题
	if str := developmentManager.TakeInterview(); str != "问关于开发的问题!" {
		t.Fail()
	}
	// 行政人员
	communityExecutive := &CommunityExecutive{}
	// 获取一个给行政人员面试的面试官
	communityExecutiveManager := NewHiringManager(communityExecutive)
	// 询问行政人员问题
	if str := communityExecutiveManager.TakeInterview(); str != "问关于行政的问题!" {
		t.Fail()
	}
}
