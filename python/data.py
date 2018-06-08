#!/usr/bin/python
# -*- coding: UTF-8 -*-

import MySQLdb
import time

while 1:
    localtime = time.localtime(time.time())
    #print "本地时间为 :", localtime
    #print localtime.tm_wday
    if localtime.tm_wday == 4:
        # 打开数据库连接
        db = MySQLdb.connect(host="127.0.0.1",port=3306,user="root",passwd="root",db="data",charset="utf8")

        # 使用cursor()方法获取操作游标
        cursor = db.cursor()

        # SQL 查询语句
        sql = "SELECT count(1) FROM order"
        try:
           # 执行SQL语句
           cursor.execute(sql)
           # 获取所有记录列表
           results = cursor.fetchall()
           for row in results:
              #print results
              fname = row[0]
              # 打印结果
              #print "fname=%s" % (fname)
              print fname
        except:
           print "Error: unable to fecth data"

        # 关闭数据库连接
        db.close()
    time.sleep(20)
