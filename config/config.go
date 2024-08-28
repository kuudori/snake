package config

import "time"

const SCORE = 0
const DEFAULT_SPEED = 100 * time.Millisecond

const StartTitle = "Press ENTER to start or H for Help"

const LogoTitle = `
_______    __   __    ______    __  __    ______    
/\  ___\  /\ "-.\ \  /\  __ \  /\ \/ /   /\  ___\   
\ \___  \ \ \ \-.  \ \ \  __ \ \ \  _"-. \ \  __\   
 \/\_____\ \ \_\\"\_\ \ \_\ \_\ \ \_\ \_\ \ \_____\ 
  \/_____/  \/_/ \/_/  \/_/\/_/  \/_/\/_/  \/_____/ 
`
const HelpTitle = `
Controls:                            
↑↓←→ -  Move the snake
P            -  Pause/Resume game
PgUp         -  Speed up game  
PgDown       -  Slow down game        
Esc          -  Quit game      
                                      
Press С to return to the menu
`
