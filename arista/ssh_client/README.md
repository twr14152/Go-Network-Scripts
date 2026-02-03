### Updated 02/03/2026
This script was updated to use PTY session rather than shell session for running commands. Output looks cleaner now.

This was more of a can i do it script. The default authentication method used for arista is keyboard interactive. This is different from Cisco by default.
This script is more of a screen scraping cli script that could be used to connect to multiple device and issue various commands (configs or show commands). 
Good to use in a lab. 
```
todd@todd-TOSHIBA-DX735:~/Code_folder/containerlab/containerlabs_sandbox/ceos_lab/lab4/scripts/go_ssh$ ./go_ssh
Number of hosts: 2

Enter host and ssh port being used. (Ex. x.x.x.x:22 or hostname:22)

Enter host: leaf2:22

cmds: enable, show ip bgp summ, show ip int brief
Enter host: spine2:22

cmds: enable, show ip bgp summ, show ip int brief
++++++++++++++++++++++++++++++++
Connected to: leaf2:22
++++++++++++++++++++++++++++++++
Pagination disabled.
BGP summary information for VRF default
Router identifier 1.1.1.3, local AS number 65002
Neighbor Status Codes: m - Under maintenance
  Neighbor      V AS           MsgRcvd   MsgSent  InQ OutQ  Up/Down State   PfxRcd PfxAcc PfxAdv
  10.1.0.4      4 65000           5390      5384    0    0    3d04h Estab   3      3      4
  10.1.0.6      4 65000           5394      5386    0    0    3d04h Estab   3      3      5
  100.100.100.1 4 65000           5538      5593    0    0    3d04h Estab   3      3      7
  100.100.100.2 4 65000           5541      5611    0    0    3d04h Estab   3      3      7
                                        
Interface   IP Address   Status Protocol
----------- ------------ ------ --------
Ethernet1   10.1.0.5/31  up     up      
Ethernet2   10.1.0.7/31  up     up      
Loopback0   1.1.1.3/32   up     up      
Loopback1   1.1.1.103/32 up     up      
Loopback69  unassigned   up     up      
Management0 10.0.0.6/24  up     up      
Vlan20      10.0.20.4/24 up     up      

                              Address
Interface              MTU    Owner  
----------------- ----------- -------
Ethernet1             1500           
Ethernet2             1500           
Loopback0            65535           
Loopback1            65535           
Loopback69           65535           
Management0           1500           
Vlan20                1500           

++++++++++++++++++++++++++++++++
Connected to: spine2:22
++++++++++++++++++++++++++++++++
Pagination disabled.
BGP summary information for VRF default
Router identifier 100.100.100.2, local AS number 65000
Neighbor Status Codes: m - Under maintenance
  Neighbor V AS           MsgRcvd   MsgSent  InQ OutQ  Up/Down State   PfxRcd PfxAcc PfxAdv
  1.1.1.2  4 65001           5560      5473    0    0    3d04h Estab   2      2      5
  1.1.1.3  4 65002           5611      5541    0    0    3d04h Estab   2      2      5
  10.1.0.3 4 65001           5388      5397    0    0    3d04h Estab   2      2      3
  10.1.0.7 4 65002           5386      5395    0    0    3d04h Estab   2      2      3
                                       
Interface    IP Address        Status  
------------ ----------------- --------
Ethernet3    10.1.0.2/31       up      
Ethernet4    10.1.0.6/31       up      
Loopback0    100.100.100.2/32  up      
Loopback69   unassigned        up      
Management0  10.0.0.3/24       up      

                                Address
Interface    Protocol      MTU  Owner  
------------ ---------- ------- -------
Ethernet3    up           1500         
Ethernet4    up           1500         
Loopback0    up          65535         
Loopback69   up          65535         
Management0  up           1500         

todd@todd-TOSHIBA-DX735:~/Code_folder/containerlab/containerlabs_sandbox/ceos_lab/lab4/scripts/go_ssh$ 


```
