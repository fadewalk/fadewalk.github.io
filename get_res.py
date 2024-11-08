import time
import requests
import os
from datetime import datetime

def keep_alive():
    log_dir = "./logs"
    if not os.path.exists(log_dir):
        os.makedirs(log_dir)
    
    log_file = os.path.join(log_dir, "res.log")
    
    while True:
        current_time = datetime.now().strftime("%Y-%m-%d %H:%M:%S")
        with open(log_file, "a") as f:
            f.write(f"[{current_time}] Keeping the Codespace alive...\n")
        print("Keeping the Codespace alive...")
        
        # 发送一个 HTTP 请求到任意网站，模拟活动
        response = requests.get('https://github.com')
        
        with open(log_file, "a") as f:
            f.write(f"[{current_time}] Response status code: {response.status_code}\n")
        print(f"Response status code: {response.status_code}")
        
        # 每隔 5 分钟发送一次请求
        time.sleep(300)

if __name__ == "__main__":
    keep_alive()