package util

// check if the person belongs to the given conversation
func UserBelongsToConversation(conv Conversation, user User) bool {
	for _, participantID := range conv.Participants {
		if participantID == user.Id {
			return true
		}
	}
	return false
}
