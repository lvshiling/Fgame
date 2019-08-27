<template>
    <div class="app-container">
        <el-tabs type="border-card">
            <el-tab-pane label="玩家扶持">
                <SupportPlayer></SupportPlayer>
            </el-tab-pane>
            <el-tab-pane label="扶持日志">
              <SupportPlayerLog></SupportPlayerLog>
            </el-tab-pane>
        </el-tabs>
    </div>
</template>
<script>
import waves from "@/directive/waves"; // 水波纹指令
import SupportPlayer from './supportplayer/supportplayer.vue'
import SupportPlayerLog from './supportplayer/supportlog.vue'
export default {
  name: "OrderStatic",
  components: {
    SupportPlayer,
    SupportPlayerLog
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

