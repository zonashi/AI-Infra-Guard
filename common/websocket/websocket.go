package websocket

import (
	"github.com/Tencent/AI-Infra-Guard/internal/gologger"
	"github.com/Tencent/AI-Infra-Guard/internal/options"
	"net/http"
)

func RunWebsocket(options *options.Options) {
	// 创建WebSocket服务器
	wsServer := NewWSServer(options)
	// 设置WebSocket路由
	http.HandleFunc("/ws", wsServer.HandleWS)
	http.HandleFunc("/show", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("show"))
	})
	// 启动HTTP服务器
	gologger.Infof("Starting WebSocket server on %s\n", options.WebSocketAddr)
	if err := http.ListenAndServe(options.WebSocketAddr, nil); err != nil {
		gologger.Fatalf("Could not start WebSocket server: %s\n", err)
	}
}
