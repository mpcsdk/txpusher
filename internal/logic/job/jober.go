package job

import (
	"time"
)

type Jober struct {
	userId string
	addr   string
	//
	chainStat map[uint64]uint64
	////
	msg    []*item
	seqId  int
	oldmsg []*item
	///
	expire time.Time
	//
}

func (s *Jober) IsExpire() bool {
	return time.Now().After(s.expire)
}

func (s *Jober) PushMsg(msg []*item) {
	s.msg = append(s.msg, msg...)
}
func (s *Jober) Ack(seqId int) {
	if s.seqId == seqId {
		s.oldmsg = nil
	}
}

// /
func (s *Jober) RecentMsg() (int, []*item) {
	if s.oldmsg != nil {
		s.oldmsg = append(s.oldmsg, s.msg...)
		s.msg = s.oldmsg
		s.oldmsg = nil
	} else {
		if s.seqId > 1000000 {
			s.seqId = 0
		}
		s.seqId += 1
		s.oldmsg = s.msg
	}
	///
	return s.seqId, s.msg
}
func (s *Jober) UpState(seqId int) {
	s.seqId = seqId
	s.oldmsg = nil
	///
}

func newJober(expire time.Duration) *Jober {
	return &Jober{
		userId: "",
		////
		msg:    []*item{},
		oldmsg: []*item{},
		seqId:  0,
		///
		expire: time.Now().Add(expire),
	}
}
