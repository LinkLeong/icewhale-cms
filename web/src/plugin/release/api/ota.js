import service from '@/utils/request'


export const build = ()=>{
    return service({
        url:'/ota/build',
        method:'post'
    })
}