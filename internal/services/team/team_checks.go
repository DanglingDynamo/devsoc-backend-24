package services

import (
	"github.com/google/uuid"

	"github.com/CodeChefVIT/devsoc-backend-24/internal/database"
)

func CheckTeamCode(code string) bool {
	query := `SELECT COUNT(*) FROM teams WHERE code = $1`
	var temp int
	err := database.DB.QueryRow(query, code).Scan(&temp)
	if err != nil || temp == 0 {
		return false
	}
	return true
}

func CheckTeamSize(team_id uuid.UUID) bool {
	query := `SELECT COUNT(*) WHERE team_id = $1`

	var no_members int

	err := database.DB.QueryRow(query, team_id).Scan(&no_members)
	if err != nil || no_members >= 5 {
		return false
	}

	return true
}

func CheckUserInTeam(id uuid.UUID, teamid uuid.UUID) bool {
	query := `SELECT COUNT(*) FROM users WHERE team_id = $1 AND id = $2`
	var check int
	err := database.DB.QueryRow(query, teamid, id).Scan(&check)
	if err != nil || check == 0 {
		return false
	}
	return true
}
