##Add nvme storage to the Elitedesk 800g2

### Configure disk
```
#lsblk
  nvme0n1 259:0    0 465.8G  0 disk


#sudo fdisk -l /dev/nvme0n1

Disk /dev/nvme0n1: 465.78 GiB, 500107862016 bytes, 976773168 sectors
Disk model: KINGSTON SNVS500G                       
Units: sectors of 1 * 512 = 512 bytes
Sector size (logical/physical): 512 bytes / 512 bytes
I/O size (minimum/optimal): 512 bytes / 512 bytes

#sudo gdisk /dev/nvme0n1

#sudo mkfs.ext4 /dev/nvme0n1

#sudo blkid /dev/nvme0n1

#sudo vi /etc/fstab
UUID=11111111-1111-1111-1111-111111111111 /mnt/nvme ext4 defaults 0 0

#sudo mkdir /mnt/nvme

#sudo mount -a

#df -hT |grep ext4
```
