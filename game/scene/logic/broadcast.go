package logic

import (
	"fgame/fgame/game/scene/scene"

	"github.com/golang/protobuf/proto"
)

func BroadcastNeighbor(so scene.SceneObject, msg proto.Message) {
	for _, nei := range so.GetNeighbors() {
		switch ts := nei.(type) {
		case scene.Player:
			//优化内挂流量
			if !ts.IsGuaJiPlayer() {
				ts.SendMsg(msg)
			}
		}
	}
}

func BroadcastNeighborIncludeSelf(so scene.SceneObject, msg proto.Message) {
	BroadcastNeighbor(so, msg)
	switch ts := so.(type) {
	case scene.Player:
		//优化内挂流量
		if !ts.IsGuaJiPlayer() {
			ts.SendMsg(msg)
		}
	}

}
