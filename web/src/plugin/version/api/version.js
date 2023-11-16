import service from '@/utils/request'

// @Tags System
// @Summary 获取日志列表
// @Security ApiKeyAuth
// @Produce  application/json
// @Param data body email_response.Email true "发送邮件必须的参数"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"发送成功"}"
// @Router /email/sendEmail [post]
export const getVersion = () => {
  return service({
    url: '/version/list',
    method: 'get',
  })
}
export const addVersion = (data) => {
  return service({
    url: '/version/add',
    method: 'post',
    data
  })
}

export const updateVersion = (data) => {
  return service({
    url: '/version/update',
    method: 'post',
    data
  })
}

export const deleteVersion = (data) => {
  return service({
    url: '/version/'+data.ID,
    method: 'delete'
  })
} 