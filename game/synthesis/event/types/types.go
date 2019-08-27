package types

import (
	synthesistypes "fgame/fgame/game/synthesis/types"
)

type SynthesisFinishEventType string

const (
	EventTypeSynthesisFinish SynthesisFinishEventType = "SynthesisFinish"
)

type SynthesisFinishEventData struct {
	typ synthesistypes.SynthesisType //
	num int32                        //合成次数
}

func (s *SynthesisFinishEventData) GetType() synthesistypes.SynthesisType {
	return s.typ
}

func (s *SynthesisFinishEventData) GetNum() int32 {
	return s.num
}

func CreateSynthesisFinishEventData(typ synthesistypes.SynthesisType, num int32) *SynthesisFinishEventData {
	return &SynthesisFinishEventData{

		typ: typ,
		num: num,
	}
}
