import service from '@/utils/request'


export const build = (data)=>{
    return service({
        url:'/ota/build',
        method:'post',
        data
    })
}

export const buildStatus = ()=>{
    return service({
        url:'/ota/build-status',
        method:'get' 
    })
}