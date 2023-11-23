<template>
    <div >
        <div class="gva-table-box flex flex-col">
            <div class="font-bold text-2xl">构建</div>
            <div class="flex flex-col gap-2">
                <div class="flex flex-col w-full gap-2">
                    <div class="flex flex-col w-1/2 gap-2">
                        <div>版本号</div>
                        <el-input v-model="release_version"></el-input>
                    </div>
                    <div class="flex flex-col w-1/2 gap-2">
                        <div>更新日志</div>
                        <el-input class="" v-model="release_note" type="textarea"  autosize />
                    </div>
                </div>
                

                <div class="">
                    <el-button
                        type="primary"
                        v-if="!isBuilding"
                        @click="handleBuildBtnClick"
                    >构建</el-button>
                    <el-button
                        v-if="isBuilding"
                        disabled>
                        构建中
                    </el-button>
                </div>

                <div class="">
                    <div>
                        构建状态: {{status.status}}
                    </div>
                    <div>
                        构建消息: {{status.message}}
                    </div>
                </div>

            </div>

            <div class="divider	divider-neutral	"/>

            <div class="flex flex-col">
                <div class="font-bold text-2xl">发布</div>
            </div>
        </div>
    </div>
</template>

<script setup>
import { ref,reactive,computed,onUnmounted } from 'vue'

import { build, buildStatus } from "@/plugin/release/api/ota.js"


const release_version = ref("v0.4.9")
const release_note = ref(`## [0.4.9.1]\n### Fixed\n- Fixed the issue of SMB mounts being lost after a reboot.\n- Fixed the download path issue for OTA updates.\n- Optimized the startup interface.\n- Disabled the display of optical drives.\n- Fixed an issue with the display of messages for inserted disks.\n`)

const status = reactive({
    status: "unstart",
    message: "未开始构建",
})

let pollingInterval = null;
const handleBuildBtnClick =()=>{
    console.log(release_note.value)
    console.log(release_version.value)
    isBuilding.value = true
    build({
        version: release_version.value,
        release_note: release_note.value
    })
    status.status = "building"
    pollingInterval = setInterval(fetchData, 2000); // 每2000毫秒轮询一次
}

const isBuilding = computed(() => {
    return status.status === "building"
})

const fetchData = async () =>{
    const result = await buildStatus()
    status.status = result.data.status
    status.message = result.data.message
    if (!isBuilding) {
        clearInterval(pollingInterval);
    }
}


onUnmounted(() => {
    if (pollingInterval) {
    clearInterval(pollingInterval);
    }
});

</script>
