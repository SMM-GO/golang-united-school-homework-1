package solution

import emoji "github.com/kyokomi/emoji"

func GetMessage() string {
	s := emoji.Sprint("Hello :world_map:!")
	return s
}
