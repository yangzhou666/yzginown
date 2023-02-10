/**
*@author:yangzhou
*@date: 2023/2/9
*@email: yangzhou2224@shengtian.com
*@description:
 */
package response

type PageResult struct {
	Data     interface{} `json:"data"`
	Total    int64       `json:"total,omitempty"`
	Page     int         `json:"current_page,omitempty"`
	Limit    int   `json:"per_page,omitempty"`     // 每页20条
}
