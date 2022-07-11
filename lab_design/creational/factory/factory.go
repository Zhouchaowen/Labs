// 工厂方法模式（Factory Method）
// 现实世界的例子:考虑招聘经理的情况。一个人不可能对每个职位进行面试。根据职位空缺，她必须决定并将面试步骤委托给不同的人。
// 简单来说:它提供了一种将实例化逻辑委托给子类的方法。
// 维基百科说:在基于类的编程中，工厂方法模式是一种创建模式，它使用工厂方法来处理创建对象的问题，而无需指定将要创建的对象的确切类
// 过调用工厂方法来创建对象来完成的 - 在接口中指定并由子类实现，或者在基类中实现并可选地由派生类覆盖 - 而不是通过调用构造函数。
package main

// 场景：以招聘经理为例。首先，我们有一个访谈者界面和回答问题方法

// Interviewer 面试接口 实现回答问题
type interviewer interface {
	AskQuestions() string
}

// Developer 开发者
type Developer struct {
}

// AskQuestions 开发者面试需要回答的问题
func (developer *Developer) AskQuestions() string {
	return "问关于开发的问题!"
}

// CommunityExecutive CommunityExecutive(行政人员)
type CommunityExecutive struct {
}

// AskQuestions CommunityExecutive(行政人员)面试需要回答的问题
func (communityExecutive *CommunityExecutive) AskQuestions() string {
	return "问关于行政的问题!"
}

// HiringManager 现在让我们创造我们的 HiringManager(招聘经理)
type HiringManager struct {
	Interviewer interviewer
}

// TakeInterview 接受面试
func (hiringManager *HiringManager) TakeInterview() string {
	return hiringManager.Interviewer.AskQuestions()
}

// NewHiringManager 去找招聘者面试()
func NewHiringManager(iv interviewer) *HiringManager {
	return &HiringManager{Interviewer: iv}
}
