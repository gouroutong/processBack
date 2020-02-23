 ### **server**
 ```
 http://0.0.0.0:3001
 ```
 ### **静态文件地址**
 ```
 http://0.0.0.0:3001/static
 ```

## **home页面**
```
 http://0.0.0.0:3001/api/xprocess/home
```
### **新增项目和修改项目**
API:`/update_project`

请求方式:`post`

 | 参数 | 数据类型 | 是否必须 | 说明 |
 |:---: | :---: | :---:   | :---: |
 | id | int | 否 |  如果没有传 id 则表示新建，否则表示修改 |
 | userId | int | 否 |  用户Id暂时不需要 |
 | projectType | string | 是 |  项目类型 |
 | projectName | string | 是 |  项目名称 |
 | lastUpdatePeople | string | 否 |  暂时不需要 |
 | level | string | 否 |  暂时不需要 |
 | role | string | 否 |  暂时不需要 |

 请求成功返回数据:
 ```go
 {
     "code":2000,
     "result":"新建成功", // 或者 “修改成功”
     "errMsg":""//为空
 }
 ```
 请求失败返回数据:
 ```go
 {
     "code":500,
     "result":"",//为空
     "errMsg":"项目名已存在" //等等其他错误信息
 }
 ```
### **获取project集合**
API:`/get_project`

请求方式:`post`

 | 参数 | 数据类型 | 是否必须 | 说明 |
 |:---: | :---: | :---:   | :---: |
 | search | string | 否 |  如果传该参数表示查询指定名称的 project |
 | type | string | 是 |  指定项目类型 |
 | current | int | 否 |  分页后当前页 |
 | limit | int | 否 |  每页条数 |
 | isAll | bool | 否 |  是否查询所有数据 |

 请求成功返回数据:
 ```go
 {
     "code":2000,
     "result":{
         "list":[
             {
                 "Id": 1,
                 "UserId": 13,
                 "ProjectType": "private",
                 "ProjectName": "测试1",
                 "CreatedAt": "2019-10-14T17:26:06+08:00",
                 "LastUpdatePeople": "hello",
                 "Level": "一般",
                 "Role": "哈哈",
                 "Del": 0
            }
         ]
     },
     "errMsg":""//为空
 }
 ```
 请求失败返回数据:
 ```go
 {
     "code":500,
     "result":"",//为空
     "errMsg":"" //错误信息
 }
 ```
### **项目操作**
API:`/project_opera`

请求方式:`post`

 | 参数 | 数据类型 | 是否必须 | 说明 |
 |:---: | :---: | :---:   | :---: |
 | project_id | int | 是 |  对应的 project的id|
 | operation | string | 是 |  进行的操作 |
`operation`详细介绍:
```javascript
"del":删除
"copy":copy
```
 请求成功返回数据:
 ```go
 {
     "code":2000,
     "result":"删除成功", // 等等
     "errMsg":"" //为空
 }
 ```
 请求失败返回数据:
 ```go
 {
     "code":500,
     "result":"",//为空
     "errMsg":"删除失败" //等等其他错误信息
 }
 ```
### **最近浏览**
API:`/browse`

请求方式:`post`

 | 参数 | 数据类型 | 是否必须 | 说明 |
 |:---: | :---: | :---:   | :---: |
 | id | int | 是 |  用户 id|

 请求成功返回数据:
 ```go
 {
     "code":2000,
     "result":[1,2,3,4,5], //数据源id的集合，长度为5的数组，索引越小表示越近
     "errMsg":"" //为空
 }
 ```
 请求失败返回数据:
 ```go
 {
     "code":500,
     "result":"",//为空
     "errMsg":"error" //错误信息
 }
 ```
### **获取回收站内容**
API:`/dustbin`

请求方式:`post`

 | 参数 | 数据类型 | 是否必须 | 说明 |
 |:---: | :---: | :---:   | :---: |
 | id | int | 是 |  用户 id|

 请求成功返回数据:
 ```go
 {
     "code":2000,
    "result":{
         "list":[
             {
                 "Id": 1,
                 "UserId": 13,
                 "ProjectType": "private",
                 "ProjectName": "测试1",
                 "CreatedAt": "2019-10-14T17:26:06+08:00",
                 "LastUpdatePeople": "hello",
                 "Level": "一般",
                 "Role": "哈哈",
                 "Del": 1
            }
         ]
     },
     "errMsg":"" //为空
 }
 ```
 请求失败返回数据:
 ```go
 {
     "code":500,
     "result":"",//为空
     "errMsg":"error" //错误信息
 }
 ``` 
### **回收站操作**
API:`/dustbin_opera`

请求方式:`post`

 | 参数 | 数据类型 | 是否必须 | 说明 |
 |:---: | :---: | :---:   | :---: |
 | id | []int | 是 |  用户id的数组|
 | operation | string | 是 |  操作类型|
 `operation`详细介绍:
```go
"del":删除指定数据
"restore":还原指定数据
"delAll":清空回收站
```
 请求成功返回数据:
 ```go
 {
     "code":2000,
     "result":"删除成功", //等其他信息
     "errMsg":"" //为空
 }
 ```
 请求失败返回数据:
 ```go
 {
     "code":500,
     "result":"",//为空
     "errMsg":"error" //错误信息
 }
 ```
## **Analysis分析页面**
```
 http://0.0.0.0:3001/api/xprocess/analysis
```
### **新增或者修改analysis**
API:`/update_analysis`

请求方式:`post`

 | 参数 | 数据类型 | 是否必须 | 说明 |
 |:---: | :---: | :---:   | :---: |
 | id | int | 否|  analysis的id，若不传表示新建|
 | projectId | int | 是 |  项目Id|
 | analysisName | string | 是 |  分析类型|
 | DataModel | string | 否 |  绑定数据源名称|
 | DataId | int | 否 |  绑定数据源id|
 请求成功返回数据:
 ```go
 {
     "code":2000,
     "result":"新建成功", //等其他信息
     "errMsg":"" //为空
 }
 ```
 请求失败返回数据:
 ```go
 {
     "code":500,
     "result":"",//为空
     "errMsg":"error" //错误信息
 }
 ```
### **获取analysis集合**
API:`/get_all`

请求方式:`post`

 | 参数 | 数据类型 | 是否必须 | 说明 |
 |:---: | :---: | :---:   | :---: |
 | projectId | int | 是 |  项目id |
 | current | int | 否 |  分页后当前页 |
 | limit | int | 否 |  每页条数 |
 | isAll | bool | 否 |  是否查询当前项目下所有分析 |

 请求成功返回数据:
 ```go
 {
     "code":2000,
     "result":{
         "list":[
             {
                 "Id": 1,
                 "ProjectId": 1,
                 "AnalysisName": "",
                 "ProjectName": "测试1",
                 "CreatedOn": "2019-10-14T17:26:06+08:00",
                 "LastChange": "2019-10-14T17:26:06+08:00",
                 "DataModel": "测试1",
                 "DataId": 1,
            }
         ]
     },
     "errMsg":""//为空
 }
 ```
 请求失败返回数据:
 ```go
 {
     "code":500,
     "result":"",//为空
     "errMsg":"" //错误信息
 }
 ```
### **删除分析**
API:`/delete`

请求方式:`post`

 | 参数 | 数据类型 | 是否必须 | 说明 |
 |:---: | :---: | :---:   | :---: |
 | id | int | 是 |  分析 id|

 请求成功返回数据:
 ```go
 {
     "code":2000,
     "result":"删除成功" 
     "errMsg":"" //为空
 }
 ```
 请求失败返回数据:
 ```go
 {
     "code":500,
     "result":"",//为空
     "errMsg":"error" //错误信息
 }
 ```
## **data页面**
```
http://0.0.0.0:3001/api/xprocess/ipt_data
```
### **导入数据（新增数据源）**
API:`/new`

请求方式:`post`

 | 参数 | 数据类型 | 是否必须 | 说明 |
 |:---: | :---: | :---:   | :---: |
 | file | (暂时不太确定) | 是 |  上传的文件 |
 | filename | string | 是 |  文件名 |

 请求成功返回数据:
 ```go
 {
     "code":2000,
     "result":"file/cy.csv" // filePath
     "errMsg":"" //为空
 }
 ```
 请求失败返回数据:
 ```go
 {
     "code":500,
     "result":"",//为空
     "errMsg":"error" //错误信息
 }
 ```
### **获取当前项目下所有data**
API:`/get_data`

请求方式:`post`

 | 参数 | 数据类型 | 是否必须 | 说明 |
 |:---: | :---: | :---:   | :---: |
 | projectId | int | 是 |  项目id |
 | current | int | 否 |  分页当前页数 |
 | limit | int | 否 |  分页每页条数 |
| isAll | bool | 否 |  是否查询当前项目下所有data |
 请求成功返回数据:
 ```go
 {
     "code":2000,
     "result":{
         "list" :[
             {
                 "Id": 1,
                 "ProjectId": 1,
                 "Name": "cy.csv",
                 "CaseNum": 10,
                 "CreateName": "",
                 "CreatedAt": "2019-10-14T17:26:15+08:00",
                 "LastUpdateName": "",
                 "UpdatedAt": "2019-10-15T11:40:52+08:00",
                 "FilePath": "file/cy.csv",
                 "Fields": "[\"Case ID\",\"Start Timestamp\",\"Complete Timestamp\",\"Activity\",\"Resource\",\"Role\"]",
                 "MapFields": "[{\"name\":\"CaseId\",\"from\":\"Case ID\"},{\"name\":\"StartAt\",\"from\":\"Start Timestamp\"},{\"name\":\"EndAt\",\"from\":\"Complete Timestamp\"},{\"name\":\"Activity\",\"from\":\"Activity\"},{\"name\":\"Role\",\"from\":\"Role\"},{\"name\":\"Labor\",\"from\":\"Resource\"}]",
                 "UsedAnalysis": "",
                 "UsedAnalysisId": ""
             },
             {
                 ...
             }
         ]
     } 
     "errMsg":"" //为空
 }
 ```
 请求失败返回数据:
 ```go
 {
     "code":500,
     "result":"",//为空
     "errMsg":"error" //错误信息
 }
 ```
### **保存新导入数据或编辑**
API:`/save`

请求方式:`post`

 | 参数 | 数据类型 | 是否必须 | 说明 |
 |:---: | :---: | :---:   | :---: |
 | id | int | 否/是 | 若新导入不需要该参数，编辑需要 |
 | projectId | int | 是/否 | 所属项目id,新建需要，编辑不修改此参数则不需要 |
 | name | string | 是/否 | 数据源name,新建需要，编辑不修改此参数则不需要 |
 | filePath | string | 是/否 | 文件上传后路径,编辑不修改此参数则不需要 |
 | mapFields | string | 是/否 | 字段映射,新建需要，编辑不修改此参数则不需要 |

 请求成功返回数据:
 ```go
 {
     "code":2000,
     "result":"修改成功" 
     "errMsg":"" //为空
 }
 ```
 请求失败返回数据:
 ```go
 {
     "code":500,
     "result":"",//为空
     "errMsg":"error" //错误信息
 }
 ```
### **预览字段映射**
API:`/preview`

请求方式:`post`

 | 参数 | 数据类型 | 是否必须 | 说明 |
 |:---: | :---: | :---:   | :---: |
 | filePath | string | 是 | 上传文件路径 |

 请求成功返回数据:
 ```go
 {
     "code":2000,
    // result 结果data是一个 [][]string,其中data[0]是字段名，之后的元素是按顺序对应字段名的数据
    // result 结果数组最长为6，即5条数据
     "result": [
                       [
                           "case_id",
                           "labor",
                           "start_timestamp",
                           "complete_timestamp",
                           "role",
                           "concept_name"
                       ],
                       [
                           "1",
                           "a",
                           "2011-02-16 14:31:00",
                           "2011-02-16 14:31:00",
                           "A",
                           "Request created"
                       ]
                ]
     "errMsg":"" //为空
 }
 ```
 请求失败返回数据:
 ```go
 {
     "code":500,
     "result":"",//为空
     "errMsg":"error" //错误信息
 }
 ```
### **删除选中数据源**
API:`/delete`

请求方式:`post`

 | 参数 | 数据类型 | 是否必须 | 说明 |
 |:---: | :---: | :---:   | :---: |
 | dataId | int | 是 | 数据源id |

 请求成功返回数据:
 ```go
 {
     "code":2000,
     "result":"success" 
     "errMsg":"" //为空
 }
 ```
 请求失败返回数据:
 ```go
 {
     "code":500,
     "result":"",//为空
     "errMsg":"error" //错误信息
 }
 ```
### **导出指定数据源为指定格式文件**
API:`/export`

请求方式:`post`

 | 参数 | 数据类型 | 是否必须 | 说明 |
 |:---: | :---: | :---:   | :---: |
 | id | int | 是 | 数据源id |
 | file_type | string | 是 | 导出的文件格式,例如".xml" |

 请求成功返回数据:
 ```go
 {
     "code":2000,
     "result":"192.168.0.1:3001/...."  // 导出文件的下载地址
     "errMsg":"" //为空
 }
 ```
 请求失败返回数据:
 ```go
 {
     "code":500,
     "result":"",//为空
     "errMsg":"error" //错误信息
 }
 ```
## **page页面**
```
http://0.0.0.0:3001/api/xprocess/custom
```
### **获取数据源**
API:`/get_datasource`

请求方式:`post`

 | 参数 | 数据类型 | 是否必须 | 说明 |
 |:---: | :---: | :---:   | :---: |
 | dataId | int | 是 | 数据源id |

 请求成功返回数据:
 ```go
 {
     "code":2000,
     "result":{
     	    "date": [
     	        {
     		    "date": "2011-01-01",
     		    "caseNum": 9,
     		    "eventNum": 42,
     		    "loop_num": 5,
     		    "variantNum": 8,
     		    "avgLeadTime": 1434,
     		    "newCase": 9,
     		    "endedCaseNum": 9
     		    }
     		]
     }
     "errMsg":"" //为空
 }
 ```
 请求失败返回数据:
 ```go
 {
     "code":500,
     "result":"",//为空
     "errMsg":"error" //错误信息
 }
 ```
### **新建或者修改page**
API:`/update`

请求方式:`post`

 | 参数 | 数据类型 | 是否必须 | 说明 |
 |:---: | :---: | :---:   | :---: |
 | id | int | 否 | page id，若不传表示新建 |
 | analysisId | int | 否 | 新建必须，修改可以不传|
 | pageName | string | 是 | page名称 |
 | dataSource | string | 否 | 数据源key 如`date` |
 | styleMap | string | 否 | 页面内card布局即配置详细信息 |

 请求成功返回数据:
 ```go
{
    "code": 2000,
    "result": "修改成功",// 或者新建成功
    "errMsg": ""
}
 ```
 请求失败返回数据:
 ```go
 {
     "code":500,
     "result":"",//为空
     "errMsg":"error" //错误信息
 }
 ```
### **获取当前项目下所有page**
API:`/get_page_data`

请求方式:`post`

 | 参数 | 数据类型 | 是否必须 | 说明 |
 |:---: | :---: | :---:   | :---: |
 | analysisId | int | 是 | 项目id|
 | current | int | 否 | 当前页数，默认为1|
 | limit | int | 否 | 每页数据条数，默认为20 |
 | isAll | bool | 否 | 是否是查询所有 |

 请求成功返回数据:
 ```go
{
    "code": 2000,
    "result": "修改成功",// 或者新建成功
    "errMsg": ""
}
 ```
 请求失败返回数据:
 ```go
 {
     "code":500,
     "result":"",//为空
     "errMsg":"error" //错误信息
 }
 ```

### **page页面删除等操作**
API:`/operation`

请求方式:`post`

 | 参数 | 数据类型 | 是否必须 | 说明 |
 |:---: | :---: | :---:   | :---: |
 | pageId | int | 是 | id|
 | operation | string | 是 | 操作类型，如 `del`|

 请求成功返回数据:
 ```go
{
    "code": 2000,
    "result": "删除成功",
    "errMsg": ""
}
 ```
 请求失败返回数据:
 ```go
 {
     "code":500,
     "result":"",//为空
     "errMsg":"error" //错误信息
 }
 ```