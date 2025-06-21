package cli

import (
	"math/rand"
	"os"
	"strings"
)

var banners = []string{
	`
  ______            ______                     ______    __                       
 /      \          /      \                   /      \  |  \                      
|  ▓▓▓▓▓▓\ ______ |  ▓▓▓▓▓▓\ ______   ______ |  ▓▓▓▓▓▓\_| ▓▓_    ______   ______  
| ▓▓ __\▓▓/      \| ▓▓   \▓▓/      \ |      \| ▓▓_  \▓▓   ▓▓ \  /      \ /      \ 
| ▓▓|    \  ▓▓▓▓▓▓\ ▓▓     |  ▓▓▓▓▓▓\ \▓▓▓▓▓▓\ ▓▓ \    \▓▓▓▓▓▓ |  ▓▓▓▓▓▓\  ▓▓▓▓▓▓\
| ▓▓ \▓▓▓▓ ▓▓  | ▓▓ ▓▓   __| ▓▓   \▓▓/      ▓▓ ▓▓▓▓     | ▓▓ __| ▓▓    ▓▓ ▓▓   \▓▓
| ▓▓__| ▓▓ ▓▓__/ ▓▓ ▓▓__/  \ ▓▓     |  ▓▓▓▓▓▓▓ ▓▓       | ▓▓|  \ ▓▓▓▓▓▓▓▓ ▓▓      
 \▓▓    ▓▓\▓▓    ▓▓\▓▓    ▓▓ ▓▓      \▓▓    ▓▓ ▓▓        \▓▓  ▓▓\▓▓     \ ▓▓      
  \▓▓▓▓▓▓  \▓▓▓▓▓▓  \▓▓▓▓▓▓ \▓▓       \▓▓▓▓▓▓▓\▓▓         \▓▓▓▓  \▓▓▓▓▓▓▓\▓▓      `,

	// Here you can add more banners, they will be randomly selected at runtime on each execution
	// ` SOME OTHER BANNER `,
}

func GetDescriptions(descriptionArg []string, _ bool) map[string]string {
	var description, banner string
	if descriptionArg != nil {
		if strings.Contains(strings.Join(os.Args[0:], ""), "-h") {
			description = descriptionArg[0]
		} else {
			description = descriptionArg[1]
		}
	} else {
		description = ""
	}
	bannerRandLen := len(banners)
	bannerRandIndex := rand.Intn(bannerRandLen)
	banner = banners[bannerRandIndex]
	return map[string]string{"banner": banner, "description": description}
}
