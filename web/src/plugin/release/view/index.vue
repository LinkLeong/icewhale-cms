<template>
    <div>
        <div class="ota-action-container">
            <div>版本号</div>
            <el-input v-model="release_version"></el-input>
            <div>更新日志</div>
            <textarea v-model="release_note"></textarea>

            <el-button
                v-if="!isBuilding"
                @click="handleBuildBtnClick"
            >构建</el-button>
            <el-button
                v-if="isBuilding"
                disabled>
                构建中
            </el-button>

        </div>
    </div>
</template>

<script setup>
import { ref } from 'vue'

import {build} from "@/plugin/release/api/ota.js"

const release_version = ref("v0.4.9")
const release_note = ref("xxxxxxxxxxx")

const isBuilding = ref(false)

const handleBuildBtnClick =()=>{
    console.log(release_note.value)
    console.log(release_version.value)
    isBuilding.value = true
    build()
}
</script>

<style scoped>
.ota-action-container{
    display: flex;
    flex-direction: column;
    margin-top: 20px;
}
</style>