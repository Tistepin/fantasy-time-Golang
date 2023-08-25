package model

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"gopkg.in/fatih/set.v0"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"sync"
	"time"
)

/**
* User:徐国纪
* Create_Time:下午 01:56
**/

// Message 消息
type Message struct {
	gorm.Model
	UserId     int64  //发送者
	TargetId   int64  //接受者
	Type       int    //发送类型  1私聊  2群聊  3心跳
	Media      int    //消息类型  1文字  2表情包 3语音 4图片 /表情包
	Content    string //消息内容
	CreateTime uint64 //创建时间
	ReadTime   uint64 //读取时间
	Pic        string
	Url        string
	Desc       string
	Amount     int //其他数字统计
}

// Node 用户连接
type Node struct {
	Conn      *websocket.Conn //连接
	DataQueue chan []byte     //消息
	GroupSets set.Interface   //好友 / 群
}

// ClientMap 映射关系 客户端
var ClientMap map[int64]*Node = make(map[int64]*Node, 0)

// 读写锁
var rwLocker sync.RWMutex

// Chat 需要 ：发送者ID ，接受者ID ，消息类型，发送的内容，发送类型
func Chat(writer http.ResponseWriter, request *http.Request) {
	// 1.获取参数ID
	token := request.URL.Query().Get("FantasyTimetoken")
	user := &FtUser{}
	Id := user.GetId(token)
	if Id == 0 {
		return
	}
	id, err := strconv.ParseInt(strconv.Itoa(int(Id)), 10, 64)
	conn, err := (&websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}).Upgrade(writer, request, nil)
	if err != nil {
		now := time.Now()
		fmt.Println(now.Format("2006-01-02 15:04:05")+"[Debug] err", err.Error())
		return
	}
	// 创建专属连接
	node := &Node{
		Conn:      conn,
		DataQueue: make(chan []byte, 50),
		GroupSets: set.New(set.ThreadSafe),
	}

	// 用户关系绑定
	// 加锁存入
	rwLocker.Lock() // 加锁
	ClientMap[id] = node
	rwLocker.Unlock() // 解锁
	// 5.完成发送逻辑
	go sendProc(node)
	// 6.完成接收逻辑
	go recvProc(node)
	sendMsg(id, []byte("欢迎进入聊天系统"))
}

// sendProc 发送消息
func sendProc(node *Node) {
	// for 循环 读取连接的管道内部消息 然后发送
	for {
		select {
		// 读取管道消息
		case msg := <-node.DataQueue:
			// 发送消息
			err := node.Conn.WriteMessage(websocket.TextMessage, msg)
			if err != nil {
				now := time.Now()
				fmt.Println(now.Format("2006-01-02 15:04:05")+"[Debug] err", err.Error())
				return
			}
		}
	}
}

// 读取消息
func recvProc(node *Node) {
	_, data, err := node.Conn.ReadMessage()
	if err != nil {
		now := time.Now()
		fmt.Println(now.Format("2006-01-02 15:04:05")+"[Debug] err", err.Error())
		return
	}
	msg := Message{}
	err = json.Unmarshal(data, &msg)
	if err != nil {
		now := time.Now()
		fmt.Println(now.Format("2006-01-02 15:04:05")+"[Debug] err", err.Error())
		return
	}
	//心跳检测 msg.Media == -1 || msg.Type == 3
	if msg.Type == 3 {
		//currentTime := uint64(time.Now().Unix())
		//node.Heartbeat(currentTime)
	} else {
		// 接收消息后发送给个人
		dispatch(data)
	}
}

// 接收消息后发送给个人
func dispatch(data []byte) {
	msg := Message{}
	err := json.Unmarshal(data, &msg)
	if err != nil {
		now := time.Now()
		fmt.Println(now.Format("2006-01-02 15:04:05")+"[Debug] err", err.Error())
		return
	}
	// 根据不同发送类型选择
	switch msg.Type {
	case 1: //私信
		fmt.Println("dispatch  data :", string(data))
		sendMsg(msg.TargetId, data)
	case 2: //群发
		//sendGroupMsg(msg.TargetId, data) //发送的群ID ，消息内容
		// case 4: // 心跳
		// 	node.Heartbeat()
		//case 4:
		//
	}
}

func sendMsg(userId int64, msg []byte) {
	rwLocker.RLock() // 加锁
	node, ok := ClientMap[userId]
	rwLocker.RUnlock() // 解锁
	if ok {
		node.DataQueue <- msg
	}
}
