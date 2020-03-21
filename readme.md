# 僵尸企业分析

> BaseUrl：www.zerokirin.online/api/jxqy

## 校验方式

请求头中加入字段：sign

内容：KJb4Zq1M1F9BHfuA

## 上传文件

- 地址：/upload

- 方法：POST

- 方式：form表单传值

- 数据说明

  - code：每次上传都不一样，作为一次上传的标识码
  - file1：第一个文件，依此类推到file4

- 成功返回示例

  ```json
  {
      "data": {
          "file1": {
              "Filename": "new 1",
              "Header": {
                  "Content-Disposition": [
                      "form-data; name=\"file1\"; filename=\"new 1\""
                  ],
                  "Content-Type": [
                      "application/octet-stream"
                  ]
              },
              "Size": 11
          },
      },
      "error": 0,
      "msg": "success"
  }
  ```

  ### 错误对照

  | error |        含义        |
  | :---: | :----------------: |
  | 40300 |   没有sign请求头   |
  | 40301 |      code重复      |
  | 50001 |    接收文件失败    |
  | 50002 |    保存文件失败    |
  | 50003 | 创建code文件夹失败 |

  

## 验证是否上传成功

- 地址：/find/{code}

- 方法：GET

- 实例：/find/0001

- 成功返回示例

  Time：上传时间

  Name：文件名

  ```json
  {
      "data": [
          {
              "Time": "2020-03-22T03:01:14.9581417+08:00",
              "Name": "new 1"
          },
          {
              "Time": "2020-03-22T03:01:14.9601333+08:00",
              "Name": "new 2"
          },
          {
              "Time": "2020-03-22T03:01:14.9811217+08:00",
              "Name": "new 3"
          },
          {
              "Time": "2020-03-22T03:01:14.9820737+08:00",
              "Name": "new 4"
          }
      ],
      "error": 0,
      "msg": "success"
  }
  ```

  ### 错误对照

  | error |           含义           |
  | :---: | :----------------------: |
  | 40300 |      没有sign请求头      |
  | 50000 | code不存在（上传没成功） |

  