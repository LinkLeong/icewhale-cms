<template>
    <div class="flex flex-col gap-10">
        <div class="font-bold text-2xl">构建</div>
        <div class="flex flex-col gap-2">
            <div class="flex w-full gap-2">
                <div class="flex flex-col w-1/2">
                    <div>版本号</div>
                    <el-input v-model="release_version"></el-input>
                </div>
                <div class="flex flex-col w-1/2 h-44">
                    <div>更新日志</div>
                    <textarea 
                        class="h-44"
                        v-model="release_note" />
                </div>
            </div>
            

            <el-button
                v-if="!isBuilding"
                @click="handleBuildBtnClick"
            >构建</el-button>
            <el-button
                v-if="isBuilding"
                disabled>
                构建中
            </el-button>

            <div>
                构建状态: {{status.status}}

                构建消息: {{status.message}}

            </div>
        </div>

        <div class="divider	divider-neutral	"/>
        <div class="flex flex-col">
            <div class="font-bold text-2xl">发布</div>
        </div>
    </div>
</template>

<script setup>
import { ref,reactive,onMounted,onUnmounted } from 'vue'

import { build, buildStatus } from "@/plugin/release/api/ota.js"


const release_version = ref("v0.4.9")
const release_note = ref(`## [0.4.9.1]\n### Fixed\n- Fixed the issue of SMB mounts being lost after a reboot.\n- Fixed the download path issue for OTA updates.\n- Optimized the startup interface.\n- Disabled the display of optical drives.\n- Fixed an issue with the display of messages for inserted disks.\n`)

const status = reactive({
    status: "unstart",
    message: "",
})

const isBuilding = ref(false)
let pollingInterval = null;
const handleBuildBtnClick =()=>{
    console.log(release_note.value)
    console.log(release_version.value)
    isBuilding.value = true
    build({
        version: release_version.value,
        release_note: release_note.value
    })
    pollingInterval = setInterval(fetchData, 2000); // 每2000毫秒轮询一次
}

const fetchData = async () =>{
    const result = await buildStatus()
    status.status = result.data.status
    status.message = result.data.message

    // stop when build finished

}


onUnmounted(() => {
    if (pollingInterval) {
    clearInterval(pollingInterval);
    }
});

</script>
