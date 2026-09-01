package model

type CharacterSkills struct {
	Skills        []CharacterSkill `json:"skills"`
	TotalSP       int64            `json:"total_sp"`
	UnallocatedSP int64            `json:"unallocated_sp"`
}

type CharacterSkill struct {
	SkillID            int64 `json:"skill_id"`
	ActiveSkillLevel   int   `json:"active_skill_level"`
	TrainedSkillLevel  int   `json:"trained_skill_level"`
	SkillpointsInSkill int64 `json:"skillpoints_in_skill"`
}
