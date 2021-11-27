## Remote management 

### Configure

#### Enable AMT
1. Press F10 during boot
2. Select advanced
3. Select Remote Managment Options
4. Enable 'Activate Management (AMT)'
5. Save & Exit


#### Setup ME
1. Press ESC during boot
2. Select "ME Setup (F6)"
3. Select Generic..
4. Set a new password
5. Select configure AMT
6. Set hostname to "k8s10..."
7. Change user consensus from KVM to None
8. Enable ME
9. Save & Exit
10. Reboot

### Install Meshcommander
```
docker run -d -p 127.0.0.1:3000:3000 --name meshcommander vga101/meshcommander
```


### Note
Remote desktop will not work without a monitor connected.   
Workaroud: Use a vga dummy adapter (buy or diy)
