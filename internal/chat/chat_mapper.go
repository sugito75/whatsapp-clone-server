package chat

func ChatModelToDTO(chats []ChatMember) []GetChatsDTO {
	result := []GetChatsDTO{}
	for _, chat := range chats {
		if chat.LastMessage.ID == 0 {
			continue
		}

		title := chat.Chat.Title
		icon := chat.Chat.Icon
		member := chat.Chat.Members[0].User

		if chat.Chat.ChatType == ChatTypePrivate {
			title = &member.DisplayName
			icon = member.ProfilePicture
		}

		c := GetChatsDTO{
			ID:       chat.ChatID,
			ChatType: chat.Chat.ChatType,
			Title:    title,
			Icon:     icon,
			LastMessage: LastMessageDTO{
				Text:     chat.LastMessage.Content,
				Status:   chat.LastMessage.Status.Status,
				SentAt:   chat.LastMessage.SentAt,
				SenderID: chat.LastMessage.SenderID,
			},
		}

		result = append(result, c)
	}

	return result
}
