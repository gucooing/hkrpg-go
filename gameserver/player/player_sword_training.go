package player

import (
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	"github.com/gucooing/hkrpg-go/protocol/proto"
	pb "google.golang.org/protobuf/proto"
)

func (g *GamePlayer) GetSwordTrainingDataCsReq(payloadMsg pb.Message) {
	rsp := &proto.GetSwordTrainingDataScRsp{
		MOOHNCHOPNH: nil,
		SwordTrainingGameInfo: &proto.SwordTrainingGameInfo{
			SkillInfo: &proto.SwordTrainingSkillInfo{
				SkillPower:      1,
				PAOJAAONCBF:     0,
				TargetSkill:     102,
				UnlockSkillList: make([]uint32, 0),
			},
			MCKIEJODKGE: &proto.ECIGNEGEAIH{
				GameStoryLineId: 1,
				JAFLOLLFNDO:     0,
				MOOHNCHOPNH:     nil,
				HCMJKFEAADP:     nil,
			},
			PowerInfo: &proto.SwordTrainingPowerInfo{
				CurMood: 100,
				SkillPowerList: []*proto.SwordTrainingSkillPowerInfo{
					{
						Value:       100,
						PCLPEHLNDAF: 0,
						Status:      proto.SwordTrainingStatus_SWORD_TRAINING_STATUS_TYPE_AGILITY,
					},
					{
						Value:       100,
						PCLPEHLNDAF: 0,
						Status:      proto.SwordTrainingStatus_SWORD_TRAINING_STATUS_TYPE_POWER,
					},
					{
						Value:       100,
						PCLPEHLNDAF: 0,
						Status:      proto.SwordTrainingStatus_SWORD_TRAINING_STATUS_TYPE_TOUGHNESS,
					},
					{
						Value:       100,
						PCLPEHLNDAF: 0,
						Status:      proto.SwordTrainingStatus_SWORD_TRAINING_STATUS_TYPE_PERCEPTION,
					},
				},
				MaxMood: 100,
			},
			IIHDBEFHEOC: nil,
			INJNGCMDBGL: &proto.KKIPILADIGB{
				MIMOGGLAFID: 101,
				MMAKPNLBHII: nil,
				NEMGFNPGEDK: &proto.KKIPILADIGB_PAOKBIDANEJ{
					PAOKBIDANEJ: &proto.MHBCFFNEBEL{
						Type: proto.SwordTrainingDailyPhaseType_SWORD_TRAINING_DAILY_PHASE_TYPE_MORNING,
					},
				},
			},
			PendingAction: &proto.KKDCJKFPLMF{
				AOPBDFIBFPH: nil,
				ABHNAFELEME: &proto.KKDCJKFPLMF_PCPAFLJBGDD{},
			},
			SkillPower:       1,
			OptionResultInfo: nil,
			DPHNBKLJEHM:      nil,
			SwordTrainingActionInfo: &proto.SwordTrainingActionInfo{
				ActionInfoList: []*proto.SwordActionInfo{
					{
						QueuePosition: 1,
						Level:         1,
					},
					{
						QueuePosition: 2,
						Level:         1,
					},
					{
						QueuePosition: 3,
						Level:         1,
					},
					{
						QueuePosition: 4,
						Level:         1,
					},
					{
						QueuePosition: 5,
						Level:         1,
					},
					{
						QueuePosition: 6,
						Level:         1,
					},
					{
						QueuePosition: 7,
						Level:         1,
					},
				},
			},
		},
		AILAGENLDGI: 0,
		DIGIDEKCKPF: nil,
		Retcode:     0,
		IMJLPHEJMBB: nil,
		DPHNBKLJEHM: nil,
		BNNIMFGFOAH: nil,
		GBDANLKCLMG: nil,
	}
	g.Send(cmd.GetSwordTrainingDataScRsp, rsp)
}

func (g *GamePlayer) SwordTrainingStartGameCsReq(payloadMsg pb.Message) {
	// req := payloadMsg.(*proto.SwordTrainingStartGameCsReq)
	g.SwordTrainingGameSyncChangeScNotify()
	rsp := &proto.SwordTrainingStartGameScRsp{
		SwordTrainingGameInfo: &proto.SwordTrainingGameInfo{
			SkillInfo: &proto.SwordTrainingSkillInfo{
				SkillPower:      250,
				PAOJAAONCBF:     0,
				TargetSkill:     102,
				UnlockSkillList: make([]uint32, 0),
			},
			MCKIEJODKGE: &proto.ECIGNEGEAIH{
				GameStoryLineId: 1,
				JAFLOLLFNDO:     0,
				MOOHNCHOPNH:     nil,
				HCMJKFEAADP:     nil,
			},
			PowerInfo: &proto.SwordTrainingPowerInfo{
				CurMood: 100,
				SkillPowerList: []*proto.SwordTrainingSkillPowerInfo{
					{
						Value:       100,
						PCLPEHLNDAF: 0,
						Status:      proto.SwordTrainingStatus_SWORD_TRAINING_STATUS_TYPE_AGILITY,
					},
					{
						Value:       100,
						PCLPEHLNDAF: 0,
						Status:      proto.SwordTrainingStatus_SWORD_TRAINING_STATUS_TYPE_POWER,
					},
					{
						Value:       100,
						PCLPEHLNDAF: 0,
						Status:      proto.SwordTrainingStatus_SWORD_TRAINING_STATUS_TYPE_TOUGHNESS,
					},
					{
						Value:       100,
						PCLPEHLNDAF: 0,
						Status:      proto.SwordTrainingStatus_SWORD_TRAINING_STATUS_TYPE_PERCEPTION,
					},
				},
				MaxMood: 100,
			},
			IIHDBEFHEOC: nil,
			INJNGCMDBGL: &proto.KKIPILADIGB{
				MIMOGGLAFID: 101,
				MMAKPNLBHII: nil,
				NEMGFNPGEDK: &proto.KKIPILADIGB_PAOKBIDANEJ{
					PAOKBIDANEJ: &proto.MHBCFFNEBEL{
						Type: proto.SwordTrainingDailyPhaseType_SWORD_TRAINING_DAILY_PHASE_TYPE_MORNING,
					},
				},
			},
			PendingAction:    nil,
			SkillPower:       250,
			OptionResultInfo: nil,
			DPHNBKLJEHM:      nil,
			SwordTrainingActionInfo: &proto.SwordTrainingActionInfo{
				ActionInfoList: []*proto.SwordActionInfo{
					{
						QueuePosition: 1,
						Level:         1,
					},
					{
						QueuePosition: 2,
						Level:         1,
					},
					{
						QueuePosition: 3,
						Level:         1,
					},
					{
						QueuePosition: 4,
						Level:         1,
					},
					{
						QueuePosition: 5,
						Level:         1,
					},
					{
						QueuePosition: 6,
						Level:         1,
					},
					{
						QueuePosition: 7,
						Level:         1,
					},
				},
			},
		},
	}
	g.Send(cmd.SwordTrainingStartGameScRsp, rsp)
}

func (g *GamePlayer) SwordTrainingGameSyncChangeScNotify() {
	notify := &proto.SwordTrainingGameSyncChangeScNotify{
		SwordTrainingChange: []*proto.SwordTrainingChange{
			{
				RogueAction: &proto.SwordTrainingAction{
					INMOIILMOJN: nil,
					KABMHIGOCHM: nil,
					Action: &proto.SwordTrainingAction_LDMGDOGJKNC{
						LDMGDOGJKNC: &proto.HNHNFFFGFJC{
							SkillPower: 1,
							PCMAAKHAEBC: &proto.KKIPILADIGB{
								MIMOGGLAFID: 101,
								MMAKPNLBHII: nil,
								NEMGFNPGEDK: &proto.KKIPILADIGB_PAOKBIDANEJ{
									PAOKBIDANEJ: &proto.MHBCFFNEBEL{
										Type: proto.SwordTrainingDailyPhaseType_SWORD_TRAINING_DAILY_PHASE_TYPE_MORNING,
									},
								},
							},
						},
					},
				},
			},
		},
	}
	g.Send(cmd.SwordTrainingGameSyncChangeScNotify, notify)
}
