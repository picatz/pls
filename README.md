# pls

Cross-platform process listing command-line application.

* pid
* ppid
* user
* running
* number of network connections	
* number of threads
* mem percentage
* cpu percentage
* exe
* cmdline
* children pids

## Install

```console
$ go get github.com/picatz/pls
```

## Usage

```console
$ pls
pid     ppid    user	running	connections	threads	mem	        cpu	                    exe	                    cmdline	    children
26743   26731   root	true	0	        6	    0.5134768	0	                    /home/vagrant/pls	    ./pls	    [26751]
504     2       root	true	0	        1	    0	        0			                                                []
1       0       root	true	17	        1	    0.5146581	0.023360224775823342	/lib/systemd/systemd	/lib/systemd/systemd --system --deserialize 19	[568 627 731 1158 1326 1329 1339 1342 1372 1387 1397 1464 1513 1550 1575 1577 1939 3376 4415 4416 4752 7505 12969 22700 26723]
2       0       root	true	0	        1	    0	        0			                                                [3 5 7 8 9 10 11 12 13 15 16 17 18 19 20 21 22 23 24 25 26 27 28 29 33 34 35 36 52 54 56 57 58 59 60 61 62 63 64 65 66 67 73 86 87 126 157 158 159 160 161 162 384 460 503 504 577 589 590 592 593 594 595 598 606 673 824 2440 2651 4165 4166 22681 22747 22793 25046 25076 25081 25465 26720 26726]
...
```

Right now it plays pretty well with tools like `wc`, `grep`, `awk` and `sed` (sorry windows)

```console
# get number of network connections for all processes
$ pls | sed -n '1!p' | awk '{ sum += $5; } END { print sum; }' "$@"
72
```

```console
# get number of processes
$ pls | sed -n '1!p' | wc -l
117
```

```console
# get number of processes
$ pls | sed -n '1!p' | wc -l
117
```

```console
# find process ids that match "fluentd"
$ pls | grep 'fluentd' |  awk '{print $1}'
1939
1944
27952
```