
#  Warning

#### THIS TOOL IS ONLY MEANT FOR EDUCATIONAL PURPOSES. I WILL NOT BE HELD RESPONSIBLE FOR ANY MISUSE OF THIS TOOL.


# Description

This script is an example of Denial of Service attack scripts. Juggernaut sends obfuscated and
valid get requests to the server which drains its resource pool.
Thus the server will not be able to meet all get requests on time
which will either force the server to throw Resource pool error or
will make the server extremely slow and the webpage will take unbearable
time to load.

The tool contains a user agents text file which ensures that the requests are
different everytime inorder to bypass weak firewalls.


# Requirements

* Python3
* Requests module

# Installation & Usage

`git clone https://github.com/Rajas2323/Juggernaut`


`pip install requests`


`python juggernaut.py`

#### The interface of this software is extremely simple so I don't think there is a need to explain that