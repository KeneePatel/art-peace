package quests

import "github.com/keep-starknet-strange/art-peace/backend/core"

var QuestChecks = map[int]func(*Quest, string) (int, int){
	AuthorityQuestType:  CheckAuthorityStatus,
	HodlQuestType:       CheckHodlStatus,
	NFTMintQuestType:    CheckNftStatus,
	PixelQuestType:      CheckPixelStatus,
	RainbowQuestType:    CheckRainbowStatus,
	TemplateQuestType:   CheckTemplateStatus,
	UnruggableQuestType: CheckUnruggableStatus,
	VoteQuestType:       CheckVoteStatus,
	FactionQuestType:    CheckFactionStatus,
	UsernameQuestType:   CheckUsernameStatus,
}

func (q *Quest) CheckStatus(user string) (progress int, needed int) {
	check := QuestChecks[q.Type]
	if check == nil {
		return 0, 0
	}
	return check(q, user)
}

func CheckAuthorityStatus(q *Quest, user string) (progress int, needed int) {
	// TODO: Implement this
	return 0, 1
}

func CheckHodlStatus(q *Quest, user string) (progress int, needed int) {
	// TODO: Implement this
	return 0, 1
}

func CheckNftStatus(q *Quest, user string) (progress int, needed int) {
	// gets the count of NFTs minted by `user`
	nfts_minted_by_user, err := core.PostgresQueryOne[int]("SELECT COUNT(*) FROM NFTs WHERE minter = $1", user)
	// if error is encountered, return 0, 1
	if err != nil {
		return 0, 1
	}
	// checks if count is greater than 0, if yes, returns 1, 1, else 0, 1
	if *nfts_minted_by_user > 0 {
		return 1, 1
	}
	return 0, 1
}

func CheckPixelStatus(q *Quest, user string) (progress int, needed int) {
	pixelQuestInputs := NewPixelQuestInputs(q.InputData)
	if pixelQuestInputs.IsDaily {
		day := pixelQuestInputs.ClaimDay
		if pixelQuestInputs.IsColor {
			// TODO: Use coalesce
			count, err := core.PostgresQueryOne[int]("SELECT COUNT(*) FROM Pixels WHERE address = $1 AND color = $2 AND day = $3", user, pixelQuestInputs.Color, day)
			if err != nil {
				return 0, 1
			}
			return *count, int(pixelQuestInputs.PixelsNeeded)
		} else {
			count, err := core.PostgresQueryOne[int]("SELECT COUNT(*) FROM Pixels WHERE address = $1 AND day = $2", user, day)
			if err != nil {
				return 0, 1
			}
			return *count, int(pixelQuestInputs.PixelsNeeded)
		}
	} else {
		if pixelQuestInputs.IsColor {
			count, err := core.PostgresQueryOne[int]("SELECT COUNT(*) FROM Pixels WHERE address = $1 AND color = $2", user, pixelQuestInputs.Color)
			if err != nil {
				return 0, 1
			}
			return *count, int(pixelQuestInputs.PixelsNeeded)
		} else {
			count, err := core.PostgresQueryOne[int]("SELECT COUNT(*) FROM Pixels WHERE address = $1", user)
			if err != nil {
				return 0, 1
			}
			return *count, int(pixelQuestInputs.PixelsNeeded)
		}
	}
}

func CheckVoteStatus(q *Quest, user string) (progress int, needed int) {
	voteQuestInputs := NewVoteQuestInputs(q.InputData)

	count, err := core.PostgresQueryOne[int]("SELECT COUNT(*) FROM ColorVotes WHERE user_address = $1 AND day_index = $2", user, voteQuestInputs.DayIndex)
	if err != nil {
		return 0, 1
	}

	return *count, 1
}

func CheckFactionStatus(q *Quest, user string) (progress int, needed int) {
	count, err := core.PostgresQueryOne[int]("SELECT COUNT(*) FROM FactionMembersInfo WHERE user_address = $1", user)
	if err != nil {
		return 0, 1
	}

	return *count, 1
}

func CheckRainbowStatus(q *Quest, user string) (progress int, needed int) {
	// TODO: Implement this
	return 0, 1
}

func CheckTemplateStatus(q *Quest, user string) (progress int, needed int) {
	// TODO: Implement this
	return 0, 1
}

func CheckUnruggableStatus(q *Quest, user string) (progress int, needed int) {
	// TODO: Implement this
	return 0, 1
}

func CheckUsernameStatus(q *Quest, user string) (progress int, needed int) {

	count, err := core.PostgresQueryOne[int]("SELECT COUNT (*) FROM Users where address = $1", user)

	if err != nil {
		return 0, 1
	} else {
		return *count, 1
	}
}
