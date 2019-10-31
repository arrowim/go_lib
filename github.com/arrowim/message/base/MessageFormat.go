package base

type MessageFormatType int16

func (self MessageFormatType) GetValue() int16 {
	return int16(self)
}

const (
	//	ERROR_FORMAT  = iota
	ERROR_FORMAT  MessageFormatType = 0
	TEXT          MessageFormatType = 1
	IMAGE         MessageFormatType = 2
	AUDIO         MessageFormatType = 3
	VEDIO         MessageFormatType = 4
	NOTIFICATION  MessageFormatType = 5
	FILE          MessageFormatType = 6
	JSON          MessageFormatType = 7
	LOCATIONPOINT MessageFormatType = 8
	AT_SOMEBODY   MessageFormatType = 9
	VIDEO_CALL    MessageFormatType = 11;
	VIDEO_CALL2   MessageFormatType = 12;
)

func (self MessageFormatType) GetPushContent(content string) string {
	data := ""
	switch self {
	case TEXT:
		data += content

	case IMAGE:
		data += "[图片]"

	case AUDIO:
		data += "[语音]"

	case VEDIO:
		data += "[视频]"

	case NOTIFICATION:
		data += "[通知]"

	case FILE:
		data += "[文件]"

		//	case message.JSON:
		//TODO 3消息类型为JSON，解析出content中的json串中的指定key值
		//content += "[JSON]"

	case LOCATIONPOINT:
		data += "[地址]"
		//
		//case AT_SOMEBODY:
		//	atsb := message.AtSomeBody{}
		//	json.Unmarshal([]byte(msg_.MessageContent), &atsb)
		//	//		fmt.Println(atsb.Content)
		//	//		fmt.Println(atsb.BeAtIds)
		//	flag := false
		//	for _, v := range atsb.BeAtIds {
		//		if v == userId {
		//			flag = true
		//		}
		//	}
		//	if flag {
		//		content = "[有人@你]" + content + atsb.Content
		//	} else {
		//		content += atsb.Content
		//	}

	default:
		break
	}

	return data
}
