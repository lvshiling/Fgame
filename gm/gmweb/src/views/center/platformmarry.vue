<template>
    <div class="app-container">
        <el-tabs type="border-card">
            <el-tab-pane label="结婚价格版本配置">
                <PlatformMarrySet></PlatformMarrySet>
            </el-tab-pane>
            <el-tab-pane label="结婚价格版本日志">
              <PlatformMarrySetLog></PlatformMarrySetLog>
            </el-tab-pane>
        </el-tabs>
    </div>
</template>
<script>
import waves from "@/directive/waves"; // 水波纹指令
import PlatformMarrySet from './platform_marry_set.vue'
import PlatformMarrySetLog from './platform_marry_log.vue'
export default {
  name: "CenterPlatformMarry",
  components: {
    PlatformMarrySet,
    PlatformMarrySetLog
  },
  directives: {
    waves
  },
  created() {},
  data() {
    return {};
  },
  methods: {}
};
</script>

