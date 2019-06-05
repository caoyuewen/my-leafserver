package internal

import (
	"fmt"
	"github.com/name5566/leaf/gate"
	"sync"
	"time"
)

const (
	autoDeleteAgentTime = time.Second * 5 // 自动删除认证中的agent

	autoDeleteAuthAgentTime = time.Second * 5 // 自动删除认证登陆的agent
)

type UserAgent struct {
	Agent  []*gate.Agent // 用slice 为了存储web client phone 多个客户端同时在线
	CNonce string        // 客户端的nonce
	SNonce string        // 服务器的nonce
	UserId string        // 用户的绑定ID 唯一标识
}

// 存放没有经过认证的用户
var globalAgents map[string]*gate.Agent

// 存放已经经过认证切已经登录的用户
var globalAuthAgents map[string]*UserAgent

var rwMutex sync.RWMutex

func init() {
	globalAgents = make(map[string]*gate.Agent, 10)
	globalAuthAgents = make(map[string]*UserAgent, 10)
}

// 存放用户没有经过认证的用户
// globalAgents
// @param key   : cNonce + sNonce
// @param agent
func saveAgent(key string, agent *gate.Agent) {
	rwMutex.Lock()

	globalAgents[key] = agent
	fmt.Println("当前保存的连接:", len(globalAgents))

	rwMutex.Unlock()

	// 保存之后五秒钟完成登陆并移除保存客户端
	go autoDeleteAgent(key)
}

func autoDeleteAgent(key string) {
	time.Sleep(autoDeleteAgentTime)
	authAgent := getAuthAgent(key)
	if authAgent == nil {
		agent := getAgent(key)
		(*agent).Close()
	}
	removeAgent(key)
}

// 获取agent
// globalAgents
// @param key   : cNonce + sNonce
func getAgent(key string) *gate.Agent {
	rwMutex.Lock()
	defer rwMutex.Unlock()
	return globalAgents[key]
}

// 获取agent
// globalAgents
// @param key   : cNonce + sNonce
func removeAgent(key string) {
	rwMutex.Lock()
	defer rwMutex.Unlock()
	delete(globalAgents, key)
}

/*---------------------------------------------*/

// 存放认证登录的用户
// globalAuthAgents
// @param key   : userId
// @param agent
func saveAuthAgent(key string, userAgent *UserAgent) {
	rwMutex.Lock()
	defer rwMutex.Unlock()
	globalAuthAgents[key] = userAgent
}

// 获取agent
// globalAuthAgents
// @param key   : userId
func getAuthAgent(key string) *UserAgent {
	rwMutex.Lock()
	defer rwMutex.Unlock()
	return globalAuthAgents[key]
}

// 获取agent
// globalAuthAgents
// @param key   : userId
func removeAuthAgent(key string) {
	rwMutex.Lock()
	defer rwMutex.Unlock()
	delete(globalAuthAgents, key)
}
