package infra

import "fmt"

const (
	banner0 = `
  _____             _    
 |  __ \           | |   
 | |__) | ___  ___ | | __
 |  _  / / _ \/ __|| |/ /
 | | \ \|  __/\__ \|   < 
 |_|  \_\\___||___/|_|\_\
`
	banner1 = `


██████╗ ███████╗███████╗██╗  ██╗
██╔══██╗██╔════╝██╔════╝██║ ██╔╝
██████╔╝█████╗  ███████╗█████╔╝ 
██╔══██╗██╔══╝  ╚════██║██╔═██╗ 
██║  ██║███████╗███████║██║  ██╗
╚═╝  ╚═╝╚══════╝╚══════╝╚═╝  ╚═╝

`
	banner2 = `
 ▄▄▄▄▄▄▄▄▄▄▄  ▄▄▄▄▄▄▄▄▄▄▄  ▄▄▄▄▄▄▄▄▄▄▄  ▄    ▄ 
▐░░░░░░░░░░░▌▐░░░░░░░░░░░▌▐░░░░░░░░░░░▌▐░▌  ▐░▌
▐░█▀▀▀▀▀▀▀█░▌▐░█▀▀▀▀▀▀▀▀▀ ▐░█▀▀▀▀▀▀▀▀▀ ▐░▌ ▐░▌ 
▐░▌       ▐░▌▐░▌          ▐░▌          ▐░▌▐░▌  
▐░█▄▄▄▄▄▄▄█░▌▐░█▄▄▄▄▄▄▄▄▄ ▐░█▄▄▄▄▄▄▄▄▄ ▐░▌░▌   
▐░░░░░░░░░░░▌▐░░░░░░░░░░░▌▐░░░░░░░░░░░▌▐░░▌    
▐░█▀▀▀▀█░█▀▀ ▐░█▀▀▀▀▀▀▀▀▀  ▀▀▀▀▀▀▀▀▀█░▌▐░▌░▌   
▐░▌     ▐░▌  ▐░▌                    ▐░▌▐░▌▐░▌  
▐░▌      ▐░▌ ▐░█▄▄▄▄▄▄▄▄▄  ▄▄▄▄▄▄▄▄▄█░▌▐░▌ ▐░▌ 
▐░▌       ▐░▌▐░░░░░░░░░░░▌▐░░░░░░░░░░░▌▐░▌  ▐░▌
 ▀   
                         
`
)

//http://patorjk.com/software/taag/#p=testall&h=0&v=0&f=Graceful&t=Resk
func init() {
	fmt.Println(banner1)
}
