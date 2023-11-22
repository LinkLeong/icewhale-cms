import service from '@/utils/request'


export const build = (data)=>{
    return service({
        url:'/ota/build',
        method:'post',
        data
    })
}