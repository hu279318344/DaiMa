# -*- coding: utf8 -*-
import ssl
import json
import smtplib
import requests
import urllib.request
import time
from email.mime.text import MIMEText
from email.header import Header
from email.utils import formataddr

ssl._create_default_https_context = ssl._create_unverified_context
userAgent = {"user-agent": "Mozilla/5.0 (X11; Linux x86_64; rv:45.0) Gecko/20100101 Firefox/45.0"} # 添加一个user-agent,访问反爬虫策略严格的网站很有用
timeOut = 0.04 # 请求超时时间设置为4秒

def sendmail(content):
        msg=MIMEText(content,'plain','utf-8')
        to_addrs = [formataddr(["Server","sikaichao@126.com"]),formataddr(["Server","yanghu@dyhqc.com"])] 
        msg['Subject'] = "站点监控"+"|"+time.strftime("%Y-%m-%d %H:%M:%S", time.localtime())
        msg['From'] = formataddr(["Server","yanghu@dyhqc.com"])
        msg['To'] = Header(",".join(to_addrs))
        try:
            server = smtplib.SMTP_SSL("smtp.mxhichina.com", 465)
            server.login('yanghu@dyhqc.com', 'yanghu1989##')
            server.sendmail("yanghu@dyhqc.com",to_addrs,msg.as_string())
            server.quit()
        except Exception as e:
            print("send mail faild %s"%e)

def getStatusCode(url,userAgent,timeOut):
        try:
            request = requests.get(url, headers = userAgent, timeout = timeOut)
            httpStatusCode = request.status_code
            return httpStatusCode
        except requests.exceptions.HTTPError as e:
            return e

def getWebTime(url): 
        # 毫秒
        return requests.get(url).elapsed.microseconds/1000

def main():
        url_list = ["http://server1.cdce.cn/","http://www.cdce.cn/","https://www.baidu.com"]
        content = []
        for url in url_list:
            try:
                content_str = "网站%s访问状态码为%s,响应耗时为%s毫秒\n"%(url,getStatusCode(url,userAgent,timeOut),getWebTime(url))
                content.append(content_str)   
            except Exception as e:
                content_str = "网站 %s 访问超过%s秒无应答\n\n"%(url,timeOut)
                content.append(content_str)
        content=''.join(content)
        sendmail(content)

if __name__ == '__main__':
        main()

