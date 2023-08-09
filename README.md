# revshell
revshell is a reverse shell tool written in Go language, it is the command line version of Reverse Shell Generator, which is designed to help security researchers and penetration testers to quickly generate the appropriate Shell when they need to establish a reverse connection with the target host, users can choose from a wide range of Shell types to generate the bounce Shell code suitable for their own needs, so as to more efficiently interact with the target host and control it in tasks such as penetration testing, vulnerability exploits, or security assessments.

Translated with www.DeepL.com/Translator (free version)

Supported operating systems: Mac, Linux, Windows.

# Installation:
`go build revshell.go`

Or download apps compiled in releases.

## Usage:

`./revshell [IPADDR] [PORT] [LANGUAGE]`

## Supported languages: 

bash|sh|nc|ruby|php|python|rcat|perl|socat|node|telnet|zsh|lua|golang|vlang|awk|crystal

<img width="1210" alt="image" src="https://github.com/BetterDefender/revshell/assets/59255707/f245955e-714f-41db-9f39-abf79c7ba4be">


## Example:

`./revshell 192.168.1.1 2222 bash`

### MacOS：
<img width="1213" alt="image" src="https://github.com/BetterDefender/revshell/assets/59255707/c5c4ea25-9191-4b5c-ab20-b3a86a604de3">

### Windows：
<img width="1087" alt="image" src="https://github.com/BetterDefender/revshell/assets/59255707/59f7ef07-01da-4271-87c7-fa286c143c94">

### Linux：
<img width="1034" alt="image" src="https://github.com/BetterDefender/revshell/assets/59255707/e29555ee-16eb-4a03-96dc-999341d174f3">



