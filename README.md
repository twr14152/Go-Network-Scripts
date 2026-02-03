# Go-Network-Scripts
### 02/03/2026
- Updated arista/ssh_client to use PTY session vs shell session for running commands. The output is now cleaner.
  - arista/ssh_client/main.go
### 01/31/2026
- This script uses json-rpc to connect to arista eAPI using hostfile.txt and host.cfg files to run commands on arista devices
  - arista/json-rpc/main.go
### 07/18/2025
- 07/20/2025 Finally got ssh using hostfile working on arista devices
- Arista uses a keyboard interactive auth which requires using a shell session with pty
- Trying to get stdin/stdout working in the same session was a difficult process
- To say nothing about using expect logic for the prompts
- To be honest I thought C with libssh was easier to write and understand for this specific problem, i say that because arista understands string logic with ";" working as a line break. Trying to do this in Go is frustrating. So I went the shell route.
- I think there's something to said about being able to look under the hood with C, and the issues with black box tendencies in Go
- Could be I haven't coded in go in a while but this was a pain in the ass. No likey
- The code can be found here:
  * arista/ssh_using_hostfile/main.go

### 07/06/2025 
Apple Silicon has made virtual labbing almost impossible. Almost, now with Orbstack with containerlabs I now have an outlet. So far cEOS, FRR, SR Linux, are the NOS I'm working with. Can actually ssh to host names..

### 01-29-2025
Loaded ceos image cEOSarm-lab-4.33.1-EFT3.tar.xz on my Mac M1. It runs a lot better than any other image I've loaded. I can connect to ceos hosts using "0.0.0.0:ports" from my terminal. I have not been successful in connecting to the native management ip of the hosts. I've tried the --net=host command in docker create. Didn't have success.  I've read that this is an issue with the Mac. I'm not going to spend much more time on it. For now the "localhost:ports" workaround to connect will work.
- Had to update .eapi.conf to use host=0.0.0.0 and the corresponding ports used in the docker create eg..port = 8443, etc...
- SSH worked by using localhost:202x command or you can update your hostfile to point localhost to ceos1 etc. Your port assignments will distinguish host

### 4/27/2024

I played around with GoEAPI today. I was able to issue show and exec commands farely easily. After some reading I was able to get the configuration functionality worked out. I don't know how heavily used this network automation library is but it seems serviceable. Once you figure out how to deploy configurations it doesnt seem there's too much more to it.

- arista/goeapi/show_commands/main.go
- arista/goeapi/change_config/main.go
- arista/goeapi/deploy_lab_cfgs/main.go

### 4/25/2024

Set up a new ceoslab environment. Had all kinds of issues with my macbook m1. Moved to a windows machine with WSL2. Working like a charm now. Decided that rewritting some of my previous scripts would help me get back into the swing of things. Then I remembered with Arista the default auth is keyboard interactive. So I had to mess around with that. I will give some credit to dgjustice he had a nice example of keyboard interactive on his gist page. That said I incorporated some of his code with mine and now have a useable lab ssh client. The other folders under the arista leave something to be desired. Hopefully will get around to it.

 -arista/ssh_client/main.go



### 4/19/2024

This repo is in need of an update. Code examples are dated and some are non functional. Will update / refresh as time permits.

### 1/30/2022
Finally got back to look at my go2run repo after about a year and found and fixed an issue with my module/packaging set up. The modules in that repo were using local package names for example runcli (which was local to my machine when I created it) when it should have been changed to github.com/twr14152/go2run/runcli when I posted it to github. It ended up breaking the automatic download function of go modules. Anyways thats been fixed and tested. If your wondering why am I posting that here, the answer is I feel its more a of an actual network library ready for individual use than this repo which was a brain dump testing ground for network scripts. So if interested check it out. 

### 02/26/2021
Had a request to create a script that would copy directory over to remote device. I was not able to find a real elegant way to accomplish it, but it works.

- scp_client/copy_dir/scp_scriptv3.go

### 12/20/2020
When I started this repository it was to help me learn Go by using it to do things I know are correct in another field. In that regard I feel the scripts were decent in that they progressed into something that would resemble practical. Practical in one space does not mean elegant or correct in another. That being said in my effort to improve with Go, I have updated some of the more relevant scripts to make them more polished while also adding go modules to help with dependency management. I'm kinda having a love hate relationship with Go.mod but for the basic stuff like pulling more up-to-date packages from the time the code was orginally written it seems to work pretty well.

That said I'm going to try updating and revising the more relevant scripts as I get time. 

Starting with ssh scripts. I'm moving the original scripts to a folder called ssh_client/pre_go.mod_stuff/.
I've added the following scripts in their place. I feel these are ready for practical use. However the testbed I'm still using is 95% Cisco Devnet.

This repository will be used to create packages for the **go2run repo** if the code functionality makes it worthwhile. Otherwise its just a learning / testing and sharing repo.

- ssh_client/runscript/ - You can use this package to make changes and validate your network devices using host and command files.
- ssh_client/runcli/ - You call this package from your main.go file and you can issue show/config commands on multiple remote devices. 
- ssh_client/ssh-cli-client - Easy to use script for connecting to multiple devices and issuing (config/show) commands.
- ssh_client/ssh_client - Configuation / validation script uses a hostfile and seperate configuration files to configure hosts
- scp_client - Script for copying files to remote host


Will update as time permits.

-------------
#  Original Post

This is an attempt to try and learn Go by creating some potentially useful network scripts. Rather than using space and memory on my mac to build a lab, I'll be using Cisco Devnet hosts for testing. At least initially. Why not? Its free and accessible over the Internet... For non-Cisco gear I will use other means.

While I know that there are some 3rd party packages out there for network automation. My goal is to try and learn by building stuff from the ground up. Now the ground I'm standing on may actually be the shoulders of those who built the standard libraries and potentially some third party packages, but you have to start from somewhere.


# SSH Client Scripts
First script was to ssh to Cisco Devnet device and print the output.
It works but its not a real useful script.
```
- ssh_client/can_I_ssh_script.go
```
Second script had some success getting multiple show commands to run on the device.
I used a map which is analogous to a dictionary in python to create my command set. Then used a for loop to run through those commands. 

However it needs to reconnect after each command. Not a fan of that, especially for config mode. 
```
- ssh_client/go_ssh_script_2.go
```
Third script involved using a slice to create the command set and then using a for loop to run through the commands.
Nothing earth shattering but it was an opportunity to create another iterable data set.
```
- ssh_client/go_ssh_script_3.go
```
The fourth SSH script will allow multi-line commands show or config. Previous script used Run() command which has a one command limit per session. This was resolved by using Shell() rather than Run(). 
```
 - ssh_client/ssh_multi_cmd.go 
 
```
The fifth script opens and reads another file to get the commands and then uses a for loop to execute the commands on the remote device.
The benefit of this method is that to make changes to the target host you no longer need to update the script rather the cmd_file.txt 
```
- ssh_client/ssh_use_cmd_file.go
- ssh_client/cmd_file.txt

```
The 6th Script uses standard input to put the commands into the script to issue commands on the remote device.
The benefit of doing this is you do not need to overwrite a cmds_file to make changes. You simply create a new file with what ever commands you want and input it into the go script using the format below.
How to use script: go run ssh_use_stdin.go < cmd_file.txt
```
- ssh_client/ssh_use_stdin.go
- ssh_client/cmd_file.txt
- ssh_client/cmd_file2.txt

```
The 7th script uses host_file to login to target hosts and a cmd_file to issue commands on them. At this point this script would be useful for running a common set of commands on multiple devices. The script will simply loop through the hosts and the commands. I've listed the same router multiple times in the host file to emulate multiple hosts.
```
- ssh_client/ssh_using_host_cmd_files.go
- ssh_client/host_file001.txt 
- ssh_client/cmd_file001.txt
- ssh_using_host_cmd_files_output.txt - shows output of script
```
The 8th script uses os.Args to allow you to enter the commands you want to run at the time you run the script.
For example you could issue "go run interactive_ssh.go show run interface loopback 0, show ip int brief, show version"
This script will split the the commands up by the commas. Update - Refined the scripts code use of os.Args by using strings.method as opposed to using a for loop to build the cmds slice. Newer script is ssh_using_cli_args.go.
```
- ssh_client/ssh_using_cli_args.go
- ssh_client/ssh_using_cli_args_output.txt

- ssh_client/interactive_ssh.go
- ssh_client/interactive_ssh_output.txt
```
The 9th script was an improvement on the 7th script. I really wish I would have labeled the scripts this way but I didnt...Basically I wanted a script that could go out to multiple devices and apply unique configs on them in one go around. This is accomplished by this script by having the main script loop through the devices then use the device hostname +.cfg (it's arbitrary could be .txt suffix) as a target for the file the script opens per device. That way you can uniquely mark and identify the device and the configs the script needs to apply.
```
- ssh_use_host_cmd_files_V2.go
- ios-xe-mgmt-latest.cisco.com\:8181.cfg <-- the \ got added when I created the file by the operating system
- ios-xe-mgmt.cisco.com\:8181.cfg <-- the \ got added when I created the file by the operating system
- ssh_using_host_cmd_files_output.txt - shows output of script
```
The 10th script is another update to the ssh_use_host_cmd_file_V2.go script.  I simply added conditional authentication variables based off the hostname to allow the script to be run against devices with different login creditionals. The test examples were run against two csr's and a nexus.
```
- ssh_client_script_10.go
- config files used
   - sbx-nxos-mgmt.cisco.com\:8181.cfg
   - ios-xe-mgmt-latest.cisco.com\:8181.cfg
   - ios-xe-mgmt.cisco.com\:8181.cfg
- hostfile001.txt <-- updated to include sbx-nxos-mgmt.cisco.com:8181
- ssh_client_script_10_output.txt
```

# HTTP Client Scripts - (Restconf)

I really like this section because it takes advantage of the Go std library for this connectivity. No 3rd party libraries required. Unfortunately restconf is still evolving and documentation could be more user friendly.

The first script uses the https transport protocol to connect to a Cisco Devnet router and GET router configs off the device. As it stands the script is more of a learning tool for Go than a useful production script. More research is needed into Yang and Restconf to truly understand the capabilities. Updated script to use application/yang-data+json format.
```
- http_client/httpGet_v1.go
- http_client/httpGet_v1_output.txt - shows output of script
```
The second script offers the same functionality as the first only it opens hostfile to get its target hosts. It then places those hosts in a slice, then uses a for loop to iterate through the slice and issues the commands on the devices. Updated script to use application/yang-data+json format.
```
- http_client/httpGet_v2.go
- http_client/host_file001.txt
- http_client/httpGet_v2_output.txt - shows output of script
```
The third script is a simple config script using restconf and the POST method. The example simply adds a loopback interface.
```
- http_client/httpPost_v1.go
- http_client/httpPost_v1_output.txt - shows output of script
```
The fourth script is a config script that pull hosts from a host_file add then loop through them to POST the configurations on each device. The configuration file is located in PostCmds.go. Its in the same folder as the main script httpPost_v2.go. So it can call the variable thats in PostCmds.go. This would be useful if you want to apply some common configuration on multiple device. The interface config was simply used to demonstrate the capability. 
```
- httpPost_v2.go - The main script
- PostCmds.go - holds the actual configs for the device 
- host_file002.txt - holds the hosts
- httpPost_v2_output.txt - displays how to run the script and the output.
```
The fifth script in this series is a http POST script that will create an interface and add an ospf routing instance to the router. Applying the interface to an ospf area. The script is using ietf-interface module and CISCO-IOS-XE-native modules. I am still fumbling around with finding the correct uri's to apply to get the script to do what I want. This script is not polished, (dry) does not apply in this instance. More of a can I do this script.
```
- httpPost_v3.go
- httpPost_v3_output.txt - shows pre and post state of device after running script.
```

# Misc Folder
Simply a folder to test Go concepts and features as it could pertain to networking. More about testing features in Go than it is about testing a network concept. 
``` 
- cli_args.go - how to put basic command line arguments in an iterable data format
- methodsTest.go - create a struct and method to print switch config details to screen
- methodsTest_output.txt - script output
- struct_switch_break_loop,go - playing with features for possible future inspiration
- struct_switch_break_loop_output.txt - script output

```
# SCP Client
This is a simple scp client built from the scp - GoDoc file.  
```
- scp_client/scp_client_v1.go
- scp_client/hello.txt

```
# ARISTA CEOS
This is a show commands script will use arista's eapi to issue commands on 3 arista CEOS devices. 
The syntax to run the script is as follows:
>>go run show_cmds.go "show running-config|json", "show ip ospf neighbor|json", "show ip route|json"
```
- arista/show_cmds.go
- arista/show_cmds_output.txt
- arista/show_cmds_output2.txt
```
The configuration script is a work in progress. The documentation surrounding this functionality using Go with eapi is limited.
The script does work but not quite where I would like it to be. Work in progress.
>>>go run config_script.go enable, configure, interface loopback70, ip address 70.70.70.70/32, description testAPI
```
- arista/config_script.go
- arista/config_script_output.txt
- arista/config_script_output2.txt
```

Minor update to the configuration script. It allows you to configure multiple devices with unique configurations. Created new device type struct to distinguish host and configurations. Still a work in progress. 
>>>go run config_script_v2.go
```
- arista/config_script_v2.go
- arista/config_script_v2_output.txt
```

# Using Go modules 
You can use go mod to keep scripts dependencies updated.
```
- using_go_mod_example.txt
```
