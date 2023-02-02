package main

import ("fmt"; "os"; "strings")

func main() {
    if (len(os.Args) == 1 || ArrayContains(os.Args[1:], "--help") || ArrayContains(os.Args[1:], "-h")) {
      fmt.Println("Usage: genshell [OPTIONS | EXECUTER] IP:PORT")
      fmt.Println("     \\ genshell 192.168.1.1:10000")
      fmt.Println("     \\ genshell powershell_3 192.168.1.1:10000") 
      fmt.Println("     \\ genshell --shell=bash -e perl_no_sh 192.168.1.1:10000\n")
    
      fmt.Println("Options")
      fmt.Println("-------")
      fmt.Println(" --help                 Show this screen")
      fmt.Println(" -s / --shell           Specify an unix shell (sh, bash, zsh, etc.) [default: sh]")
      fmt.Println(" -e / --exec            Program responsible for the shell (netcat, python3, etc.) [default: python3_short]") 
      fmt.Println(" --show-exec            Show all programs available\n")

      fmt.Println("Most shells come from https://www.revshells.com/ >> Make sure to check it out.") 
      os.Exit(0) 
  }
 
  if (ArrayContains(os.Args, "--show-exec")) {
    fmt.Println("bash_i, bash_196, bash_readline, bash_5, bash_udp, nc_mkfifo, nc_e, nc_exe, busybox_nc_3, nc_c, ncat_e, ncat.exe_e, ncat_udp, rustcat, c_lang, c_windows, csharp_tcp_client, c_sharp_bash, haskell_1, perl, perl_no_sh, php_cmd, php_cmd_2, php_cmd_small, php_exe, php_shell_exec, php_system, php_passthru, php_r, php_popen, php_proc_open, windows_con_pty, powershell_1, powershell_2, powershell_3, powershell_4_TLS, python_1, python_2, python3_1, python3_2, python3_windows, python3_short, ruby_1, ruby_no_sh, socat_1, node_js, node_js_2, java_1, java_2, java_3, javascript, groovy, telnet, zsh, lua_1, lua_2, golang, vlang, awk, dart") 
    os.Exit(0)
  }

  address := os.Args[len(os.Args) - 1]
  ip, port := "", ""

  for i := 0; i < len(address); i++ {
    if string(address[i]) == ":" {port = address[i + 1:]; break;}
    ip += string(address[i])
  }

  options := os.Args[1:len(os.Args) - 1]
  shell := GetArg(options, [2]string {"--shell", "-s"}, "sh")

  exec := ""
  
  if len(options) == 1 && !strings.Contains(options[0], "=") {exec = string(options[0])
  } else {exec = GetArg(options, [2]string {"--exec", "-e"}, "python3_short")}
  
  data, _ := os.ReadFile("/etc/genshell.conf")
  payloads := string(data)
  
  title := "### " + exec + " ###"
  title_index := strings.Index(payloads, title)
  
  if (title_index == -1) {
    fmt.Println("error: exec not found. Type --show-exec to see all values.")
    os.Exit(1)
  }
  
  command := ""

  for i := title_index + len(title) + 1; i < len(payloads); i++ {
    command += string(payloads[i])
    if string(payloads[i+1]) + string(payloads[i+2]) + string(payloads[i+3]) == "###" {break;}
  }

  command = strings.TrimSpace(command)
  command = strings.ReplaceAll(command, "{ip}", ip)
  command = strings.ReplaceAll(command, "{port}", port)
  command = strings.ReplaceAll(command, "{shell}", shell)
  fmt.Println(command)
}

func ArrayContains(s []string, str string) bool {
  for _, v := range s {if v == str {return true}}
  return false
}

func GetArg(args []string, flags [2]string, default_value string) string {
  value := default_value

  for i := 0; i < len(args); i++ {
    have_arg := strings.Contains(args[i], flags[0]) || strings.Contains(args[i], flags[1])
    
    if have_arg {
      equal_index := strings.Index(args[i], "=")
      
      if equal_index != -1 {
        value = ""
        for j := equal_index + 1; j < len(args[i]); j++ {value += string(args[i][j]);}
      } else {value += args[i + 1];}
      
    }
    break;
  }
  return value;
}

