package gdconf

import (
	"strings"

	"github.com/gucooing/hkrpg-go/pkg/constant"
	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/pkg/text"
)

type GoppAbility struct {
	GoppMazeSkill map[uint32]map[int]*GoppMazeSkill // avatar/index
}

type GoppMazeSkill struct {
	AvatarId           uint32
	Index              int
	TriggerBattle      bool
	ActionsList        []*MazeSkillAction
	AdventureModifiers map[string]struct{}
}

type MazeSkillAction struct {
	Type     constant.SkillAction
	Id       uint32
	Duration int
}

func (g *GameDataConfig) goppAbility() {
	g.GoppAbility = &GoppAbility{
		GoppMazeSkill: g.newGoppMazeSkill(),
	}

	logger.Info(text.GetText(17), len(g.GoppAbility.GoppMazeSkill), "MazeSkill")
}

func (g *GameDataConfig) newGoppMazeSkill() map[uint32]map[int]*GoppMazeSkill {
	list := make(map[uint32]map[int]*GoppMazeSkill)
	abilityConf := GetAvatarAbilityMap()
	for avatarId, abilityList := range abilityConf {
		if list[avatarId] == nil {
			list[avatarId] = make(map[int]*GoppMazeSkill)
		}
		for _, ability := range abilityList.AbilityList {
			var skill *GoppMazeSkill
			if strings.Contains(ability.Name, "NormalAtk") {
				if list[avatarId][1] == nil {
					list[avatarId][1] = &GoppMazeSkill{
						AvatarId:           avatarId,
						Index:              1,
						TriggerBattle:      true,
						ActionsList:        make([]*MazeSkillAction, 0),
						AdventureModifiers: make(map[string]struct{}),
					}
				}
				skill = list[avatarId][1]
			} else if strings.Contains(ability.Name, "MazeSkill") {
				if list[avatarId][2] == nil {
					list[avatarId][2] = &GoppMazeSkill{
						AvatarId:           avatarId,
						Index:              2,
						TriggerBattle:      true,
						ActionsList:        make([]*MazeSkillAction, 0),
						AdventureModifiers: make(map[string]struct{}),
					}
				}
				skill = list[avatarId][2]
				summData := GetSummonUnitData(avatarId*10 + 1)
				if summData != nil && !summData.IsClient {
					skill.ActionsList = append(skill.ActionsList, &MazeSkillAction{
						Type:     constant.SummonUnit,
						Id:       summData.ID,
						Duration: 20,
					})
				}
			} else {
				logger.Warn(text.GetText(22), ability.Name)
				continue
			}

			for _, task := range ability.OnStart {
				actionList := parseTask(skill, task)
				skill.ActionsList = append(skill.ActionsList, actionList...)
			}
		}
	}
	return list
}

func parseTask(skill *GoppMazeSkill, task *TaskInfo) []*MazeSkillAction {
	actionList := make([]*MazeSkillAction, 0)
	if strings.Contains(task.Type, "AddMazeBuff") {
		actionList = append(actionList, &MazeSkillAction{
			Type: constant.AddMazeBuff,
			Id:   task.ID,
		})
	} else if strings.Contains(task.Type, "RemoveMazeBuff") {

	} else if strings.Contains(task.Type, "AdventureModifyTeamPlayerHP") {
		actionList = append(actionList, &MazeSkillAction{
			Type: constant.AddTeamPlayerHP,
			Id:   1500,
		})
	} else if strings.Contains(task.Type, "AdventureModifyTeamPlayerSP") {
		actionList = append(actionList, &MazeSkillAction{
			Type: constant.AddTeamPlayerSp,
			Id:   5000,
		})
	} else if strings.Contains(task.Type, "CreateSummonUnit") && !task.IsClientOnly {
		skill.TriggerBattle = false
	} else if strings.Contains(task.Type, "AddAdventureModifier") {
		skill.AdventureModifiers[task.ModifierName] = struct{}{}
	} else if strings.Contains(task.Type, "AdventureSetAttackTargetMonsterDie") {
		actionList = append(actionList, &MazeSkillAction{
			Type: constant.SetMonsterDie,
		})
	} else if task.SuccessTaskList != nil {
		for _, t := range task.SuccessTaskList {
			actionList = append(actionList, parseTask(skill, t)...)
		}
	} else if strings.Contains(task.Type, "AdventureTriggerAttack") {
		if skill.Index == 2 {
			skill.TriggerBattle = task.TriggerBattle
		}
		if task.OnAttack != nil {
			for _, t := range task.OnAttack {
				actionList = append(actionList, parseTask(skill, t)...)
			}
		}
		if task.OnBattle != nil {
			for _, t := range task.OnBattle {
				actionList = append(actionList, parseTask(skill, t)...)
			}
		}
	} else if strings.Contains(task.Type, "AdventureFireProjectile") {
		if task.OnProjectileHit != nil {
			for _, t := range task.OnProjectileHit {
				actionList = append(actionList, parseTask(skill, t)...)
			}
		}
		if task.OnProjectileLifetimeFinish != nil {
			for _, t := range task.OnProjectileLifetimeFinish {
				actionList = append(actionList, parseTask(skill, t)...)
			}
		}
	} else {
		// logger.Warn("task Type:%s", task.Type)
	}

	return actionList
}

func GetGoppMazeSkill(avatar uint32, index int) *GoppMazeSkill {
	if getConf().GoppAbility.GoppMazeSkill[avatar] == nil {
		return nil
	}
	return getConf().GoppAbility.GoppMazeSkill[avatar][index]
}
