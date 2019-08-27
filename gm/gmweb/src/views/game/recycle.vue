<template>
    <div class="app-container">
        <el-tabs type="border-card">
            <el-tab-pane label="设置回收">
                <RecycleState></RecycleState>
            </el-tab-pane>
            <el-tab-pane label="回收日志">
              <RecycleLog></RecycleLog>
            </el-tab-pane>
        </el-tabs>
    </div>
</template>
<script>
import waves from "@/directive/waves"; // 水波纹指令
import RecycleState from './recycle/recycleset.vue'
import RecycleLog from './recycle/recyclelog.vue'
export default {
  name: "recycle",
  components: {
    RecycleState,
    RecycleLog
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

