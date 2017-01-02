#!/usr/bin/python
# -*- coding: UTF-8 -*-

import json
import urllib

class RequestModel:
    ModuleName = ""
    MethodName = ""
    Data = []

    '''
    转换为字典
    '''
    def Convert_To_Dic(this):
        result={}
        result["ModuleName"] = this.ModuleName
        result["MethodName"] = this.MethodName
        result["Data"] = this.Data

        return result

# 服务器地址
serverAddr = "http://127.0.0.1:20016/Api"

# 调用接口
def Request(moduleName,methodName,value):
    requestObj = RequestModel()

    requestObj.ModuleName = moduleName
    requestObj.MethodName = methodName
    requestObj.Data = value

    requestData = requestObj.Convert_To_Dic()    
    f = urllib.urlopen(serverAddr, json.dumps(requestData).decode('utf-8'), proxies={})

    responseResult = f.read()
    print responseResult.decode('utf-8')
    

# 测试接口测试
Request("Test","Hello",["123"])

