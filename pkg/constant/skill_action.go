package constant

type SkillAction int

const (
	AddMazeBuff     SkillAction = 1
	AddTeamPlayerHP SkillAction = 2
	AddTeamPlayerSp SkillAction = 3
	SetMonsterDie   SkillAction = 4
	SummonUnit      SkillAction = 5
)
